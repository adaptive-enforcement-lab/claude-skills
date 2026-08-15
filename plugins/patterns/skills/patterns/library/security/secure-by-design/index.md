# Secure-by-Design Pattern Library

Source: https://adaptive-enforcement-lab.com/patterns/security/secure-by-design/

Building security into architecture from the ground up, not bolting it on afterward. These patterns enforce security properties at the application, network, and admission control layers, making violations visible and costly.

## Pattern Categories

This library covers four fundamental security principles:

### [Zero Trust Patterns](zero-trust.md)

Zero trust rejects implicit trust. Every service, workload, and request proves its identity and intent.

- Service Mesh mTLS with certificate rotation
- Mutual authentication for all inter-service communication
- Network-level verification instead of implicit trust

### [Defense in Depth](defense-in-depth.md)

Defense in depth layers multiple security controls. Compromise at one layer does not compromise the system.

- Pod security contexts with restrictive capabilities
- Network policies with default-deny rules
- Resource limits and read-only filesystems

### [Least Privilege](least-privilege.md)

Least privilege grants only the minimum permissions required for a task.

- Scoped ServiceAccounts with minimal RBAC
- Resource-level permission granularity
- Cross-namespace isolation by default

### [Fail Secure](fail-secure.md)

Fail secure means the system defaults to denying access. Failures default to safe states.

- Admission control with webhook failure modes
- Policy enforcement before object admission
- Audit logging of all decisions

## Integration Patterns

### [End-to-End Deployment](integration.md)

Complete example combining all patterns:

- Zero trust mTLS communication
- Defense in depth pod hardening
- Least privilege RBAC configuration
- Fail secure admission controls

### [Security Audit Checklist](integration.md#security-audit-checklist)

Verification checklist before deployment:

- [ ] Zero Trust: mTLS policies in place
- [ ] Defense in Depth: Pod security contexts enforced
- [ ] Network Policies: Default-deny rules configured
- [ ] Least Privilege: Minimal RBAC permissions
- [ ] Fail Secure: Admission webhooks with failurePolicy: Fail

## Quick Reference

### Security Properties Matrix

| Pattern | Primary Control | Threat Mitigated | Implementation Cost |
|---------|----------------|------------------|---------------------|
| **Zero Trust mTLS** | Network encryption | MITM attacks, lateral movement | Medium (service mesh) |
| **Pod Security Context** | Process isolation | Privilege escalation | Low (YAML config) |
| **Network Policies** | Network isolation | Unauthorized access | Low (YAML config) |
| **Scoped RBAC** | Permission control | Lateral movement | Medium (design effort) |
| **Admission Webhooks** | Policy enforcement | Configuration bypass | High (webhook service) |

> **Security Is Not Optional**
>
> These patterns are not suggestions. They're requirements for production systems. Skipping defense-in-depth or fail-secure controls creates exploitable vulnerabilities.
>

### Common Anti-Patterns

- **Privilege escalation for convenience**: `allowPrivilegeEscalation: true` defeats most controls
- **PERMISSIVE mTLS mode**: Leaves window for plaintext traffic
- **Cluster-admin role bindings**: Workloads should never have cluster-admin
- **failurePolicy: Ignore**: Causes bypass of policies if webhook unavailable
- **Wildcard RBAC permissions**: `verbs: ["*"]` violates least privilege

## References

- [Kubernetes Security Best Practices](https://kubernetes.io/docs/concepts/security/pod-security-standards/)
- [Istio Security Policies](https://istio.io/latest/docs/concepts/security/)
- [NIST Cybersecurity Framework](https://www.nist.gov/cyberframework)
- [CIS Kubernetes Benchmarks](https://www.cisecurity.org/benchmark/kubernetes)
- [OWASP Top 10 for Kubernetes](https://www.cisa.gov/kubernetes-hardening-guidance)

---

*Security by design, not by accident. Enforce properties through architecture, not documentation.*
