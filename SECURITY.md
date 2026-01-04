# Security Policy

## Supported Versions

| Version | Supported          |
| ------- | ------------------ |
| Latest  | :white_check_mark: |

## Reporting a Vulnerability

**Do not open public issues for security vulnerabilities.**

To report a security vulnerability:

1. **Email**: Send details to [security contact needed]
2. **Response Time**: Expect acknowledgment within 48 hours
3. **Disclosure**: We follow coordinated disclosure practices

## What to Include

- Description of the vulnerability
- Steps to reproduce
- Potential impact
- Suggested fix (if available)

## Scope

This repository contains:
- Go code for skill generation
- Generated Claude Code skills (text content)
- GitHub Actions workflows

Security concerns include:
- Code injection in generator
- Unsafe file operations
- Workflow permission escalation
- Secrets exposure in generated content

## Out of Scope

- Generated skill content quality (use GitHub Issues)
- Documentation errors (use GitHub Issues)
- Claude Code client behavior (report to Anthropic)

## Security Best Practices

When contributing:
- Never commit secrets or credentials
- Validate all file paths before operations
- Sanitize markdown content before processing
- Use least-privilege GitHub App permissions
- Review workflow permissions carefully
