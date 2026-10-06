# /sdd:new

## PURPOSE

Start a component or a new evolution.

## INVOCATION

```text
/sdd:new [name]
```

## ACTIONS

The agent must:

1. Inspect the project.
2. Inspect `.spec/`.
3. Determine whether this is:

   * a new component;
   * an evolution of an existing component.
4. Preserve the history of `LOCKED` states.
5. Create the minimal necessary context.
6. Establish or update the Objective only with supplied or accepted information.
7. Advance towards Domain and Specification only when decisions are derivable.

## FORBIDDEN

Never invent functional decisions.

Never automatically create a Plan or Tasks.

## BLOCK

If Domain or Specification require a decision that cannot be derived:

```text
BLOCKED
```

and request the necessary decision.
