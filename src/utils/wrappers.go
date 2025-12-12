package utils

import (
	"os"
	"path/filepath"
)

const (
	base_path = ".moldr"
)

// Wrapper for directory operations
func DirWrapper(dir string, fn func(string)) {
	DirWrapperWithError(dir, func(dir string) error {
		fn(dir)
		return nil
	})
}

// Wrapper for directory operations that returns an error
func DirWrapperWithError(dir string, fn func(string) error) error {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	return fn(filepath.Join(home_dir, base_path, dir))
}
