package analysis

type ChangeAnalysisHandler struct {
	source     AnalysisSource
	repository AnalysisRepository
}

func NewChangeAnalysisHandler(source AnalysisSource, repository AnalysisRepository) *ChangeAnalysisHandler {
	return &ChangeAnalysisHandler{source: source, repository: repository}
}

func (handler *ChangeAnalysisHandler) ChangeAnalysis(name string) {
	analysis := handler.source.Query()

	handler.repository.Set(name, analysis)
}
