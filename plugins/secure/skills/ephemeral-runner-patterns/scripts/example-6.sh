#!/bin/bash
# /opt/runner-orchestrator/vm-startup.sh
# GCP VM startup script for ephemeral runner

set -euo pipefail

echo "==> Configuring ephemeral runner VM"

# Install runner
mkdir -p /opt/actions-runner && cd /opt/actions-runner
curl -o actions-runner-linux-x64-2.311.0.tar.gz \
  -L https://github.com/actions/runner/releases/download/v2.311.0/actions-runner-linux-x64-2.311.0.tar.gz
tar xzf actions-runner-linux-x64-2.311.0.tar.gz
rm actions-runner-linux-x64-2.311.0.tar.gz

# Fetch registration token from Secret Manager
RUNNER_TOKEN=$(gcloud secrets versions access latest --secret=github-runner-token)
RUNNER_NAME="vm-ephemeral-$(hostname)-$(date +%s)"
RUNNER_LABELS="self-hosted,ephemeral,vm,gcp"

# Register runner (ephemeral mode)
./config.sh \
  --url https://github.com/my-org/my-repo \
  --token "${RUNNER_TOKEN}" \
  --name "${RUNNER_NAME}" \
  --labels "${RUNNER_LABELS}" \
  --ephemeral \
  --unattended

# Run single job
./run.sh

# Self-destruct after job completion
echo "==> Job complete, destroying VM"
gcloud compute instances delete "$(hostname)" --zone="$(gcloud compute instances list --filter="name=$(hostname)" --format="value(zone)")" --quiet