# Test cluster-admin prevention (should fail for unapproved subject)
kubectl create clusterrolebinding test-admin \
  --clusterrole=cluster-admin \
  --user=attacker@example.com
# Expected: Admission denied by cluster-admin.yaml

# Test privileged verb block (should fail with escalate verb)
cat <<EOF | kubectl apply -f -
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: escalate-test
rules:
  - apiGroups: ["rbac.authorization.k8s.io"]
    resources: ["clusterroles"]
    verbs: ["escalate"]
EOF
# Expected: Admission denied by privileged-verbs.yaml

# Test wildcard prevention (should fail with resources: ["*"])
cat <<EOF | kubectl apply -f -
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: wildcard-test
  namespace: default
rules:
  - apiGroups: ["*"]
    resources: ["*"]
    verbs: ["get", "list"]
EOF
# Expected: Admission denied by wildcards.yaml

# Test compliant role (should succeed)
cat <<EOF | kubectl apply -f -
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: compliant-role
  namespace: default
rules:
  - apiGroups: [""]
    resources: ["pods", "configmaps"]
    verbs: ["get", "list", "watch"]
EOF
# Expected: Admission allowed by all policies