---
name: opa-policy-templates - Reference
description: Complete reference for OPA Policy Templates
---

# OPA Policy Templates - Reference

This is the complete reference documentation extracted from the source.

# OPA Policy Templates

> **Deploy in Audit Mode First**
>
> Use `enforcementAction: dryrun` initially. Existing resources may violate constraints. Monitor violations for 48 hours using `kubectl get constraints`, fix non-compliant resources, then switch to `deny`.
>

Production-ready OPA/Gatekeeper constraint templates for Kubernetes admission control. **20 policies** covering pod security, image
validation, RBAC, and resource governance. Each template includes complete Rego implementation, constraint examples, customization
options, validation commands, and real-world use cases.

---

## Available Templates

### Pod Security (5 Policies)

Prevent privileged containers, block host namespace access, enforce capability drops, require secure contexts, and prevent privilege escalation.

**Files:**

- **[Privileged & Host Namespaces →](pod-security/overview.md)** (2 policies)
- **[Capabilities Drop →](pod-security/capabilities.md)** (1 policy)
- **[Security Contexts →](pod-security/contexts.md)** (1 policy)
- **[Privilege Escalation Prevention →](pod-security/escalation.md)** (1 policy)

Key policies:

- Privileged Container Prevention (block `privileged: true`)
- Host Namespace Restrictions (block hostNetwork, hostPID, hostIPC, hostPort)
- Required Capabilities Drop (enforce `drop: ["ALL"]`, restrict dangerous capabilities)
- Security Context Requirements (runAsNonRoot, readOnlyRootFilesystem, UID restrictions)
- Privilege Escalation Prevention (block `allowPrivilegeEscalation: true`)

---

### Image Security (5 Policies)

Control container images with registry allowlists, tag requirements, digest enforcement, signature verification annotations, and base image governance.

**Files:**

- **[Registry & Tag Validation →](image/security.md)** (2 policies)
- **[Digest Enforcement →](image/digest.md)** (1 policy)
- **[Signature Verification →](image/verification.md)** (1 policy)
- **[Base Image Enforcement →](image/base.md)** (1 policy)

Key policies:

- Registry Allowlist (enforce approved registries, block public Docker Hub)
- Tag Requirements (block `latest` tags, require specific tag patterns)
- Digest Enforcement (require SHA256 digest references, block tag-only images)
- Image Signature Verification Annotations (require proof of cosign verification in CI/CD)
- Base Image Enforcement (require approved base images via annotations, block deprecated)

---

### RBAC (5 Policies)

Restrict service accounts, prevent cross-namespace role bindings, block cluster-admin assignments, restrict privileged verbs, and prevent wildcard permissions.

**Files:**

- **[Service Accounts & Role Bindings →](rbac/overview.md)** (2 policies)
- **[Cluster-Admin Prevention →](rbac/cluster-admin.md)** (1 policy)
- **[Privileged Verbs Restrictions →](rbac/privileged-verbs.md)** (1 policy)
- **[Wildcard Prevention →](rbac/wildcards.md)** (1 policy)

Key policies:

- Service Account Restrictions (block default SA usage, require dedicated SAs, prevent auto-mount tokens)
- Role Binding Namespace Enforcement (prevent cross-namespace subjects in RoleBindings)
- Cluster-Admin Prevention (block cluster-admin and system:masters role assignments)
- Privileged Verbs Restrictions (block escalate, impersonate, bind verbs)
- Wildcard Resource Prevention (block `*` in resources, apiGroups, verbs)

---

### Resource Governance (5 Policies)

Enforce resource limits and requests, require namespace quotas, mandate LimitRanges, control ephemeral storage, and restrict storage classes and PVC sizes.

**Files:**

- **[Resource Limits & Quotas →](resource/governance.md)** (2 policies)
- **[LimitRange & Ephemeral Storage →](resource/limitrange.md)** (2 policies)
- **[Storage Class & PVC Constraints →](resource/storage.md)** (2 policies)

Key policies:

- Resource Limits and Requests Enforcement (CPU, memory limits required, max limits enforcement)
- Resource Quota Requirements (namespace quotas required, prevent unbounded consumption)
- LimitRange Requirements (default limits in namespaces, prevent extreme requests)
- Ephemeral Storage Limits (ephemeral-storage limits required, prevent disk exhaustion)
- Storage Class Restrictions (allowlist/blocklist, cost control, migration enforcement)
- PVC Size Constraints (min/max sizes, approval workflow for large volumes)

---

## OPA vs Kyverno

Choosing between OPA/Gatekeeper and Kyverno depends on your team's expertise and requirements:

### Use OPA/Gatekeeper When

