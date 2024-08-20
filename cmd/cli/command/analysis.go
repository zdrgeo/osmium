package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewAnalysisCommand(createAnalysisCommand, changeAnalysisCommand, deleteAnalysisCommand *cobra.Command) *cobra.Command {
	command := &cobra.Command{
		Use:   "analysis",
		Short: "Analysis",
		Long:  `Analysis.`,
	}

	command.PersistentFlags().String("analysis-name", "", "Name of the analysis")
	command.MarkPersistentFlagRequired("analysis-name")
	viper.BindPFlag("analysisname", command.PersistentFlags().Lookup("analysis-name"))

	command.AddCommand(createAnalysisCommand)
	command.AddCommand(changeAnalysisCommand)
	command.AddCommand(deleteAnalysisCommand)

	return command
}
