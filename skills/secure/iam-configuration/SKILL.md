---
name: iam-configuration
description: >-
  Least-privilege IAM roles for GKE nodes and workloads. Workload Identity Federation for external authentication and comprehensive audit logging for visibility.
---

# IAM Configuration

## When to Use This Skill

This section covers identity and access management for GKE clusters:

- **Service Account Roles**: Fine-grained IAM permissions for nodes, admins, and developers
- **Workload Identity Federation**: External identity provider integration (GitHub, OIDC)
- **Audit Logging**: Complete visibility into cluster management and API access


## Prerequisites

- GCP project with billing enabled
- Terraform 1.0+
- Appropriate IAM permissions (Security Admin or Project Editor)


## Implementation

Identity and access management controls who can do what in your cluster. Least-privilege service accounts minimize blast radius. Workload Identity Federation enables external identity integration. Audit logging provides complete visibility.

> **IAM Security Layers**
>
>
> 1. **[Least Privilege Roles](least-privilege-roles.md)** - Minimal permissions for service accounts
> 2. **[Workload Identity Federation](workload-identity-federation.md)** - GitHub Actions and external auth
> 3. **[Audit Logging](audit-logging.md)** - Comprehensive activity tracking
>

## Overview

This section covers identity and access management for GKE clusters:

- **Service Account Roles**: Fine-grained IAM permissions for nodes, admins, and developers
- **Workload Identity Federation**: External identity provider integration (GitHub, OIDC)
- **Audit Logging**: Complete visibility into cluster management and API access

## Security Principles

### Least Privilege

Grant only the minimum IAM roles required for each service account:

- Node service accounts: Logging, monitoring only
- Application service accounts: Specific GCP resource access
- Developer accounts: Read-only cluster access
- Admin accounts: Full cluster management (limited users)

### External Identity Integration

Workload Identity Federation enables pods and external systems to authenticate without static credentials:

- GitHub Actions: OIDC token exchange
- External CI/CD: Custom identity providers
- Multi-cloud workloads: Cross-cloud authentication

### Complete Audit Trail

Comprehensive audit logging captures all cluster activity:

- API server requests (create, update, delete)
- Authentication attempts and failures
- IAM policy changes
- Service account usage

## Prerequisites

- GCP project with billing enabled
- Terraform 1.0+
- Appropriate IAM permissions (Security Admin or Project Editor)

## Related Configuration

- **[Cluster Configuration](../cluster-configuration/index.md)** - Private GKE, Workload Identity, Shielded Nodes
- **[Network Security](../network-security/index.md)** - VPC-native networking, Network Policies
- **[Runtime Security](../runtime-security/index.md)** - Pod Security Standards, admission controllers

### Overview

This section covers identity and access management for GKE clusters:

- **Service Account Roles**: Fine-grained IAM permissions for nodes, admins, and developers
- **Workload Identity Federation**: External identity provider integration (GitHub, OIDC)
- **Audit Logging**: Complete visibility into cluster management and API access

### Security Principles

### Least Privilege

Grant only the minimum IAM roles required for each service account:

- Node service accounts: Logging, monitoring only
- Application service accounts: Specific GCP resource access
- Developer accounts: Read-only cluster access
- Admin accounts: Full cluster management (limited users)

### External Identity Integration

Workload Identity Federation enables pods and external systems to authenticate without static credentials:

- GitHub Actions: OIDC token exchange
- External CI/CD: Custom identity providers
- Multi-cloud workloads: Cross-cloud authentication

### Complete Audit Trail

Comprehensive audit logging captures all cluster activity:

- API server requests (create, update, delete)
- Authentication attempts and failures
- IAM policy changes
- Service account usage

### Prerequisites

- GCP project with billing enabled
- Terraform 1.0+
- Appropriate IAM permissions (Security Admin or Project Editor)

### Related Configuration

- **[Cluster Configuration](../cluster-configuration/index.md)** - Private GKE, Workload Identity, Shielded Nodes
- **[Network Security](../network-security/index.md)** - VPC-native networking, Network Policies
- **[Runtime Security](../runtime-security/index.md)** - Pod Security Standards, admission controllers


## Key Principles

### Least Privilege

Grant only the minimum IAM roles required for each service account:

- Node service accounts: Logging, monitoring only
- Application service accounts: Specific GCP resource access
- Developer accounts: Read-only cluster access
- Admin accounts: Full cluster management (limited users)

### External Identity Integration

Workload Identity Federation enables pods and external systems to authenticate without static credentials:

- GitHub Actions: OIDC token exchange
- External CI/CD: Custom identity providers
- Multi-cloud workloads: Cross-cloud authentication

### Complete Audit Trail

Comprehensive audit logging captures all cluster activity:

- API server requests (create, update, delete)
- Authentication attempts and failures
- IAM policy changes
- Service account usage


## Related Patterns

- Cluster Configuration
- Network Security
- Runtime Security

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/secure/cloud-native/)
- [AEL Secure](https://adaptive-enforcement-lab.com/secure/)
