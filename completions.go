package main

import "fmt"

// printCompletions outputs bare command names for shell completion scripts.
func printCompletions() {
	path, err := detectRunme()
	if err != nil {
		// No RUNME.sh — output nothing silently
		return
	}

	// Try usage output first, fall back to source parsing
	commands, err := parseFromUsage(path)
	if err != nil || len(commands) == 0 {
		commands, _ = parseFromSource(path)
	}

	for _, c := range commands {
		fmt.Println(c.Name)
	}
}
