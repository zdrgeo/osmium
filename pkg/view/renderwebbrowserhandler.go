package view

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
)

type RenderWebBrowserHandler struct {
	repository ViewRepository
	basePath   string
}

func NewRenderWebBrowserHandler(repository ViewRepository) *RenderWebBrowserHandler {
	return &RenderWebBrowserHandler{repository: repository}
}

func (handler *RenderWebBrowserHandler) RenderWebBrowser(analysisName, viewName, spanName string) {
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

	fileName := filepath.Join(viewPath, "view.html")

	renderViewToHTMLFile(view, spanName, "pkg/view/template/view.template", fileName)
}

func renderViewToHTMLFile(view *AnalysisView, spanName, templateFileName, fileName string) {
	file, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	funcMap := template.FuncMap{
		"inc": func(i int) int { return i + 1 },
		"dec": func(i int) int { return i - 1 },
	}

	_ = funcMap

	// template, err := template.New(filepath.Base(templateFileName)).Funcs(funcMap).ParseFiles(templateFileName)
	template, err := template.New(filepath.Base(templateFileName)).ParseFiles(templateFileName)

	if err != nil {
		log.Fatal(err)
	}

	err = template.Execute(file, view)

	if err != nil {
		log.Fatal(err)
	}
}
