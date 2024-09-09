package analysis

import "log"

type ChangeGitHandler struct {
	source     GitAnalysisSource
	repository AnalysisRepository
}

func NewChangeGitHandler(source GitAnalysisSource, repository AnalysisRepository) *ChangeGitHandler {
	return &ChangeGitHandler{source: source, repository: repository}
}

func (handler *ChangeGitHandler) ChangeGit(name string, spanSize int, repositoryURL, repositoryPath string) {
	analysis, err := handler.source.Query(spanSize, repositoryURL, repositoryPath)

	if err != nil {
		log.Print(err)
	}

	handler.repository.Set(name, analysis)
}
