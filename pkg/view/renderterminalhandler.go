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
	const nodeStartIndex = 0
	const edgeNodeStartIndex = 0

	const count = 25

	nodeNames := view.NodeNames[nodeStartIndex : nodeStartIndex+count]
	edgeNodeNames := view.NodeNames[edgeNodeStartIndex : edgeNodeStartIndex+count]

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
			if nodeIndex != edgeNodeIndex {
				fmt.Printf("\033[38;5;%dm%2d\033[0m ", valueColor(spanView.MinValue, spanView.MaxValue, spanView.Values[nodeIndex][edgeNodeIndex]), spanView.Values[nodeIndex][edgeNodeIndex])
			} else {
				// fmt.Printf("%2d ", spanView.Values[nodeIndex][edgeNodeIndex])
				// fmt.Printf("\033[7m%2d\033[0m ", spanView.Values[nodeIndex][edgeNodeIndex])
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

	fmt.Fprintf(writer, "   Nodes (%d - %d)\t   Edge Nodes (%d - %d)\n", nodeStartIndex, nodeStartIndex+count, edgeNodeStartIndex, edgeNodeStartIndex+count)

	for index := range count {
		fmt.Fprintf(writer, "%2d (%d) %s\t%2d (%d) %s\n", index, nodeStartIndex+index, nodeNames[index], index, edgeNodeStartIndex+index, edgeNodeNames[index])
	}

	writer.Flush()
}
