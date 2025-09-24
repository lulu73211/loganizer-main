package reporter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func ExportJSON(path string, payload any) (string, error) {
	// Crée les dossiers si besoin (bonnes pratiques, tu peux retirer si tu veux être ultra strict)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return "", fmt.Errorf("create dirs for %s: %w", path, err)
	}

	f, err := os.Create(path)
	if err != nil {
		return "", fmt.Errorf("create %s: %w", path, err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(payload); err != nil {
		return "", fmt.Errorf("write json: %w", err)
	}
	return path, nil
}
