## 1. Create RUNME.sh

- [x] 1.1 Create `RUNME.sh` at project root with RUNME.sh boilerplate and `make_command` declarations for `build`, `test`, `fmt`, `vet`
- [x] 1.2 Implement `build` function with `go build -ldflags "-s -w -X main.version=0.1.0" -o rme .`
- [x] 1.3 Implement `test` function with `go test ./...`
- [x] 1.4 Implement `fmt` function with `gofmt -w .`
- [x] 1.5 Implement `vet` function with `go vet ./...`
- [x] 1.6 Make `RUNME.sh` executable (`chmod +x`)

## 2. Verify

- [x] 2.1 Run `rme --completions` in project directory and confirm `build`, `test`, `fmt`, `vet` appear
- [x] 2.2 Run each command and verify it works
