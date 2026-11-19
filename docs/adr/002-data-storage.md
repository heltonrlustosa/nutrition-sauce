# ADR 002: Nutrition Data Storage

**Status:** Accepted  
**Date:** 2024-11-19

## Context

The server needs to load and serve a nutrition dataset. The data is expected to be mostly static and small enough to fit in memory. Options considered were JSON files, a relational database, a document database, and a persistent JSON dataset with an in-memory index.

## Decision

Store the source dataset as JSON files and load it into memory when the server starts. Build in-memory indexes needed for common lookups. Keep file access and JSON parsing behind an infrastructure adapter so use cases depend on application or domain interfaces rather than storage details.

## Consequences

- JSON is easy to review, version, and distribute with the project.
- In-memory reads avoid a separate database service and suit the expected static dataset.
- Data changes require updating the source files and restarting the server; runtime writes and hot reload are not part of this decision.
- Memory use grows with the dataset. Revisit the storage choice if dataset size, update frequency, query needs, or concurrent write requirements change materially.

## References

- [Go `encoding/json` package](https://pkg.go.dev/encoding/json)
