---
name: sdd-cook
description: Specification-Driven AI Engineering. Controls AI software work through explicit artifacts, gates, traceability, decision control, change management, validation and locked evolutions.
---

# SDD Cook

SDD Cook is a Specification-Driven AI Engineering methodology. The agent executes the workflow through the `/sdd:*` commands.

## Commands

- `/sdd:init` — initialize or integrate SDD into a project
- `/sdd:new` — start a component or evolution
- `/sdd:plan` — create Plan, Tests and Tasks from accepted Specification
- `/sdd:run` — execute authorized Tasks
- `/sdd:validate` — validate implementation and lock accepted state
- `/sdd:change` — evolve accepted or locked artifacts through Change Management
- `/sdd:status` — show the real project state

## Core workflow

OBJECTIVE → DOMAIN → SPECIFICATION → PLAN → TESTS → TASKS → IMPLEMENTATION → VALIDATION → LOCKED

LOCKED is a lifecycle state, not a phase.

## Master decision rule

If a decision is derivable from accepted artifacts, decide and continue.

If it is not derivable, STOP and request a decision.

Never convert uncertainty into behavior.

## Authority by responsibility

- Objective: purpose, scope, constraints and acceptance criteria.
- Domain: conceptual meaning, entities, relationships, conceptual states and invariants.
- Specification: behavior and rules.
- Plan: execution and verification strategy.
- Tests: evidence of specified behavior.
- Tasks: executable work units.
- Implementation: technical realization.
- Validation: contextual evidence.
- Locked: accepted validated state.

No downstream artifact silently becomes authority over an upstream responsibility.

## Gates

Possible outcomes: PASS, FAIL, BLOCKED, INCONCLUSIVE.

PASS → continue.
FAIL → diagnose and route to earliest affected phase.
BLOCKED → stop and request the missing decision or artifact.
INCONCLUSIVE → investigate only with a new verifiable hypothesis.

## Anti-loop

Retry a correction only when a new verifiable hypothesis justifies it. Repeating the same failed action without new evidence is forbidden.

## Change control

Any change affecting behavior, scope, architecture or an accepted decision requires Change Management:

LOCKED → CHANGE → IMPACT ANALYSIS → DECISION → NEW EVOLUTION → VALIDATION → LOCKED

Never silently overwrite accepted state.

## Traceability

Maintain, where applicable:

Requirement → Domain → Specification → Test → Task → Implementation → Validation

## Adaptive rigor

Use the minimum methodology necessary for sufficient confidence and traceability. Auxiliary artifacts are created only when their value justifies their cost.

## Project state

`.spec/` is the persistent SDD state. Read actual artifacts before acting. Never infer project state from memory alone.

## Optional CLI

An optional deterministic CLI lives in `cli/`. It manages `.spec/` state: scaffold (`sdd init`), status (`sdd status`), gates (`sdd check`), evolution start (`sdd new`), acceptance (`sdd accept`), lock (`sdd lock`) and artifact display (`sdd show`). Chain artifacts carry a YAML frontmatter block with `status` (DRAFT/ACCEPTED/LOCKED). Use CLI state as evidence, never as a replacement for reading actual artifacts. All AI work still belongs to the agent.
