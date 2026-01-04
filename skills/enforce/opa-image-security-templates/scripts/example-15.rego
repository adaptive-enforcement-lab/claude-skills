# Pseudo-code: Full implementation in security.yaml
deny[msg] {
  scan_result := input.metadata.annotations["trivy.scan.result"]
  criticals := count(scan_result.Results[_].Vulnerabilities[_] | _.Severity == "CRITICAL")
  criticals > 0
  msg := sprintf("Image has %d critical vulnerabilities", [criticals])
}