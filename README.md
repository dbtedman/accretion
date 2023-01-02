# [Accretion](https://github.com/dbtedman/accretion)

[![CI GitHub Pipeline](https://img.shields.io/github/actions/workflow/status/dbtedman/accretion/ci.yml?branch=main&style=for-the-badge&logo=github&label=ci)](https://github.com/dbtedman/accretion/actions/workflows/ci.yml?query=branch%3Amain)
[![SAST GitHub Pipeline](https://img.shields.io/github/actions/workflow/status/dbtedman/accretion/sast.yml?branch=main&style=for-the-badge&logo=github&label=sast)](https://github.com/dbtedman/accretion/actions/workflows/sast.yml)

> ⚠️ WARNING! This project is in early development, and is not ready for production use.

Manage internal technical documentation that is enriched with live data accreted from your environment.

-   [Getting Started](#getting-started)
-   [Verification](#verification)
-   [Design](#design)
-   [Security](#security)
-   [References](#references)
-   [License](#license)

## Getting Started

```shell
nvm use && make
```

### Run Local Environment

```shell
make local
```

## Verification

### Linting

```shell
make lint
```

These rules can then be automatically applied:

```shell
make format
```

### Unit Testing

```shell
make test
```

## Design

### Repository Structure

| Directory           | Purpose                                          |
| ------------------- | ------------------------------------------------ |
| `.github/`          | Dependabot configuration and pipeline workflows. |
| `.husky/`           | Husky git hook configuration.                    |
| `api/`              |                                                  |
| `cmd/`              | Entrypoint to go application.                    |
| `internal/domain/`  | Core application logic.                          |
| `internal/gateway/` | Encapsulated access to external systems.         |
| `public/`           | Static assets included directly in ui.           |
| `ui/`               | Vue based user interface.                        |
| `web/`              |                                                  |

### Light and Dark Modes

The user interface will automatically switch between dark and light mode based on the operating systems chosen mode.

### Data

Primarily represented as a directed graph, persisted as a series of discrete events. Events allow for time traveling
though the data.

### Domain

(Proposed)

-   Source Code
    -   Git Repository
    -   Git Reference
    -   Git Commit
    -   Merge Request
-   Work Item
    -   Issue
    -   Board
-   External Dependency
-   Product

### Providers

(Proposed)

-   Bitbucket
-   Confluence
-   GitHub
-   GitLab
-   Jira
-   NPM

### Stack

-   [✅ Cobra (cobra.dev)](https://cobra.dev)
-   [✅ CodeQL (codeql.github.com)](https://codeql.github.com)
-   [✅ Dependabot (docs.github.com)](https://docs.github.com/en/code-security/dependabot)
-   [✅ Docker (docker.com)](https://www.docker.com)
-   [✅ Docker Compose (docs.docker.com)](https://docs.docker.com/compose/)
-   [✅ ESLint (eslint.org)](https://eslint.org/)
-   [✅ GitHub Actions (github.com)](https://github.com/features/actions)
-   [✅ GitHub Security (github.com)](https://github.com/security)
-   [✅ Go (go.dev)](https://go.dev)
-   [✅ Go Embed (pkg.go.dev)](https://pkg.go.dev/embed)
-   [✅ Go Modules (go.dev)](https://go.dev/ref/mod)
-   [✅ Happy Dom (github.com)](https://github.com/capricorn86/happy-dom)
-   [✅ Husky (typicode.github.io)](https://typicode.github.io/husky/#/)
-   [✅ Node Version Manager (github.com)](https://github.com/nvm-sh/nvm)
-   [✅ NodeJS (nodejs.org)](https://nodejs.org/en/)
-   [✅ OSV Scanning (osv.dev)](https://osv.dev/)
-   [✅ PNPM (pnpm.io)](https://pnpm.io/)
-   [✅ Pinia (pinia.vuejs.org)](https://pinia.vuejs.org/)
-   [✅ Prettier (prettier.io)](https://prettier.io/)
-   [✅ Security Scorecards (securityscorecards.dev)](https://securityscorecards.dev/)
-   [✅ Snyk (snyk.io)](https://snyk.io)
-   [✅ Testify (github.com)](https://github.com/stretchr/testify)
-   [✅ TypeScript (typescriptlang.org)](https://www.typescriptlang.org/)
-   [✅ Vite (vitejs.dev)](https://vitejs.dev/)
-   [✅ Vitest (vitest.dev)](https://vitest.dev/)
-   [✅ Vue Test Utils (test-utils.vuejs.org)](https://test-utils.vuejs.org/)
-   [✅ VueJS (vuejs.org)](https://vuejs.org/)

## Security

Read our [Security Guide](SECURITY.md) to learn how security is considered during the development and operation of this
plugin.

## References

> 💡 Resources referenced during the development of this project.

-   [Assigning permissions to jobs (docs.github.com)](https://docs.github.com/en/actions/using-jobs/assigning-permissions-to-jobs)
-   [Cache-Control (developer.mozilla.org)](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cache-Control)
-   [Changing the input and output directory in Vite (stackoverflow.com)](https://stackoverflow.com/questions/66863200)
-   [Configuration options for the dependabot.yml file (docs.github.com)](https://docs.github.com/en/code-security/dependabot/dependabot-version-updates/configuration-options-for-the-dependabot.yml-file#package-ecosystem)
-   [Go by Example: Regular Expressions (gobyexample.com)](https://gobyexample.com/regular-expressions)
-   [How to Hash and Verify Passwords With Argon2 in Go](https://www.alexedwards.net/blog/how-to-hash-and-verify-passwords-with-argon2-in-go)
-   [How to Use go:embed in Go (blog.jetbrains.com)](https://blog.jetbrains.com/go/2021/06/09/how-to-use-go-embed-in-go-1-16/)
-   [How to properly seed random number generator (stackoverflow.com)](https://stackoverflow.com/questions/12321133/how-to-properly-seed-random-number-generator)
-   [Password Storage Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html)
-   [Serve embedded filesystem from root path of URL (stackoverflow.com)](https://stackoverflow.com/questions/66248258)
-   [Testing Your (HTTP) Handlers in Go (blog.questionable.services)](https://blog.questionable.services/article/testing-http-handlers-go/)
-   [Using a nonce with CSP (content-security-policy.com)](https://content-security-policy.com/nonce/)
-   [What's the best way to bundle static resources in a Go program? (stackoverflow.com)](https://stackoverflow.com/questions/13904441)
-   [\<script setup\> (vuejs.org)](https://vuejs.org/api/sfc-script-setup.html#script-setup)

## License

See [LICENSE.md](./LICENSE.md) for details.
