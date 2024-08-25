package view

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type RenderCSVHandler struct {
	repository ViewRepository
	basePath   string
}

func NewRenderCSVHandler(repository ViewRepository) *RenderCSVHandler {
	return &RenderCSVHandler{repository: repository}
}

func (handler *RenderCSVHandler) RenderCSV(analysisName, viewName, spanName string) {
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

	fileName := filepath.Join(viewPath, "view.csv")

	renderViewToCSVFile(view, spanName, fileName)
}

func renderViewToCSVFile(view *AnalysisView, spanName, fileName string) {
	file, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fmt.Fprint(file, "SpanName,YNodeIndex,YNodeName,XNodeIndex,XNodeName,Value\n")

	for spanName, spanView := range view.SpanViews {
		for yNodeIndex, yNodeName := range view.NodeNames {
			for xNodeIndex, xNodeName := range view.NodeNames {
				fmt.Fprintf(file, "%s,%d,%s,%d,%s,%d\n", spanName, yNodeIndex, yNodeName, xNodeIndex, xNodeName, spanView.Values[yNodeIndex][xNodeIndex])
			}
		}
	}
}
