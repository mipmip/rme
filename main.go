package main

import (
	"fmt"
	"os"
	"syscall"
)

var version = "dev"

func main() {
	args := os.Args[1:]

	// No args or non-flag args: dispatch to RUNME.sh
	if len(args) == 0 {
		dispatch(nil)
		return
	}

	first := args[0]

	// rme's own flags
	switch first {
	case "--help":
		printHelp()
		return
	case "--version":
		fmt.Println("rme " + version)
		return
	case "--completions":
		printCompletions()
		return
	case "completion":
		handleCompletion(args[1:])
		return
	}

	// Everything else: pass through to RUNME.sh
	dispatch(args)
}

func dispatch(args []string) {
	path, err := detectRunme()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	argv := append([]string{path}, args...)
	env := os.Environ()

	if err := syscall.Exec(path, argv, env); err != nil {
		fmt.Fprintf(os.Stderr, "rme: failed to execute RUNME.sh: %v\n", err)
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Print(`rme — a command launcher for RUNME.sh

Usage:
  rme [command]          Run a RUNME.sh command
  rme                    Show RUNME.sh usage
  rme completion <shell> Print completion script (fish, bash, zsh)
  rme completion install Auto-install shell completions

Flags:
  --help                 Show this help
  --version              Show version
`)
}
