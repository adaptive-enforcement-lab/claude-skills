---
name: policy-as-code-operations - Examples
description: Code examples for Policy-as-Code Operations
---

# Policy-as-Code Operations - Examples


## Example 1: example-1.yaml


```yaml
# security-policy/charts/security-policy/templates/require-network-policy.yaml
apiVersion: kyverno.io/v1
kind: ClusterPolicy
metadata:
  name: require-network-policy
  annotations:
    policies.kyverno.io/title: Require NetworkPolicy
    policies.kyverno.io/category: Security
    policies.kyverno.io/severity: high
spec:
  validationFailureAction: {{ .Values.policies.networkPolicy.validationFailureAction }}
  background: true
  rules:
    - name: check-network-policy-exists
      match:
        resources:
          kinds:
            - Namespace
      validate:
        message: "Namespace must have a NetworkPolicy"
        deny:
          conditions:
            - key: "{{ request.object.metadata.name }}"
              operator: AnyNotIn
              value: [" kube-system", "kube-public", "kube-node-lease"]
```



## Example 2: example-2.yaml


```yaml
# security-policy/charts/security-policy/values.yaml
policies:
  networkPolicy:
    enabled: true
    validationFailureAction: Audit  # Start with Audit
```



## Example 3: example-3.sh


```bash
docker run --rm -v $(pwd):/workspace policy-platform:latest bash -c '\
  helm template security /repos/security-policy/charts/security-policy \
    -f /repos/security-policy/charts/security-policy/values.yaml \
  > /tmp/policies.yaml &&\
  kyverno apply /tmp/policies.yaml --resource /workspace/test-namespace.yaml\
'
```



## Example 4: example-4.sh


```bash
# Update policy-platform container (rebuild with new policy)
docker build -t policy-platform:v1.0.3 -f ci/Dockerfile .
docker push policy-platform:v1.0.3

# Deploy to dev cluster
helm upgrade security-policy /repos/security-policy/charts/security-policy \
  --namespace kyverno \
  --values /repos/security-policy/cd/dev/values.yaml
```



## Example 5: example-5.sh


```bash
kubectl get policyreport -A

# Check for violations
kubectl get policyreport polr-ns-default -o yaml
```



## Example 6: example-6.yaml


```yaml
# security-policy/cd/prd/values.yaml
policies:
  networkPolicy:
    validationFailureAction: Enforce  # Now block violations
```



## Example 7: example-7.yaml


```yaml
validate:
  pattern:
    spec:
      containers:
        - resources:
            limits:
              memory: "?*"
              cpu: "?*"
            requests:              # Requires both limits AND requests
              memory: "?*"
              cpu: "?*"
```



## Example 8: example-8.yaml


```yaml
validate:
  pattern:
    spec:
      containers:
        - resources:
            limits:
              memory: "?*"
              cpu: "?*"
          # Requests optional
```



