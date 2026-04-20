package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const upperMarker = "##### PLACE YOUR COMMANDS ABOVE #####"

func handleAdd(args []string) {
	if err := addCommand(args); err != nil {
		fmt.Fprintf(os.Stderr, "rme: %v\n", err)
		os.Exit(1)
	}
}

func addCommand(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: rme add <name> <description>")
	}

	name := args[0]
	desc := args[1]

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("cannot determine current directory: %w", err)
	}

	path := filepath.Join(cwd, runmeFile)

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("no RUNME.sh found in current directory")
		}
		return fmt.Errorf("cannot read RUNME.sh: %w", err)
	}

	// Check for duplicate command
	existing, _ := parseFromSource(path)
	for _, c := range existing {
		if c.Name == name {
			return fmt.Errorf("command %q already exists in RUNME.sh", name)
		}
	}

	content := string(data)
	idx := strings.Index(content, upperMarker)
	if idx == -1 {
		return fmt.Errorf("RUNME.sh is missing the \"%s\" marker", upperMarker)
	}

	block := fmt.Sprintf("make_command \"%s\" \"%s\"\n%s(){\n  echo \"%s\"\n}\n\n", name, desc, name, name)

	newContent := content[:idx] + block + content[idx:]

	if err := os.WriteFile(path, []byte(newContent), 0755); err != nil {
		return fmt.Errorf("cannot write RUNME.sh: %w", err)
	}

	fmt.Printf("Added command %q to RUNME.sh\n", name)
	return nil
}
