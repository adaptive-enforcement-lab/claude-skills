# Install OPA CLI
brew install opa  # macOS
# or download from https://www.openpolicyagent.org/docs/latest/#running-opa

# Test Rego policy locally
opa test constraint-template.yaml test-cases.yaml -v

# Example test case
# test-cases.yaml
package k8sblockprivileged

test_privileged_container_blocked {
  violation[{"msg": msg}] with input as {
    "review": {
      "object": {
        "spec": {
          "containers": [{
            "name": "test",
            "securityContext": {"privileged": true}
          }]
        }
      }
    }
  }
}