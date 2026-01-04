---
name: error-handling - Examples
description: Code examples for Error Handling
---

# Error Handling - Examples


## Example 1: example-1.mermaid


```mermaid
flowchart TD
    A["API Call"] --> B{"HTTP Status"}

    B -->|"200 OK"| C["Success"]
    B -->|"401 Unauthorized"| D["Token Expired"]
    B -->|"403 Forbidden"| E["Permission Error"]
    B -->|"429 Rate Limited"| F["Rate Limit"]
    B -->|"5xx Server Error"| G["Transient Error"]

    D --> D1["Refresh Token"]
    D1 --> D2["Retry Request"]
    D2 --> B

    E --> E1{"Installation<br/>Exists?"}
    E1 -->|"No"| E2["Install App"]
    E1 -->|"Yes"| E3["Grant Permissions"]
    E2 --> H["Configuration Required"]
    E3 --> H

    F --> F1["Check Headers"]
    F1 --> F2["Wait for Reset"]
    F2 --> B

    G --> G1["Exponential Backoff"]
    G1 --> G2{"Max Retries?"}
    G2 -->|"No"| B
    G2 -->|"Yes"| I["Fail"]

    %% Ghostty Hardcore Theme
    style A fill:#515354,stroke:#ccccc7,stroke-width:2px,color:#ccccc7
    style B fill:#fd971e,stroke:#e6db74,stroke-width:2px,color:#1b1d1e
    style C fill:#a7e22e,stroke:#a7e22e,stroke-width:2px,color:#1b1d1e
    style D fill:#f92572,stroke:#ff669d,stroke-width:2px,color:#1b1d1e
    style E fill:#f92572,stroke:#ff669d,stroke-width:2px,color:#1b1d1e
    style F fill:#9e6ffe,stroke:#9e6ffe,stroke-width:2px,color:#1b1d1e
    style G fill:#66d9ef,stroke:#66d9ef,stroke-width:2px,color:#1b1d1e
    style I fill:#f92572,stroke:#ff669d,stroke-width:2px,color:#1b1d1e
    style H fill:#fd971e,stroke:#e6db74,stroke-width:2px,color:#1b1d1e
    style D1 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style D2 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style E1 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style E2 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style E3 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style F1 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style F2 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style G1 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style G2 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
```



## Example 2: example-2.yaml


```yaml
- name: Generate token
  id: app_token
  uses: actions/create-github-app-token@v2
  with:
    app-id: ${{ secrets.CORE_APP_ID }}
    private-key: ${{ secrets.CORE_APP_PRIVATE_KEY }}
    owner: adaptive-enforcement-lab
  continue-on-error: true

- name: Check token generation
  if: steps.app_token.outcome == 'failure'
  run: |
    echo "::error::Token generation failed"
    echo "::error::Check:"
    echo "  - App ID is correct: ${{ secrets.CORE_APP_ID != '' }}"
    echo "  - Private key is configured"
    echo "  - App is installed on: adaptive-enforcement-lab"
    echo "  - Installation is not suspended"
    exit 1
```



## Example 3: example-3.yaml


```yaml
- name: API call with expiration handling
  env:
    GH_TOKEN: ${{ steps.app_token.outputs.token }}
  run: |
    # Capture both stdout and stderr
    if ! response=$(gh api user --jq .login 2>&1); then
      if echo "$response" | grep -q "401\|Bad credentials"; then
        echo "::error::Token expired or invalid"
        echo "::error::Token age may exceed 1 hour"
        exit 1
      else
        echo "::error::API call failed: $response"
        exit 1
      fi
    fi

    echo "Authenticated as: $response"
```



## Example 4: example-4.yaml


```yaml
- name: API call with auto-refresh on expiration
  env:
    APP_ID: ${{ secrets.CORE_APP_ID }}
    PRIVATE_KEY: ${{ secrets.CORE_APP_PRIVATE_KEY }}
  run: |
    # Function to generate fresh token
    generate_token() {
      gh api /app/installations \
        --jq '.[0].id' | xargs -I {} \
        gh api /app/installations/{}/access_tokens \
        -X POST --jq .token
    }

    # Function to call API with auto-refresh on 401
    api_call_with_refresh() {
      local endpoint="$1"
      local max_attempts=2
      local attempt=1

      while [ $attempt -le $max_attempts ]; do
        # Attempt API call
        if response=$(gh api "$endpoint" 2>&1); then
          echo "$response"
          return 0
        fi

        # Check if error is 401 (expired token)
        if echo "$response" | grep -q "401\|Bad credentials"; then
          if [ $attempt -lt $max_attempts ]; then
            echo "::warning::Token expired, refreshing (attempt $attempt/$max_attempts)"

            # Refresh token
            export GH_TOKEN=$(generate_token)
            echo "::notice::Token refreshed successfully"

            ((attempt++))
            sleep 2
          else
            echo "::error::Failed to refresh token after $max_attempts attempts"
            return 1
          fi
        else
          # Non-401 error - fail immediately
          echo "::error::API call failed: $response"
          return 1
        fi
      done
    }

    # Initial token
    export GH_TOKEN=$(generate_token)

    # Make API calls with auto-refresh
    api_call_with_refresh "user"
    api_call_with_refresh "orgs/adaptive-enforcement-lab/repos"
```



## Example 5: example-5.yaml


```yaml
- name: Operation with permission validation
  env:
    GH_TOKEN: ${{ steps.app_token.outputs.token }}
  run: |
    endpoint="/repos/adaptive-enforcement-lab/example-repo/collaborators"

    # Attempt operation and capture error
    if ! response=$(gh api "$endpoint" 2>&1); then
      if echo "$response" | grep -q "403\|Forbidden"; then
        echo "::error::Permission denied for: $endpoint"
        echo "::error::Required permissions:"
        echo "  - App permission: 'members' (read)"
        echo "  - Installation scope: 'adaptive-enforcement-lab/example-repo'"
        echo ""
        echo "::error::Verify app configuration at:"
        echo "  https://github.com/organizations/adaptive-enforcement-lab/settings/apps"
        exit 1
      else
        echo "::error::API call failed: $response"
        exit 1
      fi
    fi

    echo "$response"
```



## Example 6: example-6.yaml


```yaml
- name: Diagnose permission error
  if: failure()
  env:
    GH_TOKEN: ${{ steps.app_token.outputs.token }}
  run: |
    echo "::group::Diagnostic Information"

    # Check token validity
    echo "Token status:"
    if gh api user --jq '.login' 2>/dev/null; then
      echo "  ✅ Token is valid"
    else
      echo "  ❌ Token is invalid or expired"
    fi

    # Check installation access
    echo ""
    echo "Installation scope:"
    gh api /app/installations \
      --jq '.[] | "  - \(.account.login) (ID: \(.id))"'

    # Attempt to identify missing permission
    echo ""
    echo "::error::Common 403 causes:"
    echo "  1. App lacks required repository/organization permissions"
    echo "  2. Installation doesn't include target repository"
    echo "  3. Repository is private but app has 'public_only' access"
    echo "  4. Organization requires approval for app installation"

    echo "::endgroup::"
```



