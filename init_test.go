package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestHandleInitDownload(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("#!/usr/bin/env bash\n# test RUNME.sh\n"))
	}))
	defer server.Close()

	dir := t.TempDir()
	origDir, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(origDir)

	err := downloadInit(server.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	path := filepath.Join(dir, "RUNME.sh")
	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("RUNME.sh not created: %v", err)
	}

	if info.Mode()&0755 != 0755 {
		t.Errorf("expected mode 0755, got %o", info.Mode())
	}

	data, _ := os.ReadFile(path)
	if string(data) != "#!/usr/bin/env bash\n# test RUNME.sh\n" {
		t.Errorf("unexpected content: %q", data)
	}
}

func TestHandleInitExistingFile(t *testing.T) {
	dir := t.TempDir()
	origDir, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(origDir)

	os.WriteFile(filepath.Join(dir, "RUNME.sh"), []byte("existing"), 0644)

	err := downloadInit("http://unused")
	if err == nil {
		t.Fatal("expected error for existing file, got nil")
	}
}

func TestHandleInitHTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	dir := t.TempDir()
	origDir, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(origDir)

	err := downloadInit(server.URL)
	if err == nil {
		t.Fatal("expected error for HTTP 404, got nil")
	}
}
