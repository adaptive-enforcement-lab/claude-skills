---
name: argo-events-setup-guide - Troubleshooting
description: Troubleshooting guide for Argo Events Setup Guide
---

# Argo Events Setup Guide - Troubleshooting

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
