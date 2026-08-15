# Kyverno Basics

Source: https://adaptive-enforcement-lab.com/enforce/policy-as-code/kyverno/

Kyverno runs as a dynamic admission controller in Kubernetes. It validates, mutates, and generates resources based on policies written in YAML.

---

## Installation

Install Kyverno using Helm:

```bash
# Add Kyverno Helm repository
helm repo add kyverno https://kyverno.github.io/kyverno/
helm repo update

# Install Kyverno
helm install kyverno kyverno/kyverno \
  --namespace kyverno \
  --create-namespace \
  --set replicaCount=3

# Verify installation
kubectl get pods -n kyverno
```

Kyverno creates webhook configurations that intercept resource creation/updates before they reach etcd.

---

## Basic Kyverno Policy

> **Quick Start**
>
> This guide is part of a modular documentation set. Refer to related guides in the navigation for complete context.
>

Require resource limits on all deployments:

```yaml
apiVersion: kyverno.io/v1
kind: ClusterPolicy
metadata:
  name: require-resource-limits
spec:
  validationFailureAction: Enforce
  background: true
  rules:
    - name: check-resource-limits
      match:
        any:
          - resources:
              kinds:
                - Deployment
      validate:
        message: "Resource limits are required for all containers"
        pattern:
          spec:
            template:
              spec:
                containers:
                  - resources:
                      limits:
                        memory: "?*"
                        cpu: "?*"
```

Try to deploy without limits:

```bash
$ kubectl apply -f deployment.yaml
Error from server: admission webhook "validate.kyverno.svc-fail" denied the request:

policy Deployment/default/api for resource violation:

require-resource-limits:
  check-resource-limits: validation error: Resource limits are required for all containers
```

Deployment blocked. Policy enforced.

---

## Audit Mode vs Enforce Mode

Roll out policies in audit mode first:

```yaml
spec:
  validationFailureAction: Audit  # Log violations, don't block
```

Check logs for violations:

```bash
kubectl get policyreport -A

NAMESPACE   NAME                          PASS   FAIL   WARN   ERROR   SKIP
default     polr-ns-default              12     3      0      0       0
production  polr-ns-production           45     1      0      0       0
```

Fix violations. Then switch to Enforce:

```yaml
spec:
  validationFailureAction: Enforce  # Block violations
```

### Gradual Rollout Strategy

1. Deploy policy in `Audit` mode
2. Monitor PolicyReports for 1 week
3. Remediate failures
4. Switch to `Enforce` mode
5. Handle exceptions with exclusions

Don't deploy straight to Enforce. Discover violations first.

---

## Policy Structure

All Kyverno policies follow this structure:

```yaml
apiVersion: kyverno.io/v1
kind: ClusterPolicy  # or Policy for namespaced
metadata:
  name: policy-name
spec:
  validationFailureAction: Enforce | Audit
  background: true | false  # Apply to existing resources
  rules:
    - name: rule-name
      match:  # What resources to check
        any:
          - resources:
              kinds: [Deployment, StatefulSet]
              namespaces: [production, staging]
      exclude:  # What to skip
        any:
          - resources:
              namespaces: [kube-system]
      validate | mutate | generate:  # What to do
        # Policy logic here
```

---

## Related Guides

- **[Policy Patterns](policy-patterns.md)** - Common validation and mutation patterns
- **[Testing and Exceptions](testing-approaches.md)** - Test policies before production
- **[CI/CD Integration](ci-cd-integration.md)** - Automate policy validation

---

*Policy deployed in audit mode. Violations logged. Teams notified. Fixes deployed. Policy switched to enforce. Zero production impact.*
