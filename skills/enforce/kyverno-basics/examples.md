---
name: kyverno-basics - Examples
description: Code examples for Kyverno Basics
---

# Kyverno Basics - Examples


## Example 1: example-1.sh


```bash
# Add Kyverno Helm repository
helm repo add kyverno https://kyverno.github.io/kyverno/
helm repo update

# Install Kyverno
helm install kyverno kyverno/kyverno \
  --namespace kyverno \
  --create-namespace \
  --set replicaCount=3

# Verify installation
kubectl get pods -n kyverno
```



## Example 2: example-2.yaml


```yaml
apiVersion: kyverno.io/v1
kind: ClusterPolicy
metadata:
  name: require-resource-limits
spec:
  validationFailureAction: Enforce
  background: true
  rules:
    - name: check-resource-limits
      match:
        any:
          - resources:
              kinds:
                - Deployment
      validate:
        message: "Resource limits are required for all containers"
        pattern:
          spec:
            template:
              spec:
                containers:
                  - resources:
                      limits:
                        memory: "?*"
                        cpu: "?*"
```



## Example 3: example-3.sh


```bash
$ kubectl apply -f deployment.yaml
Error from server: admission webhook "validate.kyverno.svc-fail" denied the request:

policy Deployment/default/api for resource violation:

require-resource-limits:
  check-resource-limits: validation error: Resource limits are required for all containers
```



## Example 4: example-4.yaml


```yaml
spec:
  validationFailureAction: Audit  # Log violations, don't block
```



## Example 5: example-5.sh


```bash
kubectl get policyreport -A

NAMESPACE   NAME                          PASS   FAIL   WARN   ERROR   SKIP
default     polr-ns-default              12     3      0      0       0
production  polr-ns-production           45     1      0      0       0
```



## Example 6: example-6.yaml


```yaml
spec:
  validationFailureAction: Enforce  # Block violations
```



## Example 7: example-7.yaml


```yaml
apiVersion: kyverno.io/v1
kind: ClusterPolicy  # or Policy for namespaced
metadata:
  name: policy-name
spec:
  validationFailureAction: Enforce | Audit
  background: true | false  # Apply to existing resources
  rules:
    - name: rule-name
      match:  # What resources to check
        any:
          - resources:
              kinds: [Deployment, StatefulSet]
              namespaces: [production, staging]
      exclude:  # What to skip
        any:
          - resources:
              namespaces: [kube-system]
      validate | mutate | generate:  # What to do
        # Policy logic here
```



