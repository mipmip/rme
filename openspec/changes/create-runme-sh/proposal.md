## Why

rme is a launcher for RUNME.sh, but its own development has no RUNME.sh. Adding one provides a single entry point for build, test, format, and lint tasks — and lets the project eat its own dog food.

## What Changes

- Add a `RUNME.sh` to the project root with development tasks:
  - `build` — compile with version injection matching flake.nix ldflags
  - `test` — run Go test suite
  - `fmt` — format Go source with gofmt
  - `vet` — static analysis with go vet

## Capabilities

### New Capabilities
- `dev-tasks`: RUNME.sh with build, test, fmt, and vet commands for rme development

### Modified Capabilities

_None._

## Impact

- Adds one new file (`RUNME.sh`) at project root
- No code changes to rme itself
- No dependency changes
