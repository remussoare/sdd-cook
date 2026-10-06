# Design in SDD

Design is transversal to the workflow and is not a separate phase.

Use it when presentation or interaction materially affects the accepted product result.

Typical placement:
- Objective: visual scope and acceptance criteria when explicitly required.
- Domain: conceptual visual states only when they are domain concepts.
- Specification: observable visual/interaction behavior when it is part of the contract.
- Plan: decide the appropriate design artifacts and verification strategy.
- Tests: visual/interaction evidence when required.
- Tasks: implementation units derived from accepted design decisions.
- Validation: verify visual/interaction acceptance in context.

The agent may choose presentation details only when they are derivable from accepted requirements or design decisions. Otherwise it must propose and request a decision.

In `/sdd:plan`, a non-derivable visual request blocks Plan/Tests/Tasks creation until the missing information (colors, typography, layout, visual states, references/mockups, branding constraints, accessibility minimum) is provided and accepted. If a visual decision changes behavior, state, scope, or architecture, route it through Specification/Change Management; otherwise record it in `design/visual.md` as an auxiliary artifact.
