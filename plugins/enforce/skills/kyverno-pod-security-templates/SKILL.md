---
name: kyverno-pod-security-templates
description: >-
  Kyverno pod security policies enforcing Pod Security Standards, privilege restrictions, and security profiles for Kubernetes workloads.
---

# Kyverno Pod Security Templates

## When to Use This Skill

Pod security policies prevent privilege escalation, restrict dangerous capabilities, and enforce security boundaries for containerized workloads.

> **Pod Security Standards Replace PSP**
>
> PodSecurityPolicy was deprecated in Kubernetes 1.21 and removed in 1.25. Use Pod Security Standards (PSS) via admission controllers or Kyverno policies instead.


## When to Apply

### Scenario 1: Block All Privileged Containers

Prevent privileged mode across the cluster:

```yaml
# Enforced by: privileges.yaml
# Result: No containers can run with privileged: true
# Impact: Eliminates most container breakout vectors

```

### Scenario 2: Require Non-Root Execution

Force all containers to run as non-root users:

```yaml
# Enforced by: profiles.yaml
# Result: Containers must define runAsNonRoot: true
# Impact: Prevents root-level filesystem access and privilege escalation

```

### Scenario 3: Enforce Seccomp Profiles

Mandate seccomp profiles for syscall filtering:

```yaml
# Enforced by: standards.yaml
# Result: Pods must define securityContext.seccompProfile
# Impact: Reduces kernel attack surface by blocking dangerous syscalls

```


## Implementation

See the full implementation guide in the [source documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/).


## Examples

See [examples.md](examples.md) for code examples.


## Full Reference

See [reference.md](reference.md) for complete documentation.


## Related Patterns

- Kyverno Templates Overview
- Kyverno Network Security
- OPA Pod Security Templates

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
