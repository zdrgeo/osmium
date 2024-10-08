package view

import (
	"log"
	"math"
	"regexp"
	"slices"

	"github.com/zdrgeo/osmium/pkg/analysis"
)

type PatternViewBuilder struct {
	options   map[string]string
	nodeNames []string
}

func NewPatternViewBuilder(options map[string]string) *PatternViewBuilder {
	return &PatternViewBuilder{options: options}
}

func (builder *PatternViewBuilder) WithNodeNames(nodeNames []string) ViewBuilder {
	builder.nodeNames = nodeNames

	return builder
}

func (builder *PatternViewBuilder) Build(analysis *analysis.Analysis) *AnalysisView {
	nodeNames := []string{}

	/*
		compiledPatterns, err := compilePatterns(builder.nodeNames)

		if err != nil {
			log.Panic(err)
		}

		for _, span := range analysis.Spans {
			for nodeName := range span.Nodes {
				if len(compiledPatterns) == 0 || matchCompiledPatterns(compiledPatterns, nodeName) {
					nodeNames = append(nodeNames, nodeName)
				}
			}
		}
	*/

	for _, span := range analysis.Spans {
		for nodeName := range span.Nodes {
			if len(builder.nodeNames) == 0 || matchPattern(builder.nodeNames, nodeName) {
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

func matchPattern(patterns []string, text string) bool {
	for _, pattern := range patterns {
		match, err := regexp.MatchString(pattern, text)

		if err != nil {
			log.Panic(err)
		}

		if match {
			return true
		}
	}

	return false
}

func compilePatterns(patterns []string) ([]*regexp.Regexp, error) {
	compiledPatterns := []*regexp.Regexp{}

	for _, pattern := range patterns {
		compiledPattern, err := regexp.Compile(pattern)

		if err != nil {
			return nil, err
		}

		compiledPatterns = append(compiledPatterns, compiledPattern)
	}

	return compiledPatterns, nil
}

func matchCompiledPatterns(compiledPatterns []*regexp.Regexp, text string) bool {
	for _, compiledPattern := range compiledPatterns {
		if compiledPattern.MatchString(text) {
			return true
		}
	}

	return false
}
