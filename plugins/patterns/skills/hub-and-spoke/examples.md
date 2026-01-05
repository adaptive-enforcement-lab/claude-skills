---
name: hub-and-spoke - Examples
description: Code examples for Hub and Spoke
---

# Hub and Spoke - Examples


## Example 1: example-1.mermaid


```mermaid
flowchart LR
    A[Orchestrator] --> B[Task 1]
    B --> C[Task 2]
    C --> D[Task 3]
    D --> E[Task 4]

    %% Ghostty Hardcore Theme
    style A fill:#65d9ef,color:#1b1d1e
    style B fill:#fd971e,color:#1b1d1e
    style C fill:#fd971e,color:#1b1d1e
    style D fill:#fd971e,color:#1b1d1e
    style E fill:#fd971e,color:#1b1d1e
```



## Example 2: example-2.mermaid


```mermaid
flowchart TD
    Hub[Hub Orchestrator]
    Hub --> S1[Spoke 1]
    Hub --> S2[Spoke 2]
    Hub --> S3[Spoke 3]
    Hub --> S4[Spoke 4]
    S1 --> Hub
    S2 --> Hub
    S3 --> Hub
    S4 --> Hub

    %% Ghostty Hardcore Theme
    style Hub fill:#9e6ffe,color:#1b1d1e
    style S1 fill:#a7e22e,color:#1b1d1e
    style S2 fill:#a7e22e,color:#1b1d1e
    style S3 fill:#a7e22e,color:#1b1d1e
    style S4 fill:#a7e22e,color:#1b1d1e
```



## Example 3: example-3.yaml


```yaml
apiVersion: argoproj.io/v1alpha1
kind: WorkflowTemplate
metadata:
  name: hub-orchestrator
spec:
  entrypoint: hub
  templates:
    - name: hub
      inputs:
        parameters:
          - name: repositories
      steps:
        # Discover work
        - - name: discover
            template: get-repositories

        # Fan out to spokes
        - - name: process-repo
            template: spawn-spoke
            arguments:
              parameters:
                - name: repo
                  value: "{{item}}"
            withParam: "{{steps.discover.outputs.result}}"

        # Collect results
        - - name: summarize
            template: collect-results

    - name: spawn-spoke
      inputs:
        parameters:
          - name: repo
      resource:
        action: create
        manifest: |
          apiVersion: argoproj.io/v1alpha1
          kind: Workflow
          metadata:
            generateName: spoke-{{inputs.parameters.repo}}-
          spec:
            workflowTemplateRef:
              name: spoke-worker
            arguments:
              parameters:
                - name: repository
                  value: "{{inputs.parameters.repo}}"
```



## Example 4: example-4.yaml


```yaml
apiVersion: argoproj.io/v1alpha1
kind: WorkflowTemplate
metadata:
  name: spoke-worker
spec:
  entrypoint: process
  arguments:
    parameters:
      - name: repository
  templates:
    - name: process
      inputs:
        parameters:
          - name: repository
      container:
        image: gcr.io/project/worker:v1
        command: ["/app/worker"]
        args:
          - "--repo={{inputs.parameters.repository}}"
          - "--action=process"
```



