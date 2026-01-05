# Test privileged container block (should fail)
kubectl run privileged-test --image=nginx --privileged=true
# Expected: Admission denied by overview.yaml

# Test capability violation (should fail with CAP_SYS_ADMIN)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: cap-test
spec:
  containers:
    - name: nginx
      image: nginx
      securityContext:
        capabilities:
          add:
            - SYS_ADMIN
EOF
# Expected: Admission denied by capabilities.yaml

# Test root execution (should fail with runAsUser: 0)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: root-test
spec:
  containers:
    - name: nginx
      image: nginx
      securityContext:
        runAsUser: 0
EOF
# Expected: Admission denied by contexts.yaml

# Test privilege escalation (should fail with allowPrivilegeEscalation: true)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: escalation-test
spec:
  containers:
    - name: nginx
      image: nginx
      securityContext:
        allowPrivilegeEscalation: true
EOF
# Expected: Admission denied by escalation.yaml

# Test compliant pod (should succeed)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: compliant-test
spec:
  securityContext:
    runAsNonRoot: true
    runAsUser: 1000
    seccompProfile:
      type: RuntimeDefault
  containers:
    - name: nginx
      image: nginx
      securityContext:
        allowPrivilegeEscalation: false
        readOnlyRootFilesystem: true
        capabilities:
          drop:
            - ALL
EOF
# Expected: Admission allowed by all policies