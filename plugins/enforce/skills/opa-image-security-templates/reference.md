---
name: opa-image-security-templates - Reference
description: Complete reference for OPA Image Security Templates
---

# OPA Image Security Templates - Reference

This is the complete reference documentation extracted from the source.

# OPA Image Security Templates

Image security policies control which container images can run in your cluster. These templates enforce registry allowlists, require immutable digests, and validate cryptographic signatures.

> **Image Tags Are Mutable**
>
> Tags like `latest` or `v1.2.3` can be overwritten by attackers who compromise registries. Use digest-based references (`sha256:...`) for immutable deployments.
>

## Why Image Security Matters

Container images are the primary attack vector for supply chain compromises:

- **Registry Poisoning** - Attackers push malicious images to public registries
- **Tag Mutation** - `latest` tags updated with backdoored code
- **Typosquatting** - Misspelled image names redirect to attacker-controlled registries
- **Unsigned Images** - No cryptographic proof of provenance

## Available Templates

### [Registry Allowlist](base.md)

Restrict container images to approved registries:

- Block public registries (Docker Hub, Quay, GCR)
- Allow only corporate registry domains
- Enforce namespace-specific registry restrictions
- Prevent deployment of untrusted images

**Apply a policy:**

```bash
kubectl apply -f base.yaml

```

### [Digest Enforcement](digest.md)

Require immutable digest references instead of mutable tags:

- Block tag-based image references (`nginx:latest`)
- Mandate digest-based references (`nginx@sha256:...`)
- Validate digest format and checksum integrity
- Prevent image tag mutation attacks

**Apply a policy:**

```bash
kubectl apply -f digest.yaml

```

### [Image Scanning Requirements](security.md)

Enforce vulnerability scanning and security assessments:

- Require scan annotations on all images
- Block images with high/critical CVEs
- Validate scan freshness (no stale scans)
- Enforce minimum scan score thresholds

**Apply a policy:**

```bash
kubectl apply -f security.yaml

```

### [Signature Verification](verification.md)

Validate cryptographic signatures using Cosign or Notary:

- Require valid signatures from trusted keys
- Block unsigned or invalidly signed images
- Enforce signature transparency logs (Rekor)
- Validate attestations for build provenance

**Apply a policy:**

```bash
kubectl apply -f verification.yaml

```

## Image Security Defense Layers

Implement multiple controls for defense in depth:

1. **Registry Allowlist** - Only approved registries (base.yaml)
2. **Digest Enforcement** - Immutable image references (digest.yaml)
3. **Vulnerability Scanning** - No high/critical CVEs (security.yaml)
4. **Signature Verification** - Cryptographic provenance (verification.yaml)

Each layer addresses different attack vectors. Use all four for production environments.

## Common Enforcement Scenarios

### Scenario 1: Block Public Registries

Prevent deployment of images from untrusted sources:

```yaml
# Enforced by: base.yaml
# Result: Only images from registry.company.com allowed
# Impact: Eliminates supply chain attacks via public registries

```

### Scenario 2: Prevent Tag Mutation

Require digest-based image references:

```yaml
# Enforced by: digest.yaml
# Result: Image references must use @sha256:... format
# Impact: Guarantees deployed image matches approved version

```

### Scenario 3: Block Vulnerable Images

Reject images with known CVEs:

```yaml
# Enforced by: security.yaml
# Result: Images must have scan results with no high/critical vulnerabilities
# Impact: Prevents deployment of exploitable container images

```

### Scenario 4: Verify Build Provenance

Validate cryptographic signatures on all images:

```yaml
# Enforced by: verification.yaml
# Result: Images must be signed by trusted key in KMS
# Impact: Ensures images originated from approved CI/CD pipelines

```

## Testing Image Security Policies

Validate enforcement without blocking legitimate workloads:

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

## Signature Verification with Cosign

Deploy Cosign-based signature verification:

1. **Generate signing keys:**

   ```bash
   cosign generate-key-pair
   # Creates cosign.key (private) and cosign.pub (public)

   ```

2. **Sign container images in CI/CD:**

   ```bash
   cosign sign --key cosign.key registry.company.com/app:v1.2.3

   ```

3. **Store public key in Kubernetes Secret:**

   ```bash
   kubectl create secret generic cosign-pub \
     --from-file=cosign.pub=./cosign.pub \
     -n opa-system

   ```

4. **Configure OPA policy to verify signatures:**
   Policy references `cosign-pub` Secret for signature validation.

5. **Validate signature verification:**

   ```bash
   kubectl run test-app --image=registry.company.com/app@sha256:...
   # OPA validates signature before admission

   ```

## Image Scanning Integration

Integrate vulnerability scanning into admission control:

### Scan in CI/CD

Scan images during build and push scan results as annotations:

```bash
# Trivy example
trivy image --format json registry.company.com/app:v1.2.3 > scan.json

# Push image with scan annotation
crane mutate registry.company.com/app:v1.2.3 \
  --annotation trivy.scan.result="$(cat scan.json)"

```

### Enforce Scan Results

OPA policies read scan annotations and block vulnerable images:

```rego
# Pseudo-code: Full implementation in security.yaml
deny[msg] {
  scan_result := input.metadata.annotations["trivy.scan.result"]
  criticals := count(scan_result.Results[_].Vulnerabilities[_] | _.Severity == "CRITICAL")
  criticals > 0
  msg := sprintf("Image has %d critical vulnerabilities", [criticals])
}

```

## Migrating from Tag-Based to Digest-Based Deployments

Transition existing workloads to use digests:

1. **Audit current image references:**

   ```bash
   kubectl get pods --all-namespaces -o jsonpath='{range .items[*]}{.spec.containers[*].image}{"\n"}{end}' | sort -u

   ```

2. **Convert tags to digests:**

   ```bash
   # Get digest for tagged image
   crane digest registry.company.com/nginx:v1.21
   # Output: sha256:abcdef123456...

   # Update deployment to use digest
   kubectl set image deployment/nginx nginx=registry.company.com/nginx@sha256:abcdef123456...

   ```

3. **Deploy digest enforcement policy in audit mode:**
   Set `enforcementAction: warn` to identify non-compliant workloads.

4. **Fix violations and enable enforcement:**
   After all workloads use digests, set `enforcementAction: deny`.

## Related Resources

- [OPA Templates Overview](../index.md)
- [OPA Pod Security](../pod-security/index.md)
- [Kyverno Image Validation](../image/index.md)

