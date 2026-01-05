docker run policy-platform:latest \
  kyverno apply /repos/security-policy/policies.yaml \
  --resource my-deployment.yaml