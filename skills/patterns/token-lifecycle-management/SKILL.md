---
name: token-lifecycle-management
description: >-
  Token expiration handling, refresh strategies, and caching patterns for long-running workflows. Manage installation token lifecycle and rate limits in GitHub Actions.
---

# Token Lifecycle Management

## When to Use This Skill

Installation token lifecycle management enables:

- **Long-running workflows** - Multi-hour operations without interruption
- **Token refresh automation** - Automatic renewal before expiration
- **Rate limit optimization** - Efficient token usage across job matrices
- **Caching strategies** - Share tokens across concurrent jobs
- **Error recovery** - Graceful handling of expired tokens

> **Token Expiration**
>


    Installation tokens expire **exactly 1 hour after generation**. Plan refresh strategies for workflows exceeding 50 minutes to account for clock drift and API latency.

##



## Implementation


See the full implementation guide in the source documentation.


## Examples

See [examples.md](examples.md) for code examples.






## References

- [Source Documentation](https://adaptive-enforcement-lab.com/patterns/github-actions/)
- [AEL Patterns](https://adaptive-enforcement-lab.com/patterns/)
