kubectl create secret generic cosign-pub \
  --from-file=cosign.pub=./cosign.pub \
  -n opa-system