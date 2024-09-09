package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zdrgeo/osmium/pkg/analysis"
)

func NewCreateGitHubCommand(handler *analysis.CreateGitHubHandler) *cobra.Command {
	command := &cobra.Command{
		Use:   "create",
		Short: "Create GitHub analysis",
		Long:  `Create GitHub analysis.`,
		Run: func(cmd *cobra.Command, args []string) {
			analysisName, err := cmd.Flags().GetString("analysis-name")

			if err != nil {
				fmt.Printf("Error retrieving analysis name: %s\n", err.Error())
			}

			spanSize, err := cmd.Flags().GetInt("span-size")

			if err != nil {
				fmt.Printf("Error retrieving span size: %s\n", err.Error())
			}

			repositoryOwner, err := cmd.Flags().GetString("repository-owner")

			if err != nil {
				fmt.Printf("Error retrieving GitHub repository owner: %s\n", err.Error())
			}

			repositoryName, err := cmd.Flags().GetString("repository-name")

			if err != nil {
				fmt.Printf("Error retrieving GitHub repository name: %s\n", err.Error())
			}

			handler.CreateGitHub(analysisName, spanSize, repositoryOwner, repositoryName)
		},
	}

	command.Flags().String("repository-owner", "", "Owner of of the GitHub repository")
	command.MarkFlagRequired("repository-owner")
	viper.BindPFlag("github_repositoryowner", command.Flags().Lookup("repository-owner"))

	command.Flags().String("repository-name", "", "Name of of the GitHub repository")
	command.MarkFlagRequired("repository-name")
	viper.BindPFlag("github_repositoryname", command.Flags().Lookup("repository-name"))

	return command
}
