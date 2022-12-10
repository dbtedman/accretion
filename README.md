# [Accretion](https://github.com/dbtedman/accretion)

[![CI GitHub Pipeline](https://img.shields.io/github/workflow/status/dbtedman/accretion/ci?style=for-the-badge&logo=github&label=ci)](https://github.com/dbtedman/accretion/actions/workflows/ci.yml?query=branch%3Amain)
[![SAST GitHub Pipeline](https://img.shields.io/github/workflow/status/dbtedman/accretion/sast?style=for-the-badge&logo=github&label=sast)](https://github.com/dbtedman/accretion/actions/workflows/sast.yml)

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

### Stack

-   [✅ Cobra (cobra.dev)](https://cobra.dev)
-   [✅ Docker (docker.com)](https://www.docker.com)
-   [✅ Docker Compose (docs.docker.com)](https://docs.docker.com/compose/)
-   [✅ ESLint (eslint.org)](https://eslint.org/)
-   [✅ GitHub Actions (github.com)](https://github.com/features/actions)
-   [✅ Go (go.dev)](https://go.dev)
-   [✅ Go Modules (go.dev)](https://go.dev/ref/mod)
-   [✅ Husky (typicode.github.io)](https://typicode.github.io/husky/#/)
-   [✅ Node Version Manager (github.com)](https://github.com/nvm-sh/nvm)
-   [✅ NodeJS (nodejs.org)](https://nodejs.org/en/)
-   [✅ Pinia (pinia.vuejs.org)](https://pinia.vuejs.org/)
-   [✅ PNPM (pnpm.io)](https://pnpm.io/)
-   [✅ Prettier (prettier.io)](https://prettier.io/)
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

## License

See [LICENSE.md](./LICENSE.md) for details.
