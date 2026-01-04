---
name: oidc-federation-patterns - Examples
description: Code examples for OIDC Federation Patterns
---

# OIDC Federation Patterns - Examples


## Example 1: example-1.text


```text
repo:adaptive-enforcement-lab/api-service:*
```



## Example 2: example-2.text


```text
repo:adaptive-enforcement-lab/api-service:ref:refs/heads/main
```



## Example 3: example-3.text


```text
repo:adaptive-enforcement-lab/api-service:environment:production
```



## Example 4: example-4.text


```text
token.actions.githubusercontent.com:sub = "repo:adaptive-enforcement-lab/api-service:ref:refs/heads/*"
```



## Example 5: example-5.sh


```bash
gcloud iam workload-identity-pools create github-pool \
  --location=global \
  --display-name="GitHub Actions Pool"
```



## Example 6: example-6.sh


```bash
gcloud iam workload-identity-pools providers create-oidc github-provider \
  --location=global \
  --workload-identity-pool=github-pool \
  --issuer-uri=https://token.actions.githubusercontent.com \
  --attribute-mapping="google.subject=assertion.sub,attribute.repository=assertion.repository,attribute.repository_owner=assertion.repository_owner,attribute.ref=assertion.ref" \
  --attribute-condition="assertion.repository_owner == 'adaptive-enforcement-lab'"
```



## Example 7: example-7.sh


```bash
gcloud iam service-accounts add-iam-policy-binding deploy@my-project.iam.gserviceaccount.com \
  --role=roles/iam.workloadIdentityUser \
  --member="principalSet://iam.googleapis.com/projects/123456789/locations/global/workloadIdentityPools/github-pool/attribute.repository/adaptive-enforcement-lab/api-service"
```



## Example 8: example-8.sh


```bash
--member="principalSet://iam.googleapis.com/projects/123456789/locations/global/workloadIdentityPools/github-pool/attribute.repository/adaptive-enforcement-lab/api-service/attribute.environment/production"
```



## Example 9: example-9.yaml


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
```



