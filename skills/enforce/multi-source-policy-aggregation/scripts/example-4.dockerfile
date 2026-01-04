# Stage 1: Pull DevOps policies
FROM europe-west6-docker.pkg.dev/ops/charts/devops-policy-repo:main AS devops_policy_repo

# Stage 2: Pull Security policies
FROM europe-west6-docker.pkg.dev/ops/charts/security-policy-repo:main AS security_policy_repo

# Stage 3: Pull Application-specific policies
FROM europe-west6-docker.pkg.dev/ops/charts/backend-applications-repo:main AS backend_applications_repo

# Final stage: Combine everything
FROM alpine:3.22.1

# Install policy tools
RUN apk add --no-cache \
    curl bash ca-certificates git tar gzip \
    libc6-compat helm yq

# Install Kyverno CLI
RUN mkdir -p ./bin/kyverno && \
    curl -sSL https://github.com/kyverno/kyverno/releases/download/v1.13.2/kyverno-cli_v1.13.2_linux_x86_64.tar.gz \
      | tar -xz -C ./bin/kyverno && \
    chmod +x ./bin/kyverno/kyverno && \
    cp ./bin/kyverno/kyverno /usr/local/bin/

# Install Pluto
RUN mkdir -p ./bin/pluto && \
    curl -sSL https://github.com/FairwindsOps/pluto/releases/download/v5.21.1/pluto_5.21.1_linux_amd64.tar.gz \
      | tar xz -C ./bin/pluto && \
    chmod +x ./bin/pluto && \
    cp ./bin/pluto/pluto /usr/local/bin/

# Install Spectral
RUN mkdir -p ./bin/spectral && \
    latest_spectral=$(curl -sSL https://api.github.com/repos/stoplightio/spectral/releases/latest | grep 'tag_name' | cut -d\" -f4) && \
    curl -sSL https://github.com/stoplightio/spectral/releases/download/${latest_spectral}/spectral-alpine-x64 \
      -o ./bin/spectral/spectral && \
    chmod +x ./bin/spectral/spectral && \
    cp ./bin/spectral/spectral /usr/local/bin/

# Copy all policy repositories
COPY --from=devops_policy_repo /repos/devops-policy/ /repos/devops-policy/
COPY --from=security_policy_repo /repos/security-policy/ /repos/security-policy/
COPY --from=backend_applications_repo /repos/backend-applications/ /repos/backend-applications/

WORKDIR /repos