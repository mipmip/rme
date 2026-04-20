## Context

rme is a small Go project with source in `./src`, built via Nix flake and `RUNME.sh`. Version `0.1.0` is hardcoded in `flake.nix` (line 21) and `RUNME.sh` build command (line 9). There's no changelog, no release automation, and no binary distribution beyond Nix.

## Goals / Non-Goals

**Goals:**
- Changelog following keepachangelog.com format
- RUNME.sh `release` command: bump version in all files, move Unreleased to new version, commit, tag
- GoReleaser config for cross-platform binaries
- GitHub Actions workflow for automated release on tag push

**Non-Goals:**
- Nix release/cache automation (use `nix build` directly)
- Publishing to package managers (homebrew, etc.)
- Signing binaries

## Decisions

**1. Version lives in files, not a VERSION file**
The release bumper uses `sed` to update the version in `flake.nix` and `RUNME.sh` directly. No separate VERSION file.
- Rationale: Only two files contain the version. Adding a VERSION file + build-time reads is over-engineering for this project size.
- The bumper takes the bump type (major/minor/patch) as argument and computes the new version from the current one in `flake.nix`.

**2. Changelog format: keepachangelog**
Standard format (`## [Unreleased]` / `## [x.y.z] - YYYY-MM-DD`). The release bumper renames `## [Unreleased]` to the new version+date and inserts a fresh Unreleased section.

**3. GoReleaser with `dir: src`**
GoReleaser supports `project_name` and building from a subdirectory via `builds[].dir`. The config points at `./src` and injects version via ldflags matching the flake pattern.

**4. GitHub Action: goreleaser-action on tag push**
Triggered by `v*` tags. Uses `goreleaser/goreleaser-action@v6`. Needs `GITHUB_TOKEN` for creating the release.

**5. Release command is bash, not Go**
The release bumper is a RUNME.sh command (bash), not compiled into `rme`. Release tooling belongs in the dev workflow, not the distributed binary.

**6. VCS detection: jj vs git**
Check for `.jj` directory to detect jj. Use `jj commit` for jj, `git commit -am` for git. Tags always use `git tag` since jj uses git as its backend. GitHub release created via `gh release create` (works regardless of VCS).

**7. GitHub release via `gh` CLI**
After commit+tag, run `gh release create v<version> --generate-notes`. This triggers the goreleaser GitHub Action. No manual push needed — `gh release create` pushes the tag.

## Risks / Trade-offs

- [Version sed patterns] → Fragile if file format changes. Acceptable for a small project with stable structure.
- [GoReleaser version pinned in Action] → Pin to major version (`v6`) for stability with updates.
- [`gh` CLI required] → The release command requires `gh` to be installed and authenticated. This is a dev-time dependency only.
- [jj detection via `.jj` dir] → Simple heuristic. Could fail if `.jj` is removed but jj is still used. Acceptable.
