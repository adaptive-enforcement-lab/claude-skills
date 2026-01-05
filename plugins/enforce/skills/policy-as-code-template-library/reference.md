---
name: policy-as-code-template-library - Reference
description: Complete reference for Policy-as-Code Template Library
---

# Policy-as-Code Template Library - Reference

This is the complete reference documentation extracted from the source.

# Policy-as-Code Template Library

**48 production-ready policies** for Kubernetes security and governance. Reduce the Rego learning curve. Copy, customize, deploy.

<!-- more -->

> **Template Library Overview**
>
> This library contains **28 Kyverno policies** and **20 OPA/Gatekeeper constraint templates** covering pod security, image validation, RBAC, resource governance, network security, mutation, and generation.
> Each template includes complete YAML/Rego, customization variables, validation commands, and real-world use cases.
>

---

## What You Get

This library provides ready-to-use policies for common security scenarios:

- **48 Total Policies**: 28 Kyverno + 20 OPA/Gatekeeper
- **Complete Implementation**: Full YAML/Rego with production-ready configuration
- **Customization Tables**: Variables, defaults, and purpose for each parameter
- **Validation Commands**: Test policies before enforcement
- **Real-World Use Cases**: 4-6 production scenarios per policy
- **Testing Guidance**: Audit mode, policy reports, troubleshooting

---

## Template Categories

### [Decision Guide →](decision-guide.md)

**Choose between OPA and Kyverno** based on team expertise, policy complexity, and operational requirements.

- Quick decision matrix (expertise, scope, complexity)
- Recommended starter paths
- [Detailed comparison →](opa-kyverno-comparison.md)
- [Migration strategies →](opa-kyverno-migration.md)

---

### [Kyverno Templates →](kyverno/index.md)

**28 production-ready Kyverno policies** for Kubernetes admission control, mutation, and resource generation.

#### [Pod Security →](kyverno/pod-security/index.md) (5 Policies)

- Pod Security Standards Enforcement
- Host Namespace Restrictions
- Privilege Escalation Prevention
- Seccomp Profile Enforcement
- AppArmor Profile Requirements

#### [Image Validation →](kyverno/image/index.md) (5 Policies)

- Image Digest Requirements
- Registry Allowlist and Tag Validation
- Cosign Image Signature Verification
- Base Image Enforcement
- CVE Scanning Gates

#### [Resource Management →](kyverno/resource/index.md) (5 Policies)

- Resource Limits and Requests Enforcement
- CPU and Memory Ratio Enforcement
- Ephemeral Storage Limits
- PVC Size Constraints
- HPA Configuration Requirements

#### [Network Security →](kyverno/network/index.md) (5 Policies)

- Require Network Policies
- Egress Restrictions
- Ingress Class Requirements
- Ingress TLS Requirements
- Service Type Restrictions

#### Mutation & Generation (7 Policies)

- [Mutation Policies →](kyverno/mutation/index.md) - Default Label Injection, Namespace Label Propagation, Logging Sidecar Injection, Monitoring Sidecar Injection
- [Generation Policies →](kyverno/generation/index.md) - Automatic ResourceQuota Generation, Default-Deny NetworkPolicy Generation, Automatic PodDisruptionBudget Generation

#### [Labels & Metadata →](kyverno/labels.md) (1 Policy)

- Mandatory Labels and Annotations

---

### [OPA/Gatekeeper Templates →](opa/index.md)

**20 production-ready OPA constraint templates** with complete Rego implementation for advanced policy enforcement.

#### [Pod Security →](opa/pod-security/index.md) (5 Policies)

- Privileged Container Prevention
- Host Namespace Restrictions
- Required Capabilities Drop
- Security Context Requirements
- Privilege Escalation Prevention

#### [Image Security →](opa/image/index.md) (5 Policies)

- Registry Allowlist
- Tag Requirements
- Digest Enforcement
- Image Signature Verification Annotations
- Base Image Enforcement

#### [RBAC →](opa/rbac/index.md) (5 Policies)

- Service Account Restrictions
- Role Binding Namespace Enforcement
- Cluster-Admin Prevention
- Privileged Verbs Restrictions
- Wildcard Resource Prevention

#### [Resource Governance →](opa/resource/index.md) (5 Policies)

