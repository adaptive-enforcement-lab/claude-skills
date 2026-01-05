# 1. Deploy constraint template (defines policy logic)
kubectl apply -f constraint-template.yaml

# 2. Deploy constraint in audit mode (dryrun)
kubectl apply -f constraint.yaml

# 3. Monitor violations
kubectl get constraints
kubectl get <constraint-kind> <constraint-name> -o yaml

# 4. Check audit results
kubectl get constraints -o json | jq '.items[].status.violations'

# 5. Fix non-compliant resources
kubectl get pods -n production --show-labels

# 6. Switch to enforcement mode after validation
kubectl patch <constraint-kind> <constraint-name> \
  --type merge \
  -p '{"spec":{"enforcementAction":"deny"}}'