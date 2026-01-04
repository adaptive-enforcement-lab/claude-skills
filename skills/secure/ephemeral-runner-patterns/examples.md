---
name: ephemeral-runner-patterns - Examples
description: Code examples for Ephemeral Runner Patterns
---

# Ephemeral Runner Patterns - Examples


## Example 1: example-1.sh


```bash
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
```



## Example 2: example-2.sh


```bash
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
```



## Example 3: example-3.ini


```ini
# /etc/systemd/system/github-runner-ephemeral@.service
# Systemd template for ephemeral container runners

[Unit]
Description=GitHub Actions Ephemeral Runner (Container %i)
After=network-online.target
Wants=network-online.target

[Service]
Type=simple
User=github-runner
Environment=RUNNER_VERSION=2.311.0
Environment=RUNNER_IMAGE=ghcr.io/actions/runner:${RUNNER_VERSION}
Environment=RUNNER_TOKEN_FILE=/etc/github-runner/token
ExecStartPre=/usr/bin/podman pull ${RUNNER_IMAGE}
ExecStart=/opt/runner-orchestrator/run-ephemeral-job.sh $(cat ${RUNNER_TOKEN_FILE})
Restart=always
RestartSec=10
TimeoutStopSec=30

# Security hardening
NoNewPrivileges=true
PrivateTmp=true
ProtectSystem=strict
ProtectHome=true
ReadOnlyPaths=/
ReadWritePaths=/opt/github-runner

[Install]
WantedBy=multi-user.target
```



## Example 4: example-4.sh


```bash
# Enable multiple concurrent ephemeral runners
systemctl enable github-runner-ephemeral@{1..5}.service
systemctl start github-runner-ephemeral@{1..5}.service
```



## Example 5: example-5.sh


```bash
#!/bin/bash
# Create GCP instance template for ephemeral runners

set -euo pipefail

PROJECT_ID="my-gcp-project"
REGION="us-central1"
ZONE="${REGION}-a"
TEMPLATE_NAME="github-runner-ephemeral-$(date +%Y%m%d-%H%M%S)"
SERVICE_ACCOUNT="github-runner@${PROJECT_ID}.iam.gserviceaccount.com"

# Create instance template with startup script
gcloud compute instance-templates create "${TEMPLATE_NAME}" \
  --project="${PROJECT_ID}" \
  --machine-type=e2-medium \
  --image-family=ubuntu-2204-lts \
  --image-project=ubuntu-os-cloud \
  --boot-disk-size=20GB \
  --boot-disk-type=pd-standard \
  --service-account="${SERVICE_ACCOUNT}" \
  --scopes=cloud-platform \
  --metadata=enable-oslogin=TRUE \
  --metadata-from-file=startup-script=/opt/runner-orchestrator/vm-startup.sh \
  --tags=github-runner,ephemeral \
  --network-interface=network=default,no-address

# Create managed instance group with autoscaling
gcloud compute instance-groups managed create github-runners-ephemeral \
  --project="${PROJECT_ID}" \
  --base-instance-name=runner \
  --template="${TEMPLATE_NAME}" \
  --size=0 \
  --zone="${ZONE}"

# Configure autoscaling based on job queue
gcloud compute instance-groups managed set-autoscaling github-runners-ephemeral \
  --project="${PROJECT_ID}" \
  --zone="${ZONE}" \
  --min-num-replicas=0 \
  --max-num-replicas=10 \
  --cool-down-period=60 \
  --mode=on \
  --scale-based-on-cpu \
  --target-cpu-utilization=0.6
```



## Example 6: example-6.sh


```bash
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
```



## Example 7: example-7.json


```json
{
  "builders": [
    {
      "type": "googlecompute",
      "project_id": "my-gcp-project",
      "source_image_family": "ubuntu-2204-lts",
      "zone": "us-central1-a",
      "image_name": "github-runner-hardened-{{timestamp}}",
      "image_family": "github-runner-hardened",
      "ssh_username": "packer",
      "machine_type": "e2-medium",
      "disk_size": 20
    }
  ],
  "provisioners": [
    {
      "type": "shell",
      "script": "scripts/hardening/os-baseline.sh"
    },
    {
      "type": "shell",
      "script": "scripts/hardening/cis-benchmarks.sh"
    },
    {
      "type": "shell",
      "script": "scripts/hardening/firewall-rules.sh"
    },
    {
      "type": "shell",
      "script": "scripts/install-runner.sh"
    },
    {
      "type": "shell",
      "inline": [
        "echo 'Hardened runner image build complete'",
        "echo 'Image includes: OS hardening, firewall, audit logging, runner software'",
        "echo 'Startup script will configure ephemeral mode at boot'"
      ]
    }
  ]
}
```



## Example 8: example-8.yaml


```yaml
# arc-controller-install.yml
# Install Actions Runner Controller using Helm
```



