# /sdd:validate

## PURPOSE

Validar la implementación contra los artefactos aceptados.

## INPUT

* Specification;
* Plan;
* Tests;
* Tasks;
* Implementation.

## ACTIONS

El agente debe:

1. Ejecutar la verificación requerida.
2. Comparar resultado esperado y resultado obtenido.
3. Detectar fallos.
4. Enrutar los fallos según su causa.
5. Registrar evidencia.
6. Determinar el estado de Validation.

## STATES

```text
PASS
FAIL
BLOCKED
INCONCLUSIVE
```

## RULES

Nunca declarar `PASS` sin evidencia suficiente.

Nunca ocultar una regresión.

Nunca cerrar una decisión no resuelta.

Una limitación del entorno no puede convertirse en `PASS`.

## LOCK

```text
PASS
↓
LOCKED
```

`FAIL`, `BLOCKED` e `INCONCLUSIVE` no pueden producir `LOCKED`.
