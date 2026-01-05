---
name: runtime-deployment-admission-control-with-kyverno - Examples
description: Code examples for Runtime Deployment: Admission Control with Kyverno
---

# Runtime Deployment: Admission Control with Kyverno - Examples


## Example 1: example-1.mermaid


```mermaid
graph TD
    K[kubectl apply] --> API[Kubernetes API Server]
    API --> ADM[Admission Controller]
    ADM --> KYV[Kyverno Webhook]
    KYV --> POL{Policy Check}
    POL -->|Pass| ETCD[(etcd)]
    POL -->|Fail| REJECT[Reject Request]

    %% Ghostty Hardcore Theme
    style ETCD fill:#a7e22e,color:#1b1d1e
    style REJECT fill:#f92572,color:#1b1d1e
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
    - name: check-cpu-memory
      match:
        resources:
          kinds:
            - Deployment
      validate:
        message: "CPU and memory limits required"
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
helm repo add kyverno https://kyverno.github.io/kyverno/
helm repo update

helm install kyverno kyverno/kyverno \
  --namespace kyverno \
  --create-namespace \
  --values kyverno-values.yaml
```



## Example 4: example-4.yaml


```yaml
features:
  logging:
    logLevel: -2  # Info level

  backgroundScan:
    backgroundScanInterval: 6h

# Clean up old reports
policyReportsCleanup:
  enabled: true

cleanupJobs:
  admissionReports:
    enabled: true
    schedule: "0 0 * * *"  # Daily

  clusterAdmissionReports:
    enabled: true
    schedule: "0 0 * * SUN"  # Weekly

  policyReports:
    enabled: true
    schedule: "0 0 * * *"  # Daily

  clusterPolicyReports:
    enabled: true
    schedule: "0 0 1 * *"  # Monthly

  resources:
    limits:
      memory: 128Mi
    requests:
      cpu: 50m
      memory: 64Mi

# Exclude system namespaces
resourceFilters:
  resourceFiltersExcludeNamespaces:
    - kube-system
    - gmp-system
    - cnrm-system

# Logging configuration
admissionController:
  container:
    extraArgs:
      "loggingFormat": "json"
      "v": "1"

backgroundController:
  enabled: true
  rbac:
    clusterRole:
      extraResources:
        - apiGroups: ["apps"]
          resources: ["deployments", "statefulsets", "daemonsets"]
          verbs: ["get", "list", "watch"]
```



## Example 5: example-5.sh


```bash
helm repo add policy-reporter https://kyverno.github.io/policy-reporter
helm repo update

helm install policy-reporter policy-reporter/policy-reporter \
  --namespace policy-reporter \
  --create-namespace \
  --values policy-reporter-values.yaml
```



## Example 6: example-6.yaml


```yaml
metrics:
  enabled: true

logging:
  encoding: json
  logLevel: -2
  development: false

api:
  logging: false

ui:
  enabled: true
  displayMode: dark

kyvernoPlugin:
  enabled: true
  metrics:
    enabled: true
```



## Example 7: example-7.sh


```bash
# Check Kyverno pods
kubectl get pods -n kyverno

# Expected output:
# kyverno-admission-controller-xxx   Running
# kyverno-background-controller-xxx  Running
# kyverno-cleanup-controller-xxx     Running
# kyverno-reports-controller-xxx     Running
```



## Example 8: example-8.sh


```bash
# Check ValidatingWebhookConfiguration
kubectl get validatingwebhookconfiguration | grep kyverno

# Check MutatingWebhookConfiguration
kubectl get mutatingwebhookconfiguration | grep kyverno
```



## Example 9: example-9.sh


```bash
# Try deploying without resource limits
kubectl run test --image=nginx

# Expected: Denied by admission webhook
```



