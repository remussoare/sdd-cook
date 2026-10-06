# /sdd:init

## PURPOSE

Initialize or integrate SDD into a project.

`/sdd:init` does not create a component.

## PROJECT START

```text
PROJECT START
↓
DISCOVERY
↓
prepare Constitution
prepare AGENTS.md
prepare needed .spec/
↓
INITIALIZATION GATE
↓
OBJECTIVE
```

## EXISTING PROJECT

```text
PROJECT EXISTS
↓
DISCOVERY
↓
AUDIT
↓
existing Constitution / AGENTS / .spec / stack
↓
detect conflicts/drift
↓
integrate without silent overwrite
↓
INITIALIZATION GATE
```

## ACTIONS

The agent must:

* inspect the repository;
* detect the stack;
* detect `AGENTS.md`;
* detect the Constitution;
* detect `.spec/`;
* detect existing decisions;
* detect conflicts;
* detect drift;
* prepare the minimal necessary structure;
* prepare the Constitution and `AGENTS.md`;
* leave the project ready for `/sdd:new`.

## RULE

Never silently overwrite existing artifacts.
