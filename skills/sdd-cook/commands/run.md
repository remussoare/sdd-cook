# /sdd:run

## PURPOSE

Ejecutar únicamente trabajo autorizado mediante Tasks planificadas.

## RULE

`/sdd:run` no debe crear silenciosamente:

* Plan;
* Specification;
* Tests;
* Tasks.

## PLAN MISSING

Si no existe Plan:

```text
PLAN: MISSING
Cannot run.
Required artifact: PLAN
Run: /sdd:plan
```

## ACTIONS

El agente debe:

1. Seleccionar Tasks pendientes.
2. Implementar las Tasks autorizadas.
3. Ejecutar los Tests requeridos.
4. Reportar:

   * Tasks completadas;
   * archivos modificados;
   * resultados de Tests;
   * drift detectado.

## BLOCK

Si durante la implementación es necesaria una decisión funcional no definida:

```text
BLOCKED
```

No debe inventarse el comportamiento necesario.
