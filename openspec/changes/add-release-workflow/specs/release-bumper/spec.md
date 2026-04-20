## ADDED Requirements

### Requirement: Release command bumps version across all files
The `release` RUNME.sh command SHALL accept a bump type argument (`major`, `minor`, or `patch`), compute the new semver version, and update the version string in `flake.nix` and `RUNME.sh`.

#### Scenario: Patch bump
- **WHEN** user runs `rme release patch` with current version `0.1.0`
- **THEN** the version in `flake.nix` and `RUNME.sh` is updated to `0.1.1`

#### Scenario: Minor bump
- **WHEN** user runs `rme release minor` with current version `0.1.0`
- **THEN** the version in `flake.nix` and `RUNME.sh` is updated to `0.2.0`

#### Scenario: Major bump
- **WHEN** user runs `rme release major` with current version `0.1.0`
- **THEN** the version in `flake.nix` and `RUNME.sh` is updated to `1.0.0`

### Requirement: Release command updates changelog
The `release` command SHALL rename the `## [Unreleased]` section in `CHANGELOG.md` to `## [<new-version>] - <date>` and insert a fresh `## [Unreleased]` section above it.

#### Scenario: Changelog updated on release
- **WHEN** user runs `rme release patch`
- **THEN** `CHANGELOG.md` has a new `## [0.1.1] - YYYY-MM-DD` section (with today's date) and a fresh empty `## [Unreleased]` section above it

### Requirement: Release command detects VCS and commits
The `release` command SHALL detect whether the project uses jj (by checking for `.jj` directory) or git, and commit accordingly: `jj commit -m "Release v<version>"` for jj, `git commit -am "Release v<version>"` for git.

#### Scenario: Commit with jj
- **WHEN** user runs `rme release patch` in a jj-managed repo
- **THEN** the changes are committed with `jj commit -m "Release v0.1.1"`

#### Scenario: Commit with git
- **WHEN** user runs `rme release patch` in a git-only repo
- **THEN** the changes are committed with `git commit -am "Release v0.1.1"`

### Requirement: Release command tags via git
The `release` command SHALL create a git tag `v<version>`. This works for both git and jj repos (jj uses git underneath).

#### Scenario: Tag created
- **WHEN** user runs `rme release patch` successfully
- **THEN** a git tag `v0.1.1` is created

### Requirement: Release command creates GitHub release
The `release` command SHALL create a GitHub release using `gh release create v<version> --generate-notes` after committing and tagging.

#### Scenario: GitHub release created
- **WHEN** user runs `rme release patch` successfully
- **THEN** a GitHub release `v0.1.1` is created via `gh`, which triggers the goreleaser action

### Requirement: Release command validates arguments
The `release` command SHALL fail with a usage message if no bump type is provided or if the bump type is invalid.

#### Scenario: Missing argument
- **WHEN** user runs `rme release` without a bump type
- **THEN** the command prints usage and exits with non-zero exit code
