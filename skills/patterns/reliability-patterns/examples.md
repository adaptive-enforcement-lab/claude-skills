---
name: reliability-patterns - Examples
description: Code examples for Reliability Patterns
---

# Reliability Patterns - Examples


## Example 1: example-1.mermaid


```mermaid
flowchart TD
    A[Event Arrives] --> B[EventSource Retry]
    B -->|Success| C[EventBus Persistence]
    C --> D[Sensor Processing]
    D -->|Trigger Fails| E[Trigger Retry]
    E -->|Exhausted| F[Dead Letter Queue]
    D -->|Success| G[Action Complete]

    %% Ghostty Hardcore Theme
    style A fill:#65d9ef,color:#1b1d1e
    style B fill:#fd971e,color:#1b1d1e
    style C fill:#515354,color:#f8f8f3
    style D fill:#f92572,color:#1b1d1e
    style E fill:#9e6ffe,color:#1b1d1e
    style F fill:#f92572,color:#1b1d1e
    style G fill:#a7e22e,color:#1b1d1e
```



## Example 2: example-2.yaml


```yaml
triggers:
  - template:
      name: deploy-with-retry
      argoWorkflow:
        operation: submit
        source:
          resource:
            # ...
    retryStrategy:
      steps: 3
      duration: 10s
      factor: 2
      jitter: 0.1
```



## Example 3: example-3.yaml


```yaml
apiVersion: argoproj.io/v1alpha1
kind: EventBus
metadata:
  name: default
spec:
  jetstream:
    version: "2.9.11"
    persistence:
      accessMode: ReadWriteOnce
      storageClassName: standard
      volumeSize: 10Gi
    replicas: 3
```



