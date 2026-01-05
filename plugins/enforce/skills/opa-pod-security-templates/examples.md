---
name: opa-pod-security-templates - Examples
description: Code examples for OPA Pod Security Templates
---

# OPA Pod Security Templates - Examples


## Example 1: example-1.sh


```bash
kubectl apply -f overview.yaml
```



## Example 2: example-2.sh


```bash
kubectl apply -f capabilities.yaml
```



## Example 3: example-3.sh


```bash
kubectl apply -f contexts.yaml
```



## Example 4: example-4.sh


```bash
kubectl apply -f escalation.yaml
```



## Example 5: example-5.yaml


```yaml
# Enforced by: overview.yaml
# Result: No containers can run with privileged: true
# Impact: Eliminates most container breakout vectors
```



## Example 6: example-6.yaml


```yaml
# Enforced by: capabilities.yaml
# Result: All containers must drop CAP_SYS_ADMIN, CAP_NET_RAW
# Impact: Prevents kernel manipulation and network sniffing
```



## Example 7: example-7.yaml


```yaml
# Enforced by: contexts.yaml
# Result: Containers must define runAsNonRoot: true and runAsUser > 0
# Impact: Prevents root-level filesystem access and privilege escalation
```



## Example 8: example-8.yaml


```yaml
# Enforced by: escalation.yaml
# Result: Containers must set allowPrivilegeEscalation: false
# Impact: Blocks setuid binaries and capability inheritance
```



## Example 9: example-9.sh


```bash
# Test privileged container block (should fail)
kubectl run privileged-test --image=nginx --privileged=true
# Expected: Admission denied by overview.yaml

# Test capability violation (should fail with CAP_SYS_ADMIN)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: cap-test
spec:
  containers:
    - name: nginx
      image: nginx
      securityContext:
        capabilities:
          add:
            - SYS_ADMIN
EOF
# Expected: Admission denied by capabilities.yaml

# Test root execution (should fail with runAsUser: 0)
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
# Expected: Admission denied by contexts.yaml

# Test privilege escalation (should fail with allowPrivilegeEscalation: true)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: escalation-test
spec:
  containers:
    - name: nginx
      image: nginx
      securityContext:
        allowPrivilegeEscalation: true
EOF
# Expected: Admission denied by escalation.yaml

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
# Expected: Admission allowed by all policies
```



## Example 10: example-10.yaml


```yaml
spec:
  securityContext:
    runAsNonRoot: true
    runAsUser: 1000
    runAsGroup: 1000
    fsGroup: 1000
    seccompProfile:
      type: RuntimeDefault
```



## Example 11: example-11.yaml


```yaml
spec:
  containers:
    - name: app
      securityContext:
        allowPrivilegeEscalation: false
        readOnlyRootFilesystem: true
        capabilities:
          drop:
            - ALL
          add:
            - NET_BIND_SERVICE  # Only if binding to port 80/443
```



## Example 12: example-12.sh


```bash
kubectl get pods --all-namespaces -o json | \
  jq '.items[] | select(.spec.containers[].securityContext.privileged == true)'
```



## Example 13: example-13.sh


```bash
kubectl apply -f overview.yaml  # Set enforcementAction: warn
```



## Example 14: example-14.sh


```bash
kubectl get constrainttemplates
kubectl get <constraint-name> -o yaml
# Check status.violations for non-compliant pods
```



