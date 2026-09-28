# sdd-cook

Specification-Driven AI Engineering as an installable toolkit: an agent skill
plus a small deterministic CLI that manages the `.spec/` state.

- **Skill** (`skills/sdd-cook/`): the `/sdd:*` commands —
  `init`, `new`, `plan`, `run`, `validate`, `change`, `status`.
- **CLI** (`sdd`): manages `.spec/` state (`init`, `new`, `status`, `check`,
  `show`, `accept`, `lock`) and installs/updates the skill
  (`install`, `update`). AI work stays with the agent.

## Install

### With your AI assistant (recommended)

Give your AI this URL:

> Install sdd-cook from `https://github.com/remussoare/sdd-cook`

The AI downloads (or builds) the `sdd` binary and runs `sdd install`, which
detects your tools, asks for scope (this repo vs home) and hosts, then writes
the skill payload. Re-running is idempotent.

### Manual

Go developers:

```sh
go install github.com/remussoare/sdd-cook@latest
sdd install
```

Everyone else (prebuilt binary, no Go needed) — macOS/Linux:

```sh
curl -fsSL https://raw.githubusercontent.com/remussoare/sdd-cook/main/install.sh | sh
```

Windows (PowerShell):

```powershell
irm https://raw.githubusercontent.com/remussoare/sdd-cook/main/install.ps1 | iex
```

Flags: `--scope repo|home`, `--hosts opencode,cursor,vscode|all`,
`--dry-run`, `--yes`. Env: `INSTALL_SCOPE=project|global`,
`INSTALL_TARGET=all|opencode,cursor,vscode`, `SDD_CONFIG_DIR` (registry override).

## Use

In your project:

```sh
sdd init      # create the minimum .spec/ structure (never overwrites)
sdd status    # show the real SDD state
```

Then work with the agent: `/sdd:new`, `/sdd:plan`, `/sdd:run`, `/sdd:validate`,
`/sdd:change`. The CLI enforces gates (`sdd check`) and lifecycle
(`sdd accept`, `sdd lock`).

## Update

```sh
sdd update              # re-apply skill payloads to tracked installs
go install github.com/remussoare/sdd-cook@latest   # upgrade the CLI first
```

`sdd update` also reports when a newer release exists.

## Supported hosts

| Host | Payload |
|---|---|
| OpenCode | `.agents/skills/sdd-cook/` (full skill) |
| Cursor | `.cursor/agents/sdd-cook.md` + `.cursor/rules/sdd-cook.mdc` |
| VS Code / Copilot | `.github/skills/sdd-cook/SKILL.md` + `AGENTS.md` (created only) |

One canonical source (`skills/`), N generated targets — adding a host means
adding one generator in `internal/transpile`.

## Development

Requires Go 1.22+.

```sh
go build ./...     # build
go vet ./...       # vet
go test ./...      # tests (gates, lifecycle, spec)
go build -o sdd ./cmd/sdd
```

Release builds stamp the version:

```sh
go build -ldflags "-X github.com/remussoare/sdd-cook/internal/version.Version=v0.1.0" ./cmd/sdd
```

CI builds and tests on macOS, Linux and Windows on every push; pushing a
`v*` tag builds the six platform binaries and publishes the GitHub release
that `install.sh` / `install.ps1` download.

## Design

Design is transversal, not a separate phase (`skills/sdd-cook/workflow/design.md`).
