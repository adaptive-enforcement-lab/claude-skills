---
name: local-development-with-policy-as-code - Examples
description: Code examples for Local Development with Policy-as-Code
---

# Local Development with Policy-as-Code - Examples


## Example 1: example-1.dockerfile


```dockerfile
# Pull policy repositories as OCI containers
FROM security-policy-repo:main AS security_policy_repo
FROM devops-policy-repo:main AS devops_policy_repo

# Alpine base with all tools
FROM alpine:3.22
RUN apk add curl bash ca-certificates git helm yq

# Install Kyverno CLI
RUN curl -sSL https://github.com/kyverno/kyverno/releases/download/v1.13.2/kyverno-cli_v1.13.2_linux_x86_64.tar.gz \
  | tar -xz -C /usr/local/bin

# Install Pluto
RUN curl -sSL https://github.com/FairwindsOps/pluto/releases/download/v5.21.1/pluto_5.21.1_linux_amd64.tar.gz \
  | tar xz -C /usr/local/bin

# Copy policies from dependent containers
COPY --from=security_policy_repo /repos/security-policy/ /repos/security-policy/
COPY --from=devops_policy_repo /repos/devops-policy/ /repos/devops-policy/

WORKDIR /repos
```



## Example 2: example-2.sh


```bash
docker run --rm \
  -v $(pwd):/workspace \
  policy-platform:latest \
  kyverno apply /repos/security-policy/ \
  --resource /workspace/deployment.yaml
```



## Example 3: example-3.text


```text
Applying 1 policy to 1 resource...

pass: 12/12
fail: 0/12
warn: 0/12
error: 0/12
skip: 0/12

All resources passed policy validation!
```



## Example 4: example-4.sh


```bash
$ docker run --rm -v $(pwd):/workspace policy-platform:latest \
  kyverno apply /repos/security-policy/ \
  --resource /workspace/bad-deployment.yaml

fail: 2/12
  require-resource-limits:
    Deployment/default/nginx: CPU and memory limits required
  disallow-latest-tag:
    Deployment/default/nginx: Container uses :latest tag
```



## Example 5: example-5.sh


```bash
# Render chart for staging environment
docker run --rm \
  -v $(pwd):/workspace \
  policy-platform:latest \
  helm template my-app /workspace/charts/my-app \
    -f /workspace/charts/my-app/values.yaml \
    -f /workspace/cd/staging/values.yaml \
  > staging-manifests.yaml

# Validate rendered manifests
docker run --rm \
  -v $(pwd):/workspace \
  policy-platform:latest \
  kyverno apply /repos/security-policy/ \
    --resource /workspace/staging-manifests.yaml
```



## Example 6: example-6.sh


```bash
for env in dev qac stg prd; do
  echo "Validating ${env} environment..."

  # Render manifests
  docker run --rm -v $(pwd):/workspace policy-platform:latest \
    helm template app /workspace/charts/app \
      -f /workspace/charts/app/values.yaml \
      -f /workspace/cd/${env}/values.yaml \
    > ${env}-manifests.yaml

  # Validate policies
  docker run --rm -v $(pwd):/workspace policy-platform:latest \
    kyverno apply /repos/security-policy/ \
      --resource /workspace/${env}-manifests.yaml \
      --audit-warn
done
```



