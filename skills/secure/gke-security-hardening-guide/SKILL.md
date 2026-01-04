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


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/secure/cloud-native/)
- [AEL Secure](https://adaptive-enforcement-lab.com/secure/)
