kubectl auth can-i patch deployments \
  --as=system:serviceaccount:argo-workflows:my-sa \
  -n target-namespace