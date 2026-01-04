for env in dev qac stg prd; do
  echo "Validating ${env} environment..."

  # Render manifests
  docker run --rm -v $(pwd):/workspace policy-platform:latest \
    helm template app /workspace/charts/app \
      -f /workspace/charts/app/values.yaml \
      -f /workspace/cd/${env}/values.yaml \
    > ${env}-manifests.yaml

  # Validate policies
  docker run --rm -v $(pwd):/workspace policy-platform:latest \
    kyverno apply /repos/security-policy/ \
      --resource /workspace/${env}-manifests.yaml \
      --audit-warn
done