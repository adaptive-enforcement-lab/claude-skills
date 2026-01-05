---
name: opa-rbac-templates
description: >-
  OPA RBAC policies preventing cluster-admin privilege escalation, restricting privileged verbs, and blocking wildcard permissions.
---

# OPA RBAC Templates

## When to Use This Skill

RBAC policies control who can perform which actions on which resources. These templates prevent privilege escalation through overly permissive roles.

> **Wildcards Grant Unrestricted Access**
>
> RBAC rules with `resources: ["*"]` or `verbs: ["*"]` grant access to all current and future resources or actions. Avoid wildcards except for break-glass admin roles.


## When to Apply

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


## Implementation

See the full implementation guide in the [source documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/).


## Techniques


### RBAC Security Patterns

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

*See [reference.md](reference.md) for additional techniques and detailed examples.*


## Examples

See [examples.md](examples.md) for code examples.


## Full Reference

See [reference.md](reference.md) for complete documentation.


## Related Patterns

- OPA Templates Overview
- OPA Pod Security
- OPA Resource Governance

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
