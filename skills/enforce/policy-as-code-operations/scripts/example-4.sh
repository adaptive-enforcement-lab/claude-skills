# Update policy-platform container (rebuild with new policy)
docker build -t policy-platform:v1.0.3 -f ci/Dockerfile .
docker push policy-platform:v1.0.3

# Deploy to dev cluster
helm upgrade security-policy /repos/security-policy/charts/security-policy \
  --namespace kyverno \
  --values /repos/security-policy/cd/dev/values.yaml