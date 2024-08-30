package command

import (
	"github.com/spf13/cobra"
)

func NewGitCommand(createGitCommand, changeGitCommand *cobra.Command) *cobra.Command {
	command := &cobra.Command{
		Use:   "git",
		Short: "Git",
		Long:  `Git.`,
	}

	command.AddCommand(createGitCommand)
	command.AddCommand(changeGitCommand)

	return command
}
