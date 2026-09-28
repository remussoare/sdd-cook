# Change Control

Accepted artifacts represent an accepted project decision.

They must never be silently replaced.

## Lifecycle

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
LOCKED vN
```

Where implementation impact exists, the new evolution may require:

```text
PLAN
 ↓
TESTS
 ↓
TASKS
 ↓
IMPLEMENTATION
 ↓
VALIDATION
```

## Change Status

Allowed statuses:

* PROPOSED
* ANALYZING
* AWAITING_DECISION
* APPROVED
* IMPLEMENTING
* VALIDATING
* ACCEPTED
* REJECTED
* CANCELLED
* BLOCKED

## History

Previous accepted evolutions remain preserved.

A change creates a new evolution rather than rewriting history.
