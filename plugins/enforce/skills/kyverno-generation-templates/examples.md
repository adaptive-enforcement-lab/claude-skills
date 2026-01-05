---
name: kyverno-generation-templates - Examples
description: Code examples for Kyverno Generation Templates
---

# Kyverno Generation Templates - Examples


## Example 1: example-1.yaml


```yaml
exclude:
  resources:
    names:
      - kube-system
      - kube-public
      - kube-node-lease

preconditions:
  all:
    - key: "{{ request.object.spec.replicas }}"
      operator: GreaterThanOrEquals
      value: 2
```



## Example 2: example-2.sh


```bash
# Check that new namespaces get ResourceQuotas
kubectl create namespace test-gen
kubectl get resourcequotas -n test-gen

# Check that multi-replica Deployments get PDBs
kubectl create deployment nginx --image=nginx --replicas=3 -n test-gen
kubectl label deployment nginx app=nginx -n test-gen
kubectl get pdb -n test-gen

# Audit resources without expected generated objects
kubectl get namespaces -o json | jq -r '.items[] | select(.metadata.name != "kube-system") | .metadata.name' | while read ns; do
  quota_count=$(kubectl get resourcequotas -n $ns --no-headers 2>/dev/null | wc -l)
  if [ $quota_count -eq 0 ]; then
    echo "WARNING: Namespace $ns has no ResourceQuota"
  fi
done
```



