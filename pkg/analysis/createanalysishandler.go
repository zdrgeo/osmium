package analysis

import "log"

type CreateAnalysisHandler struct {
	source     AnalysisSource
	repository AnalysisRepository
}

func NewCreateAnalysisHandler(source AnalysisSource, repository AnalysisRepository) *CreateAnalysisHandler {
	return &CreateAnalysisHandler{source: source, repository: repository}
}

func (handler *CreateAnalysisHandler) CreateAnalysis(name, source string, sourceOptions map[string]string) {
	analysis, err := handler.source.Query(sourceOptions)

	if err != nil {
		log.Print(err)
	}

	handler.repository.Add(name, analysis)
}
