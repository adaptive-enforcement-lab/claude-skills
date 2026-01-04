# Trivy example
trivy image --format json registry.company.com/app:v1.2.3 > scan.json

# Push image with scan annotation
crane mutate registry.company.com/app:v1.2.3 \
  --annotation trivy.scan.result="$(cat scan.json)"