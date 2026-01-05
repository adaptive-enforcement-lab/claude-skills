# Install kyverno CLI
brew install kyverno/kyverno/kyverno

# Test JMESPath expression
kyverno jp query -i manifest.yaml 'spec.template.spec.containers[*].name'