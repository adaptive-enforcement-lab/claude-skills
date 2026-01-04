#!/bin/bash
# /opt/runner-orchestrator/run-ephemeral-job.sh
# Ephemeral runner using Podman rootless containers

set -euo pipefail

RUNNER_VERSION="2.311.0"
RUNNER_IMAGE="ghcr.io/actions/runner:${RUNNER_VERSION}"
RUNNER_TOKEN="${1:?Runner registration token required}"
RUNNER_NAME="ephemeral-$(date +%s)-$(openssl rand -hex 4)"
RUNNER_LABELS="self-hosted,ephemeral,container"

echo "==> Starting ephemeral runner: ${RUNNER_NAME}"

# Pull latest runner image
podman pull "${RUNNER_IMAGE}"

# Run container with strict isolation
podman run \
  --rm \
  --name "${RUNNER_NAME}" \
  --read-only \
  --tmpfs /tmp:rw,noexec,nosuid,nodev,size=2G \
  --tmpfs /opt/runner/_work:rw,noexec,nosuid,nodev,size=8G \
  --security-opt no-new-privileges=true \
  --security-opt label=type:runner_t \
  --cap-drop ALL \
  --network slirp4netns:allow_host_loopback=false \
  --env RUNNER_TOKEN="${RUNNER_TOKEN}" \
  --env RUNNER_NAME="${RUNNER_NAME}" \
  --env RUNNER_LABELS="${RUNNER_LABELS}" \
  --env RUNNER_EPHEMERAL=true \
  "${RUNNER_IMAGE}"

echo "==> Runner ${RUNNER_NAME} completed and destroyed"