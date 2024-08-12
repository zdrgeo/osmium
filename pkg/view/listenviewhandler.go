package view

import (
	"log"
	"net/http"
	"os"
)

type ListenViewHandler struct {
	basePath string
}

func NewListenViewHandler() *ListenViewHandler {
	return &ListenViewHandler{}
}

func (handler *ListenViewHandler) ListenView(analysisName, name, address string) {
	basePath := handler.basePath

	if basePath == "" {
		userHomePath, err := os.UserHomeDir()

		if err != nil {
			log.Fatal(err)
		}

		basePath = userHomePath
	}

	viewPath := viewPath(basePath, analysisName, name)

	if err := os.MkdirAll(viewPath, 0750); err != nil {
		log.Fatal(err)
	}

	fileServer := http.FileServer(http.Dir(viewPath))

	// http.Handle("analysis/view/", http.StripPrefix("/analysis/view/", fileServer))
	http.Handle("/", fileServer)

	log.Printf("Listening on %s", address)

	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal(err)
	}
}

/*
exit := make(chan struct{})

go func() {
	log.Printf("Listening on %s", address)

	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal(err)
	}

	exit <- struct{}{}
}()

var err error

switch runtime.GOOS {
case "windows":
	err = exec.Command("rundll32", "url.dll,FileProtocolHandler", "http://localhost:3000").Start()
case "darwin":
	err = exec.Command("open", "http://localhost:3000").Start()
case "linux":
	err = exec.Command("xdg-open", "http://localhost:3000").Start()
}

if err != nil {
		log.Fatal(err)
}

<-exit
*/
