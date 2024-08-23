package view

import (
	"fmt"
)

type RenderTerminalHandler struct {
	repository ViewRepository
}

func NewRenderTerminalHandler(repository ViewRepository) *RenderTerminalHandler {
	return &RenderTerminalHandler{repository: repository}
}

func (handler *RenderTerminalHandler) RenderTerminal(analysisName, name, spanName string) {
	view := handler.repository.Get(analysisName, name)

	renderViewToTerminal(view, spanName)
}

var valueColors = map[int]int{
	0: 39,
	1: 75,
	2: 111,
	3: 147,
	4: 183,
	5: 219,
}

// https://en.wikipedia.org/wiki/ANSI_escape_code#8-bit
func valueColor(minValue, maxValue, value int) int {
	key := int(float32(value-minValue) / (float32(maxValue-minValue+1) / float32(len(valueColors))))

	return valueColors[key]
}

// https://www.w3.org/TR/xml-entity-names/025.html
func renderViewToTerminal(view *AnalysisView, spanName string) {
	const sinceIndex = 20
	const untilIndex = 60

	nodeNames := view.NodeNames[sinceIndex:untilIndex]

	fmt.Printf("Analysis: %s\n", view.Name)

	spanView := view.SpanViews[spanName]

	fmt.Printf("Span: %s\n", spanView.Name)

	fmt.Print("\n")

	fmt.Print("   ")
	for nodeIndex := range nodeNames {
		fmt.Printf("%2d ", sinceIndex+nodeIndex)
	}
	fmt.Print("\n")

	fmt.Print("  ┌")
	for i := 0; i < len(nodeNames)-1; i++ {
		fmt.Print("──┬")
		// fmt.Printf("%.*s┬", 2, "──────────")
	}
	fmt.Print("──┐\n")
	for nodeIndex, nodeName := range nodeNames {
		fmt.Printf("%2d ", sinceIndex+nodeIndex)
		for edgeNodeIndex := range nodeNames {
			if nodeIndex != edgeNodeIndex {
				fmt.Printf("\033[38;5;%dm%2d\033[0m ", valueColor(spanView.MinValue, spanView.MaxValue, spanView.Values[nodeIndex][edgeNodeIndex]), spanView.Values[nodeIndex][edgeNodeIndex])
			} else {
				// fmt.Printf("%2d ", spanView.Values[nodeIndex][edgeNodeIndex])
				// fmt.Printf("\033[7m%2d\033[0m ", spanView.Values[nodeIndex][edgeNodeIndex])
				fmt.Print("▒▒ ")
			}
		}
		fmt.Printf("%2d %s\n", sinceIndex+nodeIndex, nodeName)
	}
	fmt.Print("  └")
	for i := 0; i < len(nodeNames)-1; i++ {
		fmt.Print("──┴")
	}
	fmt.Print("──┘\n")

	fmt.Print("   ")
	for nodeIndex := range nodeNames {
		fmt.Printf("%2d ", sinceIndex+nodeIndex)
	}
	fmt.Print("\n")
}
