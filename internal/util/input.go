package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadInput(day int, name string) (string, error) {
	dayDir := fmt.Sprintf("day%02d", day)
	root := moduleRoot()
	path := filepath.Join(root, "data", dayDir, name)
	b, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read input %s: %w", path, err)
	}
	return string(b), nil
}

func MustReadInput(day int, name string) string {
	s, err := ReadInput(day, name)
	if err != nil {
		panic(err)
	}
	return s
}

// moduleRoot attempts to locate the repository/module root by walking up
// the directory tree until a go.mod file is found. Falls back to the
// current working directory if not found.
func moduleRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		return "."
	}
	for i := 0; i < 16 && dir != "" && dir != "/" && dir != "\\"; i++ {
		if _, statErr := os.Stat(filepath.Join(dir, "go.mod")); statErr == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "."
}

func Input(day int) string {
	return MustReadInput(day, "task.txt")
}
