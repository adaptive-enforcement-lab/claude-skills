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