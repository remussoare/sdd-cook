# Tests

## PURPOSE

Provide evidence that the behavior defined in the Specification can be verified.

Tests do not define behavior.

## INPUT

* accepted Specification;
* accepted Plan.

## ANALYSIS

Derive test cases directly from the Specification.

As applicable:

* positive;
* negative;
* edge cases;
* transitions;
* regression;
* fixtures;
* test data;
* parametrization.

## ALLOWED

The agent may:

* derive test cases;
* create fixtures;
* create data;
* parametrize tests;
* organize suites;
* add regression tests derived from existing behavior;
* improve test determinism.

## FORBIDDEN

The agent may not:

* invent behavior;
* introduce requirements;
* modify the Specification to make a test pass;
* delete failures;
* modify expected results to fit the implementation.

## FAILURE ROUTING

When a Test fails:

```text
Implementation bug
→ Implementation

Test bug
→ Tests

Specification issue
→ Change Request
```

## OUTPUT

```text
TESTS
STATUS
TEST CASES
COVERAGE
FIXTURES
TEST DATA
REGRESSION SET
TRACEABILITY
FAILURES
OPEN DECISIONS
NEXT ACTION
```

## GATE

> Do the Tests provide sufficient evidence of the accepted behavior at the required rigor level?
