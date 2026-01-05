---
name: opa-image-security-templates - Examples
description: Code examples for OPA Image Security Templates
---

# OPA Image Security Templates - Examples


## Example 1: example-1.sh


```bash
kubectl apply -f base.yaml
```



## Example 2: example-2.sh


```bash
kubectl apply -f digest.yaml
```



## Example 3: example-3.sh


```bash
kubectl apply -f security.yaml
```



## Example 4: example-4.sh


```bash
kubectl apply -f verification.yaml
```



## Example 5: example-5.yaml


```yaml
# Enforced by: base.yaml
# Result: Only images from registry.company.com allowed
# Impact: Eliminates supply chain attacks via public registries
```



## Example 6: example-6.yaml


```yaml
# Enforced by: digest.yaml
# Result: Image references must use @sha256:... format
# Impact: Guarantees deployed image matches approved version
```



## Example 7: example-7.yaml


```yaml
# Enforced by: security.yaml
# Result: Images must have scan results with no high/critical vulnerabilities
# Impact: Prevents deployment of exploitable container images
```



## Example 8: example-8.yaml


```yaml
# Enforced by: verification.yaml
# Result: Images must be signed by trusted key in KMS
# Impact: Ensures images originated from approved CI/CD pipelines
```



## Example 9: example-9.sh


```bash
# Test registry allowlist (should fail for Docker Hub)
kubectl run docker-hub-test --image=nginx:latest
# Expected: Blocked by registry allowlist policy

# Test digest requirement (should fail for tag-based reference)
kubectl run tag-test --image=registry.company.com/nginx:v1.21
# Expected: Blocked by digest enforcement policy

# Test digest-based reference (should succeed if from approved registry)
kubectl run digest-test --image=registry.company.com/nginx@sha256:abcdef123456...
# Expected: Allowed by registry and digest policies

# Test unsigned image (should fail without valid signature)
kubectl run unsigned-test --image=registry.company.com/app@sha256:123456...
# Expected: Blocked by signature verification policy

# Test signed image (should succeed with valid Cosign signature)
# First, sign the image:
# cosign sign --key cosign.key registry.company.com/app@sha256:123456...
kubectl run signed-test --image=registry.company.com/app@sha256:123456...
# Expected: Allowed after signature verification
```



## Example 10: example-10.sh


```bash
cosign generate-key-pair
# Creates cosign.key (private) and cosign.pub (public)
```



## Example 11: example-11.sh


```bash
cosign sign --key cosign.key registry.company.com/app:v1.2.3
```



## Example 12: example-12.sh


```bash
kubectl create secret generic cosign-pub \
  --from-file=cosign.pub=./cosign.pub \
  -n opa-system
```



## Example 13: example-13.sh


```bash
kubectl run test-app --image=registry.company.com/app@sha256:...
# OPA validates signature before admission
```



## Example 14: example-14.sh


```bash
# Trivy example
trivy image --format json registry.company.com/app:v1.2.3 > scan.json

# Push image with scan annotation
crane mutate registry.company.com/app:v1.2.3 \
  --annotation trivy.scan.result="$(cat scan.json)"
```



## Example 15: example-15.rego


```rego
# Pseudo-code: Full implementation in security.yaml
deny[msg] {
  scan_result := input.metadata.annotations["trivy.scan.result"]
  criticals := count(scan_result.Results[_].Vulnerabilities[_] | _.Severity == "CRITICAL")
  criticals > 0
  msg := sprintf("Image has %d critical vulnerabilities", [criticals])
}
```



## Example 16: example-16.sh


```bash
kubectl get pods --all-namespaces -o jsonpath='{range .items[*]}{.spec.containers[*].image}{"\n"}{end}' | sort -u
```



## Example 17: example-17.sh


```bash
# Get digest for tagged image
crane digest registry.company.com/nginx:v1.21
# Output: sha256:abcdef123456...

# Update deployment to use digest
kubectl set image deployment/nginx nginx=registry.company.com/nginx@sha256:abcdef123456...
```



