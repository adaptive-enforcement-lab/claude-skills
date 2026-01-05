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