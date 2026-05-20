# Agent Instructions

## Package Manager

Use **Go Modules**: `go mod tidy`, `go test ./...`
Use **Mise**: `mise run [task]` (e.g. `mise run build`, `mise run lint`)

## Key Conventions

- **Testing**: Use standard `testing` and `github.com/google/go-cmp/cmp` **only**. No frameworks (Testify, Ginkgo, etc.).
- **Linting**: Run `golangci-lint` v1. Configure via project-specific `.golangci.yml`.
- **Build**: Run `mise run build` for a full CI build. See `mise tasks` for all available tasks.

## Module Structure

```
go.mod          # VIN SDK library (no CLI deps)
cmd/vin/        # Standalone vin CLI binary
tools/cmd/      # Data-processing tool binaries (not part of CI build)
proto/          # Protobuf definitions + buf codegen
```

## Data Tasks

Data download and processing tasks are available but are **not** part of `mise run build`:

| Task                    | Purpose                                   |
| ----------------------- | ----------------------------------------- |
| `mise run vpic`         | Download NHTSA vPIC database              |
| `mise run kba-wmi-pdf`  | Download KBA WMI PDF and split into pages |
| `mise run kba-excel`    | Download KBA Excel statistics files       |
| `mise run acea`         | Download ACEA registration data           |
| `mise run docker-build` | Build vin CLI Docker image locally        |
| `mise run docker-push`  | Push vin CLI Docker image to ghcr.io      |

## Local Skills

- **Way Go Style**: `.agents/skills/way-go-style/SKILL.md`
