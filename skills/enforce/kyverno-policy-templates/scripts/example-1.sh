# Apply policy in audit mode first
kubectl apply -f policy.yaml

# Monitor policy violations
kubectl logs -f -n kyverno deployment/kyverno

# Check policy reports
kubectl get polr -A  # PolicyReports
kubectl get cpolr    # ClusterPolicyReports

# Switch to enforce mode after validation
kubectl patch clusterpolicy <policy-name> \
  --type merge \
  -p '{"spec":{"validationFailureAction":"enforce"}}'