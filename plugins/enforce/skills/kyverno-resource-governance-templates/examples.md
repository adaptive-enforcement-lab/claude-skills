---
name: kyverno-resource-governance-templates - Examples
description: Code examples for Kyverno Resource Governance Templates
---

# Kyverno Resource Governance Templates - Examples


## Example 1: example-1.sh


```bash
kubectl apply -f limits.yaml
```



## Example 2: example-2.sh


```bash
kubectl apply -f hpa.yaml
```



## Example 3: example-3.sh


```bash
kubectl apply -f storage.yaml
```



## Example 4: example-4.yaml


```yaml
# Enforced by: limits.yaml
# Result: All containers must define resources.limits.cpu and resources.limits.memory
# Impact: Prevents single pod from consuming entire node capacity
```



## Example 5: example-5.yaml


```yaml
# Enforced by: hpa.yaml
# Result: Deployments in prod-* namespaces must have corresponding HPA
# Impact: Ensures production services scale automatically under load
```



## Example 6: example-6.yaml


```yaml
# Enforced by: storage.yaml
# Result: PVCs cannot exceed 100Gi in dev namespaces
# Impact: Prevents accidental provisioning of expensive storage volumes
```



## Example 7: example-7.sh


```bash
# Test resource limit requirement (should fail without limits)
kubectl run no-limits --image=nginx
# Expected: Blocked by policy requiring resource limits

# Test excessive resource request (should fail if beyond policy limits)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: excessive-request
spec:
  containers:
    - name: nginx
      image: nginx
      resources:
        requests:
          cpu: "100"
          memory: "1000Gi"
EOF
# Expected: Blocked by policy restricting maximum requests

# Test HPA requirement (should fail without HPA)
kubectl create deployment test-app --image=nginx --replicas=3 -n production
# Expected: Blocked by policy requiring HPA for production Deployments

# Test storage size restriction (should fail for excessive PVC)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: large-pvc
  namespace: dev
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 500Gi
EOF
# Expected: Blocked by policy restricting dev namespace PVC sizes

# Test compliant workload (should succeed)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: compliant-pod
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
# Expected: Allowed by resource limit policies
```



## Example 8: example-8.sh


```bash
kubectl get resourcequota -n production -o yaml
# Review quota limits vs current usage
```



## Example 9: example-9.sh


```bash
kubectl top pods --all-namespaces --sort-by=memory
kubectl top pods --all-namespaces --sort-by=cpu
```



## Example 10: example-10.sh


```bash
kubectl get polr -A  # Policy Reports
# Review workloads blocked by resource policies
```



