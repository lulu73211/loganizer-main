package analyzer

import (
	"bufio"
	"context"
	"math/rand"
	"os"
	"time"
)

type Result struct {
	LogID        string `json:"log_id"`
	FilePath     string `json:"file_path"`
	Status       string `json:"status"` // "OK" | "FAILED"
	Message      string `json:"message"`
	ErrorDetails string `json:"error_details"`
}

func AnalyzeOne(ctx context.Context, id, path, typ string) Result {
	res := Result{LogID: id, FilePath: path}

	f, err := os.Open(path)
	if err != nil {
		res.Status = "FAILED"
		res.Message = "Fichier introuvable ou inaccessible."
		res.ErrorDetails = (&FileError{Path: path, Err: err}).Error()
		return res
	}
	defer f.Close()

	// Simuler un temps d’analyse
	rand.Seed(time.Now().UnixNano())
	select {
	case <-time.After(time.Duration(50+rand.Intn(150)) * time.Millisecond):
	case <-ctx.Done():
		res.Status = "FAILED"
		res.Message = "Analyse annulée."
		res.ErrorDetails = ctx.Err().Error()
		return res
	}

	// Vérifier chaque ligne (détection "INVALID_LINE")
	sc := bufio.NewScanner(f)
	line := 0
	for sc.Scan() {
		line++
		txt := sc.Text()
		if len(txt) >= 12 && txt[:12] == "INVALID_LINE" {
			res.Status = "FAILED"
			res.Message = "Erreur de parsing."
			res.ErrorDetails = (&ParseError{Line: line, Snippet: txt}).Error()
			return res
		}
	}
	if err := sc.Err(); err != nil {
		res.Status = "FAILED"
		res.Message = "Erreur de lecture."
		res.ErrorDetails = (&FileError{Path: path, Err: err}).Error()
		return res
	}

	if line == 0 {
		res.Status = "OK"
		res.Message = "Fichier vide, aucune ligne à analyser."
	} else {
		res.Status = "OK"
		res.Message = "Analyse terminée avec succès."
	}
	return res
}
