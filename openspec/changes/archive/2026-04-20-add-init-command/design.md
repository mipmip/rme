## Context

rme is a lightweight Go command launcher for RUNME.sh files. It currently detects and dispatches to an existing RUNME.sh but has no way to create one. Users must manually download RUNME.sh from the mipmip/RUNME.sh GitHub repository. Adding an `init` subcommand completes the bootstrapping story.

## Goals / Non-Goals

**Goals:**
- Provide a single command to bootstrap a RUNME.sh in the current directory
- Download the latest version from GitHub (raw content URL)
- Prevent accidental overwrites of existing RUNME.sh files
- Keep zero external dependencies (use Go stdlib `net/http`)

**Non-Goals:**
- Scaffolding commands inside RUNME.sh (just download the template)
- Version pinning or selecting specific RUNME.sh releases
- Updating an existing RUNME.sh (`rme update` could be a future command)

## Decisions

**1. Download source: raw.githubusercontent.com**
- Fetch from `https://raw.githubusercontent.com/mipmip/RUNME.sh/master/RUNME.sh`
- Rationale: Direct raw URL is simplest, no GitHub API auth needed, no rate limiting for raw content
- Alternative considered: GitHub API releases endpoint — adds complexity, RUNME.sh doesn't use GitHub releases

**2. New file `init.go`**
- Follows existing pattern: each concern in its own file (`detect.go`, `parse.go`, `completion.go`)
- Contains `handleInit()` function called from `main.go`

**3. Fail-fast on existing RUNME.sh**
- Check for file existence before downloading (not after)
- Rationale: Don't waste network round-trip if we'll refuse to write anyway

**4. File permissions: 0755**
- RUNME.sh must be executable; matches what `detectRunme()` already expects
- Write with `os.WriteFile` using mode `0755`

## Risks / Trade-offs

- [Network failure] → Print clear error message with the URL so user can manually download
- [GitHub URL changes] → URL is hardcoded; if the repo moves, a new rme release is needed. Acceptable for simplicity.
- [No offline mode] → `rme init` requires internet. This is expected for a download command.
