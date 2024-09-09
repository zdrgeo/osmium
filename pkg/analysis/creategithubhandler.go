package analysis

import (
	"log"
)

type CreateGitHubHandler struct {
	source     GitHubAnalysisSource
	repository AnalysisRepository
}

func NewCreateGitHubHandler(source GitHubAnalysisSource, repository AnalysisRepository) *CreateGitHubHandler {
	return &CreateGitHubHandler{source: source, repository: repository}
}

func (handler *CreateGitHubHandler) CreateGitHub(name string, spanSize int, repositoryOwner, repositoryName string) {
	analysis, err := handler.source.Query(spanSize, repositoryOwner, repositoryName)

	if err != nil {
		log.Print(err)
	}

	handler.repository.Add(name, analysis)
}
