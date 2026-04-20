package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const runmeFile = "RUNME.sh"

// detectRunme looks for RUNME.sh in the current working directory.
// Returns the absolute path or an error.
func detectRunme() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("rme: cannot determine current directory: %w", err)
	}

	path := filepath.Join(cwd, runmeFile)

	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return "", fmt.Errorf("rme: no RUNME.sh found in current directory")
	}
	if err != nil {
		return "", fmt.Errorf("rme: cannot stat RUNME.sh: %w", err)
	}

	if info.Mode()&0111 == 0 {
		return "", fmt.Errorf("rme: RUNME.sh is not executable (chmod +x RUNME.sh)")
	}

	return path, nil
}
