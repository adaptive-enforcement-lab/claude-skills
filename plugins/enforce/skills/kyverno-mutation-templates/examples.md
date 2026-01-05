---
name: kyverno-mutation-templates - Examples
description: Code examples for Kyverno Mutation Templates
---

# Kyverno Mutation Templates - Examples


## Example 1: example-1.sh


```bash
kubectl apply -f labels.yaml
```



## Example 2: example-2.sh


```bash
kubectl apply -f sidecar.yaml
```



## Example 3: example-3.sh


```bash
# Test label mutation
kubectl apply --dry-run=server -f test-deployment.yaml -o yaml | grep -A5 labels

# Test sidecar injection
kubectl apply --dry-run=server -f test-pod.yaml -o yaml | grep -A10 containers
```



