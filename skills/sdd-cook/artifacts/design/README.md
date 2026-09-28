# Design

Design es una **capacidad transversal**, no una fase nueva del workflow.

Se utiliza cuando existen decisiones relacionadas con:

* requisitos visuales;
* decisiones de diseño;
* colores;
* tipografía;
* layout;
* estados visuales;
* interacción;
* UX.

## Decision Rule

```text id="5v8h2q"
¿La decisión visual es derivable?

YES → decide and continue

NO → propose → decision
```

Cuando una decisión visual no es derivable, el agente puede proponerla, pero la elección aceptada se convierte en una decisión de diseño.

## Traceability

```text id="j2h8p4"
SPECIFICATION
      ↓
DESIGN DECISION
      ↓
VISUAL / INTERACTION TEST
      ↓
TASK
      ↓
IMPLEMENTATION
      ↓
VALIDATION
```

Design no modifica el workflow principal.
