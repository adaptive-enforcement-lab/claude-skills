#!/bin/bash
# Podman ephemeral runner with strict isolation
podman run --rm \
  --security-opt=no-new-privileges:true \
  --cap-drop=ALL \
  --read-only \
  --tmpfs /tmp:rw,noexec,nosuid,size=1g \
  --network=slirp4netns:enable_ipv6=false \
  -e RUNNER_EPHEMERAL=true \
  -e GITHUB_TOKEN="${GITHUB_TOKEN}" \
  ghcr.io/myorg/runner:latest