# Runtime Security

Source: https://adaptive-enforcement-lab.com/secure/cloud-native/gke-hardening/runtime-security/

Runtime security enforces policies on running workloads. Pod Security Standards prevent privilege escalation. Admission controllers validate manifests before deployment. Runtime monitoring detects anomalous behavior.

> **Runtime Security Layers**
>
>
> 1. **[Pod Security Standards](pod-security-standards.md)** - Baseline and restricted policies
> 2. **[Admission Controllers](admission-controllers.md)** - Pre-deployment validation
> 3. **[Runtime Monitoring](runtime-monitoring.md)** - Behavioral analysis and alerting
>

## Overview

This section covers runtime security for GKE clusters:

- **Pod Security Standards**: Namespace-level security policies (baseline, restricted)
- **Admission Controllers**: Pre-deployment validation and policy enforcement
- **Runtime Monitoring**: Behavioral detection with Falco or GKE Cloud Logging

## Security Principles

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

## Prerequisites

- GCP project with billing enabled
- Terraform 1.0+
- kubectl configured for cluster access

## Related Configuration

- **[Cluster Configuration](../cluster-configuration/index.md)** - Private GKE, Workload Identity
- **[Network Security](../network-security/index.md)** - VPC networking, Network Policies
- **[IAM Configuration](../iam-configuration/index.md)** - Least-privilege IAM
