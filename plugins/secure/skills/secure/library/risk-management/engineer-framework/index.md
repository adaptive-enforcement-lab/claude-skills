# Risk Prioritization Framework for Engineers

Source: https://adaptive-enforcement-lab.com/secure/risk-management/engineer-framework/

Making fast, defensible decisions about vulnerability remediation under pressure. This framework translates security metrics into actionable engineering decisions.

> **Key Insight**
>
> Risk = (Impact × Likelihood × Exploitability) - (Remediation Cost). Prioritize ruthlessly based on exposure, not noise.
>

## Overview

Most teams have an unlimited list of vulnerabilities but finite resources. The difference between effective security and security theater is how you make triage decisions.

This framework gives you:

1. **Objective metrics** to compare disparate vulnerabilities
2. **Decision trees** for patch-now vs patch-later choices
3. **Cost-benefit analysis** for remediation tradeoffs
4. **Real-world examples** with concrete decisions

The goal: **Spend your security budget where exposure is highest**.

## Framework Components

This framework is organized into focused modules:

### [Risk Assessment Matrix](risk-assessment.md)

Establish baseline risk across three dimensions:

- Impact scoring (1-4 scale)
- Likelihood assessment
- Exploitability evaluation
- Risk score calculation and interpretation

### [CVSS Score Interpretation](cvss-interpretation.md)

Translate CVSS scores to engineering decisions:

- CVSS 3.1 score ranges and thresholds
- Key components (Attack Vector, Complexity, Privileges)
- Real-world CVSS vector examples
- When CVSS doesn't tell the whole story

### [Exploitability Analysis](exploitability-analysis.md)

Determine if vulnerability is actually weaponized:

- Exploit maturity spectrum
- Public exploit databases
- Tools for checking exploit status
- Timeline from PoC to active exploitation

### [Blast Radius Calculation](blast-radius.md)

Calculate infrastructure impact:

- System coverage assessment
- Dependency mapping (direct and transitive)
- User and data exposure calculation
- Blast radius multipliers (0.2 to 5.0)

### [Decision Trees](decision-trees.md)

Fast, repeatable decision frameworks:

- Patch now vs. later decision tree
- Mitigate vs. accept vs. transfer decision tree
- Emergency vs. standard patching workflow
- Implementation checklists

### [Real-World Scenarios](real-world-scenarios.md)

Complete worked examples:

- Log4Shell (CVE-2021-44228)
- Node.js session vulnerability
- Kubernetes privilege escalation
- Transitive dependency challenges

### [Remediation Cost Analysis](remediation-cost.md)

Balance risk vs. effort:

- Cost calculation framework
- Priority scoring
- Metrics to track (MTTD, MTBP)
- Implementation checklists

## Quick Reference

### Severity Thresholds

| Risk Score | Label | Action Timeline |
|-----------|-------|----------------|
| 45+ | **CRITICAL** | 24 hours |
| 30-44 | **HIGH** | 1 week |
| 15-29 | **MEDIUM** | 30 days |
| 5-14 | **LOW** | Next maintenance window |
| <5 | **MINIMAL** | Opportunistic |

### Key Metrics

- **MTTD** (Mean Time to Detect): < 24 hours
- **MTBP** (Mean Time to Patch - Critical): < 4 hours
- **MTBP** (Mean Time to Patch - High): < 72 hours
- **MTBP** (Mean Time to Patch - Medium): < 30 days

## References

- [CVSS Specification](https://www.first.org/cvss/v3.1/specification-document)
- [NVD CVE Database](https://nvd.nist.gov/vuln)
- [OWASP Risk Rating](https://owasp.org/www-community/OWASP_Risk_Rating_Methodology)
- [NIST Security Controls](https://csrc.nist.gov/pubs/sp/800/53/r5/upd1/final)

---

*Risk prioritization is a skill. Practice making fast, defensible decisions under pressure.*
