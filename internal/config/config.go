package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type LogTarget struct {
	ID   string `json:"id"`
	Path string `json:"path"`
	Type string `json:"type"`
}

func Load(path string) ([]LogTarget, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open config %s: %w", path, err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("read config %s: %w", path, err)
	}

	var targets []LogTarget
	if err := json.Unmarshal(data, &targets); err != nil {
		return nil, fmt.Errorf("parse config JSON: %w", err)
	}

	for i, t := range targets {
		if t.ID == "" || t.Path == "" || t.Type == "" {
			return nil, fmt.Errorf("invalid config entry at index %d: fields id, path, type are required", i)
		}
	}
	return targets, nil
}
