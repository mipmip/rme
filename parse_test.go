package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseUsageOutput(t *testing.T) {
	output := `
Usage: ./RUNME.sh [command]

Commands:
  deploy          Deploy to production
  test            Run the test suite
  clean           Remove build artifacts

`
	commands := parseUsageOutput(output)

	if len(commands) != 3 {
		t.Fatalf("expected 3 commands, got %d", len(commands))
	}

	expected := []struct {
		name string
		desc string
	}{
		{"deploy", "Deploy to production"},
		{"test", "Run the test suite"},
		{"clean", "Remove build artifacts"},
	}

	for i, e := range expected {
		if commands[i].Name != e.name {
			t.Errorf("command %d: expected name %q, got %q", i, e.name, commands[i].Name)
		}
		if commands[i].Desc != e.desc {
			t.Errorf("command %d: expected desc %q, got %q", i, e.desc, commands[i].Desc)
		}
	}
}

func TestParseUsageOutputEmpty(t *testing.T) {
	commands := parseUsageOutput("")
	if len(commands) != 0 {
		t.Fatalf("expected 0 commands, got %d", len(commands))
	}
}

func TestParseFromSource(t *testing.T) {
	content := `#!/usr/bin/env bash
CMDS=(); DESC=(); NARGS=$#; ARG1=$1
make_command "deploy" "Deploy to production"
deploy(){ echo deploying; }
make_command "test" "Run tests"
test(){ echo testing; }
runme
`
	dir := t.TempDir()
	path := filepath.Join(dir, "RUNME.sh")
	if err := os.WriteFile(path, []byte(content), 0755); err != nil {
		t.Fatal(err)
	}

	commands, err := parseFromSource(path)
	if err != nil {
		t.Fatal(err)
	}

	if len(commands) != 2 {
		t.Fatalf("expected 2 commands, got %d", len(commands))
	}

	if commands[0].Name != "deploy" {
		t.Errorf("expected name %q, got %q", "deploy", commands[0].Name)
	}
	if commands[1].Name != "test" {
		t.Errorf("expected name %q, got %q", "test", commands[1].Name)
	}
}
