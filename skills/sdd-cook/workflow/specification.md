# Specification

## PURPOSE

Definir el contrato de comportamiento de forma explícita y determinista.

Specification establece:

* reglas;
* condiciones;
* entradas;
* salidas;
* precondiciones;
* postcondiciones;
* estados;
* transiciones;
* invariantes;
* errores;
* edge cases;
* dependencias de comportamiento.

## INPUT

* Objective aceptado;
* Domain aceptado;
* decisiones aceptadas.

## ANALYSIS

El agente debe convertir el modelo conceptual en comportamiento verificable.

Debe identificar:

* qué ocurre;
* cuándo ocurre;
* bajo qué condiciones;
* qué entradas son necesarias;
* qué resultado se espera;
* qué estados existen;
* cómo se producen las transiciones;
* qué errores son posibles;
* qué casos límite deben contemplarse.

## ALLOWED

El agente puede:

* hacer explícitas reglas derivables;
* dividir comportamiento complejo;
* eliminar ambigüedad lingüística;
* estructurar condiciones;
* definir casos derivados directamente del comportamiento aceptado.

## FORBIDDEN

El agente no puede:

* usar la implementación como autoridad;
* inventar comportamiento;
* inventar estados;
* inventar transiciones;
* introducir requisitos;
* modificar el comportamiento para adaptarlo a la implementación;
* resolver una ambigüedad funcional mediante una suposición.

## ACTIONS

1. Leer Objective y Domain.
2. Identificar comportamiento.
3. Definir reglas deterministas.
4. Definir inputs y outputs.
5. Definir estados y transiciones cuando correspondan.
6. Definir invariantes.
7. Definir errores y edge cases.
8. Identificar dependencias.
9. Registrar decisiones abiertas.
10. Preparar Specification para planificación y verificación.

## OUTPUT

```text
SPECIFICATION
STATUS
RULES
INPUTS
OUTPUTS
PRECONDITIONS
POSTCONDITIONS
STATES
TRANSITIONS
INVARIANTS
ERRORS
EDGE CASES
DEPENDENCIES
OPEN DECISIONS
CONFLICTS
TRACEABILITY
NEXT ACTION
```

## GATE

> ¿El comportamiento puede convertirse en una verificación determinista sin introducir suposiciones no aceptadas?
