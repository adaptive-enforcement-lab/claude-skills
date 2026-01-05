---
name: kyverno-resource-governance-templates - Reference
description: Complete reference for Kyverno Resource Governance Templates
---

# Kyverno Resource Governance Templates - Reference

This is the complete reference documentation extracted from the source.

# Kyverno Resource Governance Templates

Resource governance policies prevent overconsumption, enforce autoscaling requirements, and control storage allocation across your cluster.

> **Resource Limits Prevent Noisy Neighbors**
>
> Without resource limits, a single pod can consume all node capacity and starve other workloads. Enforce limits to guarantee fair resource allocation.
>

## Why Resource Governance Matters

Kubernetes does not enforce resource limits by default. This creates operational risks:

- **Node Exhaustion** - Pods without limits can consume all CPU/memory
- **OOMKilled Pods** - Memory requests too low cause evictions
- **Autoscaling Failures** - HPA requires resource metrics from limits/requests
- **Cost Overruns** - Uncontrolled storage provisioning inflates cloud bills

## Available Templates

### [Resource Limits and Requests](limits.md)

Enforce CPU and memory limits/requests on all workloads:

- Require resource requests for scheduling decisions
- Mandate resource limits to prevent node exhaustion
- Validate requests ≤ limits for all containers
- Block workloads with excessive resource claims

**Apply a policy:**

```bash
kubectl apply -f limits.yaml
```

### [Horizontal Pod Autoscaler Requirements](hpa.md)

Mandate HPA for production workloads:

- Require HPA for Deployments in production namespaces
- Validate HPA min/max replica bounds
- Ensure HPA targets exist and are valid
- Block HPAs without resource-based metrics

**Apply a policy:**

```bash
kubectl apply -f hpa.yaml
```

### [Storage Constraints](storage.md)

Control PersistentVolume and PersistentVolumeClaim allocation:

- Restrict PVC sizes to prevent excessive storage claims
- Require specific StorageClasses for production data
- Block dynamic provisioning in restricted namespaces
- Validate volume access modes and reclaim policies

**Apply a policy:**

```bash
kubectl apply -f storage.yaml
```

## Resource Management Patterns

### Resource Quotas vs Limits

Use both mechanisms for defense in depth:

- **ResourceQuota** - Namespace-level caps (total CPU/memory across all pods)
- **LimitRange** - Default and max values for individual pods
- **Kyverno Policies** - Validation and enforcement of resource configuration

Kyverno policies complement quotas by validating workload-level configuration before admission.

### Right-Sizing Workloads

Set appropriate resource values to balance cost and reliability:

- **Requests too low** → Pods scheduled on undersized nodes → OOMKilled
- **Requests too high** → Wasted capacity → Increased costs
- **Limits too low** → Pods throttled → Performance degradation
- **Limits too high** → Noisy neighbor problems → Node instability

Use Vertical Pod Autoscaler (VPA) recommendations to identify optimal values.

### Autoscaling Strategies

Choose the right autoscaling mechanism for your workload:

- **HPA (Horizontal)** - Scale replicas based on CPU/memory/custom metrics
- **VPA (Vertical)** - Adjust resource requests/limits automatically
- **Cluster Autoscaler** - Add/remove nodes based on pending pods

Kyverno policies enforce HPA presence and configuration validity.

## Common Enforcement Scenarios

### Scenario 1: Prevent Unbounded Resource Consumption

Require resource limits on all containers:

```yaml
# Enforced by: limits.yaml
# Result: All containers must define resources.limits.cpu and resources.limits.memory
# Impact: Prevents single pod from consuming entire node capacity
```

### Scenario 2: Mandate Autoscaling for Production

Require HPA for production Deployments:

```yaml
# Enforced by: hpa.yaml
# Result: Deployments in prod-* namespaces must have corresponding HPA
# Impact: Ensures production services scale automatically under load
```

### Scenario 3: Control Storage Costs

Restrict PVC size to prevent excessive allocations:

```yaml
# Enforced by: storage.yaml
# Result: PVCs cannot exceed 100Gi in dev namespaces
# Impact: Prevents accidental provisioning of expensive storage volumes
```

## Testing Resource Policies

Validate enforcement without disrupting workloads:

```bash
# Test resource limit requirement (should fail without limits)
kubectl run no-limits --image=nginx
# Expected: Blocked by policy requiring resource limits

# Test excessive resource request (should fail if beyond policy limits)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: excessive-request
spec:
  containers:
    - name: nginx
      image: nginx
      resources:
        requests:
          cpu: "100"
          memory: "1000Gi"
EOF
# Expected: Blocked by policy restricting maximum requests

# Test HPA requirement (should fail without HPA)
kubectl create deployment test-app --image=nginx --replicas=3 -n production
# Expected: Blocked by policy requiring HPA for production Deployments

# Test storage size restriction (should fail for excessive PVC)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: large-pvc
  namespace: dev
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 500Gi
EOF
# Expected: Blocked by policy restricting dev namespace PVC sizes

# Test compliant workload (should succeed)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: compliant-pod
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
# Expected: Allowed by resource limit policies
```

## Resource Budgeting

Plan cluster capacity using policy-enforced constraints:

1. **Calculate namespace budgets:**

   ```bash
   kubectl get resourcequota -n production -o yaml
   # Review quota limits vs current usage
   ```

2. **Identify outliers:**

   ```bash
   kubectl top pods --all-namespaces --sort-by=memory
   kubectl top pods --all-namespaces --sort-by=cpu
   ```

3. **Validate policy alignment:**
   - Do LimitRanges match policy-enforced limits?
   - Are ResourceQuotas enforced at namespace level?
   - Do HPA min/max bounds align with capacity planning?

4. **Monitor policy violations:**

   ```bash
   kubectl get polr -A  # Policy Reports
   # Review workloads blocked by resource policies
   ```

## Cost Optimization

Use resource policies to reduce cloud infrastructure costs:

- **Right-size workloads** - Block oversized resource requests
- **Prevent storage sprawl** - Restrict PVC sizes in non-production
- **Enforce autoscaling** - Scale down during off-peak hours with HPA
- **Use cheaper storage classes** - Require specific StorageClasses for dev/test

## Related Resources

- [Kyverno Templates Overview](../index.md)
- [Kyverno Pod Security](../pod-security/index.md)
- [OPA Resource Governance](../resource/index.md)

