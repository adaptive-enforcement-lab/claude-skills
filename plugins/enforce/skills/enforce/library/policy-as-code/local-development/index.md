# Local Development with Policy-as-Code

Source: https://adaptive-enforcement-lab.com/enforce/policy-as-code/local-development/

Run the same policy validation locally that runs in CI. Catch issues in seconds, not hours.

## Overview

The policy-platform container includes all tools needed for local policy validation:

- **Kyverno CLI** - Policy validation and testing
- **Pluto** - Deprecated API detection
- **Helm** - Chart rendering and linting
- **Spectral** - OpenAPI/values schema validation
- **yq** - YAML processing

> **Zero Local Setup Required**
>
> One container contains all policies and tools. No local installations. Pull the container, run validations. Same environment as CI.
>

---

## The Local Development Container

### What's Inside

The policy-platform container is a multi-stage build that aggregates policies from multiple repositories:

```dockerfile
# Pull policy repositories as OCI containers
FROM security-policy-repo:main AS security_policy_repo
FROM devops-policy-repo:main AS devops_policy_repo

# Alpine base with all tools
FROM alpine:3.24
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

**Key Insight**: Policy repos are also OCI containers. Multi-stage build pulls them automatically.

---

## Basic Usage

### Running Policy Validation

Validate a Kubernetes manifest against all policies:

```bash
docker run --rm \
  -v $(pwd):/workspace \
  policy-platform:latest \
  kyverno apply /repos/security-policy/ \
  --resource /workspace/deployment.yaml
```

**Output**:

```text
Applying 1 policy to 1 resource...

pass: 12/12
fail: 0/12
warn: 0/12
error: 0/12
skip: 0/12

All resources passed policy validation!
```

### Validation Failure Example

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

> **Fix Before Committing**
>
> Local validation catches issues before CI. Fix violations now - no 20-minute CI feedback loop.
>

---

## Helm Chart Validation

### Rendering Charts with Environment Values

Real-world Helm charts need environment-specific values:

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

**This matches exactly what CI does.**

### Multi-Environment Validation

Validate across all environments before pushing:

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

**Catch environment-specific issues locally.**

---

## Next Steps

- **[Advanced Validation](advanced-validation.md)** - Policy reports, deprecated API detection, schema validation
- **[Workflow Integration](workflow-integration.md)** - Pre-commit hooks, Make targets, troubleshooting
- **[CI Integration](../ci-integration/index.md)** - Automate policy checks in pipelines
