## Context

rme uses the RUNME.sh convention (via `make_command`) for task declaration. The project builds with Go and packages with Nix. Currently there's no standardized way to run dev tasks — developers run `go build`, `go test`, etc. ad hoc.

## Goals / Non-Goals

**Goals:**
- Single entry point for all dev tasks via `RUNME.sh`
- Version injection in `build` matching flake.nix ldflags (`-s -w -X main.version=<version>`)
- Standard Go toolchain tasks: build, test, fmt, vet

**Non-Goals:**
- Release workflow (deferred)
- Nix build integration (use `nix build` directly)
- CI/CD integration

## Decisions

**Use hardcoded version in RUNME.sh matching flake.nix**
The version string (`0.1.0`) is duplicated between `flake.nix` and `RUNME.sh`. This is acceptable for now — the project is small and a single-source-of-truth version file adds complexity without clear benefit yet. Can revisit when release workflow is added.

**Output binary to project root as `rme`**
`go build -o rme .` keeps the binary in the working directory, matching typical Go dev workflow. The `.gitignore` should exclude it (but that's out of scope for this change).

**Use gofmt, not goimports or gofumpt**
`gofmt` is the standard formatter shipped with Go. No extra tooling needed.

## Risks / Trade-offs

- **Version drift**: flake.nix and RUNME.sh both contain the version string. Low risk given project size.
- **No goimports**: Missing import management. Acceptable — the project has few imports.
