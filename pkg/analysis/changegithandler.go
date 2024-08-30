package analysis

import "log"

type ChangeGitHandler struct {
	source     GitAnalysisSource
	repository AnalysisRepository
}

func NewChangeGitHandler(source GitAnalysisSource, repository AnalysisRepository) *ChangeGitHandler {
	return &ChangeGitHandler{source: source, repository: repository}
}

func (handler *ChangeGitHandler) ChangeGit(name, repositoryURL, repositoryPath string) {
	analysis, err := handler.source.Query(repositoryURL, repositoryPath)

	if err != nil {
		log.Print(err)
	}

	handler.repository.Set(name, analysis)
}
