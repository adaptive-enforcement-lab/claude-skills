# Initialize Terraform
terraform init

# Apply cluster configuration
terraform apply \
  -var="gcp_project=$PROJECT_ID" \
  -var="cluster_name=prod-cluster" \
  -var="environment=prd" \
  -var="team=platform" \
  -var="cost_center=engineering" \
  -var="admin_cidr_block=203.0.113.0/24"

# Get cluster credentials
gcloud container clusters get-credentials prod-cluster \
  --region us-central1 \
  --project $PROJECT_ID

# Verify private cluster
gcloud container clusters describe prod-cluster \
  --region us-central1 \
  --format="value(privateClusterConfig.enablePrivateNodes)"