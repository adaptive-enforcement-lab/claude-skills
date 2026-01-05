---
name: opa-pod-security-templates - Reference
description: Complete reference for OPA Pod Security Templates
---

# OPA Pod Security Templates - Reference

This is the complete reference documentation extracted from the source.

# OPA Pod Security Templates

Pod security policies written in Rego prevent privilege escalation and enforce security boundaries for containerized workloads.

> **Capabilities Bypass Security Boundaries**
>
> Linux capabilities grant fine-grained privileges. A container with `CAP_SYS_ADMIN` can bypass most kernel security mechanisms. Drop all capabilities by default.
>

## Why Pod Security Matters

Container isolation relies on kernel namespaces, cgroups, and capabilities. Misconfigurations break this isolation:

- **Privileged Mode** - Disables all security boundaries
- **Host Namespaces** - Shares PID/IPC/Network with node
- **Dangerous Capabilities** - `CAP_SYS_ADMIN`, `CAP_NET_RAW`, `CAP_SYS_PTRACE`
- **Root Execution** - Unnecessary privileges for most workloads

## Available Templates

### [Privileged Container Prevention](overview.md)

Block privileged containers and host access:

- Prevent `privileged: true` in container security contexts
- Block hostPath, hostPID, hostIPC, hostNetwork usage
- Restrict host port bindings
- Prevent sharing node resources with containers

**Apply a policy:**

```bash
kubectl apply -f overview.yaml

```

### [Capability Restrictions](capabilities.md)

Control Linux capabilities granted to containers:

- Require dropping ALL capabilities by default
- Block dangerous capabilities (SYS_ADMIN, NET_RAW, SYS_PTRACE)
- Allow safe capabilities only (NET_BIND_SERVICE, CHOWN)
- Validate capability drops in security contexts

**Apply a policy:**

```bash
kubectl apply -f capabilities.yaml

```

### [Security Context Enforcement](contexts.md)

Mandate security context configuration:

- Require non-root user execution
- Enforce read-only root filesystems
- Require runAsNonRoot: true
- Block privilege escalation via allowPrivilegeEscalation

**Apply a policy:**

```bash
kubectl apply -f contexts.yaml

```

### [Privilege Escalation Prevention](escalation.md)

Block privilege escalation mechanisms:

- Prevent `allowPrivilegeEscalation: true`
- Block setuid/setgid binaries in containers
- Enforce seccomp profiles to restrict syscalls
- Validate AppArmor/SELinux profiles

**Apply a policy:**

```bash
kubectl apply -f escalation.yaml

```

## Pod Security Defense Layers

Implement overlapping controls for defense in depth:

1. **Privileged Prevention** - Block privileged mode and host access (overview.yaml)
2. **Capability Restrictions** - Drop dangerous Linux capabilities (capabilities.yaml)
3. **Non-Root Execution** - Require runAsNonRoot (contexts.yaml)
4. **Escalation Prevention** - Block privilege escalation paths (escalation.yaml)

Each layer addresses different attack vectors. Production workloads should pass all four.

## Common Enforcement Scenarios

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

## Testing Pod Security Policies

Validate enforcement without disrupting workloads:

```bash
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

```

## Understanding Linux Capabilities

Capabilities split root privileges into fine-grained permissions:

### Safe Capabilities (Usually Allowed)

- `NET_BIND_SERVICE` - Bind to ports < 1024
- `CHOWN` - Change file ownership
- `DAC_OVERRIDE` - Bypass file permission checks
- `SETGID` / `SETUID` - Change user/group IDs

### Dangerous Capabilities (Always Blocked)

- `SYS_ADMIN` - Virtually unlimited kernel access
- `NET_RAW` - Create raw sockets (packet sniffing)
- `SYS_PTRACE` - Debug arbitrary processes (credential theft)
- `SYS_MODULE` - Load kernel modules
- `CAP_SYS_BOOT` - Reboot system

**Best practice:** Drop ALL capabilities, then add only required safe capabilities.

## Security Context Configuration

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

```yaml
spec:
  containers:
    - name: app
      securityContext:
        allowPrivilegeEscalation: false
        readOnlyRootFilesystem: true
        capabilities:
          drop:
            - ALL
          add:
            - NET_BIND_SERVICE  # Only if binding to port 80/443

```

## Migration from Permissive to Restrictive Policies

Transition existing workloads to secure configurations:

1. **Audit current security contexts:**

   ```bash
   kubectl get pods --all-namespaces -o json | \
     jq '.items[] | select(.spec.containers[].securityContext.privileged == true)'

   ```

2. **Deploy policies in audit mode:**
   Use OPA audit mode to log violations without blocking:

   ```bash
   kubectl apply -f overview.yaml  # Set enforcementAction: warn

   ```

3. **Review violations:**

   ```bash
   kubectl get constrainttemplates
   kubectl get <constraint-name> -o yaml
   # Check status.violations for non-compliant pods

   ```

4. **Fix workload security contexts:**
   Update Deployments/StatefulSets to add security context fields.

5. **Enable enforcement:**
   Change `enforcementAction: deny` after validation period.

## Related Resources

- [OPA Templates Overview](../index.md)
- [OPA RBAC Policies](../rbac/index.md)
- [Kyverno Pod Security Templates](../pod-security/index.md)

