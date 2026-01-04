---
name: argo-workflows-patterns - Troubleshooting
description: Troubleshooting guide for Argo Workflows Patterns
---

# Argo Workflows Patterns - Troubleshooting

### Workflow Stuck in Pending

1. Check service account permissions: `kubectl describe rolebinding -n argo-workflows`
2. Verify resource quotas: `kubectl describe quota -n argo-workflows`
3. Check node resources: `kubectl top nodes`
4. Look for mutex waits: `kubectl get workflows -l workflows.argoproj.io/sync-id`

### Workflow Failed with RBAC Error

1. Verify ServiceAccount exists in workflow namespace
2. Check ClusterRoleBinding subjects match namespace
3. Use `kubectl auth can-i` to test permissions:

```bash
kubectl auth can-i patch deployments \
  --as=system:serviceaccount:argo-workflows:my-sa \
  -n target-namespace
```

### Mutex Deadlock

1. Find workflows waiting on mutex: `kubectl get workflows -l workflows.argoproj.io/sync-id`
2. Identify the workflow holding the lock
3. Check if the holding workflow is stuck or failed
4. Terminate stuck workflows to release mutex

---

> **Prerequisites**
>
> Argo Workflows must be installed in your cluster. See the [official installation guide](https://argo-workflows.readthedocs.io/en/latest/quick-start/) for setup instructions.
>

---
