package analysis

type CreateAnalysisHandler struct {
	source     AnalysisSource
	repository AnalysisRepository
}

func NewCreateAnalysisHandler(source AnalysisSource, repository AnalysisRepository) *CreateAnalysisHandler {
	return &CreateAnalysisHandler{source: source, repository: repository}
}

func (handler *CreateAnalysisHandler) CreateAnalysis(name string) {
	analysis := handler.source.Query()

	handler.repository.Add(name, analysis)
}
