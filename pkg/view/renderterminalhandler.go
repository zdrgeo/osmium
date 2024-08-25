package view

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type RenderTerminalHandler struct {
	repository ViewRepository
}

func NewRenderTerminalHandler(repository ViewRepository) *RenderTerminalHandler {
	return &RenderTerminalHandler{repository: repository}
}

func (handler *RenderTerminalHandler) RenderTerminal(analysisName, name, spanName string, nodeStart, edgeNodeStart, nodeCount int) {
	view := handler.repository.Get(analysisName, name)

	renderViewToTerminal(view, spanName, nodeStart, edgeNodeStart, nodeCount)
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
func renderViewToTerminal(view *AnalysisView, spanName string, xNodeStart, yNodeStart, nodeCount int) {
	if nodeCount < 0 || nodeCount > 100 || nodeCount > len(view.NodeNames) {
		nodeCount = len(view.NodeNames)
	}

	if xNodeStart < 0 {
		xNodeStart = 0
	}

	if xNodeStart+nodeCount > len(view.NodeNames) {
		xNodeStart = len(view.NodeNames) - nodeCount
	}

	if yNodeStart < 0 {
		yNodeStart = 0
	}

	if yNodeStart+nodeCount > len(view.NodeNames) {
		yNodeStart = len(view.NodeNames) - nodeCount
	}

	yNodeNames := view.NodeNames[yNodeStart : yNodeStart+nodeCount]
	xNodeNames := view.NodeNames[xNodeStart : xNodeStart+nodeCount]

	spanView := view.SpanViews[spanName]

	fmt.Printf("Analysis: %s\n", view.Name)
	fmt.Printf("Span: %s\n", spanView.Name)

	fmt.Println()

	fmt.Print("   ")
	for xNodeIndex := range xNodeNames {
		fmt.Printf("%2d ", xNodeIndex)
	}
	fmt.Println()

	fmt.Print("  ┌")
	for range len(xNodeNames) - 1 {
		fmt.Print("──┬")
	}
	fmt.Println("──┐")
	for yNodeIndex := range yNodeNames {
		fmt.Printf("%2d ", yNodeIndex)
		for xNodeIndex := range xNodeNames {
			if xNodeStart+xNodeIndex != yNodeStart+yNodeIndex {
				fmt.Printf("\033[38;5;%dm%2d\033[0m ", valueColor(spanView.MinValue, spanView.MaxValue, spanView.Values[yNodeStart+yNodeIndex][xNodeStart+xNodeIndex]), spanView.Values[yNodeStart+yNodeIndex][xNodeStart+xNodeIndex])
			} else {
				// fmt.Printf("%2d ", spanView.Values[yNodeStart+yNodeIndex][xNodeStart+xNodeIndex])
				// fmt.Printf("\033[7m%2d\033[0m ", spanView.Values[yNodeStart+yNodeIndex][xNodeStart+xNodeIndex])
				fmt.Print("▒▒ ")
			}
		}
		fmt.Printf("%2d\n", yNodeIndex)
	}
	fmt.Print("  └")
	for range len(xNodeNames) - 1 {
		fmt.Print("──┴")
	}
	fmt.Println("──┘")

	fmt.Print("   ")
	for xNodeIndex := range xNodeNames {
		fmt.Printf("%2d ", xNodeIndex)
	}
	fmt.Println()

	fmt.Println()

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

	fmt.Fprintf(writer, "\tY Nodes (%d - %d)\t\tX Nodes (%d - %d)\n", yNodeStart, yNodeStart+nodeCount, xNodeStart, xNodeStart+nodeCount)

	for nodeIndex := range nodeCount {
		fmt.Fprintf(writer, "%2d\t(%d) %s\t%2d\t(%d) %s\n", nodeIndex, yNodeStart+nodeIndex, yNodeNames[nodeIndex], nodeIndex, xNodeStart+nodeIndex, xNodeNames[nodeIndex])
	}

	writer.Flush()
}
