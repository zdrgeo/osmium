package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zdrgeo/osmium/pkg/analysis"
)

func NewCreateAnalysisCommand(handler *analysis.CreateAnalysisHandler) *cobra.Command {
	command := &cobra.Command{
		Use:   "create",
		Short: "Create analysis",
		Long:  `Create analysis.`,
		Run: func(cmd *cobra.Command, args []string) {
			analysisName, err := cmd.Flags().GetString("analysis-name")

			if err != nil {
				fmt.Printf("Error retrieving analysis name: %s\n", err.Error())
			}

			source, err := cmd.Flags().GetString("source")

			if err != nil {
				fmt.Printf("Error retrieving source: %s\n", err.Error())
			}

			sourceOptions, err := cmd.Flags().GetStringToString("source-option")

			if err != nil {
				fmt.Printf("Error retrieving source options: %s\n", err.Error())
			}

			handler.CreateAnalysis(analysisName, source, sourceOptions)
		},
	}

	command.Flags().StringP("source", "s", "", "Source of the analysis")
	// command.Flags().VarP(&source, "source", "s", "Source of the analysis")
	command.MarkFlagRequired("source")
	viper.BindPFlag("source", command.Flags().Lookup("source"))
	viper.SetDefault("source", "github:pullrequest")

	command.Flags().StringToStringP("source-option", "o", map[string]string{}, "Options of the source")
	viper.BindPFlag("sourceoptions", command.Flags().Lookup("source-option"))

	return command
}
