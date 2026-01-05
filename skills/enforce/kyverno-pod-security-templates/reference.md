---
name: kyverno-pod-security-templates - Reference
description: Complete reference for Kyverno Pod Security Templates
---

# Kyverno Pod Security Templates - Reference

This is the complete reference documentation extracted from the source.

# Kyverno Pod Security Templates

Pod security policies prevent privilege escalation, restrict dangerous capabilities, and enforce security boundaries for containerized workloads.

> **Pod Security Standards Replace PSP**
>
> PodSecurityPolicy was deprecated in Kubernetes 1.21 and removed in 1.25. Use Pod Security Standards (PSS) via admission controllers or Kyverno policies instead.
>

## Why Pod Security Matters

Containers inherit privileges from their configuration. Without enforcement, workloads can:

- Run as root with unrestricted filesystem access
- Mount host paths to access node secrets or modify system files
- Escalate privileges using dangerous Linux capabilities
- Break out of container isolation through privileged mode

## Available Templates

### [Pod Security Standards](standards.md)

Enforce Kubernetes Pod Security Standards (Baseline, Restricted):

- Block privileged containers and hostPath volumes
- Require non-root execution and read-only root filesystems
- Enforce seccomp, AppArmor, and SELinux profiles

**Apply a policy:**

```bash
kubectl apply -f standards.yaml

```

### [Privilege Restrictions](privileges.md)

Prevent privilege escalation and dangerous execution modes:

- Block `privileged: true` containers
- Prevent `allowPrivilegeEscalation: true`
- Restrict host namespaces (PID, IPC, Network)
- Block host port bindings

**Apply a policy:**

```bash
kubectl apply -f privileges.yaml

```

### [Security Profiles](profiles.md)

Enforce security profiles and runtime restrictions:

- Require seccomp profiles (RuntimeDefault or custom)
- Mandate AppArmor annotations for workloads
- Enforce SELinux contexts for pod isolation
- Block containers running as UID 0 (root)

**Apply a policy:**

```bash
kubectl apply -f profiles.yaml

```

## Pod Security Standards Levels

Kubernetes defines three PSS levels. Choose based on risk tolerance.

### Privileged (Unrestricted)

No restrictions. Only use for trusted system components.

- **Use cases:** CNI plugins, storage drivers, monitoring agents with node access
- **Risk:** Full cluster compromise if container is exploited
- **Recommendation:** Avoid. Use Restricted where possible.

### Baseline (Minimize Known Privilege Escalations)

Prevents known privilege escalation vectors:

- No privileged containers
- No host namespace sharing
- No host path mounts
- Limited capabilities (drops `ALL`, allows safe subset)

**Use for:** Most production workloads without special requirements.

### Restricted (Hardened for High-Security Environments)

Enforces current security best practices:

- Non-root execution (`runAsNonRoot: true`)
- Read-only root filesystem
- Seccomp profile required
- Drops all capabilities
- No privilege escalation

**Use for:** Internet-facing services, multi-tenant clusters, compliance requirements.

## Common Enforcement Scenarios

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

## Testing Pod Security Policies

Validate enforcement without disrupting production:

```bash
# Test privileged container block (should fail)
kubectl run privileged-test --image=nginx --privileged=true
# Expected: Blocked by privilege restriction policy

# Test root user block (should fail)
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
# Expected: Blocked by non-root requirement policy

# Test hostPath mount block (should fail)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: hostpath-test
spec:
  containers:
    - name: nginx
      image: nginx
      volumeMounts:
        - name: host
          mountPath: /host
  volumes:
    - name: host
      hostPath:
        path: /
EOF
# Expected: Blocked by Pod Security Standards policy

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
# Expected: Allowed by all policies

```

## Migration from PodSecurityPolicy

Replace deprecated PSPs with Kyverno policies:

1. **Audit current PSP usage:**

   ```bash
   kubectl get psp
   kubectl get pods --all-namespaces -o jsonpath='{range .items[*]}{.metadata.name}{"\t"}{.metadata.annotations.kubernetes\.io/psp}{"\n"}{end}'

   ```

2. **Map PSP rules to Kyverno policies:**
   - `privileged: false` → Use `privileges.yaml`
   - `allowPrivilegeEscalation: false` → Use `privileges.yaml`
   - `runAsUser` rules → Use `profiles.yaml`
   - `volumes` restrictions → Use `standards.yaml`

3. **Deploy Kyverno policies in audit mode:**

   ```bash
   kubectl apply -f standards.yaml
   kubectl apply -f privileges.yaml
   kubectl apply -f profiles.yaml

   ```

4. **Review policy reports for violations:**

   ```bash
   kubectl get polr -A  # Policy Reports
   kubectl describe polr <report-name> -n <namespace>

   ```

5. **Switch to enforce mode after validation:**
   Update `validationFailureAction: Enforce` in policies.

## Related Resources

- [Kyverno Templates Overview](../index.md)
- [Kyverno Network Security](../network/index.md)
- [OPA Pod Security Templates](../pod-security/index.md)

