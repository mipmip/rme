## 1. Changelog

- [x] 1.1 Create `CHANGELOG.md` with keepachangelog header, `## [Unreleased]` section, and `## [0.1.0]` initial entry

## 2. GoReleaser

- [x] 2.1 Create `.goreleaser.yaml` with builds from `./src`, ldflags version injection, linux/darwin amd64/arm64 targets

## 3. GitHub Action

- [x] 3.1 Create `.github/workflows/release.yml` triggered on `v*` tags, using `goreleaser/goreleaser-action@v6`

## 4. Release Bumper

- [x] 4.1 Add `release` command to `RUNME.sh` that: validates bump type arg, reads current version from `flake.nix`, computes new semver, updates version in `flake.nix` and `RUNME.sh`, updates `CHANGELOG.md` (renames Unreleased, inserts fresh section), detects VCS (jj via `.jj` dir, else git), commits, creates git tag, and creates GitHub release via `gh release create`

## 5. Verify

- [x] 5.1 Verify `CHANGELOG.md` format is correct
- [x] 5.2 Dry-run: test the release bumper logic by inspecting the sed commands against current file content
