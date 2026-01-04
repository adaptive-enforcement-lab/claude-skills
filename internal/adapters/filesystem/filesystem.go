package filesystem

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// FileSystem implements ports.FileSystem using the OS filesystem.
type FileSystem struct{}

// NewFileSystem creates a new OS filesystem adapter.
func NewFileSystem() *FileSystem {
	return &FileSystem{}
}

// ReadFile reads the entire file content at the given path.
func (f *FileSystem) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// WriteFile writes data to a file atomically.
// Writes to a temp file first, then renames to avoid partial writes.
func (f *FileSystem) WriteFile(path string, data []byte, perm int) error {
	// Validate path to prevent directory traversal
	cleanPath := filepath.Clean(path)

	// Create parent directories if they don't exist
	dir := filepath.Dir(cleanPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	// Write to temp file first for atomic operation
	tempPath := cleanPath + ".tmp"
	if err := os.WriteFile(tempPath, data, fs.FileMode(perm)); err != nil {
		return fmt.Errorf("failed to write temp file %s: %w", tempPath, err)
	}

	// Atomic rename
	if err := os.Rename(tempPath, cleanPath); err != nil {
		os.Remove(tempPath) // Clean up temp file on error
		return fmt.Errorf("failed to rename %s to %s: %w", tempPath, cleanPath, err)
	}

	return nil
}

// MkdirAll creates all directories in the path with the given permissions.
func (f *FileSystem) MkdirAll(path string, perm int) error {
	return os.MkdirAll(path, fs.FileMode(perm))
}

// Glob returns all files matching the pattern.
func (f *FileSystem) Glob(pattern string) ([]string, error) {
	return filepath.Glob(pattern)
}

// Exists returns true if the path exists on the filesystem.
func (f *FileSystem) Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// IsDir returns true if the path is a directory.
func (f *FileSystem) IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// FindIndexFiles recursively finds all index.md files in the given categories.
// Categories are subdirectories under rootPath (e.g., "patterns", "enforce", "build", "secure").
func FindIndexFiles(filesystem *FileSystem, rootPath string, categories []string) ([]string, error) {
	var indexFiles []string

	for _, category := range categories {
		categoryPath := filepath.Join(rootPath, category)

		// Skip if category directory doesn't exist
		if !filesystem.Exists(categoryPath) {
			continue
		}

		// Walk the category directory
		err := filepath.WalkDir(categoryPath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			// Skip if not a file
			if d.IsDir() {
				return nil
			}

			// Only match index.md files
			if strings.ToLower(d.Name()) == "index.md" {
				absPath, err := filepath.Abs(path)
				if err != nil {
					return err
				}
				indexFiles = append(indexFiles, absPath)
			}

			return nil
		})

		if err != nil {
			return nil, fmt.Errorf("failed to walk category %s: %w", category, err)
		}
	}

	return indexFiles, nil
}

// DetermineCategory extracts the category name from a file path.
// Example: "/docs/patterns/idempotency/index.md" -> "patterns"
func DetermineCategory(path string, categories []string) string {
	cleanPath := filepath.Clean(path)
	parts := strings.Split(cleanPath, string(filepath.Separator))

	// Find the category in the path
	for i, part := range parts {
		for _, category := range categories {
			if part == category {
				return category
			}
		}

		// Also check if this is right after "docs"
		if i > 0 && parts[i-1] == "docs" {
			for _, category := range categories {
				if part == category {
					return category
				}
			}
		}
	}

	return ""
}
