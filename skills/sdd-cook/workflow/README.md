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

`LOCKED` es un estado del ciclo de vida, no una fase.

Cada fase define:

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

La autoridad de cada fase está definida en `principles/authority.md`.
