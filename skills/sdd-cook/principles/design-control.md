# Design Control

Design is a transversal capability, not a workflow phase.

Use design artifacts when visual presentation, interaction, readability or UX forms part of the product contract or when an explicit design decision materially affects implementation or validation.

## Decision rule

If a design decision is derivable from accepted artifacts, the agent may decide and continue.

If it is not derivable and affects the accepted result, the agent must propose options and request a decision. In `/sdd:plan` this means BLOCKED before creating the Plan: request the missing visual information first.

A proposed color, layout, typography, visual state or interaction is not an accepted requirement until explicitly accepted or derivable from accepted artifacts.

## Required visual information

When blocking for visual input, request only what is missing and material to the result:

- colors / palette (including dark/light when applicable);
- typography;
- layout and structure;
- visual states;
- references, mockups, or existing branding constraints;
- accessibility minimum (contrast, readability).

Record the accepted answers as design decisions (e.g. `design/visual.md` DRAFT) and reference them from the Plan. Do not invent them.

## Scope

Typical design concerns:
- colors
- line styles and thickness
- opacity
- typography
- labels
- zones
- panels
- icons
- layout
- visibility
- interaction
- accessibility/readability

## Constraint

Design artifacts must not silently change functional behavior. If a visual decision changes behavior, state, scope or architecture, use the normal decision/change-control process.
