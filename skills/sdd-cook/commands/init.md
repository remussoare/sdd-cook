# /sdd:init

## Purpose

Initialize SDD Cook in a project or integrate it into an existing project without silently overwriting existing work.

## Procedure

1. Inspect the repository before modifying anything.
2. Detect project type, stack, structure, commands, tests, documentation and existing agent instructions.
3. Detect existing `.spec/`, Constitution, `AGENTS.md`, decisions and architecture documentation.
4. For a new project, prepare the minimum SDD structure required.
5. For an existing project, audit current artifacts and detect conflicts or drift before integration.
6. Create or update only what is necessary; never silently overwrite existing accepted information.
7. Prepare Constitution and `AGENTS.md` only when needed and preserve existing content unless an explicit change is authorized.
8. Run the initialization gate.

## Minimum project structure

```text
.spec/
├── objective/
├── domain/
├── specifications/
├── plan/
├── tests/
├── tasks/
├── validation/
├── changes/
└── traceability/
```

Do not create unnecessary empty structures merely because the methodology mentions them.

## Initialization gate

PASS only when the project can proceed with SDD without silent conflicts or missing prerequisites.

## Output

Report:
- project type
- discovery findings
- existing SDD artifacts
- created/updated artifacts
- conflicts/drift
- initialization status
- next action

If blocked, state the exact reason and required decision.
