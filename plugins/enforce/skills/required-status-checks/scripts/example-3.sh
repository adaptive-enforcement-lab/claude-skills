# Get workflow runs for a PR
gh api repos/org/repo/actions/runs \
  --jq '.workflow_runs[] | select(.head_branch=="feature-branch") |
    {name: .name, conclusion, created_at}'