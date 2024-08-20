package command

import (
	"github.com/spf13/cobra"
)

func NewCSVCommand(renderCSVCommand *cobra.Command) *cobra.Command {
	command := &cobra.Command{
		Use:   "csv",
		Short: "CSV",
		Long:  `CSV.`,
	}

	command.AddCommand(renderCSVCommand)

	return command
}
