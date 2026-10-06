# /sdd:change

## PURPOSE

Manage changes over an accepted or `LOCKED` state.

## FLOW

```text
LOCKED
↓
CHANGE
↓
IMPACT ANALYSIS
↓
DECISION
↓
NEW EVOLUTION
↓
VALIDATION
↓
LOCKED v2
```

When the change requires implementation:

```text
NEW EVOLUTION
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
LOCKED v2
```

## ACTIONS

1. Identify the requested change.
2. Classify its impact.
3. If it affects behavior, scope, architecture or an accepted decision, create a Change Request.
4. Perform the Impact Analysis.
5. Request a decision when the change is not derivable.
6. Never silently replace an accepted artifact.
7. Update while maintaining traceability.
8. Replan, test, implement and validate as applicable.
9. Preserve the previous evolution.

## STATUSES

```text
PROPOSED
ANALYZING
AWAITING_DECISION
APPROVED
IMPLEMENTING
VALIDATING
ACCEPTED
REJECTED
CANCELLED
BLOCKED
```
