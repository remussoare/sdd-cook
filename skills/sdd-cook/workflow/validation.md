# Validation

## PURPOSE

Obtain sufficient, reliable and traceable evidence that the implementation satisfies the Specification and the acceptance criteria in the relevant context.

The fundamental difference is:

```text
IMPLEMENTATION
= did I build it correctly?

VALIDATION
= do we have sufficient evidence that it works correctly in context?
```

## INPUT

* Objective;
* Domain;
* Specification;
* Plan;
* Tests;
* Tasks;
* Implementation.

## ANALYSIS

Compare:

```text
EXPECTED
vs
ACTUAL
```

Validation must use the level of evidence appropriate to the context.

It may include, as applicable:

* unit;
* integration;
* scenarios;
* regression;
* datasets;
* visual;
* performance;
* backtest.

## ALLOWED

The agent may execute the verifications needed to obtain sufficient evidence.

## FORBIDDEN

The agent may not:

* declare PASS without sufficient evidence;
* hide regressions;
* close unresolved decisions;
* turn an environment limitation into PASS;
* declare success based solely on an insufficient check.

## STATES

```text
PASS
FAIL
BLOCKED
INCONCLUSIVE
```

## OUTPUT

```text
VALIDATION
STATUS
EVIDENCE
TEST RESULTS
ACCEPTANCE CRITERIA
REGRESSION
FAILURES
TRACEABILITY
DECISION
NEXT ACTION
```

## FAILURE ROUTING

```text
FAIL
 ↓
DIAGNOSE
 ↓
CORRECT
 ↓
REVALIDATE
```

## GATE

PASS may only be declared when sufficient and reliable evidence exists.

```text
PASS
 ↓
LOCKED
```

`FAIL`, `BLOCKED` and `INCONCLUSIVE` cannot produce LOCKED.
