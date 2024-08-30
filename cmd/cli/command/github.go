package command

import (
	"github.com/spf13/cobra"
)

func NewGitHubCommand(createGitHubCommand, changeGitHubCommand *cobra.Command) *cobra.Command {
	command := &cobra.Command{
		Use:   "github",
		Short: "GitHub",
		Long:  `GitHub.`,
	}

	command.AddCommand(createGitHubCommand)
	command.AddCommand(changeGitHubCommand)

	return command
}
