# Test label mutation
kubectl apply --dry-run=server -f test-deployment.yaml -o yaml | grep -A5 labels

# Test sidecar injection
kubectl apply --dry-run=server -f test-pod.yaml -o yaml | grep -A10 containers