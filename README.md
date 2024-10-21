# [Accretion](https://github.com/dbtedman/accretion)

> ⚠️ This is pre-release software, please do not use in production.

[![CI GitHub Pipeline](https://img.shields.io/github/actions/workflow/status/dbtedman/accretion/ci.yml?branch=main&style=for-the-badge&logo=github&label=ci)](https://github.com/dbtedman/accretion/actions/workflows/ci.yml?query=branch%3Amain)
[![Latest Release](https://img.shields.io/github/v/release/dbtedman/accretion?style=for-the-badge&logo=github&color=43cc11)](https://github.com/dbtedman/accretion/releases)

Evaluate application security best practice implementation for running applications.

- [Usage Instructions](#usage-instructions)
- [Design](#design)
- [License](#license)

## Usage Instructions

### Install

```shell
brew install dbtedman/tap/accretion
```

## Design

### Interceptor

Component that intercepts traffic and collects metrics.

### Collector

Component that accepts metrics from the interceptor on a regular schedule.

### Notifier

Component that publishes alerts based on defined policies.

### Navigator

Component that provides a user interface for navigating collected metrics.

## License

The [MIT License](./LICENSE.md) is used by this project.
