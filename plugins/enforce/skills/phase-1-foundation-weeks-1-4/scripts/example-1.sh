# Test pre-commit secrets detection
echo "AWS_KEY=AKIAIOSFODNN7EXAMPLE" > .env
git add .env && git commit -m "test"
# Expected: Commit blocked by TruffleHog

# Verify branch protection admin enforcement
gh api repos/org/repo/branches/main/protection | jq '.enforce_admins.enabled'
# Expected: true

# Count repositories with protection
gh repo list org --limit 1000 --json name --jq '.[].name' | while read repo; do
  gh api repos/org/$repo/branches/main/protection >/dev/null 2>&1 && echo "✅"
done | grep "✅" | wc -l