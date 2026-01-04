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


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/secure/cloud-native/)
- [AEL Secure](https://adaptive-enforcement-lab.com/secure/)
