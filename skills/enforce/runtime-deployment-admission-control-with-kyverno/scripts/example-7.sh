# Check Kyverno pods
kubectl get pods -n kyverno

# Expected output:
# kyverno-admission-controller-xxx   Running
# kyverno-background-controller-xxx  Running
# kyverno-cleanup-controller-xxx     Running
# kyverno-reports-controller-xxx     Running