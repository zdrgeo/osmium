package view

import "github.com/zdrgeo/osmium/pkg/analysis"

type ChangeViewHandler struct {
	analysisRepository analysis.AnalysisRepository
	repository         ViewRepository
}

func NewChangeViewHandler(analysisRepository analysis.AnalysisRepository, repository ViewRepository) *ChangeViewHandler {
	return &ChangeViewHandler{analysisRepository: analysisRepository, repository: repository}
}

func (handler *ChangeViewHandler) ChangeView(analysisName, name string, nodeNames []string, builder string, builderOptions map[string]string) {
	analysis := handler.analysisRepository.Get(analysisName)

	viewBuilder := viewBuilderFactory(builder, builderOptions)

	view := viewBuilder.
		WithNodeNames(nodeNames).
		Build(analysis)

	handler.repository.Set(analysisName, name, view)
}
