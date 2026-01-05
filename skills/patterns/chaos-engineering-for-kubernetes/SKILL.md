---
name: chaos-engineering-for-kubernetes
description: >-
  Chaos engineering for Kubernetes with Chaos Mesh and LitmusChaos. Pod deletion, network chaos, resource chaos, blast radius control, and validation patterns for reliability testing.
---

# Chaos Engineering for Kubernetes

## When to Use This Skill

Chaos engineering transforms reliability from a passive afterthought into an active practice. Instead of waiting for failures to happen, you intentionally inject faults into your systems under controlled conditions. This reveals weaknesses before they become production incidents.

The discipline requires three things: intent, control, and measurement. You run deliberate experiments to test system resilience, limit blast radius to prevent cascade failures, and validate that your observability actually detects the problems you've designed for.

This guide provides production-proven experiment patterns using Chaos Mesh and LitmusChaos, complete with YAML configurations, success criteria, and rollback procedures.


## Implementation

*See [examples.md](examples.md) for detailed code examples.*

> **Start Small, Scale Systematically**
>
> Begin with single-pod experiments in staging. Progress to production only after validating success criteria, rollback procedures, and observability coverage.
## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/reliability/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
