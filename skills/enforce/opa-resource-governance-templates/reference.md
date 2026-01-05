---
name: opa-resource-governance-templates - Reference
description: Complete reference for OPA Resource Governance Templates
---

# OPA Resource Governance Templates - Reference

This is the complete reference documentation extracted from the source.

# OPA Resource Governance Templates

Resource governance policies prevent overconsumption, enforce quotas, and validate LimitRange compliance across your cluster.

> **ResourceQuota vs LimitRange vs OPA**
>
> ResourceQuota caps total namespace consumption. LimitRange sets defaults and bounds for individual pods. OPA validates configuration before admission. Use all three for comprehensive governance.
>

## Why Resource Governance Matters

Kubernetes does not enforce resource limits by default. This creates operational and financial risks:

- **Node Exhaustion** - Unbounded pods consume all CPU/memory
- **Quota Violations** - Deployments exceed namespace ResourceQuota
- **Cost Overruns** - Excessive storage provisioning or oversized VMs
- **Cluster Instability** - OOMKilled pods cascade across nodes

## Available Templates

### [Resource Governance](governance.md)

Enforce resource limits, requests, and quota compliance:

- Require resource requests for all containers
- Mandate resource limits to prevent node exhaustion
- Validate requests ≤ limits for CPU and memory
- Block workloads exceeding namespace quotas

**Apply a policy:**

```bash
kubectl apply -f governance.yaml

```

### [LimitRange Validation](limitrange.md)

Enforce LimitRange compliance for pods and containers:

- Validate pod resource requests against LimitRange defaults
- Block pods exceeding LimitRange max values
- Require LimitRange in all non-system namespaces
- Enforce container-level and pod-level limits

**Apply a policy:**

```bash
kubectl apply -f limitrange.yaml

```

### [Storage Constraints](storage.md)

Control PersistentVolume and PersistentVolumeClaim allocation:

- Restrict PVC sizes based on namespace or StorageClass
- Require approved StorageClasses for production data
- Block dynamic provisioning in restricted namespaces
- Validate volume access modes and reclaim policies

**Apply a policy:**

```bash
kubectl apply -f storage.yaml

```

## Resource Governance Patterns

### Three-Layer Resource Control

Implement overlapping controls for comprehensive governance:

1. **OPA Policies** - Validate resource configuration at admission time
2. **LimitRange** - Set namespace-level defaults and max values
3. **ResourceQuota** - Cap total namespace consumption

### Right-Sizing Workloads

Balance cost and reliability with appropriate resource values:

| Resource | Too Low | Too High | Sweet Spot |
|----------|---------|----------|------------|
| **Requests** | OOMKilled pods | Wasted capacity | Actual usage (P95) |
| **Limits** | Throttling | Noisy neighbors | 2x requests |
| **Storage** | Out of space | High costs | Actual data + 30% |

Use Vertical Pod Autoscaler (VPA) to identify optimal values.

### Quota Enforcement Strategy

Define quotas based on team size and workload type:

```yaml
# Development namespace quota
apiVersion: v1
kind: ResourceQuota
metadata:
  name: dev-quota
  namespace: dev-team
spec:
  hard:
    requests.cpu: "10"
    requests.memory: "20Gi"
    limits.cpu: "20"
    limits.memory: "40Gi"
    persistentvolumeclaims: "10"
    requests.storage: "100Gi"

```

OPA policies validate workloads against these quotas before admission.

## Common Enforcement Scenarios

### Scenario 1: Prevent Unbounded Resource Consumption

Require resource limits on all containers:

```yaml
# Enforced by: governance.yaml
# Result: All containers must define resources.limits.cpu and resources.limits.memory
# Impact: Prevents single pod from consuming entire node capacity

```

### Scenario 2: Enforce LimitRange Compliance

Block pods exceeding namespace LimitRange maximums:

```yaml
# Enforced by: limitrange.yaml
# Result: Pods cannot request more CPU/memory than LimitRange allows
# Impact: Ensures fair resource distribution across namespace

```

### Scenario 3: Control Storage Costs

Restrict PVC sizes based on environment:

```yaml
# Enforced by: storage.yaml
# Result: PVCs in dev-* namespaces cannot exceed 50Gi
# Impact: Prevents accidental provisioning of expensive storage volumes

```

## Testing Resource Governance Policies

Validate enforcement without blocking legitimate workloads:

