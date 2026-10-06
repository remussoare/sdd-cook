# Domain

## PURPOSE

Build the conceptual model of the problem.

Domain defines:

* concepts;
* entities;
* value objects;
* relationships;
* conceptual states;
* invariants;
* domain boundaries;
* common language.

DDD depth adapts to the real complexity of the problem.

## INPUT

* accepted Objective;
* available domain knowledge;
* accepted decisions.

## ANALYSIS

Identify:

* core concepts;
* entities;
* value objects;
* relationships;
* boundaries;
* conceptual states;
* invariants;
* potential semantic conflicts.

## ALLOWED

The agent may:

* identify concepts derived from the Objective;
* structure entities;
* identify relationships;
* identify value objects;
* establish the conceptual structure;
* derive conceptual invariants when they are clearly implicit;
* adapt the depth of the model to the complexity.

## FORBIDDEN

The agent may not:

* invent behavior rules;
* invent behavioral states;
* invent transitions;
* resolve semantic ambiguities through an assumption when they affect behavior;
* introduce functional requirements.

## ACTIONS

1. Read the Objective.
2. Identify concepts.
3. Identify entities and value objects.
4. Establish relationships.
5. Identify conceptual states and invariants.
6. Detect conflicts.
7. Record open decisions.
8. Prepare the domain for Specification.

## OUTPUT

```text
DOMAIN
STATUS
COMPLEXITY
CONCEPTS
ENTITIES
VALUE OBJECTS
RELATIONSHIPS
STATES
INVARIANTS
BOUNDARIES
OPEN DECISIONS
CONFLICTS
TRACEABILITY
NEXT ACTION
```

## GATE

> Is there sufficient conceptual clarity to specify behavior without inventing concepts or decisions?
