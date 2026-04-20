package main

import (
	"bufio"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// command represents a parsed RUNME.sh command.
type command struct {
	Name string
	Desc string
}

// parseFromUsage runs ./RUNME.sh with no args and parses the Commands: section.
func parseFromUsage(path string) ([]command, error) {
	cmd := exec.Command(path)
	out, _ := cmd.Output() // no-args always exits non-zero (usage), ignore error

	return parseUsageOutput(string(out)), nil
}

// parseUsageOutput extracts commands from RUNME.sh usage text.
func parseUsageOutput(output string) []command {
	var commands []command
	inCommands := false
	scanner := bufio.NewScanner(strings.NewReader(output))

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "Commands:") {
			inCommands = true
			continue
		}

		if inCommands {
			// Empty line ends the commands section
			if strings.TrimSpace(line) == "" {
				break
			}

			// Parse "  name           description"
			trimmed := strings.TrimLeft(line, " ")
			parts := regexp.MustCompile(`\s{2,}`).Split(trimmed, 2)
			if len(parts) >= 1 && parts[0] != "" {
				c := command{Name: parts[0]}
				if len(parts) == 2 {
					c.Desc = parts[1]
				}
				commands = append(commands, c)
			}
		}
	}

	return commands
}

// parseFromSource reads the RUNME.sh file and extracts make_command calls via regex.
var makeCommandRe = regexp.MustCompile(`make_command\s+"([^"]+)"\s+"([^"]*)"`)

func parseFromSource(path string) ([]command, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var commands []command
	matches := makeCommandRe.FindAllStringSubmatch(string(data), -1)
	for _, m := range matches {
		commands = append(commands, command{Name: m[1], Desc: m[2]})
	}

	return commands, nil
}
