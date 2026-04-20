#!/usr/bin/env bash
#(C)2019-2022 Pim Snel - https://github.com/mipmip/RUNME.sh
CMDS=();DESC=();NARGS=$#;ARG1=$1;make_command(){ CMDS+=($1);DESC+=("$2");};usage(){ printf "\nUsage: %s [command]\n\nCommands:\n" $0;line="              ";for((i=0;i<=$(( ${#CMDS[*]} -1));i++));do printf "  %s %s ${DESC[$i]}\n" ${CMDS[$i]} "${line:${#CMDS[$i]}}";done;echo;};runme(){ if test $NARGS -eq 1;then eval "$ARG1"||usage;else usage;fi;}

##### PLACE YOUR COMMANDS BELOW #####

make_command "build" "Build rme binary with version injection"
build(){
  cd src && go build -ldflags "-s -w -X main.version=0.2.0" -o ../rme .
}

make_command "run_tests" "Run Go test suite"
run_tests(){
  cd src && go test ./...
}

make_command "fmt" "Format Go source files"
fmt(){
  cd src && gofmt -w .
}

make_command "vet" "Run static analysis"
vet(){
  cd src && go vet ./...
}

_do_release(){
  local bump_type="$1"

  # Read current version from flake.nix
  local current
  current=$(grep 'version = "' flake.nix | head -1 | sed 's/.*version = "\([^"]*\)".*/\1/')
  if [ -z "$current" ]; then
    echo "Error: could not read version from flake.nix"
    return 1
  fi

  # Compute new version
  local major minor patch
  IFS='.' read -r major minor patch <<< "$current"
  case "$bump_type" in
    major) major=$((major + 1)); minor=0; patch=0 ;;
    minor) minor=$((minor + 1)); patch=0 ;;
    patch) patch=$((patch + 1)) ;;
  esac
  local new_version="${major}.${minor}.${patch}"

  echo "Bumping version: ${current} -> ${new_version}"

  # Update version in flake.nix
  sed -i "s/version = \"${current}\"/version = \"${new_version}\"/" flake.nix

  # Update version in RUNME.sh
  sed -i "s/main.version=${current}/main.version=${new_version}/" RUNME.sh

  # Update CHANGELOG.md
  local today
  today=$(date +%Y-%m-%d)
  sed -i "s/## \[Unreleased\]/## [Unreleased]\n\n## [${new_version}] - ${today}/" CHANGELOG.md

  echo "Updated flake.nix, RUNME.sh, and CHANGELOG.md"

  # Detect VCS and commit
  if [ -d ".jj" ]; then
    echo "Detected jj"
    jj commit -m "Release v${new_version}"
  else
    echo "Detected git"
    git add -A
    git commit -m "Release v${new_version}"
  fi

  # Tag (git works for both git and jj)
  git tag "v${new_version}"

  echo "Created tag v${new_version}"

  # Create GitHub release (pushes tag and triggers goreleaser)
  gh release create "v${new_version}" --generate-notes

  echo "Released v${new_version}"
}

make_command "release_major" "Release: bump major version (x.0.0)"
release_major(){ _do_release major; }

make_command "release_minor" "Release: bump minor version (0.x.0)"
release_minor(){ _do_release minor; }

make_command "release_patch" "Release: bump patch version (0.0.x)"
release_patch(){ _do_release patch; }

##### PLACE YOUR COMMANDS ABOVE #####

runme
