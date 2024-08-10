package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewOsmiumCommand(analysisCommand, viewCommand *cobra.Command) *cobra.Command {
	osmiumCommand := &cobra.Command{
		Use:   "osmium",
		Short: "A brief description of your application",
		Long:  `A longer description that spans multiple lines.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("osmium called")
		},
	}

	osmiumCommand.AddCommand(analysisCommand)
	osmiumCommand.AddCommand(viewCommand)

	return osmiumCommand
}
