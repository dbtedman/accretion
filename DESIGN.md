# [Accretion](https://github.com/dbtedman/accretion) Design Guide

## Principals

- Secure
- Verifiable
- Open
- Private
- Accessible
- Immutable
- Elegant
- Standards

## References

🚧 Placeholder

## Event Sourced

### Commands

Represents a request for a change to be made to the state of the system, e.g. `Add Product`.

### Queries

Read only interaction with the system used to inspect its current state, e.g. `List Products`.

### Events

The systems source of truth is an immutable stream of idempotent Events which represent changes that occur to the
system, e.g. `Product Added`.

> 💡 Sensitive data will be encrypted at rest, using a secret key associated with the owner of the data, which can be
> purged on request, see [GDPR Compliant Event Sourcing With HashiCorp Vault (hashicorp.com)](https://www.hashicorp.com/resources/gdpr-compliant-event-sourcing-with-hashicorp-vault) for reference.

### Projections

Queries won't be performed on the event log, but rather on projections derived from this log, e.g. `Search Index`.

## Data Flow

A harvester populates data into the append only event log from each of the configured external services.

![](./doc/data-flow.svg)

```shell
d2 -w -t 200 --layout=elk ./doc/data-flow.d2 ./doc/data-flow.svg
```

## Simple Infrastructure

A single go binary with embedded web application supporting multiple replicas for redundancy and horizontal scalability.

## Discover UI

Implemented as a HUD over a results grid that is optimised for fast interaction using keyboard, mouse, touch screen or assistive devices.

### Elements

- **Search** using boolean expressions and keywords.
- **Sort** by any sortable properties.
- **Filter** Results based on attributes contextually related to the current results list.
- **Time Travel** back in time to see how results would appear in the past.
- **Pivot** on attributes to show related results.
- **Scroll** through a grid of search results.

### Scenarios

- Do we have an api that allows us to email a customer?
- Which projects depend on an outdated version of the x package?
- Are there any code repositories that have not been updated in over a year?
- Who can I contact about the x.y.z service?
- Show me all projects who are looked after by team abc, then filter that list by those updated within the last 30 days, but show me how this search would look 10 days ago.

## Contribute UI

### Elements

- Contribute card overlays discover UI based on calls to action.

### Scenarios

- The a, d, and f services are part of the same project called g.
- The abc service should not be used for new projects.
