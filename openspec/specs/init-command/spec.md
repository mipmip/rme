# init-command Specification

## Purpose
TBD - created by archiving change add-init-command. Update Purpose after archive.
## Requirements
### Requirement: Init command downloads RUNME.sh
The `rme init` command SHALL download the latest RUNME.sh from `https://raw.githubusercontent.com/mipmip/RUNME.sh/master/RUNME.sh` and save it as `./RUNME.sh` in the current working directory with executable permissions (0755).

#### Scenario: Successful init in empty project
- **WHEN** user runs `rme init` in a directory without RUNME.sh
- **THEN** the system downloads RUNME.sh from GitHub, saves it as `./RUNME.sh` with mode 0755, and prints a success message

#### Scenario: RUNME.sh is functional after init
- **WHEN** user runs `rme init` successfully
- **THEN** the downloaded `./RUNME.sh` SHALL be detected by `rme` commands (e.g., `rme --completions`)

### Requirement: Init command refuses to overwrite existing RUNME.sh
The `rme init` command SHALL fail with a non-zero exit code and an error message if a `RUNME.sh` file already exists in the current working directory.

#### Scenario: RUNME.sh already exists
- **WHEN** user runs `rme init` in a directory that already contains RUNME.sh
- **THEN** the system prints an error message indicating RUNME.sh already exists and exits with a non-zero exit code without modifying the existing file

### Requirement: Init command handles network errors
The `rme init` command SHALL print a descriptive error message and exit with a non-zero exit code if the download fails.

#### Scenario: Network unreachable
- **WHEN** user runs `rme init` and the HTTP request fails
- **THEN** the system prints an error message and exits with a non-zero exit code without creating any file

#### Scenario: Non-200 HTTP response
- **WHEN** user runs `rme init` and GitHub returns a non-200 status code
- **THEN** the system prints an error message including the HTTP status and exits with a non-zero exit code without creating any file

