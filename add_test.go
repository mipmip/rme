package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const testRunme = `#!/usr/bin/env bash
CMDS=();DESC=();NARGS=$#;ARG1=$1;make_command(){ CMDS+=($1);DESC+=("$2");};usage(){ :;};runme(){ :;}

##### PLACE YOUR COMMANDS BELOW #####

make_command "existing" "An existing command"
existing(){
  echo "existing"
}

##### PLACE YOUR COMMANDS ABOVE #####

runme
`

func chdir(t *testing.T, dir string) {
	t.Helper()
	orig, _ := os.Getwd()
	os.Chdir(dir)
	t.Cleanup(func() { os.Chdir(orig) })
}

func TestAddCommandSuccess(t *testing.T) {
	dir := t.TempDir()
	chdir(t, dir)
	os.WriteFile(filepath.Join(dir, "RUNME.sh"), []byte(testRunme), 0755)

	err := addCommand([]string{"deploy", "Deploy to production"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, _ := os.ReadFile(filepath.Join(dir, "RUNME.sh"))
	content := string(data)

	if !strings.Contains(content, `make_command "deploy" "Deploy to production"`) {
		t.Error("missing make_command declaration")
	}
	if !strings.Contains(content, "deploy(){\n  echo \"deploy\"\n}") {
		t.Error("missing function stub")
	}
	// Verify it's before the marker
	markerIdx := strings.Index(content, upperMarker)
	deployIdx := strings.Index(content, `make_command "deploy"`)
	if deployIdx > markerIdx {
		t.Error("command block should be before the upper marker")
	}
}

func TestAddCommandMissingArgs(t *testing.T) {
	err := addCommand([]string{})
	if err == nil {
		t.Fatal("expected error for missing args")
	}

	err = addCommand([]string{"deploy"})
	if err == nil {
		t.Fatal("expected error for missing description")
	}
}

func TestAddCommandNoRunme(t *testing.T) {
	dir := t.TempDir()
	chdir(t, dir)

	err := addCommand([]string{"deploy", "Deploy"})
	if err == nil {
		t.Fatal("expected error for missing RUNME.sh")
	}
}

func TestAddCommandDuplicate(t *testing.T) {
	dir := t.TempDir()
	chdir(t, dir)
	os.WriteFile(filepath.Join(dir, "RUNME.sh"), []byte(testRunme), 0755)

	err := addCommand([]string{"existing", "Another existing"})
	if err == nil {
		t.Fatal("expected error for duplicate command")
	}
	if !strings.Contains(err.Error(), "already exists") {
		t.Errorf("expected 'already exists' error, got: %v", err)
	}
}

func TestAddCommandMissingMarker(t *testing.T) {
	dir := t.TempDir()
	chdir(t, dir)
	os.WriteFile(filepath.Join(dir, "RUNME.sh"), []byte("#!/bin/bash\necho hi\n"), 0755)

	err := addCommand([]string{"deploy", "Deploy"})
	if err == nil {
		t.Fatal("expected error for missing marker")
	}
	if !strings.Contains(err.Error(), "marker") {
		t.Errorf("expected 'marker' error, got: %v", err)
	}
}
