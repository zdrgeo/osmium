package main

import (
	"log"

	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/spf13/viper"
	"github.com/zdrgeo/osmium/cmd/cli/command"
	"github.com/zdrgeo/osmium/pkg/analysis"
	"github.com/zdrgeo/osmium/pkg/repository"
	"github.com/zdrgeo/osmium/pkg/source/github"
	"github.com/zdrgeo/osmium/pkg/view"
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

	analysisSource := github.NewPullRequestAnalysisSource(client, viper.GetString("GITHUB_REPOSITORY_OWNER"), viper.GetString("GITHUB_REPOSITORY_NAME"))

	analysisRepository := repository.NewFileAnalysisRepository("") // Empty means user home
	viewRepository := repository.NewFileViewRepository("")         // Empty means user home

	createAnalysisHandler := analysis.NewCreateAnalysisHandler(analysisSource, analysisRepository)
	changeAnalysisHandler := analysis.NewChangeAnalysisHandler(analysisSource, analysisRepository)
	deleteAnalysisHandler := analysis.NewDeleteAnalysisHandler(analysisRepository)

	createViewHandler := view.NewCreateViewHandler(analysisRepository, viewRepository)
	changeViewHandler := view.NewChangeViewHandler(analysisRepository, viewRepository)
	deleteViewHandler := view.NewDeleteViewHandler(viewRepository)
	renderViewHandler := view.NewRenderViewHandler(viewRepository)
	listenViewHandler := view.NewListenViewHandler()

	createAnalysisCommand := command.NewCreateAnalysisCommand(createAnalysisHandler)
	changeAnalysisCommand := command.NewChangeAnalysisCommand(changeAnalysisHandler)
	deleteAnalysisCommand := command.NewDeleteAnalysisCommand(deleteAnalysisHandler)

	analysisCommand := command.NewAnalysisCommand(createAnalysisCommand, changeAnalysisCommand, deleteAnalysisCommand)

	createViewCommand := command.NewCreateViewCommand(createViewHandler)
	changeViewCommand := command.NewChangeViewCommand(changeViewHandler)
	deleteViewCommand := command.NewDeleteViewCommand(deleteViewHandler)
	renderViewCommand := command.NewRenderViewCommand(renderViewHandler)
	listenViewCommand := command.NewListenViewCommand(listenViewHandler)

	viewCommand := command.NewViewCommand(createViewCommand, changeViewCommand, deleteViewCommand, renderViewCommand, listenViewCommand)

	osmiumCommand := command.NewOsmiumCommand(analysisCommand, viewCommand)

	if err := osmiumCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
