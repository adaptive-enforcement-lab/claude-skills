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