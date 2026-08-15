# Chaos Engineering for Kubernetes

Source: https://adaptive-enforcement-lab.com/patterns/reliability/chaos-engineering/

Chaos engineering transforms reliability from a passive afterthought into an active practice. Instead of waiting for failures to happen, you intentionally inject faults into your systems under controlled conditions. This reveals weaknesses before they become production incidents.

The discipline requires three things: intent, control, and measurement. You run deliberate experiments to test system resilience, limit blast radius to prevent cascade failures, and validate that your observability actually detects the problems you've designed for.

This guide provides production-proven experiment patterns using Chaos Mesh and LitmusChaos, complete with YAML configurations, success criteria, and rollback procedures.

## Why Chaos Engineering Matters

Traditional testing validates happy paths. Chaos engineering validates failure handling: the code paths that matter most when systems break.

### Common discovery patterns

- **Graceful degradation failures**: Service stops responding instead of falling back to defaults
- **Cascading timeouts**: One slow dependency freezes the entire request tree
- **Resource starvation**: Memory leaks or unbounded connections exhaust limits under sustained load
- **Unbalanced blast radius**: Single pod deletion crashes unrelated services due to hard dependencies
- **Silent observability gaps**: Actual failures do not trigger alerts because monitoring missed the edge case

Chaos experiments expose these patterns in controlled test windows before they cause customer impact.

## Navigation

### Core Concepts

- **[Tools Comparison](tools-comparison.md)**: Chaos Mesh vs LitmusChaos capabilities and selection guidance
- **[Blast Radius Control](blast-radius.md)**: Targeting strategies, progressive intensity, and automatic rollback
- **[Validation Patterns](validation.md)**: SLI monitoring, incident detection testing, and auto-remediation verification

### Practical Implementation

- **[Experiment Catalog](experiments.md)**: Pod deletion, network latency, memory pressure, and dependency failure scenarios
- **[Running Experiments Safely](operations.md)**: Pre-experiment checklist, execution best practices, and post-experiment analysis
- **[Observability Integration](observability.md)**: Key metrics, alert rules, and common pitfalls

## Quick Start

```yaml
# Example: Simple pod deletion experiment
apiVersion: chaos-mesh.org/v1alpha1
kind: PodChaos
metadata:
  namespace: chaos-testing
  name: pod-deletion-staging
spec:
  action: pod-kill
  mode: fixed
  value: 1
  selector:
    namespaces:
      - staging
    labelSelectors:
      app: api-gateway
  duration: 2m
  schedule:
    cron: "0 2 * * 1-4"  # 2 AM, Monday-Thursday
```

> **Start Small, Scale Systematically**
>
> Begin with single-pod experiments in staging. Progress to production only after validating success criteria, rollback procedures, and observability coverage.
>

## Scaling Chaos Programs

Start small, systematize, scale:

**Phase 1: Experiment pilots** (Week 1 to 2)

- Single service, single experiment type
- Manual execution, documented runbook
- Build team confidence

**Phase 2: Recurring schedule** (Week 3 to 4)

- Weekly chaos window same time
- Automated via Argo Workflows
- Team on call rotation established

**Phase 3: Hypothesis driven experiments** (Month 2)

- Design experiments based on incident postmortems
- Validate fixes with chaos before deploying
- Track mean time to failure improvements

**Phase 4: GameDays** (Month 3 and beyond)

- Entire team participates
- Multi service scenarios
- Incident response training
- Cross team collaboration

**Phase 5: Continuous chaos** (Month 6 and beyond)

- Steady state fault injection
- Detection validation on every deployment
- Automatic experiment catalog updates
- Chaos engineering as standard practice

## References and Further Reading

- **Chaos Mesh**: Complete documentation and experiment types at chaos-mesh.org/docs
- **LitmusChaos**: Orchestration and experiment library at litmuschaos.io
- **Principles of Chaos Engineering**: Foundational concepts at principlesofchaos.org
- **SLO/SLI/SLA primer**: Track what matters during chaos
- **Incident Postmortems**: Use them to design targeted experiments
