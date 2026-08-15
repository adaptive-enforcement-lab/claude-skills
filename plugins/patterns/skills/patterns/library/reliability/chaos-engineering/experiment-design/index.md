# Chaos Experiment Design

Source: https://adaptive-enforcement-lab.com/patterns/reliability/chaos-engineering/experiment-design/

Chaos without validation is just breaking things. Proper experiment design transforms fault injection into reliability engineering.

> **Core Principle**
>
> Every chaos experiment must have a hypothesis, measurable success criteria, controlled blast radius, and automated validation. If you can't measure it, you can't learn from it.
>

## Navigation

This section covers the complete methodology for designing and executing chaos experiments:

### Core Topics

- **[Hypothesis Formation](hypothesis.md)** - Structure hypotheses, define observable outcomes, start specific
- **[Success Criteria](success-criteria.md)** - SLI-based validation, observable metrics, recovery verification
- **[Blast Radius Control](blast-radius.md)** - Targeting strategies, progressive intensity, automatic rollback
- **[SLI Monitoring](sli-monitoring.md)** - Baseline measurement, live tracking, recovery validation
- **[Validation Patterns](validation.md)** - Incident detection testing, auto-remediation verification, experiment catalog

## Quick Reference

### Hypothesis Template

```text
Given: [Normal operating conditions]
When: [Specific failure injected]
Then: [Expected system behavior]
And: [Observable metrics that validate behavior]
```

### Success Criteria Checklist

- [ ] Baseline metrics captured before chaos
- [ ] Live metrics tracked during chaos
- [ ] Recovery metrics measured after chaos
- [ ] Comparison shows system returned to baseline
- [ ] No degradation persists after experiment ends

### Blast Radius Constraints

- Start with 1 pod, 30 seconds
- Progress to 10% after 2 weeks
- Require compensating controls for production
- Configure automatic rollback on threshold breach

### Pre-Experiment Checklist

- [ ] Experiment documented in runbook with owner
- [ ] On-call team notified of chaos window
- [ ] Blast radius explicitly validated
- [ ] Rollback procedure tested in staging
- [ ] SLI dashboards visible and alert thresholds set
- [ ] No ongoing production incidents
- [ ] Low-traffic window selected
- [ ] Escalation path established

## Related Patterns

- **[Chaos Engineering Overview](../index.md)** - Framework introduction
- **[Observability](../observability.md)** - Monitoring, metrics, and SLOs
- **[Experiments](../experiments.md)** - Complete experiment catalog

---

*Hypothesis formed. Success criteria defined. Blast radius controlled. Validation automated. Chaos is science, not randomness.*
