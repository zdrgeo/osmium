package analysis

import "log"

type ChangeGitHubHandler struct {
	source     GitHubAnalysisSource
	repository AnalysisRepository
}

func NewChangeGitHubHandler(source GitHubAnalysisSource, repository AnalysisRepository) *ChangeGitHubHandler {
	return &ChangeGitHubHandler{source: source, repository: repository}
}

func (handler *ChangeGitHubHandler) ChangeGitHub(name string, spanSize int, repositoryOwner, repositoryName string) {
	analysis, err := handler.source.Query(spanSize, repositoryOwner, repositoryName)

	if err != nil {
		log.Print(err)
	}

	handler.repository.Set(name, analysis)
}
