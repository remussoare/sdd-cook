# Plan

## PURPOSE

Turn an accepted Specification into a complete, executable path through Tests and Tasks.

The Plan defines how the work will be performed.

It does not define what behavior the system must have.

## INPUT

* accepted Objective;
* accepted Domain;
* accepted Specification;
* accepted decisions.

## ANALYSIS

Determine:

* work order;
* dependencies;
* grouping;
* decomposition;
* verification strategy;
* risks;
* rigor level;
* need for auxiliary artifacts;
* visual coverage: whether the prompt or Specification requests visual presentation or interaction (UI, dashboard, landing, colors, layout, states) and whether those visual decisions are derivable from accepted Objective, Domain, Specification, or accepted design decisions.

## ALLOWED

The agent may decide:

* implementation order;
* technical dependencies;
* work grouping;
* Task decomposition;
* test strategy;
* required rigor level;
* which auxiliary artifacts are useful.

## FORBIDDEN

The agent may not:

* modify the Specification;
* invent behavior;
* introduce requirements;
* resolve non-accepted functional decisions;
* invent non-derivable visual decisions (colors, typography, layout, visual states, interaction) instead of requesting them.

## ADAPTIVE RIGOR

### Simple

Minimal plan with focused tests and validation.

### Medium

Add dependencies, affected tests and relevant regression.

### Complex / High Risk

When necessary:

* integration;
* regression;
* datasets;
* performance;
* backtest;
* visual validation;
* security;
* migration;
* model evaluation.

No empty artifacts are created just to follow the methodology.

## ACTIONS

1. Verify the Specification.
2. Check visual coverage: if a visual or interaction outcome is requested and is not derivable from accepted artifacts, STOP with BLOCKED and request the missing visual information before creating the Plan (see `principles/design-control.md`). Do not create Plan, Tests, or Tasks until visual decisions are derivable or accepted.
3. Determine dependencies.
4. Determine order.
5. Determine the verification strategy (include visual/interaction evidence when design decisions are part of the accepted result).
6. Determine rigor.
7. Identify risks.
8. Identify necessary auxiliary artifacts (use `design/visual.md` DRAFT and/or a prototype only when visual uncertainty justifies the cost).
9. Create Tests.
10. Create Tasks.
11. Establish traceability (Specification → Design Decision → Test → Task when applicable).

## OUTPUT

```text
PLAN
WORK ORDER
DEPENDENCIES
VERIFICATION STRATEGY
RIGOR
RISKS
AUXILIARY ARTIFACTS
VISUAL INPUTS
TESTS
TASKS
TRACEABILITY
NEXT ACTION
```

## GATE

> Does a complete, executable path exist from Specification through Tests and Tasks without modifying behavior and with visual decisions derivable or accepted?