- You need **maximum flexibility** in policy logic (Rego is Turing-complete)
- Your team has **Rego expertise** or investment in OPA across multiple systems
- You require **cross-platform policy** (Kubernetes, Terraform, Envoy, etc.)
- Policies involve **complex conditional logic** or multi-resource validation
- You're building a **policy platform** for enterprise governance

### Use Kyverno When

- You want **Kubernetes-native YAML** policies (no DSL learning curve)
- You need **mutation and generation** features (OPA is validation-only)
- Your team prefers **JMESPath** over Rego for data extraction
- You want **faster time-to-value** with simpler policies
- You're **new to policy-as-code** and want quick adoption

**See [Decision Guide →](../decision-guide.md)** for detailed comparison and migration strategies.

---

## Policy Deployment

All OPA constraint templates follow the same two-step deployment pattern:

### Step 1: Deploy ConstraintTemplate

The `ConstraintTemplate` defines the policy logic in Rego:

```yaml
apiVersion: templates.gatekeeper.sh/v1beta1
kind: ConstraintTemplate
metadata:
  name: k8sblockprivileged
spec:
  crd:
    spec:
      names:
        kind: K8sBlockPrivileged
  targets:
    - target: admission.k8s.gatekeeper.sh
      rego: |
        package k8sblockprivileged
        violation[{"msg": msg}] {
          # Rego policy logic here
        }
```

### Step 2: Deploy Constraint

The `Constraint` activates the template with specific parameters:

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sBlockPrivileged
metadata:
  name: block-privileged-containers
spec:
  enforcementAction: dryrun  # Use 'deny' for enforcement
  match:
    kinds:
      - apiGroups: [""]
        kinds: ["Pod"]
    namespaces:
      - "production"
      - "staging"
```

---

## Quick Start

Standard deployment workflow for all templates:

```bash
# 1. Deploy constraint template (defines policy logic)
kubectl apply -f constraint-template.yaml

# 2. Deploy constraint in audit mode (dryrun)
kubectl apply -f constraint.yaml

# 3. Monitor violations
kubectl get constraints
kubectl get <constraint-kind> <constraint-name> -o yaml

# 4. Check audit results
kubectl get constraints -o json | jq '.items[].status.violations'

# 5. Fix non-compliant resources
kubectl get pods -n production --show-labels

# 6. Switch to enforcement mode after validation
kubectl patch <constraint-kind> <constraint-name> \
  --type merge \
  -p '{"spec":{"enforcementAction":"deny"}}'
```

---

## Policy Customization

Every template includes a customization table with these common parameters:

| Parameter | Default | Purpose |
|-----------|---------|---------|
| `enforcementAction` | `dryrun` | Use `dryrun` for testing, `deny` for enforcement |
| `match.kinds` | Varies | Target specific Kubernetes resource types |
| `match.namespaces` | `[]` | Target specific namespaces (empty = all) |
| `match.excludedNamespaces` | `["kube-system"]` | Exempt system namespaces |
| `match.labelSelector` | None | Target resources with specific labels |

Template-specific parameters (e.g., `exemptImages`, `allowedRegistries`, `maxCPU`) are documented in each policy's customization table.

---

## Constraint Status

Monitor policy violations and audit results:

```bash
# List all constraints
kubectl get constraints

# Get detailed status for a specific constraint
kubectl get k8sblockprivileged block-privileged-containers -o yaml

# Extract violations from constraint status
kubectl get k8sblockprivileged block-privileged-containers \
  -o jsonpath='{.status.violations[*].message}' | jq

# Count total violations across all constraints
kubectl get constraints -o json | \
  jq '[.items[].status.totalViolations] | add'
```

---

## Rego Testing

All templates include unit test examples for Rego policies:

```bash
# Install OPA CLI
brew install opa  # macOS
# or download from https://www.openpolicyagent.org/docs/latest/#running-opa

# Test Rego policy locally
opa test constraint-template.yaml test-cases.yaml -v

# Example test case
# test-cases.yaml
package k8sblockprivileged

test_privileged_container_blocked {
  violation[{"msg": msg}] with input as {
    "review": {
      "object": {
        "spec": {
          "containers": [{
            "name": "test",
            "securityContext": {"privileged": true}
          }]
        }
      }
    }
  }
}
```

See **[Privilege Escalation Prevention →](pod-security/escalation.md#rego_unit_testing)** for complete testing guide.

---

## Related Resources

- **[Kyverno Templates →](../kyverno/index.md)** - 28 Kyverno policies for comparison
- **[Decision Guide →](../decision-guide.md)** - OPA vs Kyverno selection criteria
- **[OPA/Kyverno Comparison →](../opa-kyverno-comparison.md)** - Detailed feature comparison
- **[Migration Guide →](../opa-kyverno-migration.md)** - Switching between OPA and Kyverno
- **[Template Library Overview →](index.md)** - Back to main page

