package domain

// Categories are the plugin collections a skill can belong to. They map
// one-to-one onto top-level directories in the docs tree and onto plugin
// directories in the generated marketplace, so a document outside these
// directories produces no skill.
var Categories = []string{"patterns", "enforce", "build", "secure"}

// IsCategory reports whether name is a known plugin collection.
func IsCategory(name string) bool {
	for _, c := range Categories {
		if c == name {
			return true
		}
	}
	return false
}
