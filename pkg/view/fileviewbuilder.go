package view

import (
	"log"
	"math"
	"path"
	"slices"

	"github.com/zdrgeo/osmium/pkg/analysis"
)

type FileViewBuilder struct {
	ViewBuilder
}

func (builder *FileViewBuilder) WithNodeNames(nodeNames []string) *FileViewBuilder {
	builder.nodeNames = nodeNames

	return builder
}

func containsFile(nodeNames []string, fileName string) bool {
	for _, nodeName := range nodeNames {
		match, err := path.Match(nodeName, fileName)

		if err != nil {
			log.Fatal(err)
		}

		if match {
			return true
		}
	}

	return false
}

func (builder *FileViewBuilder) Build(analysis *analysis.Analysis) *AnalysisView {
	nodeNames := []string{}

	for _, span := range analysis.Spans {
		for nodeName := range span.Nodes {
			if len(builder.nodeNames) == 0 || containsFile(builder.nodeNames, nodeName) {
				nodeNames = append(nodeNames, nodeName)
			}
		}
	}

	slices.Sort(nodeNames)

	spanViews := make(map[string]*SpanView, len(analysis.Spans))

	for spanName, span := range analysis.Spans {
		values := make([][]int, len(nodeNames))
		minValue := math.MaxInt
		maxValue := math.MinInt

		for nodeIndex, nodeName := range nodeNames {
			edgeValues := make([]int, len(nodeNames))

			for edgeNodeIndex, edgeNodeName := range nodeNames {
				value := 0

				if edge, ok := analysis.Spans[""].Nodes[nodeName].Edges[edgeNodeName]; ok {
					value = len(edge.ChangeNames)
				}

				if nodeIndex != edgeNodeIndex {
					if minValue > value {
						minValue = value
					}

					if maxValue < value {
						maxValue = value
					}
				}

				edgeValues[edgeNodeIndex] = value
			}

			values[nodeIndex] = edgeValues
		}

		spanViews[spanName] = &SpanView{Name: span.Name, Size: span.Size, Values: values, MinValue: minValue, MaxValue: maxValue}
	}

	return &AnalysisView{NodeNames: nodeNames, SpanViews: spanViews}
}
