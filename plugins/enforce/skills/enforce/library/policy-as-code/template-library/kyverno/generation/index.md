# Kyverno Generation Templates

Source: https://adaptive-enforcement-lab.com/enforce/policy-as-code/template-library/kyverno/generation/

Kyverno generation policies automatically create supporting resources when specific conditions are met. Instead of relying on manual configuration or documentation, generation policies enforce security and resilience by default.

> **Manual Configuration Creates Gaps**
>
> Relying on documentation to ensure every namespace has ResourceQuotas or every HA workload has a PodDisruptionBudget creates inconsistencies. Generation policies enforce these requirements automatically at creation time.
>

## What is Generation?

Generation policies create new Kubernetes resources in response to triggers. When a namespace is created, a generation policy can automatically add ResourceQuotas, NetworkPolicies, and LimitRanges. When a high-availability Deployment appears, a generation policy can create a PodDisruptionBudget.

This eliminates the gap between "we should have this" and "we do have this."

---

## Key Concepts

### Triggers

Generation policies react to resource creation or changes:

- **Namespace creation** → Generate default-deny NetworkPolicies and ResourceQuotas
- **Deployment with 2+ replicas** → Generate PodDisruptionBudget
- **Production workload** → Generate stricter quotas and policies
- **Label changes** → Synchronize generated resources with new requirements

### Synchronization

When `synchronize: true` is set, Kyverno keeps generated resources in sync with the source. If you change a namespace label, the generated ResourceQuota updates to match. If you scale a Deployment to 1 replica, the PodDisruptionBudget is removed.

This enforces consistency without manual intervention.

### Exclusions and Preconditions

Generation policies use exclusions to skip system namespaces and preconditions to enforce requirements:

```yaml
exclude:
  resources:
    names:
      - kube-system
      - kube-public
      - kube-node-lease

preconditions:
  all:
    - key: "{{ request.object.spec.replicas }}"
      operator: GreaterThanOrEquals
      value: 2
```

Only workloads meeting specific criteria get generated resources.

---

## Template Categories

### Namespace Resources

Automatically create security and resource governance for new namespaces:

- **[ResourceQuotas and NetworkPolicies](namespace.md)** - Default quotas and default-deny networking for every new namespace

**Use cases:**

- Prevent resource exhaustion from unconstrained namespaces
- Enforce zero-trust networking by default
- Automatically apply environment-specific quotas (dev vs production)
- Ensure DNS egress is allowed while blocking other traffic

---

### Workload Resources

Automatically create resilience and availability controls for high-availability workloads:

- **[PodDisruptionBudgets](workload.md)** - Automatic PDBs for Deployments and StatefulSets with multiple replicas

**Use cases:**

- Prevent downtime during cluster upgrades and node maintenance
- Enforce SLA compliance for critical services
- Protect against mass pod evictions during autoscaling
- Maintain service availability during partial failures

---

## When to Use Generation

**Use generation when:**

- You want to enforce security-by-default for all new resources
- Manual configuration creates gaps and inconsistencies
- You need automatic synchronization with changing requirements
- Supporting resources should follow workload lifecycle (create, update, delete)

**Do not use generation when:**

- Resources require unique, per-workload customization
- Generated resources would conflict with existing resources
- You need human approval before resource creation
- The triggering resource does not contain enough context to generate correctly

---

## Validation Strategy

After deploying generation policies, validate that resources are created correctly:

```bash
# Check that new namespaces get ResourceQuotas
kubectl create namespace test-gen
kubectl get resourcequotas -n test-gen

# Check that multi-replica Deployments get PDBs
kubectl create deployment nginx --image=nginx --replicas=3 -n test-gen
kubectl label deployment nginx app=nginx -n test-gen
kubectl get pdb -n test-gen

# Audit resources without expected generated objects
kubectl get namespaces -o json | jq -r '.items[] | select(.metadata.name != "kube-system") | .metadata.name' | while read ns; do
  quota_count=$(kubectl get resourcequotas -n $ns --no-headers 2>/dev/null | wc -l)
  if [ $quota_count -eq 0 ]; then
    echo "WARNING: Namespace $ns has no ResourceQuota"
  fi
done
```

---

## Related Resources

- **[Kyverno Templates Overview](../index.md)** - Back to Kyverno templates
- **[Template Library Overview](../index.md)** - Back to main template library
