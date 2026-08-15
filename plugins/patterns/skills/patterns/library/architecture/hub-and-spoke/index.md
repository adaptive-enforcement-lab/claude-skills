# Hub and Spoke

Source: https://adaptive-enforcement-lab.com/patterns/architecture/hub-and-spoke/

One hub coordinates. Many spokes execute. The hub doesn't do the work. It distributes, tracks, and summarizes.

This pattern scales horizontally. Add workers without touching the orchestrator.

## The Problem

Sequential processing doesn't scale:

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

Total time: sum of all tasks. Can't parallelize. Bottleneck at the orchestrator.

## The Pattern

> **Quick Start**
>
> This guide is part of a modular documentation set. Refer to related guides in the navigation for complete context.
>

Hub coordinates, spokes execute in parallel:

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

Total time: longest single task. Linear scaling. Hub unchanged as spokes grow.

## Argo Workflows Implementation

Hub workflow spawns children:

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

Hub discovers repositories, spawns a spoke workflow for each, then summarizes results.

## Spoke Worker Template

Each spoke is independent:

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

Spoke doesn't know about the hub. Just does its work and exits.
