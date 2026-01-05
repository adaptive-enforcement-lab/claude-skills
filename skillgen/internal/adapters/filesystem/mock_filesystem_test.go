package filesystem

import (
	"fmt"
	"io/fs"
	"strings"
	"sync"

	"github.com/adaptive-enforcement-lab/claude-skills/skillgen/internal/ports"
)

// MockFileSystem is an in-memory implementation of ports.FileSystem for testing.
type MockFileSystem struct {
	files map[string][]byte
	dirs  map[string]bool
	mu    sync.RWMutex

	// Error injection
	readFileError  error
	writeFileError error
	mkdirAllError  error
}

// NewMockFileSystem creates a new in-memory filesystem.
func NewMockFileSystem() *MockFileSystem {
	return &MockFileSystem{
		files: make(map[string][]byte),
		dirs:  make(map[string]bool),
	}
}

// Ensure MockFileSystem implements ports.FileSystem
var _ ports.FileSystem = (*MockFileSystem)(nil)

// ReadFile reads a file from the in-memory filesystem.
func (m *MockFileSystem) ReadFile(path string) ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.readFileError != nil {
		return nil, m.readFileError
	}

	content, ok := m.files[path]
	if !ok {
		return nil, fmt.Errorf("file not found: %s", path)
	}

	return content, nil
}

// WriteFile writes a file to the in-memory filesystem.
func (m *MockFileSystem) WriteFile(path string, data []byte, perm int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.writeFileError != nil {
		return m.writeFileError
	}

	m.files[path] = data
	return nil
}

// MkdirAll creates directories in the in-memory filesystem.
func (m *MockFileSystem) MkdirAll(path string, perm int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.mkdirAllError != nil {
		return m.mkdirAllError
	}

	m.dirs[path] = true
	return nil
}

// Glob finds files matching a pattern (simplified implementation for tests).
func (m *MockFileSystem) Glob(pattern string) ([]string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var matches []string
	for path := range m.files {
		// Simple pattern matching - just check if pattern is a prefix
		if strings.HasPrefix(path, strings.TrimSuffix(pattern, "*")) {
			matches = append(matches, path)
		}
	}

	return matches, nil
}

// Exists checks if a path exists.
func (m *MockFileSystem) Exists(path string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	_, fileExists := m.files[path]
	_, dirExists := m.dirs[path]
	return fileExists || dirExists
}

// IsDir checks if a path is a directory.
func (m *MockFileSystem) IsDir(path string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.dirs[path]
}

// AddFile seeds the filesystem with a file for testing.
func (m *MockFileSystem) AddFile(path string, content []byte) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.files[path] = content
}

// GetFile retrieves written content for assertions.
func (m *MockFileSystem) GetFile(path string) []byte {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.files[path]
}

// InjectReadError causes ReadFile to return an error.
func (m *MockFileSystem) InjectReadError(err error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.readFileError = err
}

// InjectWriteError causes WriteFile to return an error.
func (m *MockFileSystem) InjectWriteError(err error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.writeFileError = err
}

// InjectMkdirAllError causes MkdirAll to return an error.
func (m *MockFileSystem) InjectMkdirAllError(err error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.mkdirAllError = err
}

// Reset clears all state for the next test.
func (m *MockFileSystem) Reset() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.files = make(map[string][]byte)
	m.dirs = make(map[string]bool)
	m.readFileError = nil
	m.writeFileError = nil
	m.mkdirAllError = nil
}

// FileMode implements ports.FileSystem.
func (m *MockFileSystem) FileMode(path string) (fs.FileMode, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if _, ok := m.files[path]; ok {
		return 0644, nil
	}
	if m.dirs[path] {
		return fs.ModeDir | 0755, nil
	}

	return 0, fmt.Errorf("path not found: %s", path)
}
