# Decision Control

Every relevant decision must be classified before being applied.

## Process

```text
IDENTIFY DECISION
       ↓
FIND AUTHORITY
       ↓
IS IT DERIVABLE?
   ┌───┴───┐
  YES      NO
   ↓        ↓
 APPLY     STOP
            ↓
      REQUEST DECISION
```

## Derivable Decision

A decision is derivable when it follows directly from accepted artifacts without introducing new assumptions about behavior, scope or requirements.

The agent may decide and continue.

## Non-Derivable Decision

A decision is non-derivable when multiple valid interpretations remain or the choice introduces new behavior, scope, priority or accepted constraints.

The agent must stop and request a decision.

## Prohibition

Never convert uncertainty into behavior.
