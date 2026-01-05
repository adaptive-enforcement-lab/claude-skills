---
name: argo-events-setup-guide
description: >-
  Deploy event-driven automation with EventSource, EventBus, and Sensor configuration. Complete setup guide for production Argo Events on Kubernetes clusters.
---

# Argo Events Setup Guide

## When to Use This Skill

This guide covers EventSource, EventBus, and Sensor configuration for event-driven automation.

---


## Implementation

This guide covers EventSource, EventBus, and Sensor configuration for event-driven automation.

---

## Components

| Component | Purpose | Guide |
| ----------- | --------- | ------- |
| **EventSource** | Connect to external systems (Pub/Sub, webhooks) | [EventSource Configuration](event-sources.md) |
| **EventBus** | Message broker for event delivery | [EventBus Configuration](event-bus.md) |
| **Sensor** | Filter events and trigger workflows | [Sensor Configuration](sensors.md) |

---

## Quick Start

1. **Deploy EventBus** - Start with [JetStream for production](event-bus.md#jetstream_eventbus_production)
2. **Configure EventSource** - Connect your [Pub/Sub topic](event-sources.md#pubsub_eventsource) or [GitHub webhooks](event-sources.md#github_webhook_eventsource)
3. **Create Sensor** - Define [event filters and triggers](sensors.md#basic_sensor)

---

> **EventBus First**
>
> Deploy the EventBus before creating EventSources or Sensors. Without a running EventBus, events have nowhere to go.
>

---

## Troubleshooting

### Events Not Arriving

1. Check EventSource logs: `kubectl logs -n argo-events -l eventsource-name=<name>`
2. Verify Pub/Sub subscription exists in GCP console
3. Confirm service account has `pubsub.subscriber` role

### Events Arriving But Not Triggering

1. Check Sensor logs: `kubectl logs -n argo-events -l sensor-name=<name>`
2. Verify filter conditions match event payload
3. Test with a simple sensor that logs all events

### Events Lost During Restarts

1. Enable [persistence on EventBus](event-bus.md#nats_with_persistence)
2. Increase `maxAge` retention
3. Monitor EventBus storage usage

---

## Related

- [Argo Workflows Patterns](../../argo-workflows/index.md) - WorkflowTemplate design and error handling
- [ConfigMap as Cache Pattern](../../../patterns/efficiency/idempotency/caches.md) - Volume mounts for zero-API reads
- [Event-Driven Deployments](../../../blog/posts/2025-12-14-event-driven-deployments-argo.md) - The journey to zero-latency automation

### Components

| Component | Purpose | Guide |
| ----------- | --------- | ------- |
| **EventSource** | Connect to external systems (Pub/Sub, webhooks) | [EventSource Configuration](event-sources.md) |
| **EventBus** | Message broker for event delivery | [EventBus Configuration](event-bus.md) |
| **Sensor** | Filter events and trigger workflows | [Sensor Configuration](sensors.md) |

---

### Quick Start

1. **Deploy EventBus** - Start with [JetStream for production](event-bus.md#jetstream_eventbus_production)
2. **Configure EventSource** - Connect your [Pub/Sub topic](event-sources.md#pubsub_eventsource) or [GitHub webhooks](event-sources.md#github_webhook_eventsource)
3. **Create Sensor** - Define [event filters and triggers](sensors.md#basic_sensor)

---

> **EventBus First**
>
> Deploy the EventBus before creating EventSources or Sensors. Without a running EventBus, events have nowhere to go.
>

---

### Troubleshooting

### Events Not Arriving

1. Check EventSource logs: `kubectl logs -n argo-events -l eventsource-name=<name>`
2. Verify Pub/Sub subscription exists in GCP console
3. Confirm service account has `pubsub.subscriber` role

### Events Arriving But Not Triggering

1. Check Sensor logs: `kubectl logs -n argo-events -l sensor-name=<name>`
2. Verify filter conditions match event payload
3. Test with a simple sensor that logs all events

### Events Lost During Restarts

1. Enable [persistence on EventBus](event-bus.md#nats_with_persistence)
2. Increase `maxAge` retention
3. Monitor EventBus storage usage

---

### Related

- [Argo Workflows Patterns](../../argo-workflows/index.md) - WorkflowTemplate design and error handling
- [ConfigMap as Cache Pattern](../../../patterns/efficiency/idempotency/caches.md) - Volume mounts for zero-API reads
- [Event-Driven Deployments](../../../blog/posts/2025-12-14-event-driven-deployments-argo.md) - The journey to zero-latency automation


## Troubleshooting

See [troubleshooting.md](troubleshooting.md) for common issues and solutions.


## Related Patterns

- Argo Workflows Patterns
- ConfigMap as Cache Pattern
- Event-Driven Deployments

## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/argo-events/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
