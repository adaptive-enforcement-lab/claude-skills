# Example: API-based runner group configuration
gh api --method PUT \
  /orgs/ORG/actions/runner-groups/GROUP_ID \
  -f name='production-runners' \
  -f visibility='selected' \
  -F selected_repository_ids='[123,456]' \
  -f allows_public_repositories=false \
  -f restricted_to_workflows=true \
  -F selected_workflows='[".github/workflows/deploy.yml@refs/heads/main"]'