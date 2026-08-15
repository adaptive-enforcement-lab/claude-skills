# Policy Packaging

Source: https://adaptive-enforcement-lab.com/enforce/policy-as-code/policy-packaging/

Build the policy-platform container that aggregates all policy sources and tools.

## Overview

The policy-platform container is a multi-stage Docker build that:

1. Pulls policy repositories as OCI containers
2. Installs policy validation tools (Kyverno, Pluto, Spectral, Helm)
3. Aggregates everything into a single distributable image

> **One Container, All Policies**
>
> The policy-platform image runs identically in local dev, CI pipelines, and reference environments. Zero configuration drift.
>

---

## Dockerfile Architecture

### Multi-Stage Build Pattern

```dockerfile
# Stage 1-3: Pull policy repos as OCI containers
FROM policy-repo-1:tag AS policy_repo_1
FROM policy-repo-2:tag AS policy_repo_2
FROM policy-repo-3:tag AS policy_repo_3

# Final stage: Aggregate and install tools
FROM alpine:3.24.1

# Install tools
RUN apk add curl bash helm yq

# Install Kyverno CLI
RUN curl -sSL ...kyverno.tar.gz | tar -xz

# Copy all policy repos
COPY --from=policy_repo_1 /repos/repo1/ /repos/repo1/
COPY --from=policy_repo_2 /repos/repo2/ /repos/repo2/
COPY --from=policy_repo_3 /repos/repo3/ /repos/repo3/
```

**Key Benefits**:

- Policy repos versioned independently
- Single container aggregates all
- Tools bundled with policies

---

## Complete Dockerfile

See [Multi-Source Policies](../multi-source-policies/index.md) for full Dockerfile example.

**Core components**:

1. **Base image**: Alpine Linux (small, secure)
2. **Tools**: Kyverno CLI v1.13.2, Pluto v5.21.1, Spectral, Helm, yq
3. **Policy repos**: Copied from dependent OCI containers

---

## Tool Selection

### Kyverno CLI

**Purpose**: Policy validation and testing

**Installation**:

```dockerfile
RUN curl -sSL https://github.com/kyverno/kyverno/releases/download/v1.13.2/kyverno-cli_v1.13.2_linux_x86_64.tar.gz \
  | tar -xz -C /usr/local/bin
```

### Pluto

**Purpose**: Deprecated API detection

**Installation**:

```dockerfile
RUN curl -sSL https://github.com/FairwindsOps/pluto/releases/download/v5.21.1/pluto_5.21.1_linux_amd64.tar.gz \
  | tar xz -C /usr/local/bin
```

### Spectral

**Purpose**: Schema validation

**Installation**:

```dockerfile
RUN latest_spectral=$(curl -sSL https://api.github.com/repos/stoplightio/spectral/releases/latest | grep 'tag_name' | cut -d\" -f4) && \
    curl -sSL https://github.com/stoplightio/spectral/releases/download/${latest_spectral}/spectral-alpine-x64 \
      -o /usr/local/bin/spectral && \
    chmod +x /usr/local/bin/spectral
```

> **Pin Tool Versions**
>
> Always pin specific tool versions in Dockerfile. Dynamic `latest` tags cause non-reproducible builds.
>

---

## Build Process

### Local Build

```bash
docker build -t policy-platform:latest -f ci/Dockerfile .
```

### CI Build

```yaml
# Bitbucket Pipelines
- step:
    name: Build Policy Platform
    services:
      - docker
    script:
      - docker build -t policy-platform:${BITBUCKET_BUILD_NUMBER} -f ci/Dockerfile .
      - docker tag policy-platform:${BITBUCKET_BUILD_NUMBER} policy-platform:latest
      - docker push policy-platform:${BITBUCKET_BUILD_NUMBER}
      - docker push policy-platform:latest
```

### GitHub Actions

```yaml
- name: Build and Push
  uses: docker/build-push-action@v5
  with:
    context: .
    file: ./ci/Dockerfile
    push: true
    tags: |
      policy-platform:${{ github.sha }}
      policy-platform:latest
```

---

## Next Steps

- **[Distribution](distribution.md)** - Versioning, testing, optimization
- **[Maintenance](maintenance.md)** - Troubleshooting and best practices
- **[Multi-Source Policies](../multi-source-policies/index.md)** - Policy aggregation patterns
