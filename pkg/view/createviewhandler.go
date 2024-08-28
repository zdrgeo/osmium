package view

import (
	"github.com/zdrgeo/osmium/pkg/analysis"
)

type CreateViewHandler struct {
	analysisRepository analysis.AnalysisRepository
	repository         ViewRepository
}

func NewCreateViewHandler(analysisRepository analysis.AnalysisRepository, repository ViewRepository) *CreateViewHandler {
	return &CreateViewHandler{analysisRepository: analysisRepository, repository: repository}
}

func (handler *CreateViewHandler) CreateView(analysisName, name string, nodeNames []string, builder string, builderOptions map[string]string) {
	analysis := handler.analysisRepository.Get(analysisName)

	viewBuilder := viewBuilderFactory(builder, builderOptions)

	view := viewBuilder.
		WithNodeNames(nodeNames).
		Build(analysis)

	handler.repository.Add(analysisName, name, view)
}
