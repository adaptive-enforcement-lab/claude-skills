---
name: scheduled-workflows
description: >-
  CronWorkflow patterns for scheduled automation: time-based execution, concurrency policies, orchestration pipelines, and GitHub Actions integration for DevSecOps.
---

# Scheduled Workflows

## When to Use This Skill

CronWorkflows run automation on a schedule: hourly builds, nightly backups, weekly reports. They combine the reliability of Kubernetes cron jobs with the power of Argo Workflows, enabling complex scheduled automation that survives cluster restarts and handles failures gracefully.

---


## Implementation

1. **Define the schedule** using cron syntax
2. **Set concurrency policy** to handle overlaps appropriately
3. **Configure history limits** to prevent resource accumulation
4. **Add monitoring** for schedule misses and failures

---


## Techniques


### Patterns

| Pattern | Description |
| --------- | ------------- |
| [Basic CronWorkflow](basic.md) | Simple scheduled execution |
| [Concurrency Policies](concurrency-policy.md) | Handling overlapping runs |
| [Orchestration](orchestration.md) | Scheduled pipelines that spawn child workflows |
| [GitHub Integration](github-integration.md) | Triggering GitHub Actions from schedules |

---


## Anti-Patterns to Avoid

| Pattern | Description |
| --------- | ------------- |
| [Basic CronWorkflow](basic.md) | Simple scheduled execution |
| [Concurrency Policies](concurrency-policy.md) | Handling overlapping runs |
| [Orchestration](orchestration.md) | Scheduled pipelines that spawn child workflows |
| [GitHub Integration](github-integration.md) | Triggering GitHub Actions from schedules |

---


## Related Patterns

- Basic CronWorkflow
- Concurrency Policies
- Orchestration
- GitHub Integration

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/argo-workflows/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
