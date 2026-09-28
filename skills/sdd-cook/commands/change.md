# /sdd:change

## PURPOSE

Gestionar cambios sobre un estado aceptado o `LOCKED`.

## FLOW

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
↓
VALIDATION
↓
LOCKED v2
```

Cuando el cambio requiere implementación:

```text
NEW EVOLUTION
↓
PLAN
↓
TESTS
↓
TASKS
↓
IMPLEMENTATION
↓
VALIDATION
↓
LOCKED v2
```

## ACTIONS

1. Identificar el cambio solicitado.
2. Clasificar su impacto.
3. Si afecta comportamiento, alcance, arquitectura o una decisión aceptada, crear Change Request.
4. Realizar Impact Analysis.
5. Solicitar una decisión cuando el cambio no sea derivable.
6. Nunca sustituir silenciosamente un artefacto aceptado.
7. Actualizar manteniendo trazabilidad.
8. Replanificar, probar, implementar y validar cuando corresponda.
9. Preservar la evolución anterior.

## STATUSES

```text
PROPOSED
ANALYZING
AWAITING_DECISION
APPROVED
IMPLEMENTING
VALIDATING
ACCEPTED
REJECTED
CANCELLED
BLOCKED
```
