# Implementation

## PURPOSE

Technically realize the accepted Tasks.

Implementation turns accepted decisions into executable code or configuration.

## INPUT

* accepted Tasks;
* Specification;
* Tests;
* technical constraints of the project.

## ANALYSIS

Determine the technical realization without changing accepted behavior.

## ALLOWED

The agent may:

* write code;
* refactor;
* optimize;
* change internal structures;
* change names;
* create abstractions;
* reuse components;
* fix implementation bugs.

## REFACTORING

A refactor is valid when:

```text
BEHAVIOR BEFORE = BEHAVIOR AFTER
```

## FORBIDDEN

The agent may not:

* change rules;
* change states;
* change transitions;
* change outputs;
* change errors;
* change priorities;
* change scope;
* change acceptance criteria;
* modify Tests to hide failures;
* modify the Specification to justify code;
* add unspecified behavior;
* resolve functional ambiguity;
* change architecture when the change requires a non-accepted decision.

## ACTIONS

1. Select authorized Tasks.
2. Implement them.
3. Run the relevant Tests.
4. Compare results with expectations.
5. Detect drift.
6. Report modified files.
7. Report completed Tasks.
8. Prepare evidence for Validation.

## OUTPUT

```text
IMPLEMENTATION
STATUS
TASKS COMPLETED
FILES CHANGED
BEHAVIOR IMPLEMENTED
TEST RESULTS
ARCHITECTURE COMPLIANCE
TRACEABILITY
DRIFT DETECTED
BLOCKERS
NEXT ACTION
```

## GATE

The implementation must satisfy:

* Tasks;
* Specification;
* accepted constraints;
* traceability;
* absence of unresolved functional drift.

Important:

```text
TESTS PASS ≠ IMPLEMENTATION GATE PASS
```

Tests passing does not by itself mean Implementation has passed its Gate.
