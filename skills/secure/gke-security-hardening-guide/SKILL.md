---
name: gke-security-hardening-guide
description: >-
  GKE security hardening guide with Pulumi. Private clusters, Workload Identity, Binary Authorization, network policies, IAM configuration, and runtime security enforcement.
---

# GKE Security Hardening Guide

## When to Use This Skill

> **Defense in Depth**
>
>
> GKE security hardening follows a layered approach:
>
> 1. **Control plane** - Private clusters, authenticated access, audit logging
> 2. **Network** - VPC-native networking, network policies, egress controls
> 3. **Identity** - Workload Identity Federation, least-privilege IAM, audit trails
> 4. **Runtime** - Pod Security Standards, admission controllers, monitoring
>

This guide uses Pulumi for Infrastructure as Code, enabling repeatable, auditable cluster deployments across environments (QAC, DEV, STG, PRD).

> **Environment Promotion Order**
>
>
> Always promote changes through: **QAC → DEV → STG → PRD → OPS**
>
> Never skip environments in the promotion pipeline.


## Prerequisites

- GCP project with billing enabled
- `gcloud` CLI installed and authenticated
- Pulumi 3.0+
- kubectl configured for cluster access
- Appropriate IAM permissions (Project Editor or Security Admin roles)

> **Production Warning**
>
>
> These configurations enforce strict security controls. Test in QAC/DEV before production deployment.


## Implementation

*See [examples.md](examples.md) for detailed code examples.*

> **Verification**
>
>
> After deployment, verify the security posture using the verification checklists in each configuration module.


## Anti-Patterns to Avoid

| Misconfiguration | Risk | Fix |
|------------------|------|-----|
| Public cluster endpoint | Exposed API server | Set `privateClusterConfig.enablePrivateNodes = true` |
| Metadata server enabled | Pod can access node credentials | Set `workloadMetadataConfig.mode = "GKE_METADATA"` |
| No network policies | All-to-all traffic | Apply default-deny + explicit policies |
| Privileged containers | Root container escape | Set `securityContext.privileged = false` |
| No admission controllers | Insecure pods deployed | Deploy validating/mutating webhooks |
| No audit logging | Compliance blind spot | Enable GKE Cloud Logging sink |
| Overpermissioned service accounts | Lateral movement | Use Workload Identity + least-privilege IAM |
| Public container registry | Image tampering | Use private Artifact Registry + Binary Auth |

> **Attack Surface Reduction**
>
>
> Each misconfiguration listed above represents a verified attack vector. Fix all items before production deployment.


## Examples

See [examples.md](examples.md) for code examples.


## Full Reference

See [reference.md](reference.md) for complete documentation.


## Related Patterns

- Enforce
- Secure
- Patterns

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/secure/cloud-native/)
- [AEL Secure](https://adaptive-enforcement-lab.com/secure/)
