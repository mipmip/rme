# rme

> `rme` is to `RUNME.sh` what `make` is to `Makefile` — except `rme` and `RUNME.sh` are simple enough to understand in 5 minutes.

`rme` is a command launcher for [RUNME.sh](https://github.com/mipmip/RUNME.sh). It detects `RUNME.sh` in the current directory and provides shell completions.

## Usage

```bash
$ cd my-project
$ rme <TAB>
deploy  test  clean  build
$ rme deploy
Deploying...
```

| Command | Behavior |
|---------|----------|
| `rme` | Show RUNME.sh usage |
| `rme deploy` | Run the `deploy` command from RUNME.sh |
| `rme <TAB>` | Auto-complete available commands |
| `rme --help` | Show rme help |
| `rme --version` | Show version |

## Install

<details>
<summary>Nix/NixOS</summary>

**Run directly:**

```bash
nix run github:mipmip/rme -- --help
```

**Temporary shell:**

```bash
nix shell github:mipmip/rme
rme --version
```

**System flake integration:**

```nix
{
  inputs.rme.url = "github:mipmip/rme";

  outputs = { self, nixpkgs, rme }: {
    nixosConfigurations.myhost = nixpkgs.lib.nixosSystem {
      modules = [{
        environment.systemPackages = [ rme.packages.x86_64-linux.default ];
      }];
    };
  };
}
```

</details>

<details>
<summary>Go</summary>

```bash
go install github.com/mipmip/rme@latest
```

</details>

## Shell Completions

```bash
# Auto-install (detects your shell)
rme completion install

# Or get the script for manual wiring
rme completion fish
rme completion bash
rme completion zsh
```

## How It Works

`rme` is a thin wrapper. It finds `RUNME.sh` in the current directory and passes commands through via `exec`. No child process, no signal forwarding — your shell talks directly to RUNME.sh.

The only thing `rme` adds is discoverability: a known command name that shell completions can hook into.

## License

MIT License - See [LICENSE](LICENSE) for details.
