package repository

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/zdrgeo/osmium/pkg/view"
)

type FileViewRepository struct {
	basePath string
}

func NewFileViewRepository(basePath string) *FileViewRepository {
	if basePath == "" {
		userHomePath, err := os.UserHomeDir()

		if err != nil {
			log.Panic(err)
		}

		basePath = userHomePath
	}

	return &FileViewRepository{basePath: basePath}
}

func viewPath(basePath, analysisName, name string) string {
	return filepath.Join(basePath, "osmium", "analysis", analysisName, "view", name)
}

func (repository *FileViewRepository) Add(analysisName, name string, view *view.AnalysisView) {
	viewPath := viewPath(repository.basePath, analysisName, name)

	if err := os.MkdirAll(viewPath, 0750); err != nil {
		log.Panic(err)
	}

	data, err := json.MarshalIndent(view, "", "  ")

	if err != nil {
		log.Panic(err)
	}

	err = os.WriteFile(filepath.Join(viewPath, "view.json"), data, 0660)

	if err != nil {
		log.Panic(err)
	}
}

func (repository *FileViewRepository) Set(analysisName, name string, view *view.AnalysisView) {
	viewPath := viewPath(repository.basePath, analysisName, name)

	if err := os.MkdirAll(viewPath, 0750); err != nil {
		log.Panic(err)
	}

	data, err := json.MarshalIndent(view, "", "  ")

	if err != nil {
		log.Panic(err)
	}

	err = os.WriteFile(filepath.Join(viewPath, "view.json"), data, 0660)

	if err != nil {
		log.Panic(err)
	}
}

func (repository *FileViewRepository) Remove(analysisName, name string) {
	viewPath := viewPath(repository.basePath, analysisName, name)

	err := os.RemoveAll(viewPath)

	if err != nil {
		log.Panic(err)
	}
}

func (repository *FileViewRepository) Get(analysisName, name string) *view.AnalysisView {
	viewPath := viewPath(repository.basePath, analysisName, name)

	data, err := os.ReadFile(filepath.Join(viewPath, "view.json"))

	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		log.Panic(err)
	}

	view := &view.AnalysisView{}

	err = json.Unmarshal(data, view)

	if err != nil {
		log.Panic(err)
	}

	return view
}
