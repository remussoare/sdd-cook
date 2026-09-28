# /sdd:new

## Purpose

Start a new component or evolution without inventing functional decisions.

## Invocation

`/sdd:new [name]`

## Procedure

1. Inspect the project and `.spec/` state.
2. Determine whether this is a new component or an evolution.
3. Preserve previous accepted/LOCKED history when evolving existing work.
4. Establish Objective only from supplied or already accepted information.
5. Build Domain and Specification only when required decisions are derivable.
6. If a required behavioral or scope decision is not derivable, BLOCK.
7. Do not create Plan, Tests or Tasks merely because `/sdd:new` was invoked.

## Forbidden

- invent requirements
- silently modify LOCKED artifacts
- decide ambiguous behavior
- change scope or priorities without a decision
