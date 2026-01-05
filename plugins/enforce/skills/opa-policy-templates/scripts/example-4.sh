# List all constraints
kubectl get constraints

# Get detailed status for a specific constraint
kubectl get k8sblockprivileged block-privileged-containers -o yaml

# Extract violations from constraint status
kubectl get k8sblockprivileged block-privileged-containers \
  -o jsonpath='{.status.violations[*].message}' | jq

# Count total violations across all constraints
kubectl get constraints -o json | \
  jq '[.items[].status.totalViolations] | add'