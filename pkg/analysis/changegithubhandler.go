package analysis

import "log"

type ChangeGitHubHandler struct {
	source     GitHubAnalysisSource
	repository AnalysisRepository
}

func NewChangeGitHubHandler(source GitHubAnalysisSource, repository AnalysisRepository) *ChangeGitHubHandler {
	return &ChangeGitHubHandler{source: source, repository: repository}
}

func (handler *ChangeGitHubHandler) ChangeGitHub(name, repositoryOwner, repositoryName string) {
	analysis, err := handler.source.Query(repositoryOwner, repositoryName)

	if err != nil {
		log.Print(err)
	}

	handler.repository.Set(name, analysis)
}
