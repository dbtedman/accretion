# Accretion

Navigate internet resources from the command line.

- [Design](#design)
- [License](#license)

## Local Development

```shell
go mod tidy
go mod vendor
```

```shell
CGO_ENABLED=0 go test
```

## Design

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
