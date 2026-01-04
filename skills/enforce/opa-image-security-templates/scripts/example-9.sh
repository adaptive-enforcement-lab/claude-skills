# Test registry allowlist (should fail for Docker Hub)
kubectl run docker-hub-test --image=nginx:latest
# Expected: Blocked by registry allowlist policy

# Test digest requirement (should fail for tag-based reference)
kubectl run tag-test --image=registry.company.com/nginx:v1.21
# Expected: Blocked by digest enforcement policy

# Test digest-based reference (should succeed if from approved registry)
kubectl run digest-test --image=registry.company.com/nginx@sha256:abcdef123456...
# Expected: Allowed by registry and digest policies

# Test unsigned image (should fail without valid signature)
kubectl run unsigned-test --image=registry.company.com/app@sha256:123456...
# Expected: Blocked by signature verification policy

# Test signed image (should succeed with valid Cosign signature)
# First, sign the image:
# cosign sign --key cosign.key registry.company.com/app@sha256:123456...
kubectl run signed-test --image=registry.company.com/app@sha256:123456...
# Expected: Allowed after signature verification