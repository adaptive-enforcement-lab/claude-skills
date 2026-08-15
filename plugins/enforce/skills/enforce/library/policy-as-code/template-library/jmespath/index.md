# JMESPath for Kyverno

Source: https://adaptive-enforcement-lab.com/enforce/policy-as-code/template-library/jmespath/

Master JMESPath to unlock advanced Kyverno policy capabilities. Query nested JSON structures, build complex validation logic, and enforce standards that simple pattern matching cannot express.

> **What You'll Learn**
>
> JMESPath extends Kyverno beyond basic pattern matching. Use it for cross-field validation, dynamic conditions, array filtering, and string transformations. Essential for enterprise-grade policy enforcement.
>

---

## Why JMESPath Matters

**Simple pattern matching fails when you need to:**

- Compare multiple fields (requests vs limits)
- Validate conditionally (if label exists, require annotation)
- Parse and transform data (extract image tags)
- Filter arrays dynamically (containers with specific images)
- Build complex boolean logic across nested structures

**JMESPath solves all of this.** It's a query language for JSON, purpose-built for navigating Kubernetes resources and extracting validation data.

---

## Documentation Structure

### Getting Started

**[JMESPath Patterns](patterns.md)**
Core patterns for common Kyverno use cases. Start here if you're new to JMESPath in policies.

- Projection and filtering
- Cross-field validation
- Array operations
- String transformations
- Boolean logic

### Advanced Techniques

**[Advanced Patterns](advanced.md)**
Sophisticated validation logic for complex scenarios.

- Label and annotation transformations
- Multi-level array operations
- Dynamic naming validation
- Conditional enforcement based on metadata

**[Enterprise Supply Chain](enterprise-supply-chain.md)**
Supply chain security patterns for production workloads.

- Image signature validation
- Provenance verification
- SBOM enforcement
- Attestation checks

**[Enterprise Patterns](enterprise.md)**
Production-grade policies for enterprise Kubernetes.

- Multi-cluster enforcement
- Compliance validation
- Resource governance
- Security hardening

### Reference Material

**[Function Reference](reference.md)**
Complete JMESPath function library for Kyverno.

- String functions (split, join, contains)
- Array operations (map, filter, sort)
- Comparison operators
- Logical expressions

**[Testing Guide](testing.md)**
Test JMESPath expressions before deploying policies.

- `kyverno jp` CLI usage
- Test case development
- Debugging patterns
- Common mistakes and fixes

---

## Quick Start

**Install Kyverno CLI for testing:**

```bash
# Install kyverno CLI
brew install kyverno/kyverno/kyverno

# Test JMESPath expression
kyverno jp query -i manifest.yaml 'spec.template.spec.containers[*].name'
```

**Simple validation example:**

```yaml
apiVersion: kyverno.io/v1
kind: ClusterPolicy
metadata:
  name: require-resource-limits
spec:
  validationFailureAction: Enforce
  rules:
  - name: validate-limits
    match:
      any:
      - resources:
          kinds:
          - Pod
    validate:
      message: "All containers must define resource limits"
      deny:
        conditions:
          any:
          - key: "{{ request.object.spec.containers[?!resources.limits.memory].name | length(@) }}"
            operator: GreaterThan
            value: 0
```

**What this does:**

- Filters containers without memory limits: `containers[?!resources.limits.memory]`
- Extracts their names: `.name`
- Counts them: `| length(@)`
- Denies if count > 0

---

## When to Use JMESPath

**Use JMESPath when:**

- Pattern matching can't express your logic
- You need conditionals or transformations
- Validation depends on multiple fields
- You're filtering or comparing arrays

**Skip JMESPath when:**

- Simple pattern matching works (`pattern`, `anyPattern`)
- You're only checking field existence
- No cross-field validation needed

> **Test Before Deploying**
>
> Always test JMESPath expressions with `kyverno jp` before adding them to policies. Syntax errors fail silently in audit mode and block resources in enforce mode.
>

---

## Learning Path

**Beginner:**

1. Read [JMESPath Patterns](patterns.md) - core techniques
2. Use [Testing Guide](testing.md) - validate your expressions
3. Reference [Function Reference](reference.md) - lookup syntax

**Intermediate:**

1. Study [Advanced Patterns](advanced.md) - complex scenarios
2. Apply [Enterprise Patterns](enterprise.md) - production use cases

**Advanced:**

1. Implement [Enterprise Supply Chain](enterprise-supply-chain.md) - security hardening
2. Build custom patterns for your environment

---

## External Resources

- [JMESPath Official Tutorial](https://jmespath.org/tutorial.html) - language fundamentals
- [Kyverno JMESPath Documentation](https://kyverno.io/docs/writing-policies/jmespath/) - policy-specific usage
- [JMESPath Playground](https://jmespath.org/) - interactive testing

---

**Next:** Start with [JMESPath Patterns](patterns.md) to learn core techniques.
