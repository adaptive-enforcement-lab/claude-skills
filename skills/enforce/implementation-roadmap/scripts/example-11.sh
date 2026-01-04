# Verify branch protection
gh api repos/org/repo/branches/main/protection

# Sample March PRs
gh api 'repos/org/repo/pulls?state=closed&base=main' \
  --jq '.[] | select(.merged_at | startswith("2025-03"))'

# Check signature coverage
./scripts/signature-coverage.sh 2025-03-01 2025-04-01