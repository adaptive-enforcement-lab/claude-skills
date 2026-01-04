---
name: event-routing - Examples
description: Code examples for Event Routing
---

# Event Routing - Examples


## Example 1: example-1.mermaid


```mermaid
flowchart LR
    A[EventSource] --> B[EventBus]
    B --> C[Sensor Filter]
    C --> D{Match?}
    D -->|Yes| E[Transform]
    D -->|No| F[Drop]
    E --> G[Trigger]

    %% Ghostty Hardcore Theme
    style A fill:#fd971e,color:#1b1d1e
    style B fill:#515354,color:#f8f8f3
    style C fill:#f92572,color:#1b1d1e
    style E fill:#9e6ffe,color:#1b1d1e
    style G fill:#a7e22e,color:#1b1d1e
    style F fill:#75715e,color:#f8f8f3
```



## Example 2: example-2.yaml


```yaml
apiVersion: argoproj.io/v1alpha1
kind: Sensor
metadata:
  name: prod-image-filter
spec:
  dependencies:
    - name: image-push
      eventSourceName: container-registry
      eventName: push
      filters:
        data:
          - path: body.tag
            type: string
            value:
              - "v*"
              - "release-*"
```