- Resource Limits and Requests Enforcement
- Resource Quota Requirements
- LimitRange Requirements
- Ephemeral Storage Limits
- Storage Class Restrictions

---

### [JMESPath Patterns →](jmespath/index.md)

**Advanced Kyverno pattern library** for complex validation logic using JMESPath.

- Pattern fundamentals (projection, filtering, multi-select)
- Cross-field validation (requests vs limits, label dependencies)
- Complex conditions (nested logic, transformations)
- [Advanced patterns →](jmespath/advanced.md) (aggregation, arithmetic, string manipulation)
- [Enterprise examples →](jmespath/enterprise.md) (registry policies, cost controls, HA requirements)
- [Testing guide →](jmespath/testing.md) (kyverno jp CLI, debugging, validation)

---

### [CI/CD Integration →](ci-cd-integration.md)

Automated policy validation in development pipelines:

- GitHub Actions pre-flight validation
- ArgoCD policy gating
- Pre-commit hooks

---

### [Usage Guide →](usage-guide.md)

Template customization workflow, validation best practices, and quick start guides:

- Customization workflow
- Validation best practices
- Quick start guides
- Troubleshooting

---

## Policy Engine Comparison

Choose the right policy engine for your team:

| Feature | Kyverno | OPA/Gatekeeper |
|---------|---------|----------------|
| **Policies** | 28 (validation, mutation, generation) | 20 (validation only) |
| **Language** | YAML + JMESPath | Rego (Go-like DSL) |
| **Learning Curve** | < 1 hour | 4-8 hours |
| **Best For** | Kubernetes-native teams, fast adoption | Multi-platform policies, complex logic |
| **Mutation** | ✅ Native support | ❌ Validation only |
| **Generation** | ✅ Auto-create resources | ❌ Validation only |

**See [Decision Guide →](decision-guide.md)** for detailed comparison and recommended starter paths.

---

## Quick Start

> **Deploy in Audit Mode First**
>
> Always start with `audit` (Kyverno) or `dryrun` (OPA) mode. Monitor violations for 48 hours before switching to enforcement. Existing workloads may violate policies.
>

### Kyverno Quick Start (5 minutes)

```bash
# 1. Install Kyverno
helm repo add kyverno https://kyverno.github.io/kyverno/
helm install kyverno kyverno/kyverno --namespace kyverno --create-namespace

# 2. Apply a policy (starts in audit mode)
kubectl apply -f https://raw.githubusercontent.com/adaptive-enforcement-lab/docs/main/kyverno-pod-security.yaml

# 3. Monitor violations
kubectl get polr -A  # PolicyReports
kubectl get cpolr    # ClusterPolicyReports

# 4. Switch to enforcement after validation
kubectl patch clusterpolicy require-pod-security \
  --type merge \
  -p '{"spec":{"validationFailureAction":"enforce"}}'
```

### OPA/Gatekeeper Quick Start (10 minutes)

```bash
# 1. Install Gatekeeper
kubectl apply -f https://raw.githubusercontent.com/open-policy-agent/gatekeeper/master/deploy/gatekeeper.yaml

# 2. Deploy constraint template (policy logic)
kubectl apply -f https://raw.githubusercontent.com/adaptive-enforcement-lab/docs/main/opa-pod-security.yaml

# 3. Deploy constraint (starts in dryrun mode)
kubectl apply -f constraint.yaml

# 4. Monitor violations
kubectl get constraints
kubectl get k8sblockprivileged -o yaml

# 5. Switch to enforcement after validation
kubectl patch k8sblockprivileged block-privileged \
  --type merge \
  -p '{"spec":{"enforcementAction":"deny"}}'
```

---

## Related Resources

- **[Kyverno Official Documentation](https://kyverno.io/docs/)** - Kyverno guides and API reference
- **[OPA/Gatekeeper Documentation](https://open-policy-agent.org/docs/latest/kubernetes-admission-control/)** - Gatekeeper deployment and Rego reference
- **[Kubernetes Pod Security Standards](https://kubernetes.io/docs/concepts/security/pod-security-standards/)** - Baseline and Restricted profiles
- **[NIST SP 800-190](https://csrc.nist.gov/publications/detail/sp/800-190/final)** - Application Container Security Guide
- **[CIS Kubernetes Benchmark](https://www.cisecurity.org/benchmark/kubernetes)** - Security configuration standards

