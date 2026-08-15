# Phase 1: Foundation (Weeks 1-4)

Source: https://adaptive-enforcement-lab.com/enforce/implementation-roadmap/hardening-checklist/phase-1/

Establish local development controls and repository protection. These controls prevent bad code from ever entering git history.

> **Real-World Impact**
>
> A fintech client deployed pre-commit hooks across 200 repositories in 2 weeks. Within 48 hours, the hooks blocked 14 attempted commits containing AWS keys, GCP service account tokens, and database credentials. None entered git history.
>

---

## Phase Overview

Phase 1 establishes the foundation of SDLC security through two critical control layers:

1. **[Pre-commit Hooks](pre-commit-hooks.md)** - Block bad code locally before git commit
2. **[Branch Protection](branch-protection.md)** - Prevent unauthorized merges at repository level

These controls work together to create defense-in-depth at the source code level.

---

## Phase Components

### Pre-commit Hooks

Local enforcement that prevents secrets, policy violations, and code quality issues from entering git history.

**Key Controls**:

- Secrets detection with TruffleHog
- YAML/JSON validation
- Language-specific linting (Go, Python, etc.)
- Custom policy enforcement hooks
- Organization-wide distribution

**[View Pre-commit Hooks Details →](pre-commit-hooks.md)**

---

### Branch Protection Rules

Repository-level enforcement that makes it impossible to merge without meeting security criteria.

**Key Controls**:

- Required pull request reviews
- Code owner approval requirements
- Required CI status checks
- Administrator enforcement (no bypasses)
- Commit signature requirements
- Force push and deletion prevention

**[View Branch Protection Details →](branch-protection.md)**

---

## Phase 1 Validation Checklist

Before moving to Phase 2, verify all foundation controls work:

- [ ] Pre-commit hooks block secrets in test commit
- [ ] Pre-commit hooks block invalid YAML/JSON
- [ ] Pre-commit hooks enforce language-specific linting
- [ ] All repositories have `.pre-commit-config.yaml`
- [ ] Branch protection requires at least 1 review
- [ ] Branch protection dismisses stale reviews
- [ ] Branch protection requires code owner approval
- [ ] Branch protection enforces admins (`enforce_admins: true`)
- [ ] Branch protection requires signed commits
- [ ] Branch protection blocks force pushes and deletions
- [ ] Required CI status checks are configured (tests, lint, security)
- [ ] Organization-wide branch protection script runs successfully

---

## Validation Commands

Test that controls are working:

```bash
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
```

---

## Next Steps

With Phase 1 complete, you have:

- Pre-commit hooks blocking secrets and lint violations locally
- Branch protection preventing unauthorized merges
- Required reviews and signatures on all commits
- Automated distribution ensuring organization-wide coverage

**[Proceed to Phase 2: Automation →](../phase-2/index.md)**

Phase 2 builds on this foundation by adding CI/CD gates, SBOM generation, vulnerability scanning, and automated evidence collection.

---

## Related Patterns

- **[Pre-commit Security Gates](../../../pre-commit-hooks/pre-commit-hooks.md)** - Detailed hook configuration
- **[Branch Protection Enforcement](../../../branch-protection/branch-protection.md)** - GitHub API automation
- **[Implementation Roadmap Overview](index.md)** - Complete roadmap
- **[Phase 2: Automation →](../phase-2/index.md)** - CI/CD gates

---

*Pre-commit hooks deployed. Secrets blocked at source. Branch protection enforced. Admins cannot bypass. Foundation is set.*
