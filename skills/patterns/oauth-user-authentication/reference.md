---
name: oauth-user-authentication - Reference
description: Complete reference for OAuth User Authentication
---

# OAuth User Authentication - Reference

This is the complete reference documentation extracted from the source.


# OAuth User Authentication

OAuth enables GitHub Apps to act on behalf of users, preserving user identity in audit logs and respecting user-level permissions. Use OAuth when operations must be attributed to users rather than automated workflows.

> **When to Use OAuth**
>
>
> OAuth is for **user-context operations only**. Use installation tokens for automation and JWT for app-level operations.
>

## Overview

OAuth authentication provides user-context access for GitHub Apps. It enables:

- **User attribution** - Actions appear as the user in audit logs
- **User permissions** - Respect individual user access levels
- **Personal repository access** - Access to user's private repositories
- **Interactive applications** - Web apps and CLI tools requiring user authorization
- **Long-lived sessions** - Tokens valid until revoked

> **OAuth Limitations**
>
>
> - Not suitable for automated workflows (no user present)
> - Requires user consent for each installation
> - Rate limits apply per user (5,000/hour)
> - More complex setup than installation tokens
>

## OAuth vs Other Methods

```mermaid
flowchart TD
    A["Need user context?"] --> B{"Who initiates<br/>the action?"}

    B -->|"Human user<br/>(web app, CLI)"| C["Use OAuth"]
    B -->|"Automated process<br/>(GitHub Actions)"| D["Use Installation Token"]

    C --> C1["User attribution required"]
    C --> C2["Personal repos access"]
    C --> C3["User-level permissions"]

    D --> D1["No user present"]
    D --> D2["Organization repos"]
    D --> D3["App-level permissions"]

    %% Ghostty Hardcore Theme
    style A fill:#515354,stroke:#ccccc7,stroke-width:2px,color:#ccccc7
    style B fill:#fd971e,stroke:#e6db74,stroke-width:2px,color:#1b1d1e
    style C fill:#a7e22e,stroke:#bded5f,stroke-width:2px,color:#1b1d1e
    style D fill:#f92572,stroke:#ff669d,stroke-width:2px,color:#1b1d1e
    style C1 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style C2 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style C3 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style D1 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style D2 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
    style D3 fill:#515354,stroke:#ccccc7,stroke-width:1px,color:#ccccc7
```

## OAuth Flow Types

GitHub Apps support two OAuth flows:

### Web Application Flow

For web applications with server-side backends.

**Characteristics**:

- User redirects to GitHub authorization page
- Server exchanges authorization code for token
- Secure token storage on server
- Suitable for web applications

### Device Flow

For CLI tools and applications without web browsers.

**Characteristics**:

- User enters code on GitHub website
- Device polls for authorization
- No redirect URI required
- Suitable for headless environments

## Web Application Flow

### Flow Diagram

```mermaid
sequenceDiagram

%% Ghostty Hardcore Theme
    participant U as User
    participant A as Your App
    participant G as GitHub
    participant R as Repository

    U->>A: Click "Login with GitHub"
    A->>A: Generate state parameter
    A->>U: Redirect to GitHub OAuth
    U->>G: Authorize application
    G->>U: Redirect with code
    U->>A: Return with code + state
    A->>A: Validate state
    A->>G: Exchange code for token
    G->>A: Return access token
    A->>A: Store token securely
    A->>R: API operations as user

    Note over U,R: Token valid until revoked

```

### Step 1: Direct User to GitHub

Generate authorization URL with required parameters.

```python
import secrets
import urllib.parse

# Generate state for CSRF protection
state = secrets.token_urlsafe(32)
# Store state in session for later validation

# Your GitHub App OAuth settings
client_id = "Iv1.your_client_id"
redirect_uri = "https://your-app.com/auth/callback"

# Authorization URL
params = {
    'client_id': client_id,
    'redirect_uri': redirect_uri,
    'state': state,
    'scope': 'repo user',  # Request needed scopes
}

auth_url = f"https://github.com/login/oauth/authorize?{urllib.parse.urlencode(params)}"

# Redirect user to auth_url
```

> **CSRF Protection Required**
>
>
> Always use the `state` parameter to prevent cross-site request forgery attacks. Generate a random value, store it in the user session, and validate it in the callback.
>

### Step 2: Handle Callback

Exchange authorization code for access token.

