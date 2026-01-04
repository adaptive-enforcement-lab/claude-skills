---
name: sdlc-hardening-implementation-roadmap - Examples
description: Code examples for SDLC Hardening Implementation Roadmap
---

# SDLC Hardening Implementation Roadmap - Examples


## Example 1: example-1.sh


```bash
# Test pre-commit hook blocks secrets
echo "AKIAIOSFODNN7EXAMPLE" > .env && git add .env && git commit -m "test"
# Expected: Commit blocked by TruffleHog

# Test admin enforcement
gh api repos/org/repo/branches/main/protection | jq '.enforce_admins.enabled'
# Expected: true
```



## Example 2: example-2.sh


```bash
# Test CI blocks failing tests
echo "func TestFail(t *testing.T) { t.Fatal() }" >> main_test.go
git push origin feature-branch
# Expected: Merge blocked by CI failure

# Verify SBOM generation
gsutil ls gs://audit-evidence/sbom/$(date +%Y-%m-%d)/
# Expected: SBOM files for today's builds
```



## Example 3: example-3.sh


```bash
# Test pod without resource limits is rejected
kubectl apply -f pod-no-limits.yaml
# Expected: Admission webhook denies request

# Test untrusted registry is blocked
kubectl apply -f pod-dockerhub.yaml
# Expected: Image source validation fails
```



## Example 4: example-4.sh


```bash
# Verify evidence archive exists
gsutil ls gs://audit-evidence/2025-01/branch-protection.json
# Expected: File exists with branch protection config

# Check OpenSSF Scorecard score
docker run gcr.io/openssf/scorecard-action:stable --repo=github.com/org/repo
# Expected: Score ≥ 8.0/10
```



