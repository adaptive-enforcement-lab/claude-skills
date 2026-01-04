---
name: opa-policy-templates - Examples
description: Code examples for OPA Policy Templates
---

# OPA Policy Templates - Examples


## Example 1: example-1.yaml


```yaml
apiVersion: templates.gatekeeper.sh/v1beta1
kind: ConstraintTemplate
metadata:
  name: k8sblockprivileged
spec:
  crd:
    spec:
      names:
        kind: K8sBlockPrivileged
  targets:
    - target: admission.k8s.gatekeeper.sh
      rego: |
        package k8sblockprivileged
        violation[{"msg": msg}] {
          # Rego policy logic here
        }
```



## Example 2: example-2.yaml


```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sBlockPrivileged
metadata:
  name: block-privileged-containers
spec:
  enforcementAction: dryrun  # Use 'deny' for enforcement
  match:
    kinds:
      - apiGroups: [""]
        kinds: ["Pod"]
    namespaces:
      - "production"
      - "staging"
```



## Example 3: example-3.sh


```bash
# 1. Deploy constraint template (defines policy logic)
kubectl apply -f constraint-template.yaml

# 2. Deploy constraint in audit mode (dryrun)
kubectl apply -f constraint.yaml

# 3. Monitor violations
kubectl get constraints
kubectl get <constraint-kind> <constraint-name> -o yaml

# 4. Check audit results
kubectl get constraints -o json | jq '.items[].status.violations'

# 5. Fix non-compliant resources
kubectl get pods -n production --show-labels

# 6. Switch to enforcement mode after validation
kubectl patch <constraint-kind> <constraint-name> \
  --type merge \
  -p '{"spec":{"enforcementAction":"deny"}}'
```



## Example 4: example-4.sh


```bash
# List all constraints
kubectl get constraints

# Get detailed status for a specific constraint
kubectl get k8sblockprivileged block-privileged-containers -o yaml

# Extract violations from constraint status
kubectl get k8sblockprivileged block-privileged-containers \
  -o jsonpath='{.status.violations[*].message}' | jq

# Count total violations across all constraints
kubectl get constraints -o json | \
  jq '[.items[].status.totalViolations] | add'
```



## Example 5: example-5.sh


```bash
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
```



