# Test NetworkPolicy requirement (should fail without policy)
kubectl create namespace test-ns
kubectl run test-pod --image=nginx -n test-ns
# Expected: Blocked by policy requiring NetworkPolicy in namespace

# Test Ingress TLS requirement (should fail without TLS)
cat <<EOF | kubectl apply -f -
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: test-ingress
  namespace: test-ns
spec:
  rules:
    - host: test.example.com
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: test-service
                port:
                  number: 80
EOF
# Expected: Blocked by policy requiring spec.tls

# Test Service type restriction (should fail for LoadBalancer)
kubectl expose deployment test-app --type=LoadBalancer --port=80 -n test-ns
# Expected: Blocked by policy restricting LoadBalancer type