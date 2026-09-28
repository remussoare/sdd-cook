# Implementation

## PURPOSE

Realizar técnicamente las Tasks aceptadas.

Implementation convierte las decisiones aceptadas en código o configuración ejecutable.

## INPUT

* Tasks aceptadas;
* Specification;
* Tests;
* restricciones técnicas del proyecto.

## ANALYSIS

Determinar la realización técnica sin cambiar el comportamiento aceptado.

## ALLOWED

El agente puede:

* escribir código;
* refactorizar;
* optimizar;
* cambiar estructuras internas;
* cambiar nombres;
* crear abstracciones;
* reutilizar componentes;
* corregir bugs de implementación.

## REFACTORING

Un refactor es válido cuando:

```text
BEHAVIOR BEFORE = BEHAVIOR AFTER
```

## FORBIDDEN

El agente no puede:

* cambiar reglas;
* cambiar estados;
* cambiar transiciones;
* cambiar outputs;
* cambiar errores;
* cambiar prioridades;
* cambiar alcance;
* cambiar criterios de aceptación;
* modificar Tests para ocultar fallos;
* modificar Specification para justificar código;
* añadir comportamiento no especificado;
* resolver ambigüedad funcional;
* cambiar arquitectura cuando el cambio requiere una decisión no aceptada.

## ACTIONS

1. Seleccionar Tasks autorizadas.
2. Implementarlas.
3. Ejecutar Tests relevantes.
4. Comparar resultados con expectativas.
5. Detectar drift.
6. Reportar archivos modificados.
7. Reportar Tasks completadas.
8. Preparar evidencia para Validation.

## OUTPUT

```text
IMPLEMENTATION
STATUS
TASKS COMPLETED
FILES CHANGED
BEHAVIOR IMPLEMENTED
TEST RESULTS
ARCHITECTURE COMPLIANCE
TRACEABILITY
DRIFT DETECTED
BLOCKERS
NEXT ACTION
```

## GATE

La implementación debe satisfacer:

* Tasks;
* Specification;
* restricciones aceptadas;
* trazabilidad;
* ausencia de drift funcional no resuelto.

Importante:

```text
TESTS PASS ≠ IMPLEMENTATION GATE PASS
```

Que los Tests pasen no significa por sí solo que Implementation haya superado su Gate.
