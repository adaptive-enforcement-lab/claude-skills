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