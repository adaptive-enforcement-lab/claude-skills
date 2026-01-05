---
name: kyverno-image-validation-templates
description: >-
  Kyverno image validation: registry allowlists, digests, signatures, and CVE scanning gates for K8s supply chain security.
---

# Kyverno Image Validation Templates

## When to Use This Skill

Enforce container image security controls before deployment. These policies validate image sources, require cryptographic signatures, enforce digest-based references, and block images with critical vulnerabilities.

---


## Implementation

> **Phased Rollout Recommended**
>
> Start with registry controls and digest requirements before adding signature verification and CVE scanning. This minimizes disruption while building security layers progressively.
>

### 1. Start with Registry Allowlists

Block untrusted registries before enforcing signatures or scans.

```bash
kubectl apply -f registry-allowlist-policy.yaml  # Registry controls first
kubectl get clusterpolicy -w   # Watch for Ready status
```

### 2. Add Digest Requirements

Enforce immutable image references.

```bash
kubectl apply -f digest-enforcement-policy.yaml  # Digest enforcement
kubectl get clusterpolicy -w
```

### 3. Implement Image Signing

Verify images come from trusted sources.

```bash
kubectl apply -f signature-verification-policy.yaml  # Signature verification
kubectl get clusterpolicy -w
```

### 4. Enforce CVE Scanning

Block vulnerable images based on scan attestations.

```bash
kubectl apply -f cve-scanning-policy.yaml  # CVE gates
kubectl get clusterpolicy -w
```

### 5. Centralize Base Images

Standardize on approved, maintained base images.

```bash
kubectl apply -f base-image-policy.yaml  # Base image enforcement
kubectl get clusterpolicy -w
```

---


## Examples

See [examples.md](examples.md) for code examples.


## Related Patterns

- Kyverno Labels →
- Kyverno Pod Security →
- Kyverno Resource Limits →
- Template Library Overview →

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
