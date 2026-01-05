---
name: workload-identity-federation-implementation - Examples
description: Code examples for Workload Identity Federation Implementation
---

# Workload Identity Federation Implementation - Examples


## Example 1: example-1.mermaid


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



## Example 2: example-2.sh


```bash
# 1. Enable Workload Identity on cluster
gcloud container clusters update my-cluster \
  --workload-pool=PROJECT_ID.svc.id.goog \
  --zone us-central1-a

# 2. Create Kubernetes ServiceAccount
kubectl create serviceaccount app-sa -n production

# 3. Create GCP service account
gcloud iam service-accounts create app-gcp \
  --display-name "App workload identity"

# 4. Grant GCP permissions
gcloud projects add-iam-policy-binding PROJECT_ID \
  --member="serviceAccount:app-gcp@PROJECT_ID.iam.gserviceaccount.com" \
  --role="roles/storage.objectViewer"

# 5. Bind Kubernetes SA to GCP SA
gcloud iam service-accounts add-iam-policy-binding \
  app-gcp@PROJECT_ID.iam.gserviceaccount.com \
  --role="roles/iam.workloadIdentityUser" \
  --member="serviceAccount:PROJECT_ID.svc.id.goog[production/app-sa]"

# 6. Annotate Kubernetes ServiceAccount
kubectl annotate serviceaccount app-sa \
  -n production \
  iam.gke.io/gcp-service-account=app-gcp@PROJECT_ID.iam.gserviceaccount.com

# 7. Deploy pod with annotated ServiceAccount
kubectl apply -f deployment.yaml
```



## Example 3: example-3.yaml


```yaml
# Kubernetes Secret with private key
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



## Example 4: example-4.yaml


```yaml
# Kubernetes ServiceAccount with annotation
apiVersion: v1
kind: ServiceAccount
metadata:
  name: app-sa
  annotations:
    iam.gke.io/gcp-service-account: app-gcp@PROJECT_ID.iam.gserviceaccount.com
```



## Example 5: example-5.py


```python
from google.cloud import storage

# Credentials automatic
client = storage.Client(project='PROJECT_ID')
bucket = client.bucket('my-bucket')
blob = bucket.blob('data.txt')
blob.download_to_filename('data.txt')
```



## Example 6: example-6.py


```python
from google.cloud import secretmanager

client = secretmanager.SecretManagerServiceClient()
secret_name = f"projects/PROJECT_ID/secrets/api-key/versions/latest"
response = client.access_secret_version(request={"name": secret_name})
api_key = response.payload.data.decode('UTF-8')
```



## Example 7: example-7.sh


```bash
# SERVICE_ACCOUNT_A in PROJECT_A can impersonate SERVICE_ACCOUNT_B in PROJECT_B
gcloud iam service-accounts add-iam-policy-binding \
  service-account-b@PROJECT_B.iam.gserviceaccount.com \
  --role="roles/iam.serviceAccountUser" \
  --member="serviceAccount:service-account-a@PROJECT_A.iam.gserviceaccount.com"
```



