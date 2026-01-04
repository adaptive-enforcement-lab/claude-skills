$ docker run --rm -v $(pwd):/workspace policy-platform:latest \
  kyverno apply /repos/security-policy/ \
  --resource /workspace/bad-deployment.yaml

fail: 2/12
  require-resource-limits:
    Deployment/default/nginx: CPU and memory limits required
  disallow-latest-tag:
    Deployment/default/nginx: Container uses :latest tag