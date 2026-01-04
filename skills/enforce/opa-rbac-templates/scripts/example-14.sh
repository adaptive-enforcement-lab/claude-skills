# Generate temporary kubeconfig with cluster-admin (expires in 1 hour)
kubectl create token break-glass-admin --duration=1h

# Use temporary token for emergency operations
kubectl --token=$(kubectl create token break-glass-admin --duration=1h) get nodes