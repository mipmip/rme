## ADDED Requirements

### Requirement: GoReleaser config builds cross-platform binaries
The `.goreleaser.yaml` SHALL configure builds for linux/darwin on amd64/arm64, building from the `./src` directory with version-injecting ldflags.

#### Scenario: GoReleaser builds from src directory
- **WHEN** goreleaser runs
- **THEN** it builds the binary from `./src` with `-X main.version={{.Version}}` ldflags for all target platforms

### Requirement: GoReleaser creates GitHub release with changelog
The `.goreleaser.yaml` SHALL configure the release to use the changelog section for the current version.

#### Scenario: Release includes changelog
- **WHEN** a tag is pushed and goreleaser runs
- **THEN** the GitHub release is created with auto-generated release notes
