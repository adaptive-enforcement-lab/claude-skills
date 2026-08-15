# Kyverno Network Security Templates

Source: https://adaptive-enforcement-lab.com/enforce/policy-as-code/template-library/kyverno/network/

Network policies control traffic between pods, namespaces, and external endpoints. These templates enforce network segmentation and prevent unauthorized communication.

> **Network Policies Require CNI Support**
>
> NetworkPolicy resources only function when your CNI plugin supports them. Verify your cluster's CNI (Calico, Cilium, Weave Net) before deploying network policies.
>

## Why Network Policy Enforcement Matters

Default Kubernetes behavior allows all pod-to-pod communication. This creates a flat network where any compromised pod can reach any other pod.

**Network policies provide:**

- Namespace isolation (prevent cross-namespace traffic)
- Service-level segmentation (databases only accessible to specific apps)
- Egress controls (block unauthorized external connections)
- Zero-trust networking (explicit allow-lists only)

## Available Templates

### [NetworkPolicy Requirements](security.md)

Require NetworkPolicy resources in every namespace:

- Mandate default-deny policies for new namespaces
- Enforce ingress and egress rules for production workloads
- Block namespaces without network segmentation

**Apply a policy:**

```bash
kubectl apply -f security.yaml
```

### [Ingress Class Validation](ingress-class.md)

Enforce approved IngressClass usage:

- Restrict Ingress resources to approved IngressClass values
- Prevent direct exposure through unapproved ingress controllers
- Validate ingress annotations for security requirements

**Apply a policy:**

```bash
kubectl apply -f ingress-class.yaml
```

### [Ingress TLS Requirements](ingress-tls.md)

Mandate TLS termination for all Ingress resources:

- Require TLS configuration on all Ingress objects
- Validate TLS secret references exist
- Enforce HTTPS-only traffic for external services

**Apply a policy:**

```bash
kubectl apply -f ingress-tls.yaml
```

### [Service Type Restrictions](services.md)

Control Service exposure and external access:

- Restrict LoadBalancer and NodePort Service types
- Require annotations for external-facing services
- Validate service selectors and port configurations

**Apply a policy:**

```bash
kubectl apply -f services.yaml
```

## Network Security Patterns

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

## Common Enforcement Scenarios

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

## Testing Network Policies

Validate NetworkPolicy enforcement without disrupting traffic:

```bash
# Test NetworkPolicy requirement (should fail without policy)
kubectl create namespace test-ns
kubectl run test-pod --image=nginx -n test-ns
# Expected: Blocked by policy requiring NetworkPolicy in namespace

# Test Ingress TLS requirement (should fail without TLS)
cat <<EOF | kubectl apply -f -
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: test-ingress
  namespace: test-ns
spec:
  rules:
    - host: test.example.com
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: test-service
                port:
                  number: 80
EOF
# Expected: Blocked by policy requiring spec.tls

# Test Service type restriction (should fail for LoadBalancer)
kubectl expose deployment test-app --type=LoadBalancer --port=80 -n test-ns
# Expected: Blocked by policy restricting LoadBalancer type
```

## Related Resources

- [Kyverno Templates Overview](../index.md)
- [Kyverno Pod Security](../pod-security/index.md)
- [Kyverno Resource Governance](../resource/index.md)
