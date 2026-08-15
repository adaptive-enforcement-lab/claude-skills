# Policy-as-Code Operations

Source: https://adaptive-enforcement-lab.com/enforce/policy-as-code/operations/

Day-to-day management, updates, and monitoring of policy enforcement.

## Overview

Operating a policy-as-code platform requires:

1. **Policy updates** - Rolling out new policies safely
2. **Monitoring** - Tracking compliance and violations
3. **Exception management** - Handling edge cases
4. **Troubleshooting** - Resolving policy issues
5. **Auditing** - Proving compliance

> **Operations at Scale**
>
> Policy-as-code operations follow GitOps principles. All changes go through Git. All deployments are tracked. All violations are logged.
>

---

## Policy Lifecycle

### Adding New Policies

**Step 1**: Create policy in policy repo

```yaml
# security-policy/charts/security-policy/templates/require-network-policy.yaml
apiVersion: kyverno.io/v1
kind: ClusterPolicy
metadata:
  name: require-network-policy
  annotations:
    policies.kyverno.io/title: Require NetworkPolicy
    policies.kyverno.io/category: Security
    policies.kyverno.io/severity: high
spec:
  validationFailureAction: {{ .Values.policies.networkPolicy.validationFailureAction }}
  background: true
  rules:
    - name: check-network-policy-exists
      match:
        resources:
          kinds:
            - Namespace
      validate:
        message: "Namespace must have a NetworkPolicy"
        deny:
          conditions:
            - key: "{{ request.object.metadata.name }}"
              operator: AnyNotIn
              value: [" kube-system", "kube-public", "kube-node-lease"]
```

**Step 2**: Add to values.yaml

```yaml
# security-policy/charts/security-policy/values.yaml
policies:
  networkPolicy:
    enabled: true
    validationFailureAction: Audit  # Start with Audit
```

**Step 3**: Test locally

```bash
docker run --rm -v $(pwd):/workspace policy-platform:latest bash -c '\
  helm template security /repos/security-policy/charts/security-policy \
    -f /repos/security-policy/charts/security-policy/values.yaml \
  > /tmp/policies.yaml &&\
  kyverno apply /tmp/policies.yaml --resource /workspace/test-namespace.yaml\
'
```

**Step 4**: Deploy to dev

```bash
# Update policy-platform container (rebuild with new policy)
docker build -t policy-platform:v1.0.3 -f ci/Dockerfile .
docker push policy-platform:v1.0.3

# Deploy to dev cluster
helm upgrade security-policy /repos/security-policy/charts/security-policy \
  --namespace kyverno \
  --values /repos/security-policy/cd/dev/values.yaml
```

**Step 5**: Monitor PolicyReports

```bash
kubectl get policyreport -A

# Check for violations
kubectl get policyreport polr-ns-default -o yaml
```

**Step 6**: Switch to Enforce after validation

```yaml
# security-policy/cd/prd/values.yaml
policies:
  networkPolicy:
    validationFailureAction: Enforce  # Now block violations
```

> **Always Start with Audit**
>
> New policies must start in Audit mode. Monitor violations for at least one week before switching to Enforce. This prevents breaking existing workloads.
>

---

## Updating Existing Policies

### Policy Refinement

Refine policy based on violations:

**Original policy** (too strict):

```yaml
validate:
  pattern:
    spec:
      containers:
        - resources:
            limits:
              memory: "?*"
              cpu: "?*"
            requests:              # Requires both limits AND requests
              memory: "?*"
              cpu: "?*"
```

**Updated policy** (allow limits-only):

```yaml
validate:
  pattern:
    spec:
      containers:
        - resources:
            limits:
              memory: "?*"
              cpu: "?*"
          # Requests optional
```

**Deployment**:

1. Update policy in repo
2. Increment version (`v2.1.2` → `v2.1.3`)
3. Rebuild policy-platform container
4. Deploy to clusters

---

## Next Steps

- **[Monitoring](monitoring.md)** - Compliance dashboards, metrics, and alerting
- **[Workflows](workflows.md)** - Updates, backup, performance tuning
- **[Runtime Deployment](../runtime-deployment/index.md)** - Kyverno deployment guide
