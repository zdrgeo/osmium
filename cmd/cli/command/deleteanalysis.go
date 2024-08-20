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
			analysisName, err := cmd.Flags().GetString("analysis-name")

			if err != nil {
				fmt.Printf("Error retrieving analysis name: %s\n", err.Error())
			}

			handler.DeleteAnalysis(analysisName)
		},
	}

	return command
}
