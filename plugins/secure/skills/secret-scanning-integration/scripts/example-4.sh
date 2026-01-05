# Developer pushes commit with secret
git push origin main
# > Error: secret detected in commit abc123
# > To bypass, visit: https://github.com/org/repo/security/bypass/abc123

# Developer bypasses with justification
# GitHub logs bypass event

# Security team reviews bypasses
gh api /repos/org/repo/secret-scanning/push-protection-bypasses