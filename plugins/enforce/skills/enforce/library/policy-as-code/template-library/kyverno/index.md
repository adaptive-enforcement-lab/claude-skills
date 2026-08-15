# Kyverno Policy Templates

Source: https://adaptive-enforcement-lab.com/enforce/policy-as-code/template-library/kyverno/

> **Start with Audit Mode**
>
> Deploy in `audit` mode first. Existing workloads may violate these policies. Monitor violations for 48 hours, fix non-compliant resources, then switch to `enforce`.
>

Production-ready Kyverno policies for Kubernetes admission control. **28 policies** covering validation, mutation, and generation patterns. Each template includes complete configuration, customization options, validation commands, and real-world use cases.

---

## Available Templates

### Pod Security (5 Policies)

Enforce pod security standards, prevent privileged containers, control host namespaces, and enforce security profiles.

**Files:**

- **[Pod Security Standards →](pod-security/standards.md)** (2 policies)
- **[Privilege Escalation Prevention →](pod-security/privileges.md)** (1 policy)
- **[Security Profiles →](pod-security/profiles.md)** (2 policies)

Key policies:

- Pod Security Standards Enforcement (Baseline/Restricted)
- Host Namespace Restrictions (hostNetwork, hostPID, hostIPC, hostPort)
- Privilege Escalation Prevention (allowPrivilegeEscalation, privileged containers)
- Seccomp Profile Enforcement (RuntimeDefault, Localhost, unconfined blocking)
- AppArmor Profile Requirements (runtime/default, custom profiles)

---

### Image Validation (5 Policies)

Control container images with digest requirements, registry allowlists, signature verification, base image enforcement, and CVE scanning gates.

**Files:**

- **[Image Digest & Registry Validation →](image/validation.md)** (2 policies)
- **[Image Signing Verification →](image/signing.md)** (1 policy)
- **[Base Image Enforcement →](image/security.md)** (1 policy)
- **[CVE Scanning Gates →](image/cve-scanning.md)** (1 policy)

Key policies:

- Image Digest Requirements (SHA256 enforcement)
- Registry Allowlist and Tag Validation (block `latest`, untrusted registries)
- Cosign Image Signature Verification (keyless and key-based)
- Base Image Enforcement (approved base images, deprecated blocklist)
- CVE Scanning Gates (Trivy attestations, severity thresholds)

---

### Resource Management (5 Policies)

Ensure resource requests and limits, enforce CPU/memory ratios, control ephemeral storage, constrain PVC sizes, and require HPA configuration.

**Files:**

- **[Resource Limits & Ratios →](resource/limits.md)** (2 policies)
- **[Storage Limits →](resource/storage.md)** (2 policies)
- **[HPA Requirements →](resource/hpa.md)** (1 policy)

Key policies:

- Resource Limits and Requests Enforcement (CPU, memory, QoS classes)
- CPU and Memory Ratio Enforcement (prevent over-provisioning)
- Ephemeral Storage Limits (ephemeral storage, emptyDir controls)
- PVC Size Constraints (min/max sizes, storage class governance)
- HPA Configuration Requirements (replica bounds, metrics validation)

---

### Network Security (5 Policies)

Enforce network policies, restrict egress traffic, require ingress class validation, mandate TLS encryption, and control service types.

**Files:**

- **[Network Policies & Egress →](network/security.md)** (2 policies)
- **[Ingress Class Requirements →](network/ingress-class.md)** (1 policy)
- **[Ingress TLS Requirements →](network/ingress-tls.md)** (1 policy)
- **[Service Type Restrictions →](network/services.md)** (1 policy)

Key policies:

- Require Network Policies (namespace coverage, default-deny enforcement)
- Egress Restrictions (destination controls, external IP blocking)
- Ingress Class Requirements (approved controllers, deprecated annotation blocking)
- Ingress TLS Requirements (encryption enforcement, cert-manager integration)
- Service Type Restrictions (LoadBalancer approval, NodePort controls)

---

### Mutation & Generation (7 Policies)

Automatically inject labels, add sidecars, generate resource quotas, create network policies, and ensure pod disruption budgets.

**Files:**

- **[Label Mutation →](mutation/labels.md)** (2 policies)
- **[Sidecar Injection →](mutation/sidecar.md)** (2 policies)
- **[Namespace Resource Generation →](generation/namespace.md)** (2 policies)
- **[Workload Resource Generation →](generation/workload.md)** (1 policy)

Key policies:

- Default Label Injection (team, environment, version, cost-center)
- Namespace Label Propagation (inherit team, compliance, SLA labels)
- Logging Sidecar Injection (Fluent Bit with Elasticsearch/Loki)
- Monitoring Sidecar Injection (Nginx exporter, JMX exporter, Prometheus)
- Automatic ResourceQuota Generation (default quotas, production quotas)
- Default-Deny NetworkPolicy Generation (default-deny ingress, strict egress)
- Automatic PodDisruptionBudget Generation (2+ replicas, critical workloads)

---

### Mandatory Labels (1 Policy)

Enforce required metadata for observability, cost tracking, and compliance auditing.

**Files:**

- **[Mandatory Labels →](labels.md)** (1 policy)

Key policy:

- Mandatory Labels and Annotations (require app, team, version, environment labels)

---

## Policy Types

Kyverno supports three policy types:

### Validation Policies

**Block** resources that violate security rules.

Examples: Pod security restrictions, image allowlists, resource limits, network security

### Mutation Policies

**Modify** resources before admission to enforce standards.

Examples: Add labels, inject sidecars, set default resource limits

### Generation Policies

**Create** new resources when triggers match.

Examples: Generate ResourceQuotas for new namespaces, create default-deny NetworkPolicies

---

## Quick Start

All templates follow the same deployment pattern:

```bash
# Apply policy in audit mode first
kubectl apply -f policy.yaml

# Monitor policy violations
kubectl logs -f -n kyverno deployment/kyverno

# Check policy reports
kubectl get polr -A  # PolicyReports
kubectl get cpolr    # ClusterPolicyReports

# Switch to enforce mode after validation
kubectl patch clusterpolicy <policy-name> \
  --type merge \
  -p '{"spec":{"validationFailureAction":"enforce"}}'
```

## Policy Customization

Every template includes a customization table:

| Variable | Default | Purpose |
|----------|---------|---------|
| `validationFailureAction` | `audit` | Use `audit` for testing, `enforce` for production |
| `background` | `true` | Scan existing resources (not just new admission requests) |
| Resource selectors | Varies | Target specific namespaces, kinds, or labels |

## Related Resources

- **[JMESPath Patterns →](../jmespath/patterns.md)** - Advanced Kyverno pattern examples
- **[OPA Templates →](../opa/index.md)** - Gatekeeper constraint templates
- **[Decision Guide →](../decision-guide.md)** - OPA vs Kyverno selection guide
- **[Template Library Overview →](index.md)** - Back to main page
