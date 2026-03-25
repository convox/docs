---
title: "CircuitBreaker"
description: "Enables ECS deployment circuit breaker to automatically roll back failing deployments."
---

# CircuitBreaker

ECS deployment circuit breaker. When enabled, ECS automatically rolls back deployments that fail to reach a steady state. However, if a deployment requires the Rack to scale up, it may trip the circuit breaker prematurely. Best used when sufficient capacity is available.

| Default value  | `No` |
| Allowed values | `Yes`, `No` |

## Use Cases

- Enable when you want failing deployments to roll back automatically without manual intervention
- Enable in environments where you have pre-provisioned sufficient compute capacity and do not need to scale up during deployments
- Keep disabled when deployments routinely trigger cluster scaling, as the scaling delay may cause the circuit breaker to trigger a false rollback

## Additional Information

The ECS deployment circuit breaker monitors the health of tasks during a deployment. If new tasks fail to reach a steady state, ECS will automatically stop the deployment and optionally roll back to the last completed deployment.

The circuit breaker can be overly aggressive in environments where new instances need to be launched to accommodate the deployment. In these cases, the time required for new EC2 instances to join the cluster may exceed the circuit breaker's failure threshold, causing a premature rollback. To avoid this, ensure your Rack has enough spare capacity before enabling this parameter, or use Fargate-based services where capacity provisioning is faster.

```bash
$ convox apps params set CircuitBreaker=Yes
```

## See Also

- [FargateServices](/reference/app-parameters/FargateServices)
- [Rack Parameters](/reference/rack-parameters)
