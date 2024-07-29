package repository

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/zdrgeo/osmium/internal/analysis"
)

type FileAnalysisRepository struct {
	basePath string
}

func NewFileAnalysisRepository(basePath string) *FileAnalysisRepository {
	if basePath == "" {
		userHomePath, err := os.UserHomeDir()

		if err != nil {
			log.Fatal(err)
		}

		basePath = userHomePath
	}

	return &FileAnalysisRepository{basePath: basePath}
}

func analysisPath(basePath, name string) string {
	return filepath.Join(basePath, "osmium", "analysis", name)
}

func (repository *FileAnalysisRepository) Add(name string, analysis *analysis.Analysis) {
	analysisPath := analysisPath(repository.basePath, name)

	if err := os.MkdirAll(analysisPath, 0750); err != nil {
		log.Fatal(err)
	}

	data, err := json.MarshalIndent(analysis, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(filepath.Join(analysisPath, "analysis.json"), data, 0660)

	if err != nil {
		log.Fatal(err)
	}
}

func (repository *FileAnalysisRepository) Set(name string, analysis *analysis.Analysis) {
	analysisPath := analysisPath(repository.basePath, name)

	if err := os.MkdirAll(analysisPath, 0750); err != nil {
		log.Fatal(err)
	}

	data, err := json.MarshalIndent(analysis, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(filepath.Join(analysisPath, "analysis.json"), data, 0660)

	if err != nil {
		log.Fatal(err)
	}
}

func (repository *FileAnalysisRepository) Remove(name string) {
	analysisPath := analysisPath(repository.basePath, name)

	err := os.RemoveAll(analysisPath)

	if err != nil {
		log.Fatal(err)
	}
}

func (repository *FileAnalysisRepository) Get(name string) *analysis.Analysis {
	analysisPath := analysisPath(repository.basePath, name)

	data, err := os.ReadFile(filepath.Join(analysisPath, "analysis.json"))

	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		log.Fatal(err)
	}

	analysis := &analysis.Analysis{}

	err = json.Unmarshal(data, analysis)

	if err != nil {
		log.Fatal(err)
	}

	return analysis
}
