## Why

Projects using RUNME.sh need a way to bootstrap their `RUNME.sh` file. Currently users must manually download it from GitHub. An `rme init` command provides a natural entry point — just like `git init` or `npm init` — making it trivial to adopt RUNME.sh in any project.

## What Changes

- Add `rme init` subcommand that downloads the latest `RUNME.sh` from `https://raw.githubusercontent.com/mipmip/RUNME.sh/master/RUNME.sh`
- The downloaded file is saved as `./RUNME.sh` in the current working directory and made executable
- The command fails with an error if `RUNME.sh` already exists in the current directory (prevents accidental overwrites)

## Capabilities

### New Capabilities
- `init-command`: The `rme init` subcommand that downloads and bootstraps a RUNME.sh file in the current directory

### Modified Capabilities

## Impact

- `main.go`: Add `init` to the subcommand dispatch logic
- New file `init.go`: Contains the download and file-write logic
- Adds `net/http` to imports (stdlib, no external dependencies)
- No breaking changes to existing functionality
