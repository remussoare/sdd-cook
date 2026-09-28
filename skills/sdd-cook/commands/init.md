# /sdd:init

## PURPOSE

Inicializar o integrar SDD en un proyecto.

`/sdd:init` no crea un componente.

## PROJECT START

```text
PROJECT START
↓
DISCOVERY
↓
prepare Constitution
prepare AGENTS.md
prepare needed .spec/
↓
INITIALIZATION GATE
↓
OBJECTIVE
```

## EXISTING PROJECT

```text
PROJECT EXISTS
↓
DISCOVERY
↓
AUDIT
↓
existing Constitution / AGENTS / .spec / stack
↓
detect conflicts/drift
↓
integrate without silent overwrite
↓
INITIALIZATION GATE
```

## ACTIONS

El agente debe:

* inspeccionar el repositorio;
* detectar el stack;
* detectar `AGENTS.md`;
* detectar Constitution;
* detectar `.spec/`;
* detectar decisiones existentes;
* detectar conflictos;
* detectar drift;
* preparar la estructura mínima necesaria;
* preparar Constitution y `AGENTS.md`;
* dejar el proyecto preparado para `/sdd:new`.

## RULE

No sobrescribir silenciosamente artefactos existentes.
