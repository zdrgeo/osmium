package analysis

import (
	"log"
)

type CreateGitHandler struct {
	source     GitAnalysisSource
	repository AnalysisRepository
}

func NewCreateGitHandler(source GitAnalysisSource, repository AnalysisRepository) *CreateGitHandler {
	return &CreateGitHandler{source: source, repository: repository}
}

func (handler *CreateGitHandler) CreateGit(name string, spanSize int, repositoryURL, repositoryPath string) {
	analysis, err := handler.source.Query(spanSize, repositoryURL, repositoryPath)

	if err != nil {
		log.Print(err)
	}

	handler.repository.Add(name, analysis)
}
