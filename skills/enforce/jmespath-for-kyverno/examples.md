---
name: jmespath-for-kyverno - Examples
description: Code examples for JMESPath for Kyverno
---

# JMESPath for Kyverno - Examples


## Example 1: example-1.sh


```bash
# Install kyverno CLI
brew install kyverno/kyverno/kyverno

# Test JMESPath expression
kyverno jp query -i manifest.yaml 'spec.template.spec.containers[*].name'
```



## Example 2: example-2.yaml


```yaml
apiVersion: kyverno.io/v1
kind: ClusterPolicy
metadata:
  name: require-resource-limits
spec:
  validationFailureAction: Enforce
  rules:
  - name: validate-limits
    match:
      any:
      - resources:
          kinds:
          - Pod
    validate:
      message: "All containers must define resource limits"
      deny:
        conditions:
          any:
          - key: "{{ request.object.spec.containers[?!resources.limits.memory].name | length(@) }}"
            operator: GreaterThan
            value: 0
```



