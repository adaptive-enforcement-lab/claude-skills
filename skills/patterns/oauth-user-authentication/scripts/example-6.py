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