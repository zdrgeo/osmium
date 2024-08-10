package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewAnalysisCommand(createAnalysisCommand, changeAnalysisCommand, deleteAnalysisCommand *cobra.Command) *cobra.Command {
	command := &cobra.Command{
		Use:   "analysis",
		Short: "Analysis",
		Long:  `Analysis.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("analysis called")
		},
	}

	command.PersistentFlags().String("name", "", "Name of the analysis")
	command.MarkPersistentFlagRequired("name")
	viper.BindPFlag("analysisname", command.PersistentFlags().Lookup("name"))

	command.AddCommand(createAnalysisCommand)
	command.AddCommand(changeAnalysisCommand)
	command.AddCommand(deleteAnalysisCommand)

	return command
}
