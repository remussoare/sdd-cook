# Spec-Driven Development

**sdd-cook** — Specification-Driven AI Engineering as an installable toolkit:
an agent skill that controls AI software work, plus a small deterministic CLI
that manages the `.spec/` state.

## What it is

AI coding assistants are powerful but unpredictable when requirements live only
in chat history. sdd-cook adds an explicit engineering layer so the human and
the AI agree on **what to build** before any code is written, verify it with
evidence, and evolve it without silently breaking what was accepted:

- **Skill** (`skills/sdd-cook/`): the `/sdd:*` commands the agent executes —
  `init`, `new`, `plan`, `run`, `validate`, `change`, `status`.
- **CLI** (`sdd`): manages `.spec/` state (`init`, `new`, `status`, `check`,
  `show`, `accept`, `lock`) and installs/updates the skill
  (`install`, `update`). AI work stays with the agent.

## The flow

The core workflow runs through explicit artifacts with gates between them:

```text
OBJECTIVE → DOMAIN → SPECIFICATION → PLAN → TESTS → TASKS → IMPLEMENTATION → VALIDATION → LOCKED
```

LOCKED is a lifecycle state, not a phase. After LOCKED, evolution goes
through Change Management:

```text
LOCKED → CHANGE → IMPACT ANALYSIS → DECISION → NEW EVOLUTION → PLAN → TESTS → TASKS → IMPLEMENTATION → VALIDATION → LOCKED vN
```

When the change has no implementation impact, the new evolution goes straight
to VALIDATION.

## Skill commands (`/sdd:*`)

| Command | Role in the flow |
|---|---|
| `/sdd:init` | Initialize SDD in a project or integrate it into an existing one without silently overwriting work. |
| `/sdd:new` | Start a component or evolution without inventing decisions; BLOCKs when a required decision is not derivable. |
| `/sdd:plan` | Transform accepted Specification into Plan → Tests → Tasks. Never implements code. |
| `/sdd:run` | Execute only authorized planned Tasks; refuses to run without an accepted Plan. |
| `/sdd:validate` | Produce traceable evidence that the implementation satisfies the Specification; on PASS the state becomes LOCKED. |
| `/sdd:change` | Evolve accepted or LOCKED artifacts through explicit Change Management with Impact Analysis. |
| `/sdd:status` | Show the real project state from `.spec/` — never invents states for missing artifacts. |

## Concepts it follows

Transversal functions applied across the whole workflow:

- **Master decision rule** — if a decision is derivable from accepted
  artifacts, decide and continue; if not, STOP and request a decision. Never
  convert uncertainty into behavior.
- **Authority by responsibility** — Objective owns purpose/scope/constraints;
  Domain owns meaning/entities/invariants; Specification owns behavior/rules;
  Plan owns execution strategy; Tests own evidence; Tasks own work units;
  Implementation owns technical realization; Validation owns contextual
  evidence. No downstream artifact silently becomes authority over an upstream
  responsibility.
- **Gates** — every transition ends in PASS (continue), FAIL (diagnose and
  route to the earliest affected phase), BLOCKED (stop, request the missing
  decision or artifact) or INCONCLUSIVE (investigate only with a new
  verifiable hypothesis). No gate is ever passed by assumption.
- **Decision control** — DERIVABLE → decide and continue. NOT DERIVABLE →
  BLOCKED. A proposal is not a decision.
- **Anti-loop** — retry a correction only when a new verifiable hypothesis
  justifies it; repeating the same failed action without new evidence is
  forbidden.
- **Change control** — any change affecting behavior, scope, architecture or
  an accepted decision requires Change Management. Never silently overwrite
  accepted state. LOCKED is versioned acceptance, not immutability.
- **Traceability** — maintain Requirement → Domain → Specification → Test →
  Task → Implementation → Validation; detect missing or orphaned links without
  inventing them.
- **Adaptive rigor** — use the minimum methodology necessary for sufficient
  confidence and traceability; auxiliary artifacts (ADRs, spikes, prototypes,
  diagrams, impact analysis) are created only when their value justifies
  their cost. They are never workflow phases.
- **Failure routing** — FAIL → DIAGNOSE → CORRECT → REVALIDATE, always routed
  to the earliest affected phase; never patch downstream artifacts to hide an
  upstream problem.
- **Design control** — design is transversal, not a phase. Dedicated artifacts
  (`design/visual.md`, `design/interaction.md`, `design/design-decisions.md`)
  exist only when presentation or interaction materially affects the accepted
  product result, and must never silently change functional behavior.

## Project state (`.spec/`)

`.spec/` is the persistent SDD state: one directory per artifact
(`objective/`, `domain/`, `specifications/`, `plan/`, `tests/`, `tasks/`,
`validation/`) plus `changes/` and `traceability/`. Chain artifacts carry a
YAML frontmatter block with `status` (DRAFT / ACCEPTED / LOCKED). The agent
reads actual artifacts before acting and never infers project state from
memory alone.

## CLI (`sdd`)

```sh
sdd init              # create the minimum .spec/ structure (never overwrites)
sdd new <name>        # start a component or (after LOCKED) an evolution
sdd status            # show the real SDD state
sdd check             # evaluate the artifact-chain gate
sdd show <artifact>   # print an artifact (objective, domain, ...)
sdd accept <artifact> # record acceptance of a DRAFT artifact
sdd lock              # lock the accepted validated state
sdd install [flags]   # install the skill for AI tools (scope, hosts)
sdd update [flags]    # re-apply skill payloads to tracked installs
sdd version           # print version and platform
```

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
go install github.com/remussoare/sdd-cook/cmd/sdd@latest
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
go install github.com/remussoare/sdd-cook/cmd/sdd@latest   # upgrade the CLI first
sdd update              # re-apply skill payloads to tracked installs
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
go test ./...      # tests (gates, lifecycle, spec, install, transpile)
go build -o sdd ./cmd/sdd
```

Release builds stamp the version:

```sh
go build -ldflags "-X github.com/remussoare/sdd-cook/internal/version.Version=v0.1.0" ./cmd/sdd
```

CI builds and tests on macOS, Linux and Windows on every push; pushing a
`v*` tag builds the six platform binaries and publishes the GitHub release
that `install.sh` / `install.ps1` download.
