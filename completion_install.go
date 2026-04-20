package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func installCompletion() {
	shell := detectShell()

	switch shell {
	case "fish":
		installFish()
	case "bash":
		installBash()
	case "zsh":
		installZsh()
	default:
		fmt.Fprintf(os.Stderr, "rme: could not detect shell (got %q)\n", shell)
		fmt.Fprintln(os.Stderr, "Use: rme completion <fish|bash|zsh> to get the script manually")
		os.Exit(1)
	}
}

func detectShell() string {
	shell := os.Getenv("SHELL")
	if shell == "" {
		return ""
	}
	base := filepath.Base(shell)
	// Handle cases like "fish", "bash", "zsh"
	for _, s := range []string{"fish", "bash", "zsh"} {
		if strings.Contains(base, s) {
			return s
		}
	}
	return base
}

func installFish() {
	home, err := os.UserHomeDir()
	if err != nil {
		fatal("cannot determine home directory: " + err.Error())
	}

	dir := filepath.Join(home, ".config", "fish", "completions")
	path := filepath.Join(dir, "rme.fish")
	writeCompletionFile(dir, path, fishCompletion)
}

func installBash() {
	home, err := os.UserHomeDir()
	if err != nil {
		fatal("cannot determine home directory: " + err.Error())
	}

	dir := filepath.Join(home, ".local", "share", "bash-completion", "completions")
	path := filepath.Join(dir, "rme")
	writeCompletionFile(dir, path, bashCompletion)
}

func installZsh() {
	// Try common writable fpath locations
	home, err := os.UserHomeDir()
	if err != nil {
		fatal("cannot determine home directory: " + err.Error())
	}

	candidates := []string{
		filepath.Join(home, ".zsh", "completions"),
		filepath.Join(home, ".zfunc"),
	}

	// Also check $fpath entries under home
	fpath := os.Getenv("FPATH")
	if fpath != "" {
		for _, p := range strings.Split(fpath, ":") {
			if strings.HasPrefix(p, home) {
				candidates = append([]string{p}, candidates...)
			}
		}
	}

	for _, dir := range candidates {
		path := filepath.Join(dir, "_rme")
		if err := writeCompletionFileIfPossible(dir, path, zshCompletion); err == nil {
			fmt.Printf("Installed zsh completion to %s\n", path)
			fmt.Println("Make sure this directory is in your $fpath and run: autoload -Uz compinit && compinit")
			return
		}
	}

	// Fallback: create ~/.zsh/completions
	dir := candidates[len(candidates)-1]
	path := filepath.Join(dir, "_rme")
	writeCompletionFile(dir, path, zshCompletion)
	fmt.Println("Add this to your .zshrc: fpath=(~/.zfunc $fpath); autoload -Uz compinit && compinit")
}

func writeCompletionFile(dir, path, content string) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		fatal("cannot create directory " + dir + ": " + err.Error())
	}

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		fatal("cannot write " + path + ": " + err.Error())
	}

	fmt.Printf("Installed completion to %s\n", path)
}

func writeCompletionFileIfPossible(dir, path, content string) error {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(content), 0644)
}

func fatal(msg string) {
	fmt.Fprintln(os.Stderr, "rme: "+msg)
	os.Exit(1)
}
