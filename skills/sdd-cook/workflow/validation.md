# Validation

## PURPOSE

Obtener evidencia suficiente, fiable y trazable de que la implementación satisface Specification y los criterios de aceptación en el contexto relevante.

La diferencia fundamental es:

```text
IMPLEMENTATION
= ¿lo he construido correctamente?

VALIDATION
= ¿tenemos suficiente evidencia de que funciona correctamente en contexto?
```

## INPUT

* Objective;
* Domain;
* Specification;
* Plan;
* Tests;
* Tasks;
* Implementation.

## ANALYSIS

Comparar:

```text
EXPECTED
vs
ACTUAL
```

La validación debe utilizar el nivel de evidencia apropiado al contexto.

Puede incluir, cuando corresponda:

* unit;
* integration;
* scenarios;
* regression;
* datasets;
* visual;
* performance;
* backtest.

## ALLOWED

El agente puede ejecutar las verificaciones necesarias para obtener evidencia suficiente.

## FORBIDDEN

El agente no puede:

* declarar PASS sin evidencia suficiente;
* ocultar regresiones;
* cerrar decisiones sin resolver;
* convertir una limitación del entorno en PASS;
* declarar éxito basándose únicamente en una comprobación insuficiente.

## STATES

```text
PASS
FAIL
BLOCKED
INCONCLUSIVE
```

## OUTPUT

```text
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
```

## FAILURE ROUTING

```text
FAIL
 ↓
DIAGNOSE
 ↓
CORRECT
 ↓
REVALIDATE
```

## GATE

Solo puede declararse:

```text
PASS
```

cuando existe evidencia suficiente y fiable.

```text
PASS
 ↓
LOCKED
```

`FAIL`, `BLOCKED` e `INCONCLUSIVE` no pueden producir LOCKED.
