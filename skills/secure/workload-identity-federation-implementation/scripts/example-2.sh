# 1. Enable Workload Identity on cluster
gcloud container clusters update my-cluster \
  --workload-pool=PROJECT_ID.svc.id.goog \
  --zone us-central1-a

# 2. Create Kubernetes ServiceAccount
kubectl create serviceaccount app-sa -n production

# 3. Create GCP service account
gcloud iam service-accounts create app-gcp \
  --display-name "App workload identity"

# 4. Grant GCP permissions
gcloud projects add-iam-policy-binding PROJECT_ID \
  --member="serviceAccount:app-gcp@PROJECT_ID.iam.gserviceaccount.com" \
  --role="roles/storage.objectViewer"

# 5. Bind Kubernetes SA to GCP SA
gcloud iam service-accounts add-iam-policy-binding \
  app-gcp@PROJECT_ID.iam.gserviceaccount.com \
  --role="roles/iam.workloadIdentityUser" \
  --member="serviceAccount:PROJECT_ID.svc.id.goog[production/app-sa]"

# 6. Annotate Kubernetes ServiceAccount
kubectl annotate serviceaccount app-sa \
  -n production \
  iam.gke.io/gcp-service-account=app-gcp@PROJECT_ID.iam.gserviceaccount.com

# 7. Deploy pod with annotated ServiceAccount
kubectl apply -f deployment.yaml