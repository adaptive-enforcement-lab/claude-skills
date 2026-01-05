---
name: opa-rbac-templates - Examples
description: Code examples for OPA RBAC Templates
---

# OPA RBAC Templates - Examples


## Example 1: example-1.sh


```bash
kubectl apply -f cluster-admin.yaml
```



## Example 2: example-2.sh


```bash
kubectl apply -f privileged-verbs.yaml
```



## Example 3: example-3.sh


```bash
kubectl apply -f wildcards.yaml
```



## Example 4: example-4.sh


```bash
kubectl apply -f overview.yaml
```



## Example 5: example-5.yaml


```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: temporary-debug-access
  annotations:
    rbac.expires: "2025-01-05T00:00:00Z"
subjects:
  - kind: User
    name: engineer@company.com
roleRef:
  kind: ClusterRole
  name: debug-read-only
```



## Example 6: example-6.yaml


```yaml
# Enforced by: cluster-admin.yaml
# Result: Only subjects in approved list can receive cluster-admin binding
# Impact: Prevents privilege escalation to cluster-admin
```



## Example 7: example-7.yaml


```yaml
# Enforced by: privileged-verbs.yaml
# Result: Roles cannot include escalate/bind/impersonate verbs
# Impact: Prevents users from granting themselves additional permissions
```



## Example 8: example-8.yaml


```yaml
# Enforced by: wildcards.yaml
# Result: Roles must specify resources: ["pods"], not resources: ["*"]
# Impact: Reduces blast radius of compromised service accounts
```



## Example 9: example-9.sh


```bash
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
```



## Example 10: example-10.sh


```bash
# List all ClusterRoleBindings to cluster-admin
kubectl get clusterrolebindings -o json | \
  jq '.items[] | select(.roleRef.name == "cluster-admin") | .metadata.name'

# List roles with wildcard permissions
kubectl get roles,clusterroles --all-namespaces -o json | \
  jq '.items[] | select(.rules[].resources[] == "*") | .metadata.name'
```



## Example 11: example-11.sh


```bash
# List service accounts with cluster-admin
kubectl get clusterrolebindings -o json | \
  jq '.items[] | select(.roleRef.name == "cluster-admin" and .subjects[].kind == "ServiceAccount")'

# Find service accounts with escalate/bind/impersonate verbs
kubectl get roles,clusterroles --all-namespaces -o json | \
  jq '.items[] | select(.rules[].verbs[] | IN("escalate", "bind", "impersonate"))'
```



## Example 12: example-12.sh


```bash
# List all human users with cluster-level access
kubectl get clusterrolebindings -o json | \
  jq '.items[] | select(.subjects[].kind == "User") | {binding: .metadata.name, user: .subjects[].name, role: .roleRef.name}'
```



## Example 13: example-13.rego


```rego
# Pseudo-code: Full implementation in cluster-admin.yaml
approved_admins := {
  "break-glass-admin@company.com",
  "oncall-sre@company.com",
  "system:masters",  # For kubeadm bootstrap
}

deny[msg] {
  input.kind == "ClusterRoleBinding"
  input.roleRef.name == "cluster-admin"
  not approved_admins[input.subjects[_].name]
  msg := "cluster-admin can only be granted to approved break-glass accounts"
}
```



## Example 14: example-14.sh


```bash
# Generate temporary kubeconfig with cluster-admin (expires in 1 hour)
kubectl create token break-glass-admin --duration=1h

# Use temporary token for emergency operations
kubectl --token=$(kubectl create token break-glass-admin --duration=1h) get nodes
```



