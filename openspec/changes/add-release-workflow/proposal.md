## Why

rme has no release process. The version is hardcoded in multiple places (`flake.nix`, `RUNME.sh`), there's no changelog, and no automated binary builds. Adding a proper release workflow makes it easy to cut releases with proper changelogs and cross-platform binaries.

## What Changes

- Add `CHANGELOG.md` with keepachangelog format and an Unreleased placeholder
- Add a `release` command to `RUNME.sh` that bumps the version (semver) across all files, updates the changelog, commits, and tags
- Add `.goreleaser.yaml` for cross-platform binary builds (pointing at `./src`)
- Add a GitHub Actions workflow that runs goreleaser on tag push

## Capabilities

### New Capabilities
- `changelog`: CHANGELOG.md with keepachangelog format and Unreleased section
- `release-bumper`: RUNME.sh command to bump version, update changelog, commit, and tag
- `goreleaser-config`: GoReleaser configuration for cross-platform builds from `./src`
- `github-release-action`: GitHub Actions workflow triggered on version tags

### Modified Capabilities

_None._

## Impact

- New files: `CHANGELOG.md`, `.goreleaser.yaml`, `.github/workflows/release.yml`
- Modified: `RUNME.sh` (new `release` command)
- Version string currently hardcoded in `flake.nix` (line 21) and `RUNME.sh` (line 9) — the release bumper must update both
