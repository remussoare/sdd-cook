# /sdd:status

## PURPOSE

Mostrar el estado real del proyecto SDD.

## ACTIONS

Leer `.spec/` y mostrar el estado existente.

## OUTPUT

```text
SDD STATUS

Component:
Evolution:
Current phase:
Objective:
Domain:
Specification:
Plan:
Tests:
Tasks:
Implementation:
Validation:
Lifecycle:
Blockers:
Drift:
Next action:
```

## RULES

Mostrar únicamente estados que puedan determinarse a partir de los artefactos existentes.

No inventar estados cuando falten artefactos.

Identificar:

* dependencias faltantes;
* bloqueos;
* drift;
* siguiente acción autorizada.
