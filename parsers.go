package code

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ParseFile reads the file at path (relative or absolute) and parses it
// according to its extension.
func ParseFile(path string) (map[string]any, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("resolve path %q: %w", path, err)
	}

	content, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	return Parse(content, strings.ToLower(filepath.Ext(absPath)))
}

// Parse decodes content in the format given by the file extension.
func Parse(content []byte, ext string) (map[string]any, error) {
	var data map[string]any

	switch ext {
	case ".json":
		if err := json.Unmarshal(content, &data); err != nil {
			return nil, fmt.Errorf("parse json: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported file format: %q", ext)
	}

	return data, nil
}
