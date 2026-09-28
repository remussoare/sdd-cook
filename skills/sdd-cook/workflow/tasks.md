# Tasks

## PURPOSE

Convertir el Plan en unidades coherentes y ejecutables de implementación.

Una Task puede cubrir varios Tests o partes relacionadas de una Specification.

## INPUT

* Plan aceptado;
* Tests definidos;
* Specification aceptada.

## ANALYSIS

Agrupar el trabajo en unidades coherentes considerando:

* alcance;
* dependencia;
* resultado esperado;
* relación con Tests.

## ALLOWED

El agente puede:

* agrupar trabajo;
* dividir trabajo;
* establecer dependencias;
* elegir una descomposición técnica coherente.

## FORBIDDEN

El agente no puede:

* introducir comportamiento;
* crear decisiones funcionales;
* modificar Specification;
* ampliar el alcance.

Toda decisión funcional de una Task debe poder trazarse a Specification aceptada.

## OUTPUT

Cada Task contiene:

```text
TASK
├── ID
├── PURPOSE
├── SCOPE
├── INPUT
├── EXPECTED RESULT
├── RELATED TESTS
└── DEPENDENCIES
```

## GATE

> ¿Las Tasks son completas, coherentes, ejecutables y trazables sin introducir nuevo comportamiento?
