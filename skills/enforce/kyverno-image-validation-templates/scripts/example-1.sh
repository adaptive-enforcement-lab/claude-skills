kubectl apply -f registry-allowlist-policy.yaml  # Registry controls first
kubectl get clusterpolicy -w   # Watch for Ready status