---
name: policy-packaging - Examples
description: Code examples for Policy Packaging
---

# Policy Packaging - Examples


## Example 1: example-1.dockerfile


```dockerfile
# Stage 1-3: Pull policy repos as OCI containers
FROM policy-repo-1:tag AS policy_repo_1
FROM policy-repo-2:tag AS policy_repo_2
FROM policy-repo-3:tag AS policy_repo_3

# Final stage: Aggregate and install tools
FROM alpine:3.22.1

# Install tools
RUN apk add curl bash helm yq

# Install Kyverno CLI
RUN curl -sSL ...kyverno.tar.gz | tar -xz

# Copy all policy repos
COPY --from=policy_repo_1 /repos/repo1/ /repos/repo1/
COPY --from=policy_repo_2 /repos/repo2/ /repos/repo2/
COPY --from=policy_repo_3 /repos/repo3/ /repos/repo3/
```



## Example 2: example-2.dockerfile


```dockerfile
RUN curl -sSL https://github.com/kyverno/kyverno/releases/download/v1.13.2/kyverno-cli_v1.13.2_linux_x86_64.tar.gz \
  | tar -xz -C /usr/local/bin
```



## Example 3: example-3.dockerfile


```dockerfile
RUN curl -sSL https://github.com/FairwindsOps/pluto/releases/download/v5.21.1/pluto_5.21.1_linux_amd64.tar.gz \
  | tar xz -C /usr/local/bin
```



## Example 4: example-4.dockerfile


```dockerfile
RUN latest_spectral=$(curl -sSL https://api.github.com/repos/stoplightio/spectral/releases/latest | grep 'tag_name' | cut -d\" -f4) && \
    curl -sSL https://github.com/stoplightio/spectral/releases/download/${latest_spectral}/spectral-alpine-x64 \
      -o /usr/local/bin/spectral && \
    chmod +x /usr/local/bin/spectral
```



## Example 5: example-5.sh


```bash
docker build -t policy-platform:latest -f ci/Dockerfile .
```



## Example 6: example-6.yaml


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



## Example 7: example-7.yaml


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



