---
name: workload-identity-federation-implementation
description: >-
  Workload Identity Federation implementation guide. GKE setup, IAM bindings, ServiceAccount configuration, migration from service account keys, and troubleshooting patterns.
---

# Workload Identity Federation Implementation

## When to Use This Skill

### Cloud Storage Access

```python
from google.cloud import storage

# Credentials automatic
client = storage.Client(project='PROJECT_ID')
bucket = client.bucket('my-bucket')
blob = bucket.blob('data.txt')
blob.download_to_filename('data.txt')
```

### Secret Manager Access

```python
from google.cloud import secretmanager

client = secretmanager.SecretManagerServiceClient()
secret_name = f"projects/PROJECT_ID/secrets/api-key/versions/latest"
response = client.access_secret_version(request={"name": secret_name})
api_key = response.payload.data.decode('UTF-8')
```

### Cross-Project Access

```bash
# SERVICE_ACCOUNT_A in PROJECT_A can impersonate SERVICE_ACCOUNT_B in PROJECT_B
gcloud iam service-accounts add-iam-policy-binding \
  service-account-b@PROJECT_B.iam.gserviceaccount.com \
  --role="roles/iam.serviceAccountUser" \
  --member="serviceAccount:service-account-a@PROJECT_A.iam.gserviceaccount.com"
```


## Implementation

Containers need cloud access. But service account keys are **static credentials** that never rotate, frequently get stolen, and live forever.

Workload Identity Federation lets containers prove their identity to cloud providers without ever storing keys. The Kubernetes cluster itself becomes a trusted identity provider.

> **Production Hardening**
>
> Workload Identity eliminates the largest attack surface in containerized environments. This is foundational. Get it right.
>

## What is Workload Identity Federation?

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

## Architecture


*See [examples.md](examples.md) for detailed code examples.*

## Implementation Guide

This guide is split into focused modules:

### Setup

- **[Cluster Configuration](cluster-configuration.md)**: Enable Workload Identity on GKE clusters and node pools
- **[Service Account Binding](service-account-binding.md)**: Create service accounts and configure IAM bindings

### Application Integration

- **[Pod Configuration](pod-configuration.md)**: Deploy workloads and common GCP service access patterns
- **[Migration Guide](migration-guide.md)**: Migrate from service account keys with zero downtime

### Operations

- **[Troubleshooting](troubleshooting.md)**: Debug auth failures, token issues, permissions

## Quick Start


*See [examples.md](examples.md) for detailed code examples.*

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

## Benefits

### Security

- **No static credentials**: Tokens expire automatically
- **Immediate revocation**: Disable service account, access stops
- **Audit trail**: Cloud Audit Logs track all impersonation
- **Least privilege**: Fine-grained IAM per workload

### Operations

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

## Migration from Service Account Keys

### Before (Static Keys)


*See [examples.md](examples.md) for detailed code examples.*

**Problems:**

- Key never expires
- If leaked, must manually revoke and rotate
- Stored in cluster (potential exposure)
- No audit trail of usage

### After (Workload Identity)

```yaml
# Kubernetes ServiceAccount with annotation
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

## Use Cases

### Cloud Storage Access

```python
from google.cloud import storage

# Credentials automatic
client = storage.Client(project='PROJECT_ID')
bucket = client.bucket('my-bucket')
blob = bucket.blob('data.txt')
blob.download_to_filename('data.txt')
```

### Secret Manager Access

```python
from google.cloud import secretmanager

client = secretmanager.SecretManagerServiceClient()
secret_name = f"projects/PROJECT_ID/secrets/api-key/versions/latest"
response = client.access_secret_version(request={"name": secret_name})
api_key = response.payload.data.decode('UTF-8')
```

### Cross-Project Access

```bash
# SERVICE_ACCOUNT_A in PROJECT_A can impersonate SERVICE_ACCOUNT_B in PROJECT_B
gcloud iam service-accounts add-iam-policy-binding \
  service-account-b@PROJECT_B.iam.gserviceaccount.com \
  --role="roles/iam.serviceAccountUser" \
  --member="serviceAccount:service-account-a@PROJECT_A.iam.gserviceaccount.com"
