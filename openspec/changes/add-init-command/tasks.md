## 1. Core Implementation

- [x] 1.1 Create `init.go` with `handleInit()` function that checks for existing RUNME.sh, downloads from GitHub raw URL, and writes to `./RUNME.sh` with 0755 permissions
- [x] 1.2 Add `"init"` case to the subcommand dispatch in `main.go`

## 2. Testing

- [x] 2.1 Add tests in `init_test.go` covering: successful download (with test HTTP server), existing file rejection, and HTTP error handling
- [x] 2.2 Verify `go build` succeeds and manually test `rme init` in a temp directory
