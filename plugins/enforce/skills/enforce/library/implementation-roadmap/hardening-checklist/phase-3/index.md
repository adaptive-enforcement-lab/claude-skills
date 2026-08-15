# Phase 3: Runtime (Weeks 9-12)

Source: https://adaptive-enforcement-lab.com/enforce/implementation-roadmap/hardening-checklist/phase-3/

Control what runs in production, not just what gets committed. Policy is enforced at admission time before pods can deploy.

> **Real-World Impact**
>
> A SaaS company deployed Kyverno policies to 5 Kubernetes clusters in 1 week. Within 72 hours, the policies blocked 34 pod deployments without resource limits, 18 images from untrusted registries, and 12 containers attempting to run as root. Zero manual intervention required.
>

---

## Phase Overview

Phase 3 extends enforcement to runtime through three critical areas:

1. **[Policy Enforcement](policy-enforcement.md)** - Core Kyverno policies for resource limits, image verification, security context
2. **[Advanced Policies](advanced-policies.md)** - Namespace quotas, pod security standards, network policies
3. **[Rollout Strategy](rollout.md)** - Audit-first deployment approach and metrics

These controls ensure only compliant workloads run in production.

---

## Phase Components

### Policy Enforcement

Core admission control policies that block non-compliant pods.

**Key Controls**:

- Kyverno deployment and configuration
- Required resource limits (CPU/memory)
- Image source verification (approved registries)
- Security context requirements (non-root, read-only)
- Policy Reporter dashboard

**[View Policy Enforcement Details →](policy-enforcement.md)**

---

### Advanced Policies

Extended runtime controls for comprehensive security.

**Key Controls**:

- Namespace resource quotas
- Pod security standards (baseline/restricted)
- Network policy requirements
- System namespace exclusions

**[View Advanced Policies Details →](advanced-policies.md)**

---

### Rollout Strategy

Safe deployment approach with audit-first methodology.

**Key Controls**:

- Audit mode monitoring (Week 1)
- Violation remediation (Week 2)
- Enforce mode activation (Week 3)
- Metrics tracking and tuning (Week 4)

**[View Rollout Strategy Details →](rollout.md)**

---

## Phase 3 Validation Checklist

Before moving to Phase 4, verify all runtime controls work:

- [ ] Kyverno is deployed and webhooks are running
- [ ] Policy Reporter UI is accessible and showing violations
- [ ] Resource limits policy blocks pods without limits
- [ ] Image source policy blocks untrusted registries
- [ ] Security context policy blocks root containers
- [ ] Pod security standards are enforced in production
- [ ] Network policies are required for namespaces
- [ ] Namespace quotas are enforced
- [ ] Policy violations are visible in dashboard
- [ ] All policies have been tested in Audit mode first

---

## Validation Commands

Test that controls are working:

```bash
# Test pod without limits is rejected
kubectl apply -f - <<EOF
apiVersion: v1
kind: Pod
metadata:
  name: test
spec:
  containers:
    - name: app
      image: nginx
EOF
# Expected: Admission webhook denies request

# Test untrusted registry is blocked
kubectl apply -f - <<EOF
apiVersion: v1
kind: Pod
metadata:
  name: test
spec:
  containers:
    - name: app
      image: docker.io/nginx:latest
EOF
# Expected: Image source validation fails

# Check policy reports
kubectl get policyreport -A
# Expected: Shows pass/fail summary
```

---

## Next Steps

With Phase 3 complete, you have:

- Kyverno enforcing policies at admission time
- Resource limits on all containers
- Image source verification
- Security context enforcement
- Policy observability dashboard

**[Proceed to Phase 4: Advanced →](../phase-4/index.md)**

Phase 4 completes the implementation with audit evidence collection, compliance validation, and OpenSSF Scorecard monitoring.

---

## Related Patterns

- **[Policy-as-Code with Kyverno](../../../policy-as-code/kyverno/index.md)** - Detailed policy configuration
- **[Pod Security Standards](../../../../secure/cloud-native/gke-hardening/runtime-security/pod-security-standards.md)** - Security context requirements
- **[Runtime Security](../../../../secure/cloud-native/gke-hardening/runtime-security/index.md)** - Resource limits and runtime controls
- **[Implementation Roadmap Overview](index.md)** - Complete roadmap
- **[Phase 2: Automation](../phase-2/index.md)** - CI/CD gates
- **[Phase 4: Advanced →](../phase-4/index.md)** - Audit evidence and compliance

---

*Kyverno deployed. Policies enforced. Pods without limits blocked. Untrusted images rejected. Root containers denied. Runtime security is enforced, not hoped for.*
