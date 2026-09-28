# Failure Routing

Failure is evidence that must be diagnosed and routed to the correct authority.

## Canonical Flow

```text
FAIL
 ↓
DIAGNOSE
 ↓
IDENTIFY EARLIEST AFFECTED AUTHORITY
 ↓
CORRECT
 ↓
REVALIDATE
```

## Routing

| Failure                     | Route                  |
| --------------------------- | ---------------------- |
| Implementation bug          | Implementation         |
| Test bug                    | Tests                  |
| Specification defect        | Change / Specification |
| Domain defect               | Change / Domain        |
| Missing functional decision | BLOCKED                |
| Environment limitation      | BLOCKED / INCONCLUSIVE |

## Retry Rule

The agent may retry only when a new verifiable hypothesis exists.

Repeatedly executing the same failed action without new evidence is forbidden.

## Objective

Resolve the earliest incorrect authority rather than patching downstream symptoms.
