---
name: separation-of-concerns-pattern-overview - Reference
description: Complete reference for Separation of Concerns Pattern Overview
---

# Separation of Concerns Pattern Overview - Reference

This is the complete reference documentation extracted from the source.


# Separation of Concerns Pattern Overview

> **One Responsibility Per Component**
>
>
> Every component should do one thing well. Orchestration logic separated from business logic. Testability through clear boundaries. This pattern is the foundation of maintainable systems.
>

## Intent

**Separate distinct responsibilities into isolated components with clear boundaries.**

Each component handles one concern. CLI presentation lives in `cmd/`. Business logic lives in `pkg/`. Tests run without external dependencies. Changes are localized. Systems remain maintainable at scale.

---

## Motivation

### When to Use This Pattern

You need separation when:

- **Testing requires external systems** - Database, Kubernetes cluster, container registry
- **Changes ripple across unrelated code** - Fixing a bug breaks unrelated features
- **New team members struggle to understand flow** - Control flow crosses multiple abstraction layers
- **Multiple concerns mix in one function** - Validation, transformation, persistence in single handler

### The Cost of Mixed Concerns

```go
// Bad: CLI, business logic, and I/O mixed together
func DeployCommand(cmd *cobra.Command, args []string) error {
    // Parsing flags (CLI concern)
    namespace, _ := cmd.Flags().GetString("namespace")
    image, _ := cmd.Flags().GetString("image")

    // Validation (business logic)
    if namespace == "" {
        return fmt.Errorf("namespace required")
    }

    // Kubernetes client creation (infrastructure)
    config, _ := clientcmd.BuildConfigFromFlags("", kubeconfig)
    clientset, _ := kubernetes.NewForConfig(config)

    // Deployment logic (business logic)
    deployment := &appsv1.Deployment{
        ObjectMeta: metav1.ObjectMeta{Name: "app"},
        Spec: appsv1.DeploymentSpec{
            Template: corev1.PodTemplateSpec{
                Spec: corev1.PodSpec{
                    Containers: []corev1.Container{{
                        Name:  "app",
                        Image: image,
                    }},
                },
            },
        },
    }

    // API call (infrastructure)
    _, err := clientset.AppsV1().Deployments(namespace).Create(ctx, deployment, metav1.CreateOptions{})
    return err
}
```

**Problems**:

- Cannot test without Kubernetes cluster
- Business logic trapped in CLI layer
- Impossible to reuse from CronJob or API
- Flag parsing mixed with deployment logic
- Error handling crosses all concerns

---

## Structure

### Directory Layout

```text
project/
├── cmd/                    # CLI layer (presentation)
│   └── deploy/
│       └── deploy.go       # Cobra command setup, flag parsing, output
├── pkg/                    # Business logic layer (portable)
│   ├── deployer/
│   │   └── deployer.go     # Deployment orchestration
│   ├── validator/
│   │   └── validator.go    # Configuration validation
│   └── k8s/
│       └── client.go       # Kubernetes client wrapper
└── internal/               # Private implementation details
    └── config/
        └── loader.go       # Config file parsing
```

### Component Responsibilities

| Layer | Responsibility | Framework Dependent? | Testable Without External Systems? |
|-------|----------------|---------------------|-----------------------------------|
| `cmd/` | Flag parsing, output formatting, exit codes | Yes (Cobra) | No |
| `pkg/` | Business logic, validation, orchestration | No | Yes |
| `internal/` | Implementation details, unexported helpers | No | Yes |

---

## Implementation

### The Orchestrator Pattern

**Separate CLI handling from business logic with an orchestrator:**

```go
// cmd/deploy/deploy.go - CLI layer
package main

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "example.com/pkg/deployer"
)

func NewDeployCommand() *cobra.Command {
    var opts deployer.Options

    cmd := &cobra.Command{
        Use:   "deploy",
        Short: "Deploy application to Kubernetes",
        RunE: func(cmd *cobra.Command, args []string) error {
            // Only CLI concerns here: flag parsing, output, exit codes
            d, err := deployer.New(opts)
            if err != nil {
                return fmt.Errorf("initializing deployer: %w", err)
            }

            // Business logic delegated to pkg/
            result, err := d.Deploy(cmd.Context())
            if err != nil {
                return err
            }

            // Output formatting (CLI concern)
            fmt.Printf("Deployed %s to namespace %s\n", result.Name, result.Namespace)
            return nil
        },
    }

    // Flag binding (CLI concern)
    cmd.Flags().StringVar(&opts.Namespace, "namespace", "default", "Kubernetes namespace")
    cmd.Flags().StringVar(&opts.Image, "image", "", "Container image")
    cmd.MarkFlagRequired("image")

    return cmd
}
```

