package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// https://www.alexedwards.net/blog/serving-static-sites-with-go
func main() {
	viper.AddConfigPath(".")
	// viper.SetConfigFile(".env")
	// viper.SetConfigName("config")
	// viper.SetConfigType("env") // "env", "json", "yaml"
	viper.SetEnvPrefix("osmium")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	basePath := ""
	analysisName := "ticketing_tixets"
	name := "app"

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

	// http.Handle("/view/", http.StripPrefix("/view/", fileServer))
	http.Handle("/", fileServer)

	log.Print("Listening on :3000...")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}

func viewPath(basePath, analysisName, name string) string {
	return filepath.Join(basePath, "osmium", "analysis", analysisName, "view", name)
}
