# Design

Design is a **transversal capability**, not a new workflow phase.

It is used when decisions related to:

* visual requirements;
* design decisions;
* colors;
* typography;
* layout;
* visual states;
* interaction;
* UX.

## Decision Rule

```text
Is the visual decision derivable?

YES → decide and continue

NO → propose → decision
```

When a visual decision is not derivable, the agent may propose it, but the accepted choice becomes a design decision. In `/sdd:plan`, a non-derivable visual request blocks Plan creation until the missing visual information is provided and accepted; it is recorded here (`design/visual.md`) as an auxiliary artifact, never as a silent Plan assumption.

## Traceability

```text
SPECIFICATION
      ↓
DESIGN DECISION
      ↓
VISUAL / INTERACTION TEST
      ↓
TASK
      ↓
IMPLEMENTATION
      ↓
VALIDATION
```

Design does not modify the main workflow.
