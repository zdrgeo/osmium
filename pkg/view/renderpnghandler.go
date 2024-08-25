package view

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

type RenderPNGHandler struct {
	repository ViewRepository
	basePath   string
}

func NewRenderPNGHandler(repository ViewRepository) *RenderPNGHandler {
	return &RenderPNGHandler{repository: repository}
}

func (handler *RenderPNGHandler) RenderPNG(analysisName, viewName, spanName string) {
	view := handler.repository.Get(analysisName, viewName)

	basePath := handler.basePath

	if basePath == "" {
		userHomePath, err := os.UserHomeDir()

		if err != nil {
			log.Fatal(err)
		}

		basePath = userHomePath
	}

	viewPath := viewPath(basePath, analysisName, viewName)

	if err := os.MkdirAll(viewPath, 0750); err != nil {
		log.Fatal(err)
	}

	fileName := filepath.Join(viewPath, "view.png")

	renderViewToPNGFile(view, spanName, fileName)
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

	for yNodeIndex := range view.NodeNames {
		for xNodeIndex := range view.NodeNames {
			fmt.Printf("%d,%d,%d\n", xNodeIndex, xNodeIndex, spanView.Values[yNodeIndex][xNodeIndex])
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
