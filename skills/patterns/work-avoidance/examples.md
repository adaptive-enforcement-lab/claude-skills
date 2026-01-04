---
name: work-avoidance - Examples
description: Code examples for Work Avoidance
---

# Work Avoidance - Examples


## Example 1: example-1.mermaid


```mermaid
flowchart LR
    subgraph trigger[Trigger]
        Event[Event Received]
    end

    subgraph detect[Detection]
        Check{Work Needed?}
    end

    subgraph action[Action]
        Skip[Skip]
        Execute[Execute]
    end

    Event --> Check
    Check -->|No| Skip
    Check -->|Yes| Execute

    %% Ghostty Hardcore Theme
    style Event fill:#65d9ef,color:#1b1d1e
    style Check fill:#fd971e,color:#1b1d1e
    style Skip fill:#5e7175,color:#f8f8f3
    style Execute fill:#a7e22e,color:#1b1d1e
```



## Example 2: example-2.yaml


```yaml
- name: Check for meaningful changes
  id: check
  run: |
    # Strip version line before comparing
    strip_version() {
      sed '/^version:.*# x-release-please-version$/d' "$1"
    }

    SOURCE=$(strip_version "source/CONFIG.md")
    TARGET=$(git show HEAD:CONFIG.md 2>/dev/null | \
      sed '/^version:.*# x-release-please-version$/d' || echo "")

    if [ "$SOURCE" = "$TARGET" ]; then
      echo "skip=true" >> $GITHUB_OUTPUT
    else
      echo "skip=false" >> $GITHUB_OUTPUT
    fi

- name: Distribute file
  if: steps.check.outputs.skip != 'true'
  run: ./distribute.sh
```



