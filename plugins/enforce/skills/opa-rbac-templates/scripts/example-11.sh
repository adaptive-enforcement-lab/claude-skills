# List service accounts with cluster-admin
kubectl get clusterrolebindings -o json | \
  jq '.items[] | select(.roleRef.name == "cluster-admin" and .subjects[].kind == "ServiceAccount")'

# Find service accounts with escalate/bind/impersonate verbs
kubectl get roles,clusterroles --all-namespaces -o json | \
  jq '.items[] | select(.rules[].verbs[] | IN("escalate", "bind", "impersonate"))'