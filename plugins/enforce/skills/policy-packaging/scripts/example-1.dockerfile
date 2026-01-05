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