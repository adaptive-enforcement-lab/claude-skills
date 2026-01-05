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