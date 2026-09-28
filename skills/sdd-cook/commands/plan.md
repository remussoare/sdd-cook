# /sdd:plan

## PURPOSE

Transformar una Specification aceptada en:

```text
PLAN
↓
TESTS
↓
TASKS
```

## INPUT

* Objective;
* Domain;
* Specification;
* decisiones aceptadas.

## ACTIONS

El agente debe:

1. Leer los artefactos aceptados.
2. Comprobar que Specification es suficientemente completa.
3. Bloquear si existe ambigüedad funcional.
4. Crear Plan.
5. Definir orden y dependencias.
6. Definir estrategia de verificación.
7. Definir el rigor necesario.
8. Crear Tests.
9. Crear Tasks.
10. Mantener trazabilidad.

## FORBIDDEN

No modificar el comportamiento definido por Specification.

No implementar.

No inventar decisiones funcionales.

## OUTPUT

Plan, Tests y Tasks listos para ejecución.
