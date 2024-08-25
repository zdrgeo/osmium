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
func renderViewToTerminal(view *AnalysisView, spanName string, nodeStart, edgeNodeStart, nodeCount int) {
	if nodeCount < 0 || nodeCount > 100 || nodeCount > len(view.NodeNames) {
		nodeCount = len(view.NodeNames)
	}

	if nodeStart < 0 {
		nodeStart = 0
	}

	if nodeStart+nodeCount > len(view.NodeNames)-1 {
		nodeStart = len(view.NodeNames) - 1 - nodeCount
	}

	if edgeNodeStart < 0 {
		edgeNodeStart = 0
	}

	if edgeNodeStart+nodeCount > len(view.NodeNames)-1 {
		edgeNodeStart = len(view.NodeNames) - 1 - nodeCount
	}

	nodeNames := view.NodeNames[nodeStart : nodeStart+nodeCount]
	edgeNodeNames := view.NodeNames[edgeNodeStart : edgeNodeStart+nodeCount]

	spanView := view.SpanViews[spanName]

	fmt.Printf("Analysis: %s\n", view.Name)
	fmt.Printf("Span: %s\n", spanView.Name)

	fmt.Println()

	fmt.Print("   ")
	for edgeNodeIndex := range edgeNodeNames {
		fmt.Printf("%2d ", edgeNodeIndex)
	}
	fmt.Println()

	fmt.Print("  ┌")
	for range len(edgeNodeNames) - 1 {
		fmt.Print("──┬")
	}
	fmt.Println("──┐")
	for nodeIndex := range nodeNames {
		fmt.Printf("%2d ", nodeIndex)
		for edgeNodeIndex := range edgeNodeNames {
			if nodeStart+nodeIndex != edgeNodeStart+edgeNodeIndex {
				fmt.Printf("\033[38;5;%dm%2d\033[0m ", valueColor(spanView.MinValue, spanView.MaxValue, spanView.Values[nodeStart+nodeIndex][edgeNodeStart+edgeNodeIndex]), spanView.Values[nodeStart+nodeIndex][edgeNodeStart+edgeNodeIndex])
			} else {
				// fmt.Printf("%2d ", spanView.Values[nodeStart+nodeIndex][edgeNodeStart+edgeNodeIndex])
				// fmt.Printf("\033[7m%2d\033[0m ", spanView.Values[nodeStart+nodeIndex][edgeNodeStart+edgeNodeIndex])
				fmt.Print("▒▒ ")
			}
		}
		fmt.Printf("%2d\n", nodeIndex)
	}
	fmt.Print("  └")
	for range len(edgeNodeNames) - 1 {
		fmt.Print("──┴")
	}
	fmt.Println("──┘")

	fmt.Print("   ")
	for edgeNodeIndex := range edgeNodeNames {
		fmt.Printf("%2d ", edgeNodeIndex)
	}
	fmt.Println()

	fmt.Println()

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

	fmt.Fprintf(writer, "\tNodes (%d - %d)\t\tEdge Nodes (%d - %d)\n", nodeStart, nodeStart+nodeCount, edgeNodeStart, edgeNodeStart+nodeCount)

	for nodeIndex := range nodeCount {
		fmt.Fprintf(writer, "%2d\t(%d) %s\t%2d\t(%d) %s\n", nodeIndex, nodeStart+nodeIndex, nodeNames[nodeIndex], nodeIndex, edgeNodeStart+nodeIndex, edgeNodeNames[nodeIndex])
	}

	writer.Flush()
}
