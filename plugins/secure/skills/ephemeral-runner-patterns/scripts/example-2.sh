#!/bin/bash
# Ephemeral runner with gVisor container runtime

set -euo pipefail

# Requires gVisor runsc runtime configured
# See: https://gvisor.dev/docs/user_guide/install/

RUNNER_VERSION="2.311.0"
RUNNER_IMAGE="ghcr.io/actions/runner:${RUNNER_VERSION}"
RUNNER_TOKEN="${1:?Runner registration token required}"
RUNNER_NAME="gvisor-ephemeral-$(date +%s)-$(openssl rand -hex 4)"

echo "==> Starting gVisor-isolated runner: ${RUNNER_NAME}"

podman run \
  --rm \
  --runtime /usr/local/bin/runsc \
  --name "${RUNNER_NAME}" \
  --read-only \
  --tmpfs /tmp:rw,size=2G \
  --tmpfs /opt/runner/_work:rw,size=8G \
  --security-opt no-new-privileges=true \
  --cap-drop ALL \
  --network slirp4netns \
  --env RUNNER_TOKEN="${RUNNER_TOKEN}" \
  --env RUNNER_NAME="${RUNNER_NAME}" \
  --env RUNNER_EPHEMERAL=true \
  "${RUNNER_IMAGE}"