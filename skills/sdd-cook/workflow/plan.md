# Plan

## PURPOSE

Transformar una Specification aceptada en un camino completo y ejecutable hasta Tests y Tasks.

Plan define cómo se realizará el trabajo.

No define qué comportamiento debe tener el sistema.

## INPUT

* Objective aceptado;
* Domain aceptado;
* Specification aceptada;
* decisiones aceptadas.

## ANALYSIS

Determinar:

* orden de trabajo;
* dependencias;
* agrupación;
* descomposición;
* estrategia de verificación;
* riesgos;
* nivel de rigor;
* necesidad de artefactos auxiliares.

## ALLOWED

El agente puede decidir:

* orden de implementación;
* dependencias técnicas;
* agrupación de trabajo;
* descomposición de Tasks;
* estrategia de pruebas;
* nivel de rigor necesario;
* qué artefactos auxiliares son útiles.

## FORBIDDEN

El agente no puede:

* modificar Specification;
* inventar comportamiento;
* introducir requisitos;
* resolver decisiones funcionales no aceptadas.

## ADAPTIVE RIGOR

### Simple

Plan mínimo con pruebas y validación enfocadas.

### Medium

Añadir dependencias, pruebas afectadas y regresión relevante.

### Complex / High Risk

Cuando sea necesario:

* integración;
* regresión;
* datasets;
* performance;
* backtest;
* validación visual;
* seguridad;
* migración;
* evaluación de modelos.

No se crean artefactos vacíos solo por seguir la metodología.

## ACTIONS

1. Verificar Specification.
2. Determinar dependencias.
3. Determinar orden.
4. Determinar estrategia de verificación.
5. Determinar rigor.
6. Identificar riesgos.
7. Identificar artefactos auxiliares necesarios.
8. Crear Tests.
9. Crear Tasks.
10. Establecer trazabilidad.

## OUTPUT

```text
PLAN
ORDER
DEPENDENCIES
VERIFICATION STRATEGY
RIGOR
RISKS
AUXILIARY ARTIFACTS
TESTS
TASKS
TRACEABILITY
NEXT ACTION
```

## GATE

> ¿Existe un camino completo y ejecutable desde Specification hasta Tests y Tasks sin modificar el comportamiento?
