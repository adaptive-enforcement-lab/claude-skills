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



## Implementation


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/enforce/policy-as-code/)
- [AEL Enforce](https://adaptive-enforcement-lab.com/enforce/)
