# Test pod without resource limits is rejected
kubectl apply -f pod-no-limits.yaml
# Expected: Admission webhook denies request

# Test untrusted registry is blocked
kubectl apply -f pod-dockerhub.yaml
# Expected: Image source validation fails