# Traceability

The canonical SDD chain is:

```text
Requirement
    ↓
Domain
    ↓
Specification
    ↓
Test
    ↓
Task
    ↓
Implementation
    ↓
Validation
```

## Rules

Every functional requirement should trace to a Specification.

Every Specification should have relevant verification evidence.

Every Task should trace to accepted behavior.

Every implementation change should trace to a Task.

Every validation result should trace to expected behavior and evidence.

## Broken Traceability

Examples:

```text
Requirement without Specification
Specification without Test
Test without Specification
Task without Test/Specification
Code without Task
Validation without evidence
```

Broken traceability must be reported and resolved before acceptance when it affects functional correctness.
