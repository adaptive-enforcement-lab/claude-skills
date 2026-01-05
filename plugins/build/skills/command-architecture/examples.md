---
name: command-architecture - Examples
description: Code examples for Command Architecture
---

# Command Architecture - Examples


## Example 1: example-1.mermaid


```mermaid
graph LR
    Start[Start] --> Check[check]
    Check -->|Cache Valid| Skip[Skip Rebuild]
    Check -->|Cache Invalid| Rebuild[rebuild]
    Rebuild --> Select[select]
    Select --> Restart[restart]
    Restart --> Done[Done]
    Skip --> Done

    %% Ghostty Hardcore Theme
    style Start fill:#5e7175,color:#f8f8f3
    style Check fill:#fd971e,color:#1b1d1e
    style Skip fill:#65d9ef,color:#1b1d1e
    style Rebuild fill:#65d9ef,color:#1b1d1e
    style Select fill:#65d9ef,color:#1b1d1e
    style Restart fill:#65d9ef,color:#1b1d1e
    style Done fill:#a7e22e,color:#1b1d1e
```



## Example 2: example-2.text


```text
myctl
├── orchestrate          # Main workflow
├── check                # Cache status
├── rebuild              # Force cache rebuild
├── select               # List deployments
├── restart              # Restart deployments
├── version              # Show version info
└── completion           # Shell completion scripts
    ├── bash
    ├── zsh
    └── fish
```



