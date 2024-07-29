package analysis

type DeleteAnalysisHandler struct {
	repository AnalysisRepository
}

func NewDeleteAnalysisHandler(repository AnalysisRepository) *DeleteAnalysisHandler {
	return &DeleteAnalysisHandler{repository: repository}
}

func (handler *DeleteAnalysisHandler) DeleteAnalysis(name string) {
	handler.repository.Remove(name)
}
