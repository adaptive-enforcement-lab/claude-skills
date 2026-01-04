---
name: github-token-permissions-overview - Troubleshooting
description: Troubleshooting guide for GITHUB_TOKEN Permissions Overview
---

# GITHUB_TOKEN Permissions Overview - Troubleshooting

**"Resource not accessible by integration"**: Add missing permission to `permissions` block.

**"Must have admin access to organization"**: Use GitHub App with org-level permissions instead of GITHUB_TOKEN.

**Token works locally but fails in Actions**: Personal tokens have broader scope than GITHUB_TOKEN. Adjust workflow permissions or use GitHub App.
