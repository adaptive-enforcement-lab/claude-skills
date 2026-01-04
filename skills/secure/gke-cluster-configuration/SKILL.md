---
name: gke-cluster-configuration
description: >-
  Private GKE cluster setup, Workload Identity, and Shielded Nodes with Binary Authorization using Terraform.
---

# GKE Cluster Configuration

## When to Use This Skill

This section covers the foundational security configurations for GKE clusters:

1. **[Private GKE Cluster](private-cluster.md)** - Private control plane, VPC networking, and encrypted etcd
2. **[Workload Identity](workload-identity.md)** - Pod-to-GCP authentication without service account keys
3. **[Binary Authorization](binary-authorization.md)** - Shielded Nodes and image verification

> **Public Cluster Risk**
>
>
> Public control planes expose your cluster API to the internet. Even with strong authentication, this increases attack surface and is not recommended for production.


## Prerequisites

- GCP project with billing enabled
- `gcloud` CLI installed and authenticated
- Terraform 1.0+
- kubectl configured for cluster access
- Appropriate IAM permissions (Project Editor or Security Admin roles)

> **Production Warning**
>
>
> These configurations enforce strict security controls. Test in QAC/DEV before production deployment.


## Implementation


See the full implementation guide in the source documentation.







## References

- [Source Documentation](https://adaptive-enforcement-lab.com/secure/cloud-native/)
- [AEL Secure](https://adaptive-enforcement-lab.com/secure/)
