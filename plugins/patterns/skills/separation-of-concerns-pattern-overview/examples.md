---
name: separation-of-concerns-pattern-overview - Examples
description: Code examples for Separation of Concerns Pattern Overview
---

# Separation of Concerns Pattern Overview - Examples


## Example 1: example-1.go


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



## Example 2: example-2.text


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



## Example 3: example-3.go


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



## Example 4: example-4.go


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



## Example 5: example-5.go


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



