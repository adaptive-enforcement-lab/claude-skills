# Test pod without limits is rejected
kubectl apply -f - <<EOF
apiVersion: v1
kind: Pod
metadata:
  name: test
spec:
  containers:
    - name: app
      image: nginx
EOF
# Expected: Admission webhook denies request

# Test untrusted registry is blocked
kubectl apply -f - <<EOF
apiVersion: v1
kind: Pod
metadata:
  name: test
spec:
  containers:
    - name: app
      image: docker.io/nginx:latest
EOF
# Expected: Image source validation fails

# Check policy reports
kubectl get policyreport -A
# Expected: Shows pass/fail summary