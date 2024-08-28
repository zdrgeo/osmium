package view

import "github.com/zdrgeo/osmium/pkg/analysis"

type ViewBuilder interface {
	WithNodeNames(nodeNames []string) ViewBuilder
	Build(analysis *analysis.Analysis) *AnalysisView
}

func viewBuilderFactory(builder string, builderOptions map[string]string) ViewBuilder {
	switch builder {
	case "filepath":
		return NewFilePathViewBuilder(builderOptions)
	case "pattern":
		return NewPatternViewBuilder(builderOptions)
	}

	return nil
}
