package command

import (
	"github.com/spf13/cobra"
)

func NewTerminalCommand(renderTerminalCommand *cobra.Command) *cobra.Command {
	command := &cobra.Command{
		Use:   "terminal",
		Short: "Terminal",
		Long:  `Terminal.`,
	}

	command.AddCommand(renderTerminalCommand)

	return command
}
