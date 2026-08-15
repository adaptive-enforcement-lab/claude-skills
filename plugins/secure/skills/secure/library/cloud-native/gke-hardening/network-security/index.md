# Network Security

Source: https://adaptive-enforcement-lab.com/secure/cloud-native/gke-hardening/network-security/

Network isolation is critical in multi-tenant clusters. VPC-native networking provides better performance and simpler network policies. Private Service Connect secures GCP service access. Cloud Armor defends against DDoS and application attacks.

> **Network Security Layers**
>
>
> 1. **[VPC-Native Networking](vpc-native.md)** - Container-native IP allocation
> 2. **[Network Policies](network-policies.md)** - Pod-to-pod traffic control
> 3. **[Private Service Connect](private-service-connect.md)** - Secure GCP service access
> 4. **[Cloud Armor](cloud-armor.md)** - DDoS protection and WAF
>

## Overview

This section covers network security configurations for GKE clusters:

- **VPC-Native Networking**: Container-native IP allocation with Alias IP ranges
- **Network Policies**: Zero-trust network model with default-deny ingress
- **Private Service Connect**: Private connectivity to GCP services
- **Cloud Armor**: Layer 7 DDoS protection and Web Application Firewall

## Security Principles

### Zero Trust Network

Implement default-deny network policies and explicitly allow traffic between services:

- All ingress traffic is blocked by default
- Only required pod-to-pod communication is permitted
- DNS and essential services are explicitly allowed
- Egress traffic is controlled per workload

### Private Connectivity

Route traffic through private endpoints for secure, isolated connectivity:

- No public IP addresses required
- Traffic stays on Google's backbone
- Simplified security policy management
- Cross-project access supported

### Layer 7 Protection

Cloud Armor provides application-level security:

- DDoS mitigation at the edge
- Geo-blocking and IP filtering
- Rate limiting and bot detection
- XSS and SQLi protection

## Prerequisites

- GCP project with billing enabled
- Terraform 1.0+
- kubectl configured for cluster access

## Related Configuration

- **[Cluster Configuration](../cluster-configuration/index.md)** - Private GKE, Workload Identity
- **[IAM Configuration](../iam-configuration/index.md)** - Least-privilege IAM
- **[Runtime Security](../runtime-security/index.md)** - Pod Security Standards
