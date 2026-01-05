---
name: secret-management-overview - Examples
description: Code examples for Secret Management Overview
---

# Secret Management Overview - Examples


## Example 1: example-1.yaml


```yaml
name: Deploy
on: [push]

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1

      # Repository secret - available to this repo only
      - run: echo "${{ secrets.DEPLOY_KEY }}" | base64 -d > deploy.key
```



## Example 2: example-2.yaml


```yaml
name: Publish
on: [release]

jobs:
  publish:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1

      # Organization secret - shared across repos
      - run: npm publish --registry=https://npm.example.com
        env:
          NPM_TOKEN: ${{ secrets.ORG_NPM_TOKEN }}
```



## Example 3: example-3.yaml


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

      # Environment secret - production only, requires approval
      - run: ./deploy.sh
        env:
          PROD_API_KEY: ${{ secrets.PROD_API_KEY }}
```



## Example 4: example-4.yaml


```yaml
env:
  DATABASE_PASSWORD: ${{ secrets.DB_PASSWORD }}
  API_TOKEN: ${{ secrets.EXTERNAL_API_TOKEN }}
```



## Example 5: example-5.yaml


```yaml
env:
  API_ENDPOINT: ${{ vars.API_URL }}
  ENVIRONMENT: ${{ vars.DEPLOY_ENV }}
```



## Example 6: example-6.yaml


```yaml
permissions:
  id-token: write  # Request OIDC token
  contents: read

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      # google-github-actions/auth v2.1.0
      - uses: google-github-actions/auth@f112390a2df9932162083945e46d439060d66ec2
        with:
          workload_identity_provider: 'projects/123/locations/global/workloadIdentityPools/github/providers/github-provider'
          service_account: 'deploy@project.iam.gserviceaccount.com'

      - run: gcloud compute instances list
```



## Example 7: example-7.yaml


```yaml
# DANGEROUS - Exposes secret in logs
- run: echo "Deploying with key ${{ secrets.DEPLOY_KEY }}"
```



## Example 8: example-8.yaml


```yaml
# Safe - secret passed via environment, not command line
- run: ./deploy.sh
  env:
    DEPLOY_KEY: ${{ secrets.DEPLOY_KEY }}
```



## Example 9: example-9.yaml


```yaml
# DANGEROUS - PR from fork can inject code
name: CI
on: [pull_request_target]  # Runs with repo secrets even for forks

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11  # v4.1.1
        with:
          ref: ${{ github.event.pull_request.head.sha }}  # Checks out PR code

      # Attacker controls test.sh via PR
      - run: ./test.sh
        env:
          API_KEY: ${{ secrets.API_KEY }}  # Exposed to attacker
```



