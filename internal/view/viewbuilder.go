package view

import (
	"math"
	"slices"

	"github.com/zdrgeo/osmium/internal/analysis"
)

type ViewBuilder struct {
	nodeNames []string
}

func (builder *ViewBuilder) WithNodeNames(nodeNames []string) *ViewBuilder {
	builder.nodeNames = nodeNames

	return builder
}

func (builder *ViewBuilder) Build(analysis *analysis.Analysis) *AnalysisView {
	nodeNames := []string{}

	for _, span := range analysis.Spans {
		for nodeName := range span.Nodes {
			if len(builder.nodeNames) == 0 || slices.Contains(builder.nodeNames, nodeName) {
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

				if edge, ok := span.Nodes[nodeName].Edges[edgeNodeName]; ok {
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
