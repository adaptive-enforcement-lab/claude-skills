# List all ClusterRoleBindings to cluster-admin
kubectl get clusterrolebindings -o json | \
  jq '.items[] | select(.roleRef.name == "cluster-admin") | .metadata.name'

# List roles with wildcard permissions
kubectl get roles,clusterroles --all-namespaces -o json | \
  jq '.items[] | select(.rules[].resources[] == "*") | .metadata.name'