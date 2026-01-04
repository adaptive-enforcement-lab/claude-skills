kubectl get nodes -o json | \
  jq '[.items[].status.allocatable] | add'