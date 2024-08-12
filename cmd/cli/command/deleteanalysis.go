package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zdrgeo/osmium/pkg/analysis"
)

func NewDeleteAnalysisCommand(handler *analysis.DeleteAnalysisHandler) *cobra.Command {
	command := &cobra.Command{
		Use:   "delete",
		Short: "Delete analysis",
		Long:  `Delete analysis.`,
		Run: func(cmd *cobra.Command, args []string) {
			name, err := cmd.Flags().GetString("name")

			if err != nil {
				fmt.Printf("Error retrieving name: %s\n", err.Error())
			}

			handler.DeleteAnalysis(name)
		},
	}

	return command
}
