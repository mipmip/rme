package main

import (
	"fmt"
	"os"
)

const fishCompletion = `# rme fish completion
complete -c rme -f -a '(rme --completions)'
`

const bashCompletion = `# rme bash completion
_rme() {
    local cur="${COMP_WORDS[COMP_CWORD]}"
    COMPREPLY=($(compgen -W "$(rme --completions)" -- "$cur"))
}
complete -F _rme rme
`

const zshCompletion = `#compdef rme
# rme zsh completion
_rme() {
    compadd $(rme --completions)
}
compdef _rme rme
`

func handleCompletion(args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Usage: rme completion <fish|bash|zsh|install>")
		os.Exit(1)
	}

	switch args[0] {
	case "fish":
		fmt.Print(fishCompletion)
	case "bash":
		fmt.Print(bashCompletion)
	case "zsh":
		fmt.Print(zshCompletion)
	case "install":
		installCompletion()
	default:
		fmt.Fprintf(os.Stderr, "rme: unknown shell: %s (use fish, bash, zsh, or install)\n", args[0])
		os.Exit(1)
	}
}
