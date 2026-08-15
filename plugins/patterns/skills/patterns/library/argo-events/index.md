# Argo Events

Source: https://adaptive-enforcement-lab.com/patterns/argo-events/

Argo Events is an event-driven workflow automation framework for Kubernetes. It connects external event sources to Argo Workflows, enabling reactive automation. For comprehensive documentation, see the [official Argo Events docs](https://argoproj.github.io/argo-events/).

---

## Core Components

```mermaid
flowchart LR
    A[External Event] --> B[EventSource]
    B --> C[EventBus]
    C --> D[Sensor]
    D --> E[Trigger]
    E --> F[Workflow]

    %% Ghostty Hardcore Theme
    style A fill:#65d9ef,color:#1b1d1e
    style B fill:#fd971e,color:#1b1d1e
    style C fill:#515354,color:#f8f8f3
    style D fill:#f92572,color:#1b1d1e
    style E fill:#9e6ffe,color:#1b1d1e
    style F fill:#a7e22e,color:#1b1d1e

```

| Component | Purpose | Argo Docs |
| ----------- | --------- | ----------- |
| **EventSource** | Connects to external systems | [EventSource Types](https://argoproj.github.io/argo-events/concepts/event_source/) |
| **EventBus** | Message broker for delivery | [EventBus](https://argoproj.github.io/argo-events/eventbus/eventbus/) |
| **Sensor** | Filters events and triggers | [Sensors](https://argoproj.github.io/argo-events/concepts/sensor/) |
| **Trigger** | Action when conditions met | [Triggers](https://argoproj.github.io/argo-events/concepts/trigger/) |

---

## Pattern Categories

| Category | Description |
| ---------- | ------------- |
| [Event Routing](routing/index.md) | Multi-action routing, filtering, and transformations |
| [Reliability Patterns](reliability/index.md) | Retry strategies, dead letter queues, and backpressure |

---

## Operational Guides

For deploying and managing Argo Events in production:

- [Setup Guide](setup/index.md) - EventSource, EventBus, and Sensor configuration
- [Troubleshooting](troubleshooting/index.md) - Debugging event flows and common issues
- [High Availability](reliability/high-availability.md) - Production HA architecture

---

> **Prerequisites**
>
> Argo Events requires Argo Workflows for workflow triggers. See the [official installation guide](https://argoproj.github.io/argo-events/installation/) for setup.
>

---

## Related Content

- [Argo Workflows Patterns](../argo-workflows/index.md) - WorkflowTemplate design and error handling
- [Event-Driven Deployments](../../blog/posts/2025-12-14-event-driven-deployments-argo.md) - The journey to zero-latency automation
- [ConfigMap as Cache Pattern](../efficiency/idempotency/caches.md) - Volume mounts for zero-API reads
