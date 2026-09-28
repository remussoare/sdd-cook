# Authority

SDD assigns a specific responsibility to every artifact.

| Artifact       | Authority                                                       |
| -------------- | --------------------------------------------------------------- |
| Objective      | Purpose, scope, constraints, acceptance criteria                |
| Domain         | Concepts, entities, relationships, conceptual states/invariants |
| Specification  | Behavior and rules                                              |
| Plan           | Execution and verification strategy                             |
| Tests          | Evidence                                                        |
| Tasks          | Implementation work units                                       |
| Implementation | Technical realization                                           |
| Validation     | Contextual evidence                                             |
| Locked         | Accepted validated state                                        |

## Authority Rule

Downstream artifacts cannot silently override upstream authority.

If implementation conflicts with Specification, Specification remains authoritative.

If Tests conflict with Specification, Tests are corrected or the Specification is changed through Change Control.

If a functional decision is absent from authoritative artifacts, the agent must stop and request a decision.

## Principle

No artifact may acquire authority merely because it was created later in the workflow.