```bash
# Test resource limit requirement (should fail without limits)
kubectl run no-limits --image=nginx
# Expected: Admission denied by governance.yaml

# Test excessive resource request (should fail above quota)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: excessive-request
  namespace: dev-team
spec:
  containers:
    - name: nginx
      image: nginx
      resources:
        requests:
          cpu: "50"  # Exceeds namespace quota of 10 CPU
          memory: "100Gi"
EOF
# Expected: Admission denied by governance.yaml (quota violation)

# Test LimitRange violation (should fail above max)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: limitrange-violation
  namespace: dev-team
spec:
  containers:
    - name: nginx
      image: nginx
      resources:
        requests:
          cpu: "100"  # Exceeds LimitRange max of 2 CPU
          memory: "200Gi"
EOF
# Expected: Admission denied by limitrange.yaml

# Test storage size restriction (should fail for excessive PVC)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: large-pvc
  namespace: dev-team
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 500Gi  # Exceeds policy max of 50Gi for dev namespaces
EOF
# Expected: Admission denied by storage.yaml

# Test compliant workload (should succeed)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: compliant-pod
  namespace: dev-team
spec:
  containers:
    - name: nginx
      image: nginx
      resources:
        requests:
          cpu: "100m"
          memory: "128Mi"
        limits:
          cpu: "500m"
          memory: "512Mi"
EOF
# Expected: Admission allowed by all policies

```

## ResourceQuota and LimitRange Integration

OPA policies complement native Kubernetes resource controls:

### ResourceQuota Example

```yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: production-quota
  namespace: production
spec:
  hard:
    requests.cpu: "100"
    requests.memory: "200Gi"
    limits.cpu: "200"
    limits.memory: "400Gi"
    persistentvolumeclaims: "50"
    requests.storage: "1Ti"

```

### LimitRange Example

```yaml
apiVersion: v1
kind: LimitRange
metadata:
  name: production-limits
  namespace: production
spec:
  limits:
    - max:
        cpu: "4"
        memory: "8Gi"
      min:
        cpu: "10m"
        memory: "64Mi"
      default:
        cpu: "500m"
        memory: "512Mi"
      defaultRequest:
        cpu: "100m"
        memory: "128Mi"
      type: Container

```

### OPA Validation

OPA policies validate that:

- Pod resource requests/limits respect LimitRange bounds
- Total namespace consumption stays within ResourceQuota
- Requests ≤ Limits for all containers
- Required fields are present (not relying on LimitRange defaults)

## Capacity Planning with Resource Policies

Use policy enforcement data for cluster sizing:

1. **Collect resource requests across namespaces:**

   ```bash
   kubectl get pods --all-namespaces -o json | \
     jq '[.items[].spec.containers[].resources.requests] | add'

   ```

2. **Compare against node capacity:**

   ```bash
   kubectl get nodes -o json | \
     jq '[.items[].status.allocatable] | add'

   ```

3. **Identify overcommitment:**

   ```bash
   # If total requests > allocatable capacity, cluster is overcommitted
   # Add nodes or reduce resource requests

   ```

4. **Adjust policies based on usage:**
   - Increase LimitRange maximums if legitimate workloads are blocked
   - Decrease quotas if namespace consumption is consistently low
   - Update OPA policies to reflect new capacity constraints

## Cost Optimization with Resource Governance

Reduce cloud costs through policy enforcement:

### CPU/Memory Optimization

- **Block oversized requests** - Prevent requesting 16 CPU for 100m usage
- **Enforce limits** - Prevent burst usage that triggers autoscaler
- **Right-size VPA recommendations** - Use VPA to identify bloated requests

### Storage Optimization

- **Restrict PVC sizes** - Cap dev/test at 50Gi, prod at 500Gi
- **Require cheaper StorageClasses** - Use `standard` for non-critical data
- **Block dynamic provisioning** - Require pre-provisioned PVs for large databases

### Autoscaling Optimization

- **Require HPA** - Scale replicas instead of oversizing pods
- **Set reasonable bounds** - Limit HPA max replicas to prevent cost spikes
- **Use cluster autoscaler** - Add nodes only when pending pods exist

## Related Resources

- [OPA Templates Overview](../index.md)
- [OPA RBAC Policies](../rbac/index.md)
- [Kyverno Resource Governance](../resource/index.md)

