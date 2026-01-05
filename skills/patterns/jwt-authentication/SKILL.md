---
name: jwt-authentication
description: >-
  Generate JWTs for GitHub App authentication. Direct JWT generation for app-level operations, installation discovery, and bootstrapping workflows.
---

# JWT Authentication

## When to Use This Skill

JWTs authenticate your GitHub App itself, not a specific installation. They enable:

- **Installation discovery** - List where your app is installed
- **App metadata retrieval** - Get app configuration and manifest
- **Installation management** - Suspend or configure installations
- **Bootstrap workflows** - Generate installation tokens dynamically

> **JWT Limitations**
>
>
> - Cannot access repository contents
> - Cannot create issues, pull requests, or commits
> - 10-minute expiration (maximum allowed)
> - App-level permissions only


## Implementation

*See [examples.md](examples.md) for detailed code examples.*


## Techniques


### JWT Generation Methods

### Method 1: GitHub CLI (Recommended for Workflows)

The GitHub CLI handles JWT generation automatically when using GitHub App credentials.

```yaml
jobs:
  list-installations:
    runs-on: ubuntu-latest
    steps:
      - name: List app installations
        env:
          GH_APP_ID: ${{ secrets.CORE_APP_ID }}
          GH_APP_PRIVATE_KEY: ${{ secrets.CORE_APP_PRIVATE_KEY }}
        run: |
          # gh CLI generates JWT automatically
          gh api /app/installations \
            --jq '.[] | {id: .id, account: .account.login}'
```

**How it works**:

- `GH_APP_ID` + `GH_APP_PRIVATE_KEY` triggers automatic JWT generation
- JWT is generated on-demand for each API call
- No manual token handling required

### Method 2: Manual JWT Generation (Advanced)

For custom implementations or languages without GitHub CLI support.

```yaml
jobs:
  manual-jwt:
    runs-on: ubuntu-latest
    steps:
      - name: Generate JWT manually
        id: jwt
        env:
          APP_ID: ${{ secrets.CORE_APP_ID }}
          PRIVATE_KEY: ${{ secrets.CORE_APP_PRIVATE_KEY }}
        run: |
          # Install JWT tool
          npm install -g jsonwebtoken

          # Create JWT generation script
          cat > generate-jwt.js << 'EOF'
          const jwt = require('jsonwebtoken');
          const fs = require('fs');

          const appId = process.env.APP_ID;
          const privateKey = process.env.PRIVATE_KEY;

          const now = Math.floor(Date.now() / 1000);
          const payload = {
            iat: now - 60,        // Issued 60 seconds in past
            exp: now + (10 * 60), // Expires in 10 minutes
            iss: appId
          };

          const token = jwt.sign(payload, privateKey, { algorithm: 'RS256' });
          console.log(token);
          EOF

          # Generate JWT
          JWT_TOKEN=$(node generate-jwt.js)
          echo "::add-mask::$JWT_TOKEN"
          echo "token=$JWT_TOKEN" >> $GITHUB_OUTPUT

      - name: Use JWT
        env:
          GITHUB_TOKEN: ${{ steps.jwt.outputs.token }}
        run: |
          curl -H "Authorization: Bearer $GITHUB_TOKEN" \
               -H "Accept: application/vnd.github+json" \
               https://api.github.com/app
```

> **Security: Mask JWT Token**
>
>
> Always use `echo "::add-mask::$JWT_TOKEN"` to prevent token exposure in logs.
>

### Method 3: Python Implementation

For Python-based workflows and automation.

```yaml
jobs:
  python-jwt:
    runs-on: ubuntu-latest
    steps:
      - name: Generate JWT with Python
        id: jwt
        env:
          APP_ID: ${{ secrets.CORE_APP_ID }}
          PRIVATE_KEY: ${{ secrets.CORE_APP_PRIVATE_KEY }}
        run: |
          pip install PyJWT cryptography

          python << 'EOF'
          import jwt
          import time
          import os

          app_id = os.environ['APP_ID']
          private_key = os.environ['PRIVATE_KEY']

          now = int(time.time())
          payload = {
              'iat': now - 60,
              'exp': now + (10 * 60),
              'iss': app_id
          }

          token = jwt.encode(payload, private_key, algorithm='RS256')

          # Mask token in logs
          print(f"::add-mask::{token}")

          # Output token
          with open(os.environ['GITHUB_OUTPUT'], 'a') as f:
              f.write(f"token={token}\n")
          EOF

      - name: Use JWT
        env:
          GITHUB_TOKEN: ${{ steps.jwt.outputs.token }}
        run: |
          curl -H "Authorization: Bearer $GITHUB_TOKEN" \
               https://api.github.com/app/installations
```

*See [reference.md](reference.md) for additional techniques and detailed examples.*


## Comparison

*See [examples.md](examples.md) for detailed code examples.*


## Examples

See [examples.md](examples.md) for code examples.


## Full Reference

See [reference.md](reference.md) for complete documentation.
## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/github-actions/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
