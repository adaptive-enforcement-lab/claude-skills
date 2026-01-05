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

Fundamental cluster security configuration covering private networking, identity federation, and image verification.

## Overview

This section covers the foundational security configurations for GKE clusters:

1. **[Private GKE Cluster](private-cluster.md)** - Private control plane, VPC networking, and encrypted etcd
2. **[Workload Identity](workload-identity.md)** - Pod-to-GCP authentication without service account keys
3. **[Binary Authorization](binary-authorization.md)** - Shielded Nodes and image verification

> **Public Cluster Risk**
>
>
> Public control planes expose your cluster API to the internet. Even with strong authentication, this increases attack surface and is not recommended for production.
>

## Security Principles

### Defense in Depth

- **Private Control Plane**: API server accessible only from authorized networks
- **Workload Identity**: Pods authenticate to GCP without static credentials
- **Shielded Nodes**: Secure boot, measured boot, and integrity monitoring
- **Binary Authorization**: Only verified container images run on the cluster

### Configuration Management

All configurations use Terraform for Infrastructure as Code, enabling:

- Repeatable deployments across environments
- Version-controlled security policies
- Automated compliance validation
- Drift detection and remediation

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
>

## Quick Start


*See [examples.md](examples.md) for detailed code examples.*

## Related Configuration

- **[Network Security](../network-security/index.md)** - VPC-native networking, Network Policies, Private Service Connect
- **[IAM Configuration](../iam-configuration/index.md)** - Least-privilege IAM, audit logging, service accounts
- **[Runtime Security](../runtime-security/index.md)** - Pod Security Standards, admission controllers, monitoring

### Overview

This section covers the foundational security configurations for GKE clusters:

1. **[Private GKE Cluster](private-cluster.md)** - Private control plane, VPC networking, and encrypted etcd
2. **[Workload Identity](workload-identity.md)** - Pod-to-GCP authentication without service account keys
3. **[Binary Authorization](binary-authorization.md)** - Shielded Nodes and image verification

> **Public Cluster Risk**
>
>
> Public control planes expose your cluster API to the internet. Even with strong authentication, this increases attack surface and is not recommended for production.

### Security Principles

### Defense in Depth

- **Private Control Plane**: API server accessible only from authorized networks
- **Workload Identity**: Pods authenticate to GCP without static credentials
- **Shielded Nodes**: Secure boot, measured boot, and integrity monitoring
- **Binary Authorization**: Only verified container images run on the cluster

### Configuration Management

All configurations use Terraform for Infrastructure as Code, enabling:

- Repeatable deployments across environments
- Version-controlled security policies
- Automated compliance validation
- Drift detection and remediation

### Prerequisites

- GCP project with billing enabled
- `gcloud` CLI installed and authenticated
- Terraform 1.0+
- kubectl configured for cluster access
- Appropriate IAM permissions (Project Editor or Security Admin roles)

> **Production Warning**
>
>
> These configurations enforce strict security controls. Test in QAC/DEV before production deployment.

### Quick Start


*See [examples.md](examples.md) for detailed code examples.*

### Related Configuration

- **[Network Security](../network-security/index.md)** - VPC-native networking, Network Policies, Private Service Connect
- **[IAM Configuration](../iam-configuration/index.md)** - Least-privilege IAM, audit logging, service accounts
- **[Runtime Security](../runtime-security/index.md)** - Pod Security Standards, admission controllers, monitoring


## Key Principles

### Defense in Depth

- **Private Control Plane**: API server accessible only from authorized networks
- **Workload Identity**: Pods authenticate to GCP without static credentials
- **Shielded Nodes**: Secure boot, measured boot, and integrity monitoring
- **Binary Authorization**: Only verified container images run on the cluster

### Configuration Management

All configurations use Terraform for Infrastructure as Code, enabling:

- Repeatable deployments across environments
- Version-controlled security policies
- Automated compliance validation
- Drift detection and remediation


## Related Patterns

- Network Security
- IAM Configuration
- Runtime Security

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/secure/cloud-native/)
- [AEL Secure](https://adaptive-enforcement-lab.com/secure/)
