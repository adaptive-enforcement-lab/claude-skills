---
name: opa-pod-security-templates
description: >-
  OPA pod security policies preventing privileged containers, restricting Linux capabilities, and enforcing security contexts in Kubernetes.
---

# OPA Pod Security Templates

## When to Use This Skill

Pod security policies written in Rego prevent privilege escalation and enforce security boundaries for containerized workloads.

> **Capabilities Bypass Security Boundaries**
>
> Linux capabilities grant fine-grained privileges. A container with `CAP_SYS_ADMIN` can bypass most kernel security mechanisms. Drop all capabilities by default.


## When to Apply

### Scenario 1: Block Privileged Containers

Prevent unrestricted container execution:

```yaml
# Enforced by: overview.yaml
# Result: No containers can run with privileged: true
# Impact: Eliminates most container breakout vectors

```

### Scenario 2: Drop Dangerous Capabilities

Remove capabilities that grant excessive privileges:

```yaml
# Enforced by: capabilities.yaml
# Result: All containers must drop CAP_SYS_ADMIN, CAP_NET_RAW
# Impact: Prevents kernel manipulation and network sniffing

```

### Scenario 3: Enforce Non-Root Execution

Require all containers to run as non-root users:

```yaml
# Enforced by: contexts.yaml
# Result: Containers must define runAsNonRoot: true and runAsUser > 0
# Impact: Prevents root-level filesystem access and privilege escalation

```

### Scenario 4: Block Privilege Escalation

Prevent containers from gaining privileges after start:

```yaml
# Enforced by: escalation.yaml
# Result: Containers must set allowPrivilegeEscalation: false
# Impact: Blocks setuid binaries and capability inheritance

```


## Implementation

Every pod should define security contexts at both pod and container levels:

### Pod-Level Security Context

```yaml
spec:
  securityContext:
    runAsNonRoot: true
    runAsUser: 1000
    runAsGroup: 1000
    fsGroup: 1000
    seccompProfile:
      type: RuntimeDefault

```

### Container-Level Security Context


*See [examples.md](examples.md) for detailed code examples.*


## Examples

See [examples.md](examples.md) for code examples.


## Full Reference

See [reference.md](reference.md) for complete documentation.


## Related Patterns

- OPA Templates Overview
- OPA RBAC Policies
- Kyverno Pod Security Templates

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
