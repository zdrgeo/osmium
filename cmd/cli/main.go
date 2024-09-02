package main

import (
	"log"

	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/spf13/viper"
	"github.com/zdrgeo/osmium/cmd/cli/command"
	"github.com/zdrgeo/osmium/pkg/analysis"
	"github.com/zdrgeo/osmium/pkg/repository"
	"github.com/zdrgeo/osmium/pkg/source/git"
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
		log.Panic(err)
	}

	// client, err := api.DefaultGraphQLClient()
	options := api.ClientOptions{AuthToken: viper.GetString("GITHUB_TOKEN")}
	client, err := api.NewGraphQLClient(options)

	if err != nil {
		log.Panic(err)
	}

	gitAnalysisSource := git.NewCommitAnalysisSource()
	gitHubAnalysisSource := github.NewPullRequestAnalysisSource(client)

	analysisRepository := repository.NewFileAnalysisRepository(viper.GetString("BASEPATH")) // Empty means user home
	viewRepository := repository.NewFileViewRepository(viper.GetString("BASEPATH"))         // Empty means user home

	deleteAnalysisHandler := analysis.NewDeleteAnalysisHandler(analysisRepository)
	createGitHandler := analysis.NewCreateGitHandler(gitAnalysisSource, analysisRepository)
	changeGitHandler := analysis.NewChangeGitHandler(gitAnalysisSource, analysisRepository)
	createGitHubHandler := analysis.NewCreateGitHubHandler(gitHubAnalysisSource, analysisRepository)
	changeGitHubHandler := analysis.NewChangeGitHubHandler(gitHubAnalysisSource, analysisRepository)

	createViewHandler := view.NewCreateViewHandler(analysisRepository, viewRepository)
	changeViewHandler := view.NewChangeViewHandler(analysisRepository, viewRepository)
	deleteViewHandler := view.NewDeleteViewHandler(viewRepository)
	renderTerminalHandler := view.NewRenderTerminalHandler(viewRepository)
	renderWebBrowserHandler := view.NewRenderWebBrowserHandler(viewRepository)
	listenWebBrowserHandler := view.NewListenWebBrowserHandler()
	renderCSVHandler := view.NewRenderCSVHandler(viewRepository)
	renderPNGHandler := view.NewRenderPNGHandler(viewRepository)

	deleteAnalysisCommand := command.NewDeleteAnalysisCommand(deleteAnalysisHandler)

	createGitCommand := command.NewCreateGitCommand(createGitHandler)
	changeGitCommand := command.NewChangeGitCommand(changeGitHandler)

	gitCommand := command.NewGitCommand(createGitCommand, changeGitCommand)

	createGitHubCommand := command.NewCreateGitHubCommand(createGitHubHandler)
	changeGitHubCommand := command.NewChangeGitHubCommand(changeGitHubHandler)

	gitHubCommand := command.NewGitHubCommand(createGitHubCommand, changeGitHubCommand)

	analysisCommand := command.NewAnalysisCommand(deleteAnalysisCommand, gitCommand, gitHubCommand)

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

	renderPNGCommand := command.NewRenderPNGCommand(renderPNGHandler)

	pngCommand := command.NewPNGCommand(renderPNGCommand)

	viewCommand := command.NewViewCommand(createViewCommand, changeViewCommand, deleteViewCommand, terminalCommand, webBrowserCommand, csvCommand, pngCommand)

	osmiumCommand := command.NewOsmiumCommand(analysisCommand, viewCommand)

	if err := osmiumCommand.Execute(); err != nil {
		log.Panic(err)
	}
}
