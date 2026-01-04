# Require 2 reviews, passing tests, and commit signatures
gh api repos/org/repo/branches/main/protection \
  --method PUT \
  --field required_pull_request_reviews[required_approving_review_count]=2 \
  --field required_status_checks[strict]=true \
  --field required_status_checks[contexts][]=test \
  --field required_status_checks[contexts][]=security-scan \
  --field required_signatures[enabled]=true