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