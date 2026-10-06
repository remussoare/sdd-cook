# /sdd:run

## PURPOSE

Execute only work authorized through planned Tasks.

## RULE

`/sdd:run` must not silently create:

* Plan;
* Specification;
* Tests;
* Tasks.

## PLAN MISSING

If no Plan exists:

```text
PLAN: MISSING
Cannot run.
Required artifact: PLAN
Run: /sdd:plan
```

## ACTIONS

The agent must:

1. Select pending Tasks.
2. Implement the authorized Tasks.
3. Run the required Tests.
4. Report:

   * completed Tasks;
   * modified files;
   * Test results;
   * detected drift.

## BLOCK

If during implementation a functional decision not defined in the artifacts is required:

```text
BLOCKED
```

The necessary behavior must not be invented.
