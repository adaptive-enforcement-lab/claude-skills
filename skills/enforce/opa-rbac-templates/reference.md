---
name: opa-rbac-templates - Reference
description: Complete reference for OPA RBAC Templates
---

# OPA RBAC Templates - Reference

This is the complete reference documentation extracted from the source.

# OPA RBAC Templates

RBAC policies control who can perform which actions on which resources. These templates prevent privilege escalation through overly permissive roles.

> **Wildcards Grant Unrestricted Access**
>
> RBAC rules with `resources: ["*"]` or `verbs: ["*"]` grant access to all current and future resources or actions. Avoid wildcards except for break-glass admin roles.
>

## Why RBAC Governance Matters

Kubernetes RBAC is additive (permissions accumulate across bindings). Without enforcement:

- **Privilege Escalation** - Users create ClusterRoleBindings granting themselves cluster-admin
- **Wildcard Abuse** - Roles grant `*` permissions instead of least privilege
- **Dangerous Verbs** - `escalate`, `impersonate`, `bind` allow RBAC manipulation
- **Permanent Admin Access** - ClusterRole `cluster-admin` granted to service accounts

## Available Templates

### [Cluster-Admin Prevention](cluster-admin.md)

Block cluster-admin role bindings except for approved subjects:

- Prevent creation of ClusterRoleBindings to `cluster-admin`
- Allow only break-glass admin accounts or system components
- Validate subject identity before granting cluster-admin
- Audit cluster-admin grants for compliance

**Apply a policy:**

```bash
kubectl apply -f cluster-admin.yaml
```

### [Privileged Verb Restrictions](privileged-verbs.md)

Block dangerous RBAC verbs that enable privilege escalation:

- Prevent `escalate` verb (bypass RBAC validation)
- Block `bind` verb (assign ClusterRoles to arbitrary subjects)
- Restrict `impersonate` verb (act as other users/service accounts)
- Limit `*` verb grants to approved roles

**Apply a policy:**

```bash
kubectl apply -f privileged-verbs.yaml
```

### [Wildcard Prevention](wildcards.md)

Restrict wildcard usage in RBAC rules:

- Block `resources: ["*"]` in Role/ClusterRole rules
- Prevent `verbs: ["*"]` except for read-only access
- Require explicit resource and verb lists
- Allow wildcards only for monitoring/observability roles

**Apply a policy:**

```bash
kubectl apply -f wildcards.yaml
```

### [RBAC Policy Overview](overview.md)

General RBAC governance and least privilege principles:

- Namespace-scoped roles preferred over ClusterRoles
- Service account permissions limited to pod requirements
- Time-bounded RoleBindings with expiration annotations
- Regular RBAC audits and privilege reviews

**Apply a policy:**

```bash
kubectl apply -f overview.yaml
```

## RBAC Security Patterns

### Least Privilege Principle

Grant minimum permissions required for each workload:

1. **Start with no permissions** - Service accounts have no default permissions
2. **Add specific resources** - `pods`, `configmaps`, not `*`
3. **Add specific verbs** - `get`, `list`, not `*`
4. **Scope to namespace** - Use Role instead of ClusterRole when possible

### Defense Against Privilege Escalation

Block RBAC manipulation verbs:

- **`escalate`** - Allows creating roles with more permissions than creator has
- **`bind`** - Allows granting roles to arbitrary subjects
- **`impersonate`** - Allows acting as other users without authentication

Only cluster admins should have these verbs.

### Time-Bounded Permissions

Use annotations to enforce temporary access:

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

OPA policies can validate expiration and block expired bindings.

## Common Enforcement Scenarios

### Scenario 1: Prevent Unauthorized Cluster-Admin

Block cluster-admin except for approved break-glass accounts:

```yaml
# Enforced by: cluster-admin.yaml
# Result: Only subjects in approved list can receive cluster-admin binding
# Impact: Prevents privilege escalation to cluster-admin
```

### Scenario 2: Block Dangerous RBAC Verbs

Prevent use of `escalate`, `bind`, `impersonate`:

```yaml
# Enforced by: privileged-verbs.yaml
# Result: Roles cannot include escalate/bind/impersonate verbs
# Impact: Prevents users from granting themselves additional permissions
```

### Scenario 3: Eliminate Wildcard Permissions

Require explicit resource and verb lists:

```yaml
# Enforced by: wildcards.yaml
# Result: Roles must specify resources: ["pods"], not resources: ["*"]
# Impact: Reduces blast radius of compromised service accounts
```

## Testing RBAC Policies

Validate enforcement without disrupting operations:

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

## RBAC Audit and Review

Regularly audit RBAC configuration for compliance:

### Identify Privileged Bindings

```bash
# List all ClusterRoleBindings to cluster-admin
kubectl get clusterrolebindings -o json | \
  jq '.items[] | select(.roleRef.name == "cluster-admin") | .metadata.name'

# List roles with wildcard permissions
kubectl get roles,clusterroles --all-namespaces -o json | \
  jq '.items[] | select(.rules[].resources[] == "*") | .metadata.name'
```

### Find Service Accounts with Excessive Permissions

```bash
# List service accounts with cluster-admin
kubectl get clusterrolebindings -o json | \
  jq '.items[] | select(.roleRef.name == "cluster-admin" and .subjects[].kind == "ServiceAccount")'

# Find service accounts with escalate/bind/impersonate verbs
kubectl get roles,clusterroles --all-namespaces -o json | \
  jq '.items[] | select(.rules[].verbs[] | IN("escalate", "bind", "impersonate"))'
```

### Validate RoleBinding Subjects

```bash
# List all human users with cluster-level access
kubectl get clusterrolebindings -o json | \
  jq '.items[] | select(.subjects[].kind == "User") | {binding: .metadata.name, user: .subjects[].name, role: .roleRef.name}'
```

## Break-Glass Admin Access

Maintain emergency access while enforcing policies:

### Approved Admin Accounts

Define break-glass accounts in OPA policy:

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

### Temporary Elevation

Use short-lived credentials instead of permanent cluster-admin:

```bash
# Generate temporary kubeconfig with cluster-admin (expires in 1 hour)
kubectl create token break-glass-admin --duration=1h

# Use temporary token for emergency operations
kubectl --token=$(kubectl create token break-glass-admin --duration=1h) get nodes
```

## Related Resources

- [OPA Templates Overview](../index.md)
- [OPA Pod Security](../pod-security/index.md)
- [OPA Resource Governance](../resource/index.md)

