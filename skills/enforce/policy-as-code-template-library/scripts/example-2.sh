# 1. Install Gatekeeper
kubectl apply -f https://raw.githubusercontent.com/open-policy-agent/gatekeeper/master/deploy/gatekeeper.yaml

# 2. Deploy constraint template (policy logic)
kubectl apply -f https://raw.githubusercontent.com/adaptive-enforcement-lab/docs/main/opa-pod-security.yaml

# 3. Deploy constraint (starts in dryrun mode)
kubectl apply -f constraint.yaml

# 4. Monitor violations
kubectl get constraints
kubectl get k8sblockprivileged -o yaml

# 5. Switch to enforcement after validation
kubectl patch k8sblockprivileged block-privileged \
  --type merge \
  -p '{"spec":{"enforcementAction":"deny"}}'