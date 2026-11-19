# ADR 001: Platform and Application Architecture

**Status:** Accepted  
**Date:** 2024-11-19

## Context

Nutrition Sauce provides structured nutrition data to AI clients. The project needs a client protocol, implementation language, and code structure that support reliable operation, testability, and independent evolution of business rules and technical integrations.

The considered options included REST, GraphQL, and the Model Context Protocol (MCP) for client communication; Go, Node.js, Python, and Rust for implementation; and Clean, Hexagonal, and Onion Architecture for organizing the code.

## Decision

- Use **MCP as the client-facing protocol** so compatible AI clients can discover and invoke the server's capabilities through a standard interface. Select and document a supported MCP transport as part of implementation; this decision does not require a custom JSON-RPC implementation or prescribe a specific transport.
- Implement the server in **Go**. The `go.mod` file defines the supported Go version. Prefer the standard library and keep dependencies focused on concrete needs.
- Organize the application using **Clean Architecture**, adapted to the server's needs. Keep domain rules independent of infrastructure, place application behavior in use cases, and make source dependencies point inward.

## Consequences

- MCP provides a standard integration boundary for compatible clients, while requiring familiarity with its protocol and transport specifications.
- Go provides a compiled deployment artifact and built-in formatting and testing tools. Contributors must account for Go's conventions and available ecosystem when choosing dependencies.
- Layer boundaries make business behavior easier to test and allow infrastructure adapters to change independently. They also add some structure and indirection, so abstractions should remain focused on actual boundaries.
- The `cmd/` package wires dependencies. `internal/domain/` contains core entities and rules, `internal/application/` contains use cases, and `internal/infrastructure/` contains protocol and data adapters. Adapt these paths to the actual implementation as it develops.

## References

- [Model Context Protocol](https://modelcontextprotocol.io/)
- [Go documentation](https://go.dev/doc/)
- [Clean Architecture, Robert C. Martin](https://blog.cleancoder.com/uncle-bob/2012/08/19/the-clean-architecture.html)
