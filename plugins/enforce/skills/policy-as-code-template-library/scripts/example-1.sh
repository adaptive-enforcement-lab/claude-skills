# 1. Install Kyverno
helm repo add kyverno https://kyverno.github.io/kyverno/
helm install kyverno kyverno/kyverno --namespace kyverno --create-namespace

# 2. Apply a policy (starts in audit mode)
kubectl apply -f https://raw.githubusercontent.com/adaptive-enforcement-lab/docs/main/kyverno-pod-security.yaml

# 3. Monitor violations
kubectl get polr -A  # PolicyReports
kubectl get cpolr    # ClusterPolicyReports

# 4. Switch to enforcement after validation
kubectl patch clusterpolicy require-pod-security \
  --type merge \
  -p '{"spec":{"validationFailureAction":"enforce"}}'