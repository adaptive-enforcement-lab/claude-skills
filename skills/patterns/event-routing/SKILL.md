---
name: event-routing
description: >-
  Control event flow from EventSources to triggers. Master filtering, transformation, and multi-action capabilities for sophisticated event-driven automation.
---

# Event Routing

## When to Use This Skill

Event routing controls how events flow from EventSources through Sensors to Triggers. Argo Events provides powerful filtering, transformation, and multi-action capabilities. For the complete reference, see the [official Sensors documentation](https://argoproj.github.io/argo-events/sensors/intro/).

---


## Implementation

See the full implementation guide in the [source documentation](https://adaptive-enforcement-lab.com/patterns/argo-events/).


## Techniques


### Routing Patterns

| Pattern | Use Case | Complexity |
| --------- | ---------- | ------------ |
| [Simple Filtering](filtering.md) | Route events based on field values | Low |
| [Multi-Trigger Actions](multi-trigger.md) | Execute multiple actions from one event | Medium |
| [Event Transformation](transformation.md) | Modify payloads before triggering | Medium |
| [Conditional Routing](conditional.md) | Complex decision trees | High |

---


## Examples

See [examples.md](examples.md) for code examples.


## Related Patterns

- Simple Filtering
- Multi-Trigger Actions
- Sensor Configuration
- Official Sensor Docs

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/argo-events/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
