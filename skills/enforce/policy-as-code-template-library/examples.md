---
name: policy-as-code-template-library - Examples
description: Code examples for Policy-as-Code Template Library
---

# Policy-as-Code Template Library - Examples


## Example 1: example-1.sh


```bash
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
```



## Example 2: example-2.sh


```bash
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
```



