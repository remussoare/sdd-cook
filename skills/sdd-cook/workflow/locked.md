# Locked

## Definition

`LOCKED` significa:

> El estado actual del componente ha sido validado y aceptado como consistente con su Specification.

LOCKED es un estado del ciclo de vida.

No es una fase adicional.

## Characteristics

Un estado LOCKED:

* representa una evolución aceptada;
* conserva su historial;
* puede evolucionar posteriormente;
* no puede modificarse silenciosamente.

## Change

Una modificación posterior sigue:

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
```

La nueva evolución obtiene su propia validación y puede alcanzar:

```text
LOCKED v2
```

sin eliminar la historia anterior.
