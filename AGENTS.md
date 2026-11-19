# Repository Guidelines

## Project Structure & Module Organization

This repository currently contains project documentation rather than implementation code. `README.md` describes the Nutrition Sauce MCP server and its proposed Go layout. Architecture decisions live in `docs/adr/` and Mermaid diagrams live in `docs/assets/diagrams/`. The `cmd/`, `internal/`, `data/`, and `config/` paths shown in those documents are planned locations; check the tree before adding code and keep documentation aligned with any implementation that is introduced.

## Build, Test, and Development Commands

There is no `go.mod`, application source, or test suite yet, so build and test commands are not available in the current checkout. When Go code is added, use the module's documented toolchain and verify changes with `go test ./...` and `go build ./...`. For documentation-only changes, inspect the Markdown diff and confirm links and code examples match repository paths and decisions.

## Coding Style & Naming Conventions

Use Clean Architecture as the architectural reference described in [ADR 001](docs/adr/001-platform-and-architecture.md): keep domain logic independent of infrastructure, put application behavior in use cases, and make dependencies point inward. Use the `uber-go-style` skill and the [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md) as the primary style references when writing or reviewing Go code, while following repository-specific conventions where they are already established. Format Go code with `gofmt` and use concise, lowercase package names.

Write documentation in concise Markdown with descriptive headings, fenced code blocks labeled by language, and repository-relative paths. Keep architectural decisions in numbered files such as `docs/adr/005-short-decision-name.md`; diagrams belong in `docs/assets/diagrams/` and use the `.mmd` extension.

## Testing Guidelines

No test framework or coverage target is configured yet. When implementation packages are added, place Go tests beside the code in `*_test.go` files and run `go test ./...`. Add focused tests for domain rules and use cases, keeping file or protocol details behind interfaces where practical.

## Commit & Pull Request Guidelines

Use [Conventional Commits](https://www.conventionalcommits.org/) as the default commit-message format:

```text
<type>[optional scope]: <description>
```

Use a short imperative description and keep each commit focused. Common types include `feat` for functionality, `fix` for bug fixes, `docs` for documentation, `refactor` for restructuring, and `chore` for maintenance. Add `!` after the type or scope for a breaking change, and include a `BREAKING CHANGE:` footer when useful.

Examples:

- `docs: document JSON storage decision`
- `feat(food): add category filtering`
- `refactor!: change nutrition response schema`

Pull requests should explain the purpose and notable decisions, link any relevant issue or ADR, and include rendered screenshots only when a visual document or diagram change needs them.

## Security & Configuration

Do not commit secrets, credentials, or private nutrition data. Keep configuration examples sanitized, and update relevant ADRs when changing data storage, transport, or dependency boundaries.
