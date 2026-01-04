gh api repos/org/repo/branches/main/protection \
  | jq '{reviews: .required_pull_request_reviews, admins: .enforce_admins}'