package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const runmeURL = "https://raw.githubusercontent.com/mipmip/RUNME.sh/master/RUNME.sh"

func handleInit() {
	if err := downloadInit(runmeURL); err != nil {
		fmt.Fprintf(os.Stderr, "rme: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Created RUNME.sh")
}

func downloadInit(url string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("cannot determine current directory: %w", err)
	}

	path := filepath.Join(cwd, runmeFile)

	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("RUNME.sh already exists in current directory")
	}

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download RUNME.sh: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download RUNME.sh: HTTP %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if err := os.WriteFile(path, body, 0755); err != nil {
		return fmt.Errorf("failed to write RUNME.sh: %w", err)
	}

	return nil
}
