# [Accretion](https://github.com/dbtedman/accretion)

Manage internal technical documentation that is enriched with live data accreted from your environment.

## Communication

- [Related Articles (tedman.dev)](https://tedman.dev/topics/accretion/)

## Principals

- Secure
- Verifiable
- Open
- Private
- Accessible
- Immutable
- Elegant
- Standards

## Commands

Represents a request for a change to me made to the state of the system, e.g. `Add Product`.

## Queries

Read only interaction with the system used to inspect its current state, e.g. `List Products`.

## Events

The systems source of truth is an immutable stream of idempotent Events which represent changes that occur to the
system, e.g. `Product Added`.

> 💡 Sensitive data will be encrypted at rest, using a secret key associated with the owner of the data, which can be
> purged on request,
> see [GDPR Compliant Event Sourcing With HashiCorp Vault (hashicorp.com)](https://www.hashicorp.com/resources/gdpr-compliant-event-sourcing-with-hashicorp-vault)
> for reference.

## Projections

Queries won't be performed on the event log, but rather on projections derived from this log, e.g. `Search Index`.

## Data Flow

```shell
d2 -w -t 200 --layout=elk ./doc/data-flow.d2 ./doc/data-flow.svg
```

![](./doc/data-flow.svg)

## Domain

```shell
d2 -w -t 200  --layout=elk ./doc/domain.d2 ./doc/domain.svg
```

![](./doc/domain.svg)

## UI

### Discover

Search, Filter, and Pivot though accreted data.

### Contribute

Add new resources or modify existing ones.
