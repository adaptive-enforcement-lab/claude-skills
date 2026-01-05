# Test CI blocks failing tests
echo "func TestFail(t *testing.T) { t.Fatal() }" >> main_test.go
git push origin feature-branch
# Expected: Merge blocked by CI failure

# Verify SBOM generation
gsutil ls gs://audit-evidence/sbom/$(date +%Y-%m-%d)/
# Expected: SBOM files for today's builds

# Verify SLSA provenance
gh release view vX.Y.Z --json assets | jq '.assets[].name' | grep intoto
# Expected: .intoto.jsonl file exists

# Verify evidence collection
gsutil ls gs://audit-evidence/2025-01/
# Expected: branch-protection.json, merged-prs.json