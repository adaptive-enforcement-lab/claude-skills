---
name: kyverno-image-validation-templates - Examples
description: Code examples for Kyverno Image Validation Templates
---

# Kyverno Image Validation Templates - Examples


## Example 1: example-1.sh


```bash
kubectl apply -f registry-allowlist-policy.yaml  # Registry controls first
kubectl get clusterpolicy -w   # Watch for Ready status
```



## Example 2: example-2.sh


```bash
kubectl apply -f digest-enforcement-policy.yaml  # Digest enforcement
kubectl get clusterpolicy -w
```



## Example 3: example-3.sh


```bash
kubectl apply -f signature-verification-policy.yaml  # Signature verification
kubectl get clusterpolicy -w
```



## Example 4: example-4.sh


```bash
kubectl apply -f cve-scanning-policy.yaml  # CVE gates
kubectl get clusterpolicy -w
```



## Example 5: example-5.sh


```bash
kubectl apply -f base-image-policy.yaml  # Base image enforcement
kubectl get clusterpolicy -w
```



