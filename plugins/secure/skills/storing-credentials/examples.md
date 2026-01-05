---
name: storing-credentials - Examples
description: Code examples for Storing Credentials
---

# Storing Credentials - Examples


## Example 1: example-1.mermaid


```mermaid
flowchart TD
    A["Where will credentials be used?"] --> B{"Execution<br/>environment?"}

    B -->|"GitHub Actions"| C["GitHub Secrets"]
    B -->|"External CI/CD"| D["Platform Secrets Manager"]
    B -->|"Kubernetes"| E["External Secrets Operator"]
    B -->|"Local Development"| F["Environment Variables<br/>+ Encrypted Vault"]

    C --> C1["Organization secrets<br/>(recommended)"]
    C --> C2["Repository secrets<br/>(single repo)"]
    C --> C3["Environment secrets<br/>(protected workflows)"]

    D --> D1["Jenkins Credentials"]
    D --> D2["GitLab CI Variables"]
    D --> D3["CircleCI Contexts"]

    E --> E1["AWS Secrets Manager"]
    E --> E2["HashiCorp Vault"]
    E --> E3["Google Secret Manager"]

    %% Ghostty Hardcore Theme
    style A fill:#515354,stroke:#ccccc7,stroke-width:2px,color:#ccccc7
    style B fill:#fd971e,stroke:#e6db74,stroke-width:2px,color:#1b1d1e
    style C fill:#a7e22e,stroke:#bded5f,stroke-width:2px,color:#1b1d1e
    style D fill:#65d9ef,stroke:#a3babf,stroke-width:2px,color:#1b1d1e
    style E fill:#9e6ffe,stroke:#9e6ffe,stroke-width:2px,color:#1b1d1e
    style F fill:#f92572,stroke:#ff669d,stroke-width:2px,color:#1b1d1e
    style C1 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style C2 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style C3 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style D1 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style D2 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style D3 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style E1 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style E2 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style E3 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
```



## Example 2: example-2.yaml


```yaml
jobs:
  deploy:
    environment: production
    steps:
      - name: Generate token
        uses: actions/create-github-app-token@v1
        with:
          app-id: ${{ secrets.CORE_APP_ID }}
          private-key: ${{ secrets.CORE_APP_PRIVATE_KEY }}
```



## Example 3: example-3.yaml


```yaml
- name: Generate token
  id: app-token
  uses: actions/create-github-app-token@v2
  with:
    app-id: ${{ secrets.CORE_APP_ID }}
    private-key: ${{ secrets.CORE_APP_PRIVATE_KEY }}

- name: Use token
  env:
    GITHUB_TOKEN: ${{ steps.app-token.outputs.token }}
  run: |
    gh api /repos/${{ github.repository }}/issues
```



## Example 4: example-4.yaml


```yaml
- name: Generate org-scoped token
  uses: actions/create-github-app-token@v2
  with:
    app-id: ${{ secrets.CORE_APP_ID }}
    private-key: ${{ secrets.CORE_APP_PRIVATE_KEY }}
    owner: my-organization
```



## Example 5: example-5.yaml


```yaml
- name: Generate multi-repo token
  uses: actions/create-github-app-token@v2
  with:
    app-id: ${{ secrets.CORE_APP_ID }}
    private-key: ${{ secrets.CORE_APP_PRIVATE_KEY }}
    repositories: |
      repo-one
      repo-two
      repo-three
```



