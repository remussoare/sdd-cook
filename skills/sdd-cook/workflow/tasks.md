# Tasks

## PURPOSE

Turn the Plan into coherent, executable implementation units.

A Task may cover several Tests or related parts of a Specification.

## INPUT

* accepted Plan;
* defined Tests;
* accepted Specification.

## ANALYSIS

Group work into coherent units considering:

* scope;
* dependency;
* expected result;
* relationship with Tests.

## ALLOWED

The agent may:

* group work;
* split work;
* establish dependencies;
* choose a coherent technical decomposition.

## FORBIDDEN

The agent may not:

* introduce behavior;
* create functional decisions;
* modify the Specification;
* expand scope.

Every functional decision in a Task must be traceable to accepted Specification.

## OUTPUT

Each Task contains:

```text
TASK
├── ID
├── PURPOSE
├── SCOPE
├── INPUT
├── EXPECTED RESULT
├── RELATED TESTS
└── DEPENDENCIES
```

## GATE

> Are the Tasks complete, coherent, executable and traceable without introducing new behavior?
