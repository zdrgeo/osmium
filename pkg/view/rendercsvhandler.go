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

	fmt.Fprint(file, "SpanName,NodeIndex,NodeName,EdgeNodeIndex,EdgeNodeName,Value\n")

	for spanName, spanView := range view.SpanViews {
		for nodeIndex, nodeName := range view.NodeNames {
			for edgeNodeIndex, edgeNodeName := range view.NodeNames {
				fmt.Fprintf(file, "%s,%d,%s,%d,%s,%d\n", spanName, nodeIndex, nodeName, edgeNodeIndex, edgeNodeName, spanView.Values[nodeIndex][edgeNodeIndex])
			}
		}
	}
}
