---
name: opa-resource-governance-templates - Examples
description: Code examples for OPA Resource Governance Templates
---

# OPA Resource Governance Templates - Examples


## Example 1: example-1.sh


```bash
kubectl apply -f governance.yaml
```



## Example 2: example-2.sh


```bash
kubectl apply -f limitrange.yaml
```



## Example 3: example-3.sh


```bash
kubectl apply -f storage.yaml
```



## Example 4: example-4.yaml


```yaml
# Development namespace quota
apiVersion: v1
kind: ResourceQuota
metadata:
  name: dev-quota
  namespace: dev-team
spec:
  hard:
    requests.cpu: "10"
    requests.memory: "20Gi"
    limits.cpu: "20"
    limits.memory: "40Gi"
    persistentvolumeclaims: "10"
    requests.storage: "100Gi"
```



## Example 5: example-5.yaml


```yaml
# Enforced by: governance.yaml
# Result: All containers must define resources.limits.cpu and resources.limits.memory
# Impact: Prevents single pod from consuming entire node capacity
```



## Example 6: example-6.yaml


```yaml
# Enforced by: limitrange.yaml
# Result: Pods cannot request more CPU/memory than LimitRange allows
# Impact: Ensures fair resource distribution across namespace
```



## Example 7: example-7.yaml


```yaml
# Enforced by: storage.yaml
# Result: PVCs in dev-* namespaces cannot exceed 50Gi
# Impact: Prevents accidental provisioning of expensive storage volumes
```



## Example 8: example-8.sh


```bash
# Test resource limit requirement (should fail without limits)
kubectl run no-limits --image=nginx
# Expected: Admission denied by governance.yaml

# Test excessive resource request (should fail above quota)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: excessive-request
  namespace: dev-team
spec:
  containers:
    - name: nginx
      image: nginx
      resources:
        requests:
          cpu: "50"  # Exceeds namespace quota of 10 CPU
          memory: "100Gi"
EOF
# Expected: Admission denied by governance.yaml (quota violation)

# Test LimitRange violation (should fail above max)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: limitrange-violation
  namespace: dev-team
spec:
  containers:
    - name: nginx
      image: nginx
      resources:
        requests:
          cpu: "100"  # Exceeds LimitRange max of 2 CPU
          memory: "200Gi"
EOF
# Expected: Admission denied by limitrange.yaml

# Test storage size restriction (should fail for excessive PVC)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: large-pvc
  namespace: dev-team
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 500Gi  # Exceeds policy max of 50Gi for dev namespaces
EOF
# Expected: Admission denied by storage.yaml

# Test compliant workload (should succeed)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: compliant-pod
  namespace: dev-team
spec:
  containers:
    - name: nginx
      image: nginx
      resources:
        requests:
          cpu: "100m"
          memory: "128Mi"
        limits:
          cpu: "500m"
          memory: "512Mi"
EOF
# Expected: Admission allowed by all policies
```



## Example 9: example-9.yaml


```yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: production-quota
  namespace: production
spec:
  hard:
    requests.cpu: "100"
    requests.memory: "200Gi"
    limits.cpu: "200"
    limits.memory: "400Gi"
    persistentvolumeclaims: "50"
    requests.storage: "1Ti"
```



## Example 10: example-10.yaml


```yaml
apiVersion: v1
kind: LimitRange
metadata:
  name: production-limits
  namespace: production
spec:
  limits:
    - max:
        cpu: "4"
        memory: "8Gi"
      min:
        cpu: "10m"
        memory: "64Mi"
      default:
        cpu: "500m"
        memory: "512Mi"
      defaultRequest:
        cpu: "100m"
        memory: "128Mi"
      type: Container
```



## Example 11: example-11.sh


```bash
kubectl get pods --all-namespaces -o json | \
  jq '[.items[].spec.containers[].resources.requests] | add'
```



## Example 12: example-12.sh


```bash
kubectl get nodes -o json | \
  jq '[.items[].status.allocatable] | add'
```



## Example 13: example-13.sh


```bash
# If total requests > allocatable capacity, cluster is overcommitted
# Add nodes or reduce resource requests
```



