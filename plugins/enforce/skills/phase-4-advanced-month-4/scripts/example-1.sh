# Verify evidence archive exists
gsutil ls gs://audit-evidence/2025-01/branch-protection.json
# Expected: File exists with branch protection config

# Check OpenSSF Scorecard score
docker run gcr.io/openssf/scorecard-action:stable --repo=github.com/org/repo
# Expected: Score ≥ 8.0/10

# Verify SLSA provenance
gh release view vX.Y.Z --json assets | jq '.assets[].name' | grep intoto
# Expected: .intoto.jsonl file exists

# Test evidence retrieval speed
time gsutil ls gs://audit-evidence/2024-*/branch-protection.json
# Expected: < 10 seconds