package cmd

import (
	"context"
	"fmt"
	"sync"

	"github.com/spf13/cobra"

	"github.com/axellelanca/go_loganizer/internal/analyzer"
	"github.com/axellelanca/go_loganizer/internal/config"
	"github.com/axellelanca/go_loganizer/internal/reporter"
)

var (
	cfgPath string
	outPath string
)

func init() {
	analyzeCmd := &cobra.Command{
		Use:   "analyze",
		Short: "Analyse les logs définis dans un fichier config JSON",
		RunE:  runAnalyze,
	}
	analyzeCmd.Flags().StringVarP(&cfgPath, "config", "c", "config.json", "Chemin du fichier de configuration JSON")
	analyzeCmd.Flags().StringVarP(&outPath, "output", "o", "", "Chemin du fichier JSON de rapport à créer (ex: report.json)")

	rootCmd.AddCommand(analyzeCmd)
}

func runAnalyze(cmd *cobra.Command, args []string) error {
	targets, err := config.Load(cfgPath)
	if err != nil {
		return err
	}

	ctx := context.Background()
	results := make([]analyzer.Result, 0, len(targets))
	ch := make(chan analyzer.Result)
	var wg sync.WaitGroup
	wg.Add(len(targets))

	for _, t := range targets {
		t := t
		go func() {
			defer wg.Done()
			r := analyzer.AnalyzeOne(ctx, t.ID, t.Path, t.Type)
			ch <- r
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for r := range ch {
		fmt.Printf("[%s] %s — %s — %s\n", r.LogID, r.FilePath, r.Status, r.Message)
		if r.ErrorDetails != "" {
			fmt.Printf("    erreur: %s\n", r.ErrorDetails)
		}
		results = append(results, r)
	}

	if outPath != "" {
		final, err := reporter.ExportJSON(outPath, results)
		if err != nil {
			return err
		}
		fmt.Printf("Rapport écrit dans %s\n", final)
	}
	return nil
}
