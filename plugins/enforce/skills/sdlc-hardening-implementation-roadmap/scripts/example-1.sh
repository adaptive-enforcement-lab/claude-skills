# Test pre-commit hook blocks secrets
echo "AKIAIOSFODNN7EXAMPLE" > .env && git add .env && git commit -m "test"
# Expected: Commit blocked by TruffleHog

# Test admin enforcement
gh api repos/org/repo/branches/main/protection | jq '.enforce_admins.enabled'
# Expected: true