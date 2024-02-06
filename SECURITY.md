# [Accretion](https://github.com/dbtedman/accretion) Security Guide

Outlines how security is considered during the development of Accretion.

-   [Dependency Vulnerability and Code Scanning](#dependency-vulnerability-and-code-scanning)
-   [Security Disclosure Policy](#security-disclosure-policy)
-   [Security Update Policy](#security-update-policy)
-   [Security Related Configuration](#security-related-configuration)
-   [Known Security Gaps and Future Enhancements](#known-security-gaps-and-future-enhancements)

## Dependency Vulnerability and Code Scanning

🚧 Placeholder

## Security Disclosure Policy

🚧 Placeholder

## Security Update Policy

🚧 Placeholder

## Security Related Configuration

### Content Security Policy

A strict [Content Security Policy (CSP) (developer.mozilla.org)](https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP) has been defined to provide a layer of defence against injection attacks. Within this CSP, a per-request nonce is defined to whitelist allowed resources. This nonce is implemented using the [Argon2 Hashing Algorithm with recommended settings (cheatsheetseries.owasp.org)](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html#argon2id).

## Known Security Gaps and Future Enhancements

🚧 Placeholder