```python
import requests

def handle_oauth_callback(code, state, session_state):
    # Validate state parameter
    if state != session_state:
        raise ValueError("Invalid state parameter - possible CSRF attack")

    # Exchange code for token
    token_url = "https://github.com/login/oauth/access_token"

    payload = {
        'client_id': 'Iv1.your_client_id',
        'client_secret': 'your_client_secret',  # From GitHub App settings
        'code': code,
        'redirect_uri': 'https://your-app.com/auth/callback',
    }

    headers = {
        'Accept': 'application/json',
    }

    response = requests.post(token_url, json=payload, headers=headers)
    response.raise_for_status()

    token_data = response.json()

    return {
        'access_token': token_data['access_token'],
        'token_type': token_data['token_type'],
        'scope': token_data['scope'],
    }
```

> **Client Secret Security**
>
>
> - Never expose client secret in frontend code
> - Store in environment variables or secrets manager
> - Rotate regularly (every 90 days minimum)
> - Use separate secrets for development/production
>

### Step 3: Use Access Token

Make authenticated API requests as the user.

```python
def create_issue_as_user(access_token, repo_owner, repo_name, title, body):
    """Create GitHub issue with user attribution"""

    url = f"https://api.github.com/repos/{repo_owner}/{repo_name}/issues"

    headers = {
        'Authorization': f'Bearer {access_token}',
        'Accept': 'application/vnd.github+json',
        'X-GitHub-Api-Version': '2022-11-28',
    }

    payload = {
        'title': title,
        'body': body,
    }

    response = requests.post(url, json=payload, headers=headers)
    response.raise_for_status()

    return response.json()

# Usage
issue = create_issue_as_user(
    access_token=user_token,
    repo_owner='adaptive-enforcement-lab',
    repo_name='example-repo',
    title='User-created issue',
    body='This issue was created by the authenticated user via OAuth',
)

print(f"Created issue #{issue['number']} as {issue['user']['login']}")
```

### Complete Web Application Example

```python
from flask import Flask, redirect, request, session, url_for
import requests
import secrets

app = Flask(__name__)
app.secret_key = 'your-secret-key-here'  # Use secure secret in production

GITHUB_CLIENT_ID = 'Iv1.your_client_id'
GITHUB_CLIENT_SECRET = 'your_client_secret'
REDIRECT_URI = 'http://localhost:5000/callback'

@app.route('/')
def index():
    if 'github_token' in session:
        return f"""
        <h1>Authenticated!</h1>
        <p>Token: {session['github_token'][:20]}...</p>
        <a href="/create-issue">Create Test Issue</a> |
        <a href="/logout">Logout</a>
        """
    else:
        return '<a href="/login">Login with GitHub</a>'

@app.route('/login')
def login():
    # Generate and store state
    state = secrets.token_urlsafe(32)
    session['oauth_state'] = state

    # Build authorization URL
    params = {
        'client_id': GITHUB_CLIENT_ID,
        'redirect_uri': REDIRECT_URI,
        'state': state,
        'scope': 'repo user',
    }

    auth_url = f"https://github.com/login/oauth/authorize"
    return redirect(f"{auth_url}?{'&'.join(f'{k}={v}' for k, v in params.items())}")

@app.route('/callback')
def callback():
    # Validate state
    if request.args.get('state') != session.get('oauth_state'):
        return 'Invalid state parameter', 400

    # Exchange code for token
    code = request.args.get('code')

    token_response = requests.post(
        'https://github.com/login/oauth/access_token',
        json={
            'client_id': GITHUB_CLIENT_ID,
            'client_secret': GITHUB_CLIENT_SECRET,
            'code': code,
            'redirect_uri': REDIRECT_URI,
        },
        headers={'Accept': 'application/json'},
    )

    token_data = token_response.json()

    # Store token in session (use secure storage in production)
    session['github_token'] = token_data['access_token']

    return redirect(url_for('index'))

@app.route('/create-issue')
def create_issue():
    if 'github_token' not in session:
        return redirect(url_for('login'))

    # Create issue as authenticated user
    response = requests.post(
        'https://api.github.com/repos/adaptive-enforcement-lab/test-repo/issues',
        json={
            'title': 'Test Issue from OAuth',
            'body': 'Created via OAuth user authentication',
        },
        headers={
            'Authorization': f"Bearer {session['github_token']}",
            'Accept': 'application/vnd.github+json',
        },
    )

    issue = response.json()
    return f"Created issue #{issue['number']}"

@app.route('/logout')
def logout():
    session.clear()
    return redirect(url_for('index'))

if __name__ == '__main__':
    app.run(debug=True)
```

## Device Flow

