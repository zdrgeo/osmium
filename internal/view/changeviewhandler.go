package view

import "github.com/zdrgeo/osmium/internal/analysis"

type ChangeViewHandler struct {
	analysisRepository analysis.AnalysisRepository
	repository         ViewRepository
}

func NewChangeViewHandler(analysisRepository analysis.AnalysisRepository, repository ViewRepository) *ChangeViewHandler {
	return &ChangeViewHandler{analysisRepository: analysisRepository, repository: repository}
}

func (handler *ChangeViewHandler) ChangeView(analysisName, name string, nodeNames []string) {
	analysis := handler.analysisRepository.Get(analysisName)

	// builder := &ViewBuilder{}
	builder := &FileViewBuilder{}

	view := builder.
		WithNodeNames(nodeNames).
		Build(analysis)

	handler.repository.Set(analysisName, name, view)
}
