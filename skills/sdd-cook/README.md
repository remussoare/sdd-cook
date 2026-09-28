# SDD Cook v0.3

Specification-Driven AI Engineering as an agent Skill.

## Commands

`/sdd:init`
`/sdd:new`
`/sdd:plan`
`/sdd:run`
`/sdd:validate`
`/sdd:change`
`/sdd:status`

The methodology is interface-independent. The agent executes the `/sdd:*` commands; an optional deterministic CLI (`cli/`) manages the `.spec/` state: scaffold, status, gates and lifecycle transitions.

Install/copy this skill according to the host agent's Skill/command mechanism. The `commands/` files define the behavior of the `/sdd:*` interface.

## CLI (optional)

Build and run from `cli/`:

```sh
go build -o sdd .
./sdd init              # create the minimum .spec/ structure (never overwrites)
./sdd new <name>        # start a component; after LOCKED, starts a change request
./sdd status            # show the real SDD state
./sdd check             # evaluate the artifact-chain gate
./sdd show <artifact>   # print an artifact
./sdd accept <artifact> # record acceptance of a DRAFT artifact
./sdd lock              # lock the accepted validated state
```

Artifacts use a YAML frontmatter block (`artifact`, `component`, `evolution`, `status`). The CLI never invents artifacts or states; AI work (Specification, Plan, Tests, implementation, Validation) remains with the agent through the `/sdd:*` commands.

## Design

Design is transversal, not a separate phase. When relevant, SDD Cook can maintain `.spec/design/visual.md`, `.spec/design/interaction.md` and `.spec/design/design-decisions.md`.
