# Objective

## PURPOSE

Define precisely which problem the agent is authorized to solve.

The Objective establishes:

* purpose;
* scope;
* out of scope;
* constraints;
* acceptance criteria.

It does not design the domain or the solution.

## INPUT

* the user request;
* accepted project context;
* existing constraints;
* previously accepted decisions when relevant.

## ANALYSIS

The agent must:

* structure the existing information;
* remove linguistic ambiguity;
* separate scope from out of scope;
* derive acceptance criteria directly;
* detect contradictions;
* detect missing necessary information;
* identify open decisions.

## ALLOWED

The agent may:

* reorganize information;
* improve linguistic precision;
* separate scope from out of scope;
* derive directly implicit acceptance criteria;
* adapt the level of detail;
* propose a clearer structure.

## FORBIDDEN

The agent may not:

* invent requirements;
* expand scope;
* reduce scope;
* change priorities;
* invent constraints;
* assume decisions;
* turn a technical recommendation into a functional requirement;
* remove an existing constraint;
* modify the objective to make implementation easier.

## ACTIONS

1. Analyze the request.
2. Identify objective and scope.
3. Separate what remains out of scope.
4. Identify constraints.
5. Derive acceptance criteria.
6. Detect open decisions.
7. Create or update the Objective.

## OUTPUT

```text
OBJECTIVE
STATUS
SCOPE
OUT OF SCOPE
CONSTRAINTS
ACCEPTANCE CRITERIA
OPEN DECISIONS
TRACEABILITY
```

## GATE

> Does the Objective define the problem, scope, constraints and acceptance criteria sufficiently well to continue without inventing decisions?
