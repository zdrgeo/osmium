package view

import (
	"context"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/coder/websocket"
)

type ListenWebBrowserHandler struct {
	basePath string
}

func NewListenWebBrowserHandler() *ListenWebBrowserHandler {
	return &ListenWebBrowserHandler{}
}

func (handler *ListenWebBrowserHandler) ListenWebBrowser(analysisName, viewName, address string) {
	basePath := handler.basePath

	if basePath == "" {
		userHomePath, err := os.UserHomeDir()

		if err != nil {
			log.Panic(err)
		}

		basePath = userHomePath
	}

	viewPath := viewPath(basePath, analysisName, viewName)

	if err := os.MkdirAll(viewPath, 0750); err != nil {
		log.Panic(err)
	}

	fileHandler := http.FileServer(http.Dir(viewPath))

	// http.Handle("/analysis/view/webbrowser", http.StripPrefix("/analysis/view/webbrowser", fileHandler))
	http.Handle("/", fileHandler)

	fileName := filepath.Join(viewPath, "view.json")

	fileChangeHandler := newFileChangeHandler(fileName)

	http.Handle("/change", http.HandlerFunc(fileChangeHandler.Handle))

	log.Printf("Listening on %s\n", address)

	if err := http.ListenAndServe(address, nil); err != nil {
		log.Panic(err)
	}
}

type fileChangeHandler struct {
	fileName string
}

func newFileChangeHandler(fileName string) *fileChangeHandler {
	return &fileChangeHandler{fileName: fileName}
}

func (handler *fileChangeHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	connection, err := websocket.Accept(writer, request, nil)

	if err != nil {
		log.Panic(err)
	}

	defer connection.CloseNow()

	timeoutContext, cancel := context.WithTimeout(request.Context(), 10*time.Minute)

	defer cancel()

	timeoutContext = connection.CloseRead(timeoutContext)

	oldStat, err := os.Stat(handler.fileName)

	if err != nil {
		log.Panic(err)
	}

	ticker := time.NewTicker(5 * time.Second)

	defer ticker.Stop()

	for {
		select {
		case <-timeoutContext.Done():
			err = connection.Close(websocket.StatusNormalClosure, "")

			if err != nil && websocket.CloseStatus(err) != websocket.StatusNormalClosure && websocket.CloseStatus(err) != websocket.StatusGoingAway {
				log.Panic(err)
			}
			return
		case <-ticker.C:
			newStat, err := os.Stat(handler.fileName)

			if err != nil {
				log.Panic(err)
			}

			if oldStat.ModTime() != newStat.ModTime() {
				err := connection.Close(websocket.StatusNormalClosure, "changed")

				if err != nil && websocket.CloseStatus(err) != websocket.StatusNormalClosure && websocket.CloseStatus(err) != websocket.StatusGoingAway {
					log.Panic(err)
				}

				return
			}
		}
	}
}

/*
done := make(chan struct{})

go func() {
	log.Printf("Listening on %s\n", address)

	if err := http.ListenAndServe(address, nil); err != nil {
		log.Panic(err)
	}

	done <- struct{}{}
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
		log.Panic(err)
}

<-done
*/
