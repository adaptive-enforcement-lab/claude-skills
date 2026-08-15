# Scheduled Workflows

Source: https://adaptive-enforcement-lab.com/patterns/argo-workflows/scheduled/

CronWorkflows run automation on a schedule: hourly builds, nightly backups, weekly reports. They combine the reliability of Kubernetes cron jobs with the power of Argo Workflows, enabling complex scheduled automation that survives cluster restarts and handles failures gracefully.

---

## Why CronWorkflows?

Kubernetes CronJobs work for simple scheduled tasks. But they have limitations:

- Single-container jobs only
- Limited failure handling
- No workflow visualization
- Basic retry logic

CronWorkflows provide the full power of Argo Workflows on a schedule. Multi-step pipelines, sophisticated retry strategies, visual debugging, and artifact management are all available for scheduled automation.

---

## Patterns

| Pattern | Description |
| --------- | ------------- |
| [Basic CronWorkflow](basic.md) | Simple scheduled execution |
| [Concurrency Policies](concurrency-policy.md) | Handling overlapping runs |
| [Orchestration](orchestration.md) | Scheduled pipelines that spawn child workflows |
| [GitHub Integration](github-integration.md) | Triggering GitHub Actions from schedules |

---

## Quick Start

1. **Define the schedule** using cron syntax
2. **Set concurrency policy** to handle overlaps appropriately
3. **Configure history limits** to prevent resource accumulation
4. **Add monitoring** for schedule misses and failures

---

## Cron Syntax Quick Reference

| Expression | Meaning |
| ------------ | --------- |
| `0 * * * *` | Every hour at minute 0 |
| `0 0 * * *` | Daily at midnight |
| `0 0 * * 0` | Weekly on Sunday at midnight |
| `0 0 1 * *` | Monthly on the 1st at midnight |
| `*/15 * * * *` | Every 15 minutes |
| `0 9-17 * * 1-5` | Hourly 9am-5pm, Mon-Fri |

> **Use UTC Unless Specified**
>
> CronWorkflows default to UTC. Use the `timezone` field for local time scheduling.
>

---

## Related

- [WorkflowTemplate Patterns](../templates/index.md): Building workflow templates to schedule
- [Concurrency Control](../concurrency/index.md): Mutex and semaphore patterns
- [Workflow Composition](../composition/index.md): Complex scheduled pipelines
