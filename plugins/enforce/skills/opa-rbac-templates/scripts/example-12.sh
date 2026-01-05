# List all human users with cluster-level access
kubectl get clusterrolebindings -o json | \
  jq '.items[] | select(.subjects[].kind == "User") | {binding: .metadata.name, user: .subjects[].name, role: .roleRef.name}'