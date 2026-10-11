// Package parsers reads configuration files and decodes them into maps.
package parsers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go.yaml.in/yaml/v3"
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

	data, err := Parse(content, filepath.Ext(absPath))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}

	return data, nil
}

// Parse decodes content in the format given by the file extension.
// Numbers are always returned as float64, regardless of the format.
func Parse(content []byte, ext string) (map[string]any, error) {
	var data map[string]any

	switch strings.ToLower(ext) {
	case ".json":
		if err := json.Unmarshal(content, &data); err != nil {
			return nil, fmt.Errorf("parse json: %w", err)
		}
	case ".yml", ".yaml":
		if err := yaml.Unmarshal(content, &data); err != nil {
			return nil, fmt.Errorf("parse yaml: %w", err)
		}
		normalizeNumbers(data)
	default:
		return nil, fmt.Errorf("unsupported file format: %q", ext)
	}

	if data == nil {
		data = map[string]any{}
	}

	return data, nil
}

// normalizeNumbers converts YAML integers to float64 in place, so that the
// same value read from JSON and YAML compares as equal.
func normalizeNumbers(value any) any {
	switch v := value.(type) {
	case int:
		return float64(v)
	case int64:
		return float64(v)
	case uint64:
		return float64(v)
	case map[string]any:
		for key, item := range v {
			v[key] = normalizeNumbers(item)
		}
	case []any:
		for i, item := range v {
			v[i] = normalizeNumbers(item)
		}
	}

	return value
}
