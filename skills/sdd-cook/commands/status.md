# /sdd:status

## PURPOSE

Show the real SDD project state.

## ACTIONS

Read `.spec/` and show the existing state.

## OUTPUT

```text
SDD STATUS

Component:
Evolution:
Current phase:
Objective:
Domain:
Specification:
Plan:
Tests:
Tasks:
Implementation:
Validation:
Lifecycle:
Blockers:
Drift:
Next action:
```

## RULES

Show only states that can be determined from existing artifacts.

Never invent states when artifacts are missing.

Identify:

* missing dependencies;
* blockers;
* drift;
* the next authorized action.
