# dev-tasks Specification

## Purpose
TBD - created by archiving change create-runme-sh. Update Purpose after archive.
## Requirements
### Requirement: Build command compiles with version injection
The `build` command SHALL compile the Go project with ldflags `-s -w -X main.version=0.1.0`, producing a binary named `rme` in the project root.

#### Scenario: Successful build
- **WHEN** user runs `rme build` (or `./RUNME.sh build`)
- **THEN** a binary `rme` is produced in the project root with the correct version embedded

#### Scenario: Version matches flake
- **WHEN** user runs the built binary with `--version`
- **THEN** output shows `rme 0.1.0`

### Requirement: Test command runs Go test suite
The `test` command SHALL run `go test ./...` and display results.

#### Scenario: Run tests
- **WHEN** user runs `rme test`
- **THEN** all Go tests in the project are executed and results are displayed

### Requirement: Fmt command formats Go source
The `fmt` command SHALL run `gofmt -w .` to format all Go files in place.

#### Scenario: Format source files
- **WHEN** user runs `rme fmt`
- **THEN** all Go source files are formatted according to gofmt conventions

### Requirement: Vet command runs static analysis
The `vet` command SHALL run `go vet ./...` to perform static analysis.

#### Scenario: Run vet
- **WHEN** user runs `rme vet`
- **THEN** go vet analyzes all packages and reports any issues

### Requirement: RUNME.sh follows RUNME.sh conventions
The file SHALL use `make_command` declarations and the standard RUNME.sh boilerplate so that `rme` can parse it for completions.

#### Scenario: Completions work
- **WHEN** user runs `rme --completions` in the project directory
- **THEN** output includes `build`, `test`, `fmt`, and `vet`

