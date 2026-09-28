# /sdd:new

## PURPOSE

Iniciar un componente o una nueva evolución.

## INVOCATION

```text
/sdd:new [name]
```

## ACTIONS

El agente debe:

1. Inspeccionar el proyecto.
2. Inspeccionar `.spec/`.
3. Determinar si se trata de:

   * nuevo componente;
   * evolución de un componente existente.
4. Preservar la historia de estados `LOCKED`.
5. Crear el contexto mínimo necesario.
6. Establecer o actualizar Objective únicamente con información suministrada o aceptada.
7. Avanzar hacia Domain y Specification únicamente cuando las decisiones sean derivables.

## FORBIDDEN

No inventar decisiones funcionales.

No crear automáticamente Plan ni Tasks.

## BLOCK

Si Domain o Specification requieren una decisión que no puede derivarse:

```text
BLOCKED
```

y solicitar la decisión necesaria.
