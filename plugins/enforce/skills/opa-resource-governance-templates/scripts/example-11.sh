kubectl get pods --all-namespaces -o json | \
  jq '[.items[].spec.containers[].resources.requests] | add'