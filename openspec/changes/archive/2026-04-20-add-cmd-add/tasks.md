## 1. Core Implementation

- [x] 1.1 Create `add.go` with `handleAdd(args)` function: validate args (need name + description), check RUNME.sh exists, check for duplicate via `parseFromSource`, find marker, insert command block, rewrite file
- [x] 1.2 Add `"add"` case to subcommand dispatch in `main.go` and update help text

## 2. Testing

- [x] 2.1 Add tests in `add_test.go` covering: successful add, missing args, no RUNME.sh, duplicate command, missing marker
- [x] 2.2 Verify `go build` succeeds and manually test `rme add` against the project's own RUNME.sh
