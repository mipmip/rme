## Context

rme can bootstrap a RUNME.sh via `rme init`, but adding commands requires manual editing. The RUNME.sh format uses a marker-based layout with `##### PLACE YOUR COMMANDS ABOVE #####` as the insertion boundary.

## Goals / Non-Goals

**Goals:**
- Append a command block (make_command + function stub) to an existing RUNME.sh
- Insert before the upper marker line so the file remains valid
- Validate that the command name doesn't already exist

**Non-Goals:**
- Removing or editing existing commands
- Generating function body content (stub is always `echo "<name>"`)
- Supporting RUNME.sh files without the standard markers

## Decisions

**1. Find the marker line, insert before it**
Read the entire RUNME.sh, find `##### PLACE YOUR COMMANDS ABOVE #####`, and insert the new block (with a blank line separator) just before it. Rewrite the whole file.
- Rationale: Simple string manipulation. The file is always small.
- Alternative considered: Append before `runme` call — fragile, the marker is more reliable.

**2. Duplicate command detection via `parseFromSource`**
Reuse the existing `parseFromSource()` function to check if a command with the same name already exists.
- Rationale: Already parses `make_command` calls; no new parsing logic needed.

**3. Function stub body: `echo "<name>"`**
The generated function contains `echo "<name>"` as a placeholder. Users replace this with real logic.
- Rationale: Makes it obvious the stub needs editing; running it is harmless.

## Risks / Trade-offs

- [No marker found] → Error message telling the user the RUNME.sh doesn't have the expected format
- [Command name with spaces/special chars] → Not validated beyond what bash allows. Acceptable for v1.
