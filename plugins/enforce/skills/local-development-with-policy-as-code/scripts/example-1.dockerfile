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