# Validation

VALIDATION
STATUS
EVIDENCE
TEST RESULTS
ACCEPTANCE CRITERIA
REGRESSION
FAILURES
TRACEABILITY
DECISION
NEXT ACTION

### Proportional Validation

Validation must be proportional to the type, scope, and risk of the change.

#### Documentation and SDD Changes

For changes to `.spec`, documentation, definitions, diagrams, naming, formatting, or other non-executable artifacts:

* Do not run Python tests.
* Do not run the full test suite.
* Validate structure, references, consistency, and affected artifacts only.
* Do not execute code unless execution is required to validate the artifact.

#### Local Implementation Changes

For a localized implementation change with limited impact:

* Run only the directly affected validation or targeted test.
* Do not run unrelated tests.
* Do not run the full test suite unless dependency impact justifies it.

#### Shared or Behavioral Changes

For changes affecting shared components, interfaces, core behavior, or multiple modules:

* Start with targeted validation.
* Expand validation only when dependency impact requires it.
* Run the full suite only when justified by the change.

#### Execution Control

* Do not execute tests speculatively.
* Do not repeat a successful validation without a concrete reason.
* If the same validation fails twice, stop and diagnose before executing again.
* Do not enter execution loops.
* Use the minimum validation capable of establishing correctness.

**Default rule:**

`Validation effort ∝ Change impact`
