# /sdd:validate

## Purpose

Produce sufficient, reliable and traceable evidence that implementation satisfies accepted Specification and acceptance criteria in the relevant context.

## Procedure

1. Read Specification, Plan, Tests, Tasks and implementation state.
2. Execute relevant unit, integration, scenario, regression, visual, performance, backtest or other evidence required by risk and accepted artifacts.
3. Compare expected and actual results.
4. Route failures to the earliest affected phase.
5. Never declare PASS without sufficient evidence.
6. On PASS, record Validation and transition the accepted state to LOCKED.
7. On FAIL, BLOCKED or INCONCLUSIVE, do not lock.

## Important

TESTS PASS does not automatically mean Validation PASS.
