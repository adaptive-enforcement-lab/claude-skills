# GKE Cluster Configuration

Source: https://adaptive-enforcement-lab.com/secure/cloud-native/gke-hardening/cluster-configuration/

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

```bash
# Initialize Terraform
terraform init

# Apply cluster configuration
terraform apply \
  -var="gcp_project=$PROJECT_ID" \
  -var="cluster_name=prod-cluster" \
  -var="environment=prd" \
  -var="team=platform" \
  -var="cost_center=engineering" \
  -var="admin_cidr_block=203.0.113.0/24"

# Get cluster credentials
gcloud container clusters get-credentials prod-cluster \
  --region us-central1 \
  --project $PROJECT_ID

# Verify private cluster
gcloud container clusters describe prod-cluster \
  --region us-central1 \
  --format="value(privateClusterConfig.enablePrivateNodes)"
```

## Related Configuration

- **[Network Security](../network-security/index.md)** - VPC-native networking, Network Policies, Private Service Connect
- **[IAM Configuration](../iam-configuration/index.md)** - Least-privilege IAM, audit logging, service accounts
- **[Runtime Security](../runtime-security/index.md)** - Pod Security Standards, admission controllers, monitoring
