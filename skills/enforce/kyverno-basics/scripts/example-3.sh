$ kubectl apply -f deployment.yaml
Error from server: admission webhook "validate.kyverno.svc-fail" denied the request:

policy Deployment/default/api for resource violation:

require-resource-limits:
  check-resource-limits: validation error: Resource limits are required for all containers