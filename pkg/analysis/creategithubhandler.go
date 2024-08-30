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

func (handler *CreateGitHubHandler) CreateGitHub(name, repositoryOwner, repositoryName string) {
	analysis, err := handler.source.Query(repositoryOwner, repositoryName)

	if err != nil {
		log.Print(err)
	}

	handler.repository.Add(name, analysis)
}
