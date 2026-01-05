---
name: opa-image-security-templates
description: >-
  OPA image security policies for container registry allowlisting, digest enforcement, and signature verification in Kubernetes.
---

# OPA Image Security Templates

## When to Use This Skill

Image security policies control which container images can run in your cluster. These templates enforce registry allowlists, require immutable digests, and validate cryptographic signatures.

> **Image Tags Are Mutable**
>
> Tags like `latest` or `v1.2.3` can be overwritten by attackers who compromise registries. Use digest-based references (`sha256:...`) for immutable deployments.


## When to Apply

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


## Implementation

See the full implementation guide in the [source documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/).


## Examples

See [examples.md](examples.md) for code examples.


## Full Reference

See [reference.md](reference.md) for complete documentation.


## Related Patterns

- OPA Templates Overview
- OPA Pod Security
- Kyverno Image Validation

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
