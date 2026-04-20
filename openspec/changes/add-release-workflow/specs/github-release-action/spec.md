## ADDED Requirements

### Requirement: GitHub Action triggers on version tags
The `.github/workflows/release.yml` workflow SHALL trigger on pushes of tags matching `v*`.

#### Scenario: Tag push triggers release
- **WHEN** a tag `v0.2.0` is pushed to GitHub
- **THEN** the release workflow runs goreleaser to build and publish binaries

### Requirement: GitHub Action uses goreleaser-action
The workflow SHALL use `goreleaser/goreleaser-action@v6` with Go set up and `GITHUB_TOKEN` provided.

#### Scenario: Workflow configuration
- **WHEN** the workflow runs
- **THEN** it checks out the code, sets up Go, and runs goreleaser with the repository's `.goreleaser.yaml`
