---
name: patterns - Examples
description: Code examples for Patterns
---

# Patterns - Examples


## Example 1: example-1.yaml


```yaml
# BAD: Creates duplicate issues on retry
- name: Create issue
  run: gh issue create --title "Alert" --body "Problem detected"
```



## Example 2: example-2.yaml


```yaml
# GOOD: Check before creating
- name: Create issue if not exists
  run: |
    existing=$(gh issue list --search "Alert" --state all --json number -q '.[0].number')
    if [ -z "$existing" ]; then
      gh issue create --title "Alert" --body "Problem detected"
    fi
```



## Example 3: example-3.yaml


```yaml
# BAD: Always creates PR even if no changes
- name: Create PR
  run: gh pr create --fill
```



## Example 4: example-4.yaml


```yaml
# GOOD: Check if changes exist first
- name: Create PR if changes exist
  run: |
    if git diff --quiet HEAD origin/main; then
      echo "No changes, skipping PR"
    else
      gh pr create --fill
    fi
```



## Example 5: example-5.yaml


```yaml
# BAD: Deploy first, check permissions later
- name: Deploy
  run: kubectl apply -f manifests/
- name: Check RBAC
  run: kubectl auth can-i create deployments
```



## Example 6: example-6.yaml


```yaml
# GOOD: Validate before deploying
- name: Prerequisite checks
  run: |
    kubectl auth can-i create deployments || exit 1
    kubectl get namespace production || exit 1
- name: Deploy
  run: kubectl apply -f manifests/
```



