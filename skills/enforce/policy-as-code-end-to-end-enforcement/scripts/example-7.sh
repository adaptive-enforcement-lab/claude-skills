docker run policy-platform:latest \
  kyverno apply /repos/security-policy/ \
  --resource deployment.yaml