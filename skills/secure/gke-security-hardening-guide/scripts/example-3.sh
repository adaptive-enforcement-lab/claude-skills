#!/bin/bash
# Comprehensive security verification

CLUSTER="prod-cluster"
REGION="us-central1"

echo "=== Cluster Configuration ==="
gcloud container clusters describe $CLUSTER --region $REGION \
  --format="value(privateClusterConfig.enablePrivateNodes)" | grep -q "True" && echo "✓ Private cluster" || echo "✗ Public cluster"

gcloud container clusters describe $CLUSTER --region $REGION \
  --format="value(workloadIdentityConfig.workloadPool)" | grep -q "svc.id.goog" && echo "✓ Workload Identity enabled" || echo "✗ Workload Identity disabled"

gcloud container clusters describe $CLUSTER --region $REGION \
  --format="value(databaseEncryption.state)" | grep -q "ENCRYPTED" && echo "✓ Database encryption enabled" || echo "✗ Database encryption disabled"

echo ""
echo "=== Network Configuration ==="
gcloud container clusters describe $CLUSTER --region $REGION \
  --format="value(networkingConfig.useIpAliases)" | grep -q "True" && echo "✓ VPC-native networking" || echo "✗ Routes-based networking"

gcloud container clusters describe $CLUSTER --region $REGION \
  --format="value(networkPolicy.enabled)" | grep -q "True" && echo "✓ Network policies enabled" || echo "✗ Network policies disabled"

echo ""
echo "=== Pod Security ==="
kubectl get ns -o jsonpath='{range .items[*]}{.metadata.name}{"\t"}{.metadata.labels.pod-security\.kubernetes\.io/enforce}{"\n"}{end}' | column -t

echo ""
echo "=== IAM Configuration ==="
gcloud iam service-accounts list --format="table(email,displayName)"

echo ""
echo "=== Audit Logging ==="
gcloud logging sinks list --filter="destination:*storage*" --format="table(name,destination)"