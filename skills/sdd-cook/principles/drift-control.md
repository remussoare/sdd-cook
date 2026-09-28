# Drift Control

Drift is divergence between accepted authority and downstream artifacts or implementation.

## Detect

The agent should detect at least:

* code without a Task;
* Task without Specification trace;
* Task without relevant Test trace;
* behavior without Specification;
* Specification without implementation;
* Test without Specification;
* Test expectation unsupported by Specification;
* tests modified to conceal failures;
* validation without evidence;
* unauthorized functional decisions;
* unauthorized design decisions;
* implementation outside accepted scope.

## Resolution

Drift must be classified.

```text
IMPLEMENTATION DRIFT
→ correct Implementation

TEST DRIFT
→ correct Tests

SPECIFICATION DRIFT
→ Change Control

DOMAIN DRIFT
→ Change Control

DECISION DRIFT
→ request decision
```

Unresolved functional drift prevents acceptance.
