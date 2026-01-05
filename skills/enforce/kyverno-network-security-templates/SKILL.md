---
name: kyverno-network-security-templates
description: >-
  Kyverno network security policies that enforce NetworkPolicy requirements, Ingress rules, and Service restrictions in Kubernetes.
---

# Kyverno Network Security Templates

## When to Use This Skill

Network policies control traffic between pods, namespaces, and external endpoints. These templates enforce network segmentation and prevent unauthorized communication.

> **Network Policies Require CNI Support**
>
> NetworkPolicy resources only function when your CNI plugin supports them. Verify your cluster's CNI (Calico, Cilium, Weave Net) before deploying network policies.


## When to Apply

### Scenario 1: Prevent Unapproved External Exposure

Block LoadBalancer services except for approved namespaces:

```yaml
# Enforced by: services.yaml
# Result: Only ingress-nginx namespace can create LoadBalancer services
# Impact: Prevents accidental exposure of internal services to the internet
```

### Scenario 2: Mandate TLS for Public Services

Require TLS configuration on all Ingress resources:

```yaml
# Enforced by: ingress-tls.yaml
# Result: All Ingress objects must define spec.tls with valid secrets
# Impact: Eliminates plaintext HTTP exposure for external services
```

### Scenario 3: Enforce Namespace Isolation

Require NetworkPolicy in every namespace before pod creation:

```yaml
# Enforced by: security.yaml
# Result: Namespaces must have NetworkPolicy resources before accepting workloads
# Impact: Prevents pods from communicating across namespace boundaries by default
```


## Implementation

See the full implementation guide in the [source documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/).


## Techniques


### Network Security Patterns

### Defense in Depth

Layer network controls across multiple boundaries:

1. **Namespace NetworkPolicies** - Default deny all traffic
2. **Service Restrictions** - Limit LoadBalancer/NodePort usage
3. **Ingress Controls** - Require TLS and approved ingress classes
4. **Egress Filtering** - Block unauthorized external connections

### Zero-Trust Networking

Never assume trust based on network location:

- Require explicit NetworkPolicy allow rules (no implicit trust)
- Mandate mTLS for service-to-service communication (use service mesh if needed)
- Validate identity at every network boundary (authentication, not IP allowlisting)

### Production vs Non-Production

Use different enforcement levels based on environment:

- **Production** - Strict NetworkPolicy requirements, TLS mandatory, LoadBalancer restricted
- **Development** - Relaxed policies, allow broader access for testing
- **Staging** - Production-like policies to catch configuration issues early


## Examples

See [examples.md](examples.md) for code examples.


## Related Patterns

- Kyverno Templates Overview
- Kyverno Pod Security
- Kyverno Resource Governance

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
