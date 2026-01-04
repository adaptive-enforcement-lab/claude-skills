# Clone infrastructure repository
git clone https://github.com/your-org/gke-infrastructure.git
cd gke-infrastructure

# Initialize Pulumi stack
pulumi stack init qac

# Configure cluster
pulumi config set gcp_project $PROJECT_ID
pulumi config set cluster_name qac-cluster
pulumi config set environment qac
pulumi config set team platform
pulumi config set cost_center engineering
pulumi config set admin_cidr_block 203.0.113.0/24

# Preview and deploy
pulumi preview
pulumi up

# Get cluster credentials
gcloud container clusters get-credentials qac-cluster \
  --region us-central1 \
  --project $PROJECT_ID