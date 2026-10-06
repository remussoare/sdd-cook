# Specification

## PURPOSE

Define the behavioral contract explicitly and deterministically.

Specification establishes:

* rules;
* conditions;
* inputs;
* outputs;
* preconditions;
* postconditions;
* states;
* transitions;
* invariants;
* errors;
* edge cases;
* behavioral dependencies.

## INPUT

* accepted Objective;
* accepted Domain;
* accepted decisions.

## ANALYSIS

The agent must turn the conceptual model into verifiable behavior.

It must identify:

* what happens;
* when it happens;
* under which conditions;
* which inputs are required;
* which result is expected;
* which states exist;
* how transitions occur;
* which errors are possible;
* which edge cases must be considered.

## ALLOWED

The agent may:

* make derivable rules explicit;
* split complex behavior;
* remove linguistic ambiguity;
* structure conditions;
* define cases derived directly from accepted behavior.

## FORBIDDEN

The agent may not:

* use the implementation as authority;
* invent behavior;
* invent states;
* invent transitions;
* introduce requirements;
* modify behavior to fit the implementation;
* resolve a functional ambiguity through an assumption.

## ACTIONS

1. Read the Objective and the Domain.
2. Identify behavior.
3. Define deterministic rules.
4. Define inputs and outputs.
5. Define states and transitions where applicable.
6. Define invariants.
7. Define errors and edge cases.
8. Identify dependencies.
9. Record open decisions.
10. Prepare the Specification for planning and verification.

## OUTPUT

```text
SPECIFICATION
STATUS
RULES
INPUTS
OUTPUTS
PRECONDITIONS
POSTCONDITIONS
STATES
TRANSITIONS
INVARIANTS
ERRORS
EDGE CASES
DEPENDENCIES
OPEN DECISIONS
CONFLICTS
TRACEABILITY
NEXT ACTION
```

## GATE

> Can the behavior be turned into deterministic verification without introducing unsupported assumptions?
