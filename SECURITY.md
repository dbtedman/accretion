# [Accretion](https://github.com/dbtedman/accretion) Security Policy

[![SAST GitHub Pipeline](https://img.shields.io/github/actions/workflow/status/dbtedman/accretion/sast.yml?branch=main&style=for-the-badge&logo=github&label=sast)](https://github.com/dbtedman/accretion/actions/workflows/sast.yml)
[![OpenSSF Scorecard](https://api.securityscorecards.dev/projects/github.com/dbtedman/accretion/badge?style=for-the-badge)](https://api.securityscorecards.dev/projects/github.com/dbtedman/accretion)

Outlines how security is considered during the development of Accretion.

-   [Dependency Vulnerability Scanning](#dependency-vulnerability-scanning)
-   [Security Disclosure Policy](#security-disclosure-policy)
-   [Security Update Policy](#security-update-policy)
-   [Security Related Configuration](#security-related-configuration)
-   [Known Security Gaps and Future Enhancements](#known-security-gaps-and-future-enhancements)

## Dependency Vulnerability Scanning

-   [GitHub code scanning (docs.github.com)](https://docs.github.com/en/code-security/code-scanning/automatically-scanning-your-code-for-vulnerabilities-and-errors/about-code-scanning) via [CodeQL (codeql.github.com)](https://codeql.github.com) and [Snyk (snyk.io)](https://snyk.io)
-   [GitHub dependabot alerts (docs.github.com)](https://docs.github.com/en/code-security/dependabot/dependabot-alerts/about-dependabot-alerts)
-   [GitHub dependabot security updates (docs.github.com)](https://docs.github.com/en/code-security/dependabot/dependabot-security-updates/about-dependabot-security-updates)
-   [GitHub secret scanning (docs.github.com)](https://docs.github.com/en/code-security/secret-scanning/about-secret-scanning)
-   [OSV Scanning (osv.dev)](https://osv.dev/)
-   [Snyk code analysis (snyk.io)](https://snyk.io/product/snyk-code/)
-   [Snyk container scanning (snyk.io)](https://snyk.io/product/container-vulnerability-management/)
-   [Snyk dependency scanning (snyk.io)](https://snyk.io/product/open-source-security-management/)

## Security Disclosure Policy

Privately report a vulnerability using [GitHub Security Advisories (github.com)](https://github.com/dbtedman/accretion/security/advisories).

## Security Update Policy

Best efforts will be taken to apply code fixes or update vulnerable packages as soon as is possible, this will usually be within a couple of days.

## Security Related Configuration

None currently.

## Known Security Gaps and Future Enhancements

Look at [GitHub issues tagged **Security** (github.com)](https://github.com/dbtedman/accretion/labels/security).