```go
// pkg/deployer/deployer.go - Business logic layer
package deployer

import (
    "context"
    "fmt"

    appsv1 "k8s.io/api/apps/v1"
    corev1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
)

// Options holds deployment configuration (no CLI framework types)
type Options struct {
    Namespace string
    Image     string
}

// Deployer orchestrates deployment operations
type Deployer struct {
    client    kubernetes.Interface  // Interface for testability
    validator Validator
    opts      Options
}

// New creates a deployer with dependency injection
func New(opts Options) (*Deployer, error) {
    client, err := getK8sClient()
    if err != nil {
        return nil, fmt.Errorf("creating client: %w", err)
    }

    return &Deployer{
        client:    client,
        validator: &DefaultValidator{},
        opts:      opts,
    }, nil
}

// Deploy executes the deployment (pure business logic)
func (d *Deployer) Deploy(ctx context.Context) (*DeploymentResult, error) {
    // Validation (business logic)
    if err := d.validator.Validate(d.opts); err != nil {
        return nil, fmt.Errorf("validation failed: %w", err)
    }

    // Deployment creation (business logic)
    deployment := d.buildDeployment()

    // Infrastructure call (delegated to client)
    created, err := d.client.AppsV1().Deployments(d.opts.Namespace).Create(
        ctx, deployment, metav1.CreateOptions{},
    )
    if err != nil {
        return nil, fmt.Errorf("creating deployment: %w", err)
    }

    return &DeploymentResult{
        Name:      created.Name,
        Namespace: created.Namespace,
    }, nil
}

// buildDeployment creates Deployment spec (business logic)
func (d *Deployer) buildDeployment() *appsv1.Deployment {
    return &appsv1.Deployment{
        ObjectMeta: metav1.ObjectMeta{
            Name: "app",
        },
        Spec: appsv1.DeploymentSpec{
            Selector: &metav1.LabelSelector{
                MatchLabels: map[string]string{"app": "app"},
            },
            Template: corev1.PodTemplateSpec{
                ObjectMeta: metav1.ObjectMeta{
                    Labels: map[string]string{"app": "app"},
                },
                Spec: corev1.PodSpec{
                    Containers: []corev1.Container{{
                        Name:  "app",
                        Image: d.opts.Image,
                    }},
                },
            },
        },
    }
}

type DeploymentResult struct {
    Name      string
    Namespace string
}
```

### Testing Benefits

```go
// pkg/deployer/deployer_test.go
package deployer

import (
    "context"
    "testing"

    "k8s.io/client-go/kubernetes/fake"
)

func TestDeploy(t *testing.T) {
    // Fake Kubernetes client - no cluster required
    fakeClient := fake.NewSimpleClientset()

    d := &Deployer{
        client:    fakeClient,
        validator: &MockValidator{},
        opts: Options{
            Namespace: "test",
            Image:     "gcr.io/project/app:v1",
        },
    }

    // Test business logic in isolation
    result, err := d.Deploy(context.Background())
    if err != nil {
        t.Fatalf("Deploy() failed: %v", err)
    }

    if result.Namespace != "test" {
        t.Errorf("got namespace %s, want test", result.Namespace)
    }

    // No Kubernetes cluster, registry, or network required
}
```

---

## Consequences

### Benefits

| Benefit | Impact |
|---------|--------|
| **Testability** | Business logic tests run in milliseconds without external dependencies |
| **Reusability** | Same logic callable from CLI, API, CronJob, or Argo Workflow |
| **Maintainability** | Changes localized to single concern (CLI changes don't affect business logic) |
| **Team velocity** | New developers understand boundaries, know where code belongs |

### Trade-offs

| Trade-off | Mitigation |
|-----------|-----------|
| More files/packages | Use clear naming conventions, documented structure |
| Interface overhead | Only create interfaces at real boundaries, not everywhere |
| Initial complexity | Complexity pays off after second feature addition |

---

## Related Patterns

- **[Usage Guide](guide.md)**: When to apply, common mistakes, anti-patterns
- **[Implementation Techniques](implementation.md)**: Interfaces, dependency injection, testing
- **[Go CLI Architecture](../../../build/go-cli-architecture/index.md)**: Complete CLI implementation example
- **[Orchestrator Pattern](../../../build/go-cli-architecture/command-architecture/orchestrator-pattern.md)**: Detailed orchestration example
- **[Fail Fast](../../error-handling/fail-fast/index.md)**: Error handling at boundaries
- **[Prerequisite Checks](../../error-handling/prerequisite-checks/index.md)**: Validation separation

---

*CLI in `cmd/`. Business logic in `pkg/`. Tests run in milliseconds. Changes stay localized. The system is maintainable.*

