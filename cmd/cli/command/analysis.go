package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewAnalysisCommand(deleteAnalysisCommand, gitCommand, gitHubCommand *cobra.Command) *cobra.Command {
	command := &cobra.Command{
		Use:   "analysis",
		Short: "Analysis",
		Long:  `Analysis.`,
	}

	command.PersistentFlags().StringP("analysis-name", "a", "", "Name of the analysis")
	command.MarkPersistentFlagRequired("analysis-name")
	viper.BindPFlag("analysisname", command.PersistentFlags().Lookup("analysis-name"))

	command.AddCommand(deleteAnalysisCommand)
	command.AddCommand(gitCommand)
	command.AddCommand(gitHubCommand)

	return command
}
