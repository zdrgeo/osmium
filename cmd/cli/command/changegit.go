package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zdrgeo/osmium/pkg/analysis"
)

func NewChangeGitCommand(handler *analysis.ChangeGitHandler) *cobra.Command {
	command := &cobra.Command{
		Use:   "change",
		Short: "Change Git analysis",
		Long:  `Change Git analysis.`,
		Run: func(cmd *cobra.Command, args []string) {
			analysisName, err := cmd.Flags().GetString("analysis-name")

			if err != nil {
				fmt.Printf("Error retrieving Git analysis name: %s\n", err.Error())
			}

			repositoryURL, err := cmd.Flags().GetString("repository-url")

			if err != nil {
				fmt.Printf("Error retrieving Git repository URL: %s\n", err.Error())
			}

			repositoryPath, err := cmd.Flags().GetString("repository-path")

			if err != nil {
				fmt.Printf("Error retrieving Git repository path: %s\n", err.Error())
			}

			handler.ChangeGit(analysisName, repositoryURL, repositoryPath)
		},
	}

	command.Flags().String("repository-url", "", "URL of the Git repository")
	command.MarkFlagRequired("repository-url")
	viper.BindPFlag("git_repositoryurl", command.Flags().Lookup("repository-url"))

	command.Flags().String("repository-path", "", "Path of the Git repository")
	command.MarkFlagRequired("repository-path")
	viper.BindPFlag("git_repositorypath", command.Flags().Lookup("repository-path"))

	return command
}
