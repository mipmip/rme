## ADDED Requirements

### Requirement: Add command appends a command block to RUNME.sh
The `rme add <name> <description>` command SHALL append a `make_command` declaration and function stub to the existing RUNME.sh in the current directory, inserted before the `##### PLACE YOUR COMMANDS ABOVE #####` marker.

#### Scenario: Successful add
- **WHEN** user runs `rme add deploy "Deploy to production"` in a directory with a valid RUNME.sh
- **THEN** the file is updated with a new block containing `make_command "deploy" "Deploy to production"` and a `deploy()` function stub, placed before the upper marker

#### Scenario: Generated block format
- **WHEN** a command is added successfully
- **THEN** the inserted block SHALL be:
  ```
  make_command "<name>" "<description>"
  <name>(){
    echo "<name>"
  }
  ```

### Requirement: Add command requires both name and description
The `rme add` command SHALL require exactly two arguments (name and description) and exit with an error if they are missing.

#### Scenario: Missing arguments
- **WHEN** user runs `rme add` or `rme add deploy` (without description)
- **THEN** the system prints a usage message and exits with a non-zero exit code

### Requirement: Add command requires existing RUNME.sh
The `rme add` command SHALL fail if no RUNME.sh exists in the current directory.

#### Scenario: No RUNME.sh
- **WHEN** user runs `rme add deploy "Deploy"` in a directory without RUNME.sh
- **THEN** the system prints an error and exits with a non-zero exit code

### Requirement: Add command rejects duplicate command names
The `rme add` command SHALL fail if a command with the same name already exists in RUNME.sh.

#### Scenario: Duplicate command
- **WHEN** user runs `rme add build "Build again"` and `build` already exists in RUNME.sh
- **THEN** the system prints an error indicating the command already exists and exits with a non-zero exit code without modifying the file

### Requirement: Add command requires marker in RUNME.sh
The `rme add` command SHALL fail if the RUNME.sh does not contain the `##### PLACE YOUR COMMANDS ABOVE #####` marker.

#### Scenario: Missing marker
- **WHEN** user runs `rme add deploy "Deploy"` and RUNME.sh lacks the marker
- **THEN** the system prints an error and exits with a non-zero exit code
