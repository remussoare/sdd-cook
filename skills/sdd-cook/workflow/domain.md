# Domain

## PURPOSE

Construir el modelo conceptual del problema.

Domain define:

* conceptos;
* entidades;
* value objects;
* relaciones;
* estados conceptuales;
* invariantes;
* límites del dominio;
* lenguaje común.

La profundidad DDD se adapta a la complejidad real del problema.

## INPUT

* Objective aceptado;
* conocimiento de dominio disponible;
* decisiones aceptadas.

## ANALYSIS

Identificar:

* conceptos principales;
* entidades;
* value objects;
* relaciones;
* límites;
* estados conceptuales;
* invariantes;
* posibles conflictos semánticos.

## ALLOWED

El agente puede:

* identificar conceptos derivados del Objective;
* estructurar entidades;
* identificar relaciones;
* identificar value objects;
* establecer estructura conceptual;
* derivar invariantes conceptuales cuando estén claramente implícitos;
* adaptar la profundidad del modelo a la complejidad.

## FORBIDDEN

El agente no puede:

* inventar reglas de comportamiento;
* inventar estados de comportamiento;
* inventar transiciones;
* resolver ambigüedades semánticas mediante una suposición cuando afecten al comportamiento;
* introducir requisitos funcionales.

## ACTIONS

1. Leer Objective.
2. Identificar conceptos.
3. Identificar entidades y value objects.
4. Establecer relaciones.
5. Identificar estados conceptuales e invariantes.
6. Detectar conflictos.
7. Registrar decisiones abiertas.
8. Preparar el dominio para Specification.

## OUTPUT

```text
DOMAIN
STATUS
COMPLEXITY
CONCEPTS
ENTITIES
VALUE OBJECTS
RELATIONSHIPS
STATES
INVARIANTS
BOUNDARIES
OPEN DECISIONS
CONFLICTS
TRACEABILITY
NEXT ACTION
```

## GATE

> ¿Existe suficiente claridad conceptual para especificar el comportamiento sin inventar conceptos ni decisiones?
