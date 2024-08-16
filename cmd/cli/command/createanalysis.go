package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zdrgeo/osmium/pkg/analysis"
)

func NewCreateAnalysisCommand(handler *analysis.CreateAnalysisHandler) *cobra.Command {
	// var source source

	command := &cobra.Command{
		Use:   "create",
		Short: "Create analysis",
		Long:  `Create analysis.`,
		Run: func(cmd *cobra.Command, args []string) {
			name, err := cmd.Flags().GetString("name")

			if err != nil {
				fmt.Printf("Error retrieving name: %s\n", err.Error())
			}

			source, err := cmd.Flags().GetString("source")

			if err != nil {
				fmt.Printf("Error retrieving source: %s\n", err.Error())
			}

			sourceOptions, err := cmd.Flags().GetStringToString("source-option")

			if err != nil {
				fmt.Printf("Error retrieving source options: %s\n", err.Error())
			}

			handler.CreateAnalysis(name, source, sourceOptions)
		},
	}

	command.Flags().String("source", "", "Source of the analysis")
	// command.Flags().Var(&source, "source", "Source of the analysis")
	command.MarkFlagRequired("source")
	viper.BindPFlag("source", command.Flags().Lookup("source"))
	viper.SetDefault("source", "github:pullrequest")

	command.Flags().StringToString("source-option", map[string]string{}, "Options of the source")
	viper.BindPFlag("sourceoptions", command.Flags().Lookup("source-option"))

	return command
}
