# Render chart for staging environment
docker run --rm \
  -v $(pwd):/workspace \
  policy-platform:latest \
  helm template my-app /workspace/charts/my-app \
    -f /workspace/charts/my-app/values.yaml \
    -f /workspace/cd/staging/values.yaml \
  > staging-manifests.yaml

# Validate rendered manifests
docker run --rm \
  -v $(pwd):/workspace \
  policy-platform:latest \
  kyverno apply /repos/security-policy/ \
    --resource /workspace/staging-manifests.yaml