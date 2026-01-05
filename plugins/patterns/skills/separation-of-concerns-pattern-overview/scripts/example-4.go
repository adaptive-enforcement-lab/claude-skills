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