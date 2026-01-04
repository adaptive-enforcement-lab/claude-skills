#!/bin/bash
# enable-secret-scanning.sh
# Enable secret scanning and push protection for all org repos

ORG="your-org"
TOKEN="${GITHUB_TOKEN}"

# Get all repositories in organization
repos=$(gh api \
  --paginate \
  "/orgs/${ORG}/repos" \
  --jq '.[].name')

for repo in $repos; do
  echo "Enabling secret scanning for ${ORG}/${repo}..."

  # Enable secret scanning
  gh api \
    --method PATCH \
    "/repos/${ORG}/${repo}" \
    -f security_and_analysis[secret_scanning][status]=enabled \
    -f security_and_analysis[secret_scanning_push_protection][status]=enabled

  echo "✓ ${repo} configured"
done