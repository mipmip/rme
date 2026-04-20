## Why

After `rme init` creates a RUNME.sh, users must manually edit the file to add commands. An `rme add <name> <description>` command completes the scaffolding story — users can add new command stubs without knowing the RUNME.sh boilerplate format.

## What Changes

- Add `rme add <name> <description>` subcommand that appends a `make_command` declaration and empty function stub to the existing `RUNME.sh` in the current directory
- The new block is inserted before the `##### PLACE YOUR COMMANDS ABOVE #####` marker
- Requires a RUNME.sh to already exist (fails otherwise)

## Capabilities

### New Capabilities
- `add-command`: The `rme add` subcommand that appends command blocks to an existing RUNME.sh

### Modified Capabilities

_None._

## Impact

- `main.go`: Add `"add"` case to subcommand dispatch
- New file `add.go`: Contains the file manipulation logic
- No new dependencies (uses `os`, `strings` from stdlib)
