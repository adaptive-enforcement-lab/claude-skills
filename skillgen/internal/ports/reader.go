package ports

import "github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/domain"

// DocumentReader reads and parses documentation files from the filesystem.
// This interface abstracts file I/O to enable testing without real files.
type DocumentReader interface {
	// ReadDocument reads and parses a single document file.
	ReadDocument(path string) (*domain.Document, error)

	// ListIndexFiles finds all index.md files in the specified root path.
	// Returns absolute paths to each index.md file found.
	ListIndexFiles(rootPath string, categories []string) ([]string, error)
}

// FileSystem abstracts file system operations for testing.
type FileSystem interface {
	// ReadFile reads the entire file content at the given path.
	ReadFile(path string) ([]byte, error)

	// WriteFile writes data to a file, creating it if necessary.
	// Uses atomic write operations (write to temp, then rename).
	WriteFile(path string, data []byte, perm int) error

	// MkdirAll creates all directories in the path with the given permissions.
	MkdirAll(path string, perm int) error

	// Glob returns all files matching the pattern.
	Glob(pattern string) ([]string, error)

	// Exists returns true if the path exists on the filesystem.
	Exists(path string) bool

	// IsDir returns true if the path is a directory.
	IsDir(path string) bool
}
