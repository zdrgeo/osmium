package command

import (
	"github.com/spf13/cobra"
)

func NewOsmiumCommand(analysisCommand, viewCommand *cobra.Command) *cobra.Command {
	osmiumCommand := &cobra.Command{
		Use:   "osmium",
		Short: "Osmium",
		Long:  `Osmium.`,
	}

	osmiumCommand.AddCommand(analysisCommand)
	osmiumCommand.AddCommand(viewCommand)

	return osmiumCommand
}
