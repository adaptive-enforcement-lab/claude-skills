---
name: runtime-security
description: >-
  Pod Security Standards and admission controllers for GKE. Runtime monitoring with Falco and behavioral analysis to detect anomalous workload activity.
---

# Runtime Security

## When to Use This Skill

This section covers runtime security for GKE clusters:

- **Pod Security Standards**: Namespace-level security policies (baseline, restricted)
- **Admission Controllers**: Pre-deployment validation and policy enforcement
- **Runtime Monitoring**: Behavioral detection with Falco or GKE Cloud Logging


## Prerequisites

- GCP project with billing enabled
- Terraform 1.0+
- kubectl configured for cluster access


## Implementation

- **[Cluster Configuration](../cluster-configuration/index.md)** - Private GKE, Workload Identity
- **[Network Security](../network-security/index.md)** - VPC networking, Network Policies
- **[IAM Configuration](../iam-configuration/index.md)** - Least-privilege IAM


## Key Principles

### Defense in Depth

Multiple layers of runtime security controls:

- Pod Security Standards enforce secure defaults
- Admission controllers block invalid configurations
- Runtime monitoring detects anomalous behavior
- Audit logging captures all activity

### Secure by Default

Production workloads must meet strict security requirements:

- Run as non-root user
- Read-only root filesystem
- Drop all Linux capabilities
- No privilege escalation
- Resource limits defined

### Continuous Monitoring

Runtime monitoring provides visibility into pod behavior:

- Process execution tracking
- File access monitoring
- Network connection detection
- System call auditing


## Related Patterns

- Cluster Configuration
- Network Security
- IAM Configuration

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/secure/cloud-native/)
- [AEL Secure](https://adaptive-enforcement-lab.com/secure/)
