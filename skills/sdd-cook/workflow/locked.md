# Locked

## Definition

`LOCKED` means:

> The current state of the component has been validated and accepted as consistent with its Specification.

LOCKED is a lifecycle state.

It is not an additional phase.

## Characteristics

A LOCKED state:

* represents an accepted evolution;
* preserves its history;
* may evolve later;
* cannot be modified silently.

## Change

A later modification follows:

```text
LOCKED
 ↓
CHANGE
 ↓
IMPACT ANALYSIS
 ↓
DECISION
 ↓
NEW EVOLUTION
```

The new evolution obtains its own validation and may reach:

```text
LOCKED v2
```

without deleting previous history.
