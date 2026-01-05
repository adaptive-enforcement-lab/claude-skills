# Try deploying without resource limits
kubectl run test --image=nginx

# Expected: Denied by admission webhook