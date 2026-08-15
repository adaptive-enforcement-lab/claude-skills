package extractor

import "strings"

// prepareReferenceBody turns an admonition-converted doc body into content
// suitable for embedding under another heading in reference.md: the doc's
// own leading "# Title" line is dropped (redundant — the caller already
// wraps the body in its own heading), and every remaining heading is
// shifted deeper by headingShift levels so nesting stays coherent.
//
// This is the entire extraction path for reference.md: each doc is used
// exactly once, verbatim, so there is no way for content to be duplicated
// the way the old SectionMapper-driven extraction was.
func prepareReferenceBody(body string, headingShift int) string {
	lines := strings.Split(body, "\n")
	shift := strings.Repeat("#", headingShift)

	var out []string
	droppedTitle := false
	for _, line := range lines {
		trimmed := strings.TrimLeft(line, " ")
		if isHeading(trimmed) {
			if !droppedTitle && headingLevel(trimmed) == 1 {
				droppedTitle = true
				continue
			}
			out = append(out, shift+line)
			continue
		}
		out = append(out, line)
	}

	return strings.TrimSpace(strings.Join(out, "\n"))
}

// isHeading reports whether a (left-trimmed) line is an ATX markdown heading.
func isHeading(trimmed string) bool {
	i := 0
	for i < len(trimmed) && trimmed[i] == '#' {
		i++
	}
	return i > 0 && i <= 6 && (i == len(trimmed) || trimmed[i] == ' ')
}

// headingLevel returns the number of leading '#' characters in a heading line.
func headingLevel(trimmed string) int {
	i := 0
	for i < len(trimmed) && trimmed[i] == '#' {
		i++
	}
	return i
}
