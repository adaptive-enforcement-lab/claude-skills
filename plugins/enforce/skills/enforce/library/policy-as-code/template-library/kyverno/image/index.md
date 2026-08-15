# Kyverno Image Validation Templates

Source: https://adaptive-enforcement-lab.com/enforce/policy-as-code/template-library/kyverno/image/

Enforce container image security controls before deployment. These policies validate image sources, require cryptographic signatures, enforce digest-based references, and block images with critical vulnerabilities.

---

## Purpose

Image validation is your first line of defense against supply chain attacks. These templates ensure:

- Only images from approved registries can deploy
- Images use immutable digest references, not mutable tags
- Images are cryptographically signed by trusted build pipelines
- Images without recent vulnerability scans are blocked

---

## Template Categories

### [Image Validation](validation.md)

Registry allowlists, digest requirements, and tag validation. Block untrusted registries and prohibit mutable `latest` tags. Enforce SHA256 digest references for immutable deployments and supply chain transparency.

Use this for:

- Preventing public Docker Hub images in production
- Requiring digest-based image references
- Blocking images without explicit version tags
- Enforcing approved internal registries

---

### [Image Signing](signing.md)

Cosign signature verification for trusted images. Verify container images are cryptographically signed before deployment. Support both key-based and keyless (OIDC) signing with SLSA provenance attestations.

Use this for:

- Verifying images come from approved build pipelines
- SLSA compliance and provenance attestations
- Preventing unauthorized image pushes
- Enforcing GitHub Actions workflow signatures

---

### [Base Image Security](security.md)

Approved base image enforcement. Restrict workloads to approved, maintained base images. Block deprecated distributions and enforce minimal base images for high-security workloads.

Use this for:

- Centralizing base image management
- Blocking vulnerable or EOL distributions
- Enforcing distroless or minimal images
- Standardizing across teams

---

### [CVE Scanning Gates](cve-scanning.md)

Vulnerability scan attestations and CVE thresholds. Require Trivy vulnerability scan attestations before deployment. Block images with critical or high severity CVEs based on environment risk tolerance.

Use this for:

- Zero-day vulnerability protection
- PCI-DSS and SOC2 compliance
- Shift-left security in CI/CD
- Different CVE thresholds per environment

---

## Implementation Strategy

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

## Related Resources

- **[Kyverno Labels →](../labels.md)** - Mandatory metadata enforcement
- **[Kyverno Pod Security →](../pod-security/standards.md)** - Security contexts and capabilities
- **[Kyverno Resource Limits →](../resource/limits.md)** - CPU and memory enforcement
- **[Template Library Overview →](../index.md)** - Back to main page

---

## External Documentation

- **[Kyverno Image Verification](https://kyverno.io/docs/writing-policies/verify-images/)** - Official Kyverno image verification guide
- **[Sigstore Cosign](https://docs.sigstore.dev/cosign/overview/)** - Container image signing and verification
- **[SLSA Framework](https://slsa.dev/)** - Supply chain security levels
- **[Trivy Scanner](https://aquasecurity.github.io/trivy/)** - Vulnerability scanning tool
