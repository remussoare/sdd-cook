# /sdd:run

## Purpose

Execute only authorized planned work.

## Critical rule

`/sdd:run` MUST NOT silently create a Plan, invent Tasks, invent Tests or invent Specification.

## Procedure

1. Read actual `.spec/` state.
2. Require an accepted Plan and executable Tasks.
3. If Plan is missing, return:

```text
PLAN: MISSING
Cannot run.
Required artifact: PLAN
Run: /sdd:plan
```

4. Execute pending authorized Tasks.
5. Run required available Tests.
6. Record files changed, results, architecture compliance and drift.
7. BLOCK if implementation requires an unaccepted functional decision.

## Refactoring

Allowed only when behavior before equals behavior after.

## Forbidden

- invent behavior
- change Specification to justify code
- change Tests to hide failures
- implement unplanned work