```

## References

- [Workload Identity Documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity)
- [IAM Conditions](https://cloud.google.com/iam/docs/conditions-overview)
- [GitHub Actions Integration](https://github.com/google-github-actions/auth)
- [Best Practices](https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity#best_practices)

## Related Content

- [GKE Hardening Guide](../gke-hardening/index.md): Comprehensive GKE security configuration
- [IAM Configuration](../gke-hardening/iam-configuration/index.md): Least-privilege IAM patterns
- [Secure](../../index.md): Security discovery and remediation

*Workload Identity eliminates static keys. Tokens rotate automatically. Access revokes immediately. Audit trail complete. Zero-trust credential model in place.*

### What is Workload Identity Federation?

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

### Architecture


*See [examples.md](examples.md) for detailed code examples.*

### Implementation Guide

This guide is split into focused modules:

### Setup

- **[Cluster Configuration](cluster-configuration.md)**: Enable Workload Identity on GKE clusters and node pools
- **[Service Account Binding](service-account-binding.md)**: Create service accounts and configure IAM bindings

### Application Integration

- **[Pod Configuration](pod-configuration.md)**: Deploy workloads and common GCP service access patterns
- **[Migration Guide](migration-guide.md)**: Migrate from service account keys with zero downtime

### Operations

- **[Troubleshooting](troubleshooting.md)**: Debug auth failures, token issues, permissions

### Quick Start


*See [examples.md](examples.md) for detailed code examples.*

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

### Benefits

### Security

- **No static credentials**: Tokens expire automatically
- **Immediate revocation**: Disable service account, access stops
- **Audit trail**: Cloud Audit Logs track all impersonation
- **Least privilege**: Fine-grained IAM per workload

### Operations

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

### Migration from Service Account Keys

### Before (Static Keys)


*See [examples.md](examples.md) for detailed code examples.*

**Problems:**

- Key never expires
- If leaked, must manually revoke and rotate
- Stored in cluster (potential exposure)
- No audit trail of usage

### After (Workload Identity)

```yaml
# Kubernetes ServiceAccount with annotation
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

### Use Cases

### Cloud Storage Access

```python
from google.cloud import storage

# Credentials automatic
client = storage.Client(project='PROJECT_ID')
bucket = client.bucket('my-bucket')
blob = bucket.blob('data.txt')
blob.download_to_filename('data.txt')
```

### Secret Manager Access

```python
from google.cloud import secretmanager

client = secretmanager.SecretManagerServiceClient()
secret_name = f"projects/PROJECT_ID/secrets/api-key/versions/latest"
response = client.access_secret_version(request={"name": secret_name})
api_key = response.payload.data.decode('UTF-8')
```

### Cross-Project Access

```bash
# SERVICE_ACCOUNT_A in PROJECT_A can impersonate SERVICE_ACCOUNT_B in PROJECT_B
gcloud iam service-accounts add-iam-policy-binding \
  service-account-b@PROJECT_B.iam.gserviceaccount.com \
  --role="roles/iam.serviceAccountUser" \
  --member="serviceAccount:service-account-a@PROJECT_A.iam.gserviceaccount.com"
```

### References

- [Workload Identity Documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity)
- [IAM Conditions](https://cloud.google.com/iam/docs/conditions-overview)
- [GitHub Actions Integration](https://github.com/google-github-actions/auth)
- [Best Practices](https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity#best_practices)

### Related Content

- [GKE Hardening Guide](../gke-hardening/index.md): Comprehensive GKE security configuration
- [IAM Configuration](../gke-hardening/iam-configuration/index.md): Least-privilege IAM patterns
- [Secure](../../index.md): Security discovery and remediation

*Workload Identity eliminates static keys. Tokens rotate automatically. Access revokes immediately. Audit trail complete. Zero-trust credential model in place.*


## Examples

See [examples.md](examples.md) for code examples.


## Full Reference

See [reference.md](reference.md) for complete documentation.


## Related Patterns

- GKE Hardening Guide
- IAM Configuration
- Secure

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/secure/cloud-native/)
- [AEL Secure](https://adaptive-enforcement-lab.com/secure/)
