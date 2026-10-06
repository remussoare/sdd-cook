# Workflow

## Canonical Lifecycle

```text
OBJECTIVE
    ↓
DOMAIN
    ↓
SPECIFICATION
    ↓
PLAN
    ↓
TESTS
    ↓
TASKS
    ↓
IMPLEMENTATION
    ↓
VALIDATION
    ↓
LOCKED
```

`LOCKED` is a lifecycle state, not a phase.

Each phase defines:

```text
PURPOSE
INPUT
ANALYSIS
ALLOWED
FORBIDDEN
ACTIONS
OUTPUT
GATE
```

The authority of each phase is defined in `principles/authority.md`.
