docker run --rm -v $(pwd):/workspace policy-platform:latest bash -c '\
  helm template security /repos/security-policy/charts/security-policy \
    -f /repos/security-policy/charts/security-policy/values.yaml \
  > /tmp/policies.yaml &&\
  kyverno apply /tmp/policies.yaml --resource /workspace/test-namespace.yaml\
'