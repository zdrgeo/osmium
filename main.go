/*
Copyright © 2024 Zdravko Georgiev <zdravko.georgiev@gmail.com>
*/
package main

import (
	"log"

	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/spf13/viper"
	"github.com/zdrgeo/osmium/cmd"
	"github.com/zdrgeo/osmium/internal/analysis"
	"github.com/zdrgeo/osmium/internal/repository"
	"github.com/zdrgeo/osmium/internal/source/github"
	"github.com/zdrgeo/osmium/internal/view"
)

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

	// client, err := api.DefaultGraphQLClient()
	options := api.ClientOptions{AuthToken: viper.GetString("GITHUB_TOKEN")}
	client, err := api.NewGraphQLClient(options)

	if err != nil {
		log.Fatal(err)
	}

	analysisSource := github.NewPullRequestAnalysisSource(client, "scaleforce", "tixets")

	analysisRepository := repository.NewFileAnalysisRepository("") // Empty means user home
	viewRepository := repository.NewFileViewRepository("")         // Empty means user home

	createAnalysisHandler := analysis.NewCreateAnalysisHandler(analysisSource, analysisRepository)
	changeAnalysisHandler := analysis.NewChangeAnalysisHandler(analysisSource, analysisRepository)
	deleteAnalysisHandler := analysis.NewDeleteAnalysisHandler(analysisRepository)

	createViewHandler := view.NewCreateViewHandler(analysisRepository, viewRepository)
	changeViewHandler := view.NewChangeViewHandler(analysisRepository, viewRepository)
	deleteViewHandler := view.NewDeleteViewHandler(viewRepository)
	renderViewHandler := view.NewRenderViewHandler(viewRepository)

	createAnalysisCommand := cmd.NewCreateAnalysisCommand(createAnalysisHandler)
	changeAnalysisCommand := cmd.NewChangeAnalysisCommand(changeAnalysisHandler)
	deleteAnalysisCommand := cmd.NewDeleteAnalysisCommand(deleteAnalysisHandler)

	analysisCommand := cmd.NewAnalysisCommand(createAnalysisCommand, changeAnalysisCommand, deleteAnalysisCommand)

	createViewCommand := cmd.NewCreateViewCommand(createViewHandler)
	changeViewCommand := cmd.NewChangeViewCommand(changeViewHandler)
	deleteViewCommand := cmd.NewDeleteViewCommand(deleteViewHandler)
	renderViewCommand := cmd.NewRenderViewCommand(renderViewHandler)

	viewCommand := cmd.NewViewCommand(createViewCommand, changeViewCommand, deleteViewCommand, renderViewCommand)

	osmiumCommand := cmd.NewOsmiumCommand(analysisCommand, viewCommand)

	if err := osmiumCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
