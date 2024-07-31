package view

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

type RenderViewHandler struct {
	repository ViewRepository
}

func NewRenderViewHandler(repository ViewRepository) *RenderViewHandler {
	return &RenderViewHandler{repository: repository}
}

func (handler *RenderViewHandler) RenderView(analysisName, name, spanName string) {
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
	fmt.Printf("Analysis: %s\n", view.Name)

	spanView := view.SpanViews[spanName]

	fmt.Printf("Span: %s\n", spanView.Name)

	fmt.Print("\n")

	fmt.Print("   ")
	for nodeIndex := range view.NodeNames {
		fmt.Printf("%2d ", nodeIndex)
	}
	fmt.Print("\n")

	fmt.Print("  ┌")
	for i := 0; i < len(view.NodeNames)-1; i++ {
		fmt.Print("──┬")
		// fmt.Printf("%.*s┬", 2, "──────────")
	}
	fmt.Print("──┐\n")
	for nodeIndex, nodeName := range view.NodeNames {
		fmt.Printf("%2d ", nodeIndex)
		for edgeNodeIndex := range view.NodeNames {
			if nodeIndex != edgeNodeIndex {
				fmt.Printf("\033[38;5;%dm%2d\033[0m ", valueColor(spanView.MinValue, spanView.MaxValue, spanView.Values[nodeIndex][edgeNodeIndex]), spanView.Values[nodeIndex][edgeNodeIndex])
			} else {
				// fmt.Printf("%2d ", spanView.Values[nodeIndex][edgeNodeIndex])
				// fmt.Printf("\033[7m%2d\033[0m ", spanView.Values[nodeIndex][edgeNodeIndex])
				fmt.Print("▒▒ ")
			}
		}
		fmt.Printf("%2d %s\n", nodeIndex, nodeName)
	}
	fmt.Print("  └")
	for i := 0; i < len(view.NodeNames)-1; i++ {
		fmt.Print("──┴")
	}
	fmt.Print("──┘\n")

	fmt.Print("   ")
	for nodeIndex := range view.NodeNames {
		fmt.Printf("%2d ", nodeIndex)
	}
	fmt.Print("\n")
}

func renderViewToCSVFile(view *AnalysisView, fileName string) {
	file, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fmt.Fprint(file, "SpanName,NodeIndex,NodeName,EdgeNodeIndex,EdgeNodeName,Value\n")

	for spanName, spanView := range view.SpanViews {
		for nodeIndex, nodeName := range view.NodeNames {
			for edgeNodeIndex, edgeNodeName := range view.NodeNames {
				fmt.Fprintf(file, "%s,%d,%s,%d,%s,%d\n", spanName, nodeIndex, nodeName, edgeNodeIndex, edgeNodeName, spanView.Values[nodeIndex][edgeNodeIndex])
			}
		}
	}
}

func renderViewToPNGFile(view *AnalysisView, spanName, fileName string) {
	spanView := view.SpanViews[spanName]

	myimage := image.NewRGBA(image.Rect(0, 0, 220, 220)) // x1,y1,  x2,y2 of background rectangle
	mygreen := color.RGBA{0, 100, 0, 255}                //  R, G, B, Alpha

	// backfill entire background surface with color mygreen
	draw.Draw(myimage, myimage.Bounds(), &image.Uniform{mygreen}, image.Point{}, draw.Src)

	red_rect := image.Rect(60, 80, 120, 160) //  geometry of 2nd rectangle which we draw atop above rectangle
	myred := color.RGBA{200, 0, 0, 255}

	// create a red rectangle atop the green surface
	draw.Draw(myimage, red_rect, &image.Uniform{myred}, image.Point{}, draw.Src)

	for nodeIndex := range view.NodeNames {
		for edgeNodeIndex := range view.NodeNames {
			fmt.Printf("%d,%d,%d\n", nodeIndex, edgeNodeIndex, spanView.Values[nodeIndex][edgeNodeIndex])
			red_rect := image.Rect(20, 80, 120, 160) //  geometry of 2nd rectangle which we draw atop above rectangle
			myred := color.RGBA{200, 0, 0, 255}
			draw.Draw(myimage, red_rect, &image.Uniform{myred}, image.Point{}, draw.Src)
		}
	}

	myfile, err := os.Create(fileName) // ... now lets save output image
	if err != nil {
		panic(err)
	}
	defer myfile.Close()
	png.Encode(myfile, myimage) // output file /tmp/two_rectangles.png
}
