# Pseudo-code: Full implementation in cluster-admin.yaml
approved_admins := {
  "break-glass-admin@company.com",
  "oncall-sre@company.com",
  "system:masters",  # For kubeadm bootstrap
}

deny[msg] {
  input.kind == "ClusterRoleBinding"
  input.roleRef.name == "cluster-admin"
  not approved_admins[input.subjects[_].name]
  msg := "cluster-admin can only be granted to approved break-glass accounts"
}