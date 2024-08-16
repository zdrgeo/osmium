package analysis

import "log"

type ChangeAnalysisHandler struct {
	source     AnalysisSource
	repository AnalysisRepository
}

func NewChangeAnalysisHandler(source AnalysisSource, repository AnalysisRepository) *ChangeAnalysisHandler {
	return &ChangeAnalysisHandler{source: source, repository: repository}
}

func (handler *ChangeAnalysisHandler) ChangeAnalysis(name, source string, sourceOptions map[string]string) {
	analysis, err := handler.source.Query(sourceOptions)

	if err != nil {
		log.Print(err)
	}

	handler.repository.Set(name, analysis)
}
