---
name: kyverno-pod-security-templates - Examples
description: Code examples for Kyverno Pod Security Templates
---

# Kyverno Pod Security Templates - Examples


## Example 1: example-1.sh


```bash
kubectl apply -f standards.yaml
```



## Example 2: example-2.sh


```bash
kubectl apply -f privileges.yaml
```



## Example 3: example-3.sh


```bash
kubectl apply -f profiles.yaml
```



## Example 4: example-4.yaml


```yaml
# Enforced by: privileges.yaml
# Result: No containers can run with privileged: true
# Impact: Eliminates most container breakout vectors
```



## Example 5: example-5.yaml


```yaml
# Enforced by: profiles.yaml
# Result: Containers must define runAsNonRoot: true
# Impact: Prevents root-level filesystem access and privilege escalation
```



## Example 6: example-6.yaml


```yaml
# Enforced by: standards.yaml
# Result: Pods must define securityContext.seccompProfile
# Impact: Reduces kernel attack surface by blocking dangerous syscalls
```



## Example 7: example-7.sh


```bash
# Test privileged container block (should fail)
kubectl run privileged-test --image=nginx --privileged=true
# Expected: Blocked by privilege restriction policy

# Test root user block (should fail)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: root-test
spec:
  containers:
    - name: nginx
      image: nginx
      securityContext:
        runAsUser: 0
EOF
# Expected: Blocked by non-root requirement policy

# Test hostPath mount block (should fail)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: hostpath-test
spec:
  containers:
    - name: nginx
      image: nginx
      volumeMounts:
        - name: host
          mountPath: /host
  volumes:
    - name: host
      hostPath:
        path: /
EOF
# Expected: Blocked by Pod Security Standards policy

# Test compliant pod (should succeed)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: compliant-test
spec:
  securityContext:
    runAsNonRoot: true
    runAsUser: 1000
    seccompProfile:
      type: RuntimeDefault
  containers:
    - name: nginx
      image: nginx
      securityContext:
        allowPrivilegeEscalation: false
        readOnlyRootFilesystem: true
        capabilities:
          drop:
            - ALL
EOF
# Expected: Allowed by all policies
```



## Example 8: example-8.sh


```bash
kubectl get psp
kubectl get pods --all-namespaces -o jsonpath='{range .items[*]}{.metadata.name}{"\t"}{.metadata.annotations.kubernetes\.io/psp}{"\n"}{end}'
```



## Example 9: example-9.sh


```bash
kubectl apply -f standards.yaml
kubectl apply -f privileges.yaml
kubectl apply -f profiles.yaml
```



## Example 10: example-10.sh


```bash
kubectl get polr -A  # Policy Reports
kubectl describe polr <report-name> -n <namespace>
```



