---
name: kyverno-network-security-templates - Examples
description: Code examples for Kyverno Network Security Templates
---

# Kyverno Network Security Templates - Examples


## Example 1: example-1.sh


```bash
kubectl apply -f security.yaml
```



## Example 2: example-2.sh


```bash
kubectl apply -f ingress-class.yaml
```



## Example 3: example-3.sh


```bash
kubectl apply -f ingress-tls.yaml
```



## Example 4: example-4.sh


```bash
kubectl apply -f services.yaml
```



## Example 5: example-5.yaml


```yaml
# Enforced by: services.yaml
# Result: Only ingress-nginx namespace can create LoadBalancer services
# Impact: Prevents accidental exposure of internal services to the internet
```



## Example 6: example-6.yaml


```yaml
# Enforced by: ingress-tls.yaml
# Result: All Ingress objects must define spec.tls with valid secrets
# Impact: Eliminates plaintext HTTP exposure for external services
```



## Example 7: example-7.yaml


```yaml
# Enforced by: security.yaml
# Result: Namespaces must have NetworkPolicy resources before accepting workloads
# Impact: Prevents pods from communicating across namespace boundaries by default
```



## Example 8: example-8.sh


```bash
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
```



