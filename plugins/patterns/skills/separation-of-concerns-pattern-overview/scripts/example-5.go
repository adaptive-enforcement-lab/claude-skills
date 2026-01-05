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