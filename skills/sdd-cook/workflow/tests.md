# Tests

## PURPOSE

Proporcionar evidencia de que el comportamiento definido en Specification puede verificarse.

Tests no definen comportamiento.

## INPUT

* Specification aceptada;
* Plan aceptado.

## ANALYSIS

Derivar casos de prueba directamente de Specification.

Según corresponda:

* positivos;
* negativos;
* edge cases;
* transiciones;
* regresión;
* fixtures;
* datos de prueba;
* parametrización.

## ALLOWED

El agente puede:

* derivar casos de prueba;
* crear fixtures;
* crear datos;
* parametrizar pruebas;
* organizar suites;
* añadir pruebas de regresión derivadas del comportamiento existente;
* mejorar determinismo de las pruebas.

## FORBIDDEN

El agente no puede:

* inventar comportamiento;
* introducir requisitos;
* modificar Specification para que una prueba pase;
* eliminar fallos;
* modificar resultados esperados para adaptarlos a la implementación.

## FAILURE ROUTING

Cuando un Test falla:

```text
Implementation bug
→ Implementation

Test bug
→ Tests

Specification issue
→ Change Request
```

## OUTPUT

```text
TESTS
STATUS
TEST CASES
COVERAGE
FIXTURES
TEST DATA
REGRESSION SET
TRACEABILITY
FAILURES
OPEN DECISIONS
NEXT ACTION
```

## GATE

> ¿Los Tests proporcionan evidencia suficiente del comportamiento aceptado con el nivel de rigor requerido?
