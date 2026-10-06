# /sdd:validate

## PURPOSE

Validate the implementation against the accepted artifacts.

## INPUT

* Specification;
* Plan;
* Tests;
* Tasks;
* Implementation.

## ACTIONS

The agent must:

1. Run the required verification.
2. Compare expected and actual results.
3. Detect failures.
4. Route failures according to their cause.
5. Record evidence.
6. Determine the Validation state.

## STATES

```text
PASS
FAIL
BLOCKED
INCONCLUSIVE
```

## RULES

Never declare `PASS` without sufficient evidence.

Never hide a regression.

Never close an unresolved decision.

An environment limitation cannot be turned into `PASS`.

## LOCK

```text
PASS
↓
LOCKED
```

`FAIL`, `BLOCKED` and `INCONCLUSIVE` cannot produce `LOCKED`.
