# [Accretion](https://github.com/dbtedman/accretion)

[![CI GitHub Pipeline](https://img.shields.io/github/actions/workflow/status/dbtedman/accretion/ci.yml?branch=main&style=for-the-badge&logo=github&label=ci)](https://github.com/dbtedman/accretion/actions/workflows/ci.yml?query=branch%3Amain)
[![Latest Release](https://img.shields.io/github/v/release/dbtedman/accretion?style=for-the-badge&logo=github&color=43cc11)](https://github.com/dbtedman/accretion/releases)

Navigate internet resources from the command line.

- [Usage Instructions](#usage-instructions)
- [Design](#design)
- [License](#license)

## Usage Instructions

```shell
brew install dbtedman/tap/accretion

accretion
```

## Design

### Reference

- [RDAP Client (client.rdap.org)](https://client.rdap.org)

### Workflow

Query on various internet resources:

```shell
accretion AS4608
accretion NO4-AP
accretion 2001:dc0::/32
accretion 202.12.31.0/24
accretion tedman.dev
```

Render output in different standard formats:

```shell
accretion tedman.dev --json
```

### Authoritative Sources

#### Registration Data Access Protocol (RDAP)

- AFRINIC https://rdap.afrinic.net
- APNIC https://rdap.apnic.net
- ARIN https://rdap.arin.net
- LACNIC https://rdap.lacnic.net
- RIPE NCC https://rdap.db.ripe.net

## License

The [MIT License](./LICENSE.md) is used by this project.
