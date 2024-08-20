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

	analysisSource := github.NewPullRequestAnalysisSource(client)

	analysisRepository := repository.NewFileAnalysisRepository(viper.GetString("BASEPATH")) // Empty means user home
	viewRepository := repository.NewFileViewRepository(viper.GetString("BASEPATH"))         // Empty means user home

	createAnalysisHandler := analysis.NewCreateAnalysisHandler(analysisSource, analysisRepository)
	changeAnalysisHandler := analysis.NewChangeAnalysisHandler(analysisSource, analysisRepository)
	deleteAnalysisHandler := analysis.NewDeleteAnalysisHandler(analysisRepository)

	createViewHandler := view.NewCreateViewHandler(analysisRepository, viewRepository)
	changeViewHandler := view.NewChangeViewHandler(analysisRepository, viewRepository)
	deleteViewHandler := view.NewDeleteViewHandler(viewRepository)
	renderTerminalHandler := view.NewRenderTerminalHandler(viewRepository)
	renderWebBrowserHandler := view.NewRenderWebBrowserHandler(viewRepository)
	listenWebBrowserHandler := view.NewListenWebBrowserHandler()
	renderCSVHandler := view.NewRenderCSVHandler(viewRepository)

	createAnalysisCommand := command.NewCreateAnalysisCommand(createAnalysisHandler)
	changeAnalysisCommand := command.NewChangeAnalysisCommand(changeAnalysisHandler)
	deleteAnalysisCommand := command.NewDeleteAnalysisCommand(deleteAnalysisHandler)

	analysisCommand := command.NewAnalysisCommand(createAnalysisCommand, changeAnalysisCommand, deleteAnalysisCommand)

	createViewCommand := command.NewCreateViewCommand(createViewHandler)
	changeViewCommand := command.NewChangeViewCommand(changeViewHandler)
	deleteViewCommand := command.NewDeleteViewCommand(deleteViewHandler)

	renderTerminalCommand := command.NewRenderTerminalCommand(renderTerminalHandler)

	terminalCommand := command.NewTerminalCommand(renderTerminalCommand)

	renderWebBrowserCommand := command.NewRenderWebBrowserCommand(renderWebBrowserHandler)
	listenWebBrowserCommand := command.NewListenWebBrowserCommand(listenWebBrowserHandler)

	webBrowserCommand := command.NewWebBrowserCommand(renderWebBrowserCommand, listenWebBrowserCommand)

	renderCSVCommand := command.NewRenderCSVCommand(renderCSVHandler)

	csvCommand := command.NewCSVCommand(renderCSVCommand)

	viewCommand := command.NewViewCommand(createViewCommand, changeViewCommand, deleteViewCommand, terminalCommand, webBrowserCommand, csvCommand)

	osmiumCommand := command.NewOsmiumCommand(analysisCommand, viewCommand)

	if err := osmiumCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
