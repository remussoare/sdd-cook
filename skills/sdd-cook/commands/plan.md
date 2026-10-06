# /sdd:plan

## PURPOSE

Transform an accepted Specification into:

```text
PLAN
↓
TESTS
↓
TASKS
```

## INPUT

* Objective;
* Domain;
* Specification;
* accepted decisions.

## ACTIONS

The agent must:

1. Read the accepted artifacts (Objective, Domain, Specification, accepted design decisions in `.spec/design/` when present).
2. Check that the Specification is sufficiently complete.
3. Block if functional ambiguity exists.
4. Block before creating the Plan if a requested visual or interaction outcome is not derivable from accepted artifacts: return BLOCKED with STATUS / REASON / AFFECTED ARTIFACT / REQUIRED DECISION / NEXT AUTHORIZED ACTION and request the missing visual information.
5. Create the Plan.
6. Define order and dependencies.
7. Define the verification strategy.
8. Define the required rigor.
9. Create Tests.
10. Create Tasks.
11. Maintain traceability.

## FORBIDDEN

Never modify the behavior defined by the Specification.

Never implement.

Never invent functional decisions.

Never invent non-derivable visual decisions (colors, typography, layout, visual states, interaction) instead of requesting them.

## OUTPUT

Plan, Tests and Tasks ready for execution.
