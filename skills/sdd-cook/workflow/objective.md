# Objective

## PURPOSE

Definir con precisión qué problema está autorizado a resolver el agente.

El Objective establece:

* propósito;
* alcance;
* fuera de alcance;
* restricciones;
* criterios de aceptación.

No diseña el dominio ni la solución.

## INPUT

* petición del usuario;
* contexto aceptado del proyecto;
* restricciones existentes;
* decisiones previamente aceptadas cuando sean relevantes.

## ANALYSIS

El agente debe:

* estructurar la información existente;
* eliminar ambigüedad lingüística;
* separar alcance y fuera de alcance;
* derivar criterios de aceptación directamente;
* detectar contradicciones;
* detectar información necesaria que falta;
* identificar decisiones abiertas.

## ALLOWED

El agente puede:

* reorganizar información;
* mejorar precisión lingüística;
* separar alcance y fuera de alcance;
* derivar criterios de aceptación directamente implícitos;
* adaptar el nivel de detalle;
* proponer una estructura más clara.

## FORBIDDEN

El agente no puede:

* inventar requisitos;
* ampliar el alcance;
* reducir el alcance;
* cambiar prioridades;
* inventar restricciones;
* asumir decisiones;
* convertir una recomendación técnica en requisito funcional;
* eliminar una restricción existente;
* modificar el objetivo para facilitar la implementación.

## ACTIONS

1. Analizar la petición.
2. Identificar objetivo y alcance.
3. Separar lo que queda fuera.
4. Identificar restricciones.
5. Derivar criterios de aceptación.
6. Detectar decisiones abiertas.
7. Crear o actualizar el Objective.

## OUTPUT

```text
OBJECTIVE
STATUS
SCOPE
OUT OF SCOPE
CONSTRAINTS
ACCEPTANCE CRITERIA
OPEN DECISIONS
TRACEABILITY
```

## GATE

> ¿El objetivo define suficientemente el problema, alcance, restricciones y criterios de aceptación para continuar sin inventar decisiones?
