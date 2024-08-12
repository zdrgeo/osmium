package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zdrgeo/osmium/pkg/analysis"
)

func NewChangeAnalysisCommand(handler *analysis.ChangeAnalysisHandler) *cobra.Command {
	command := &cobra.Command{
		Use:   "change",
		Short: "Change analysis",
		Long:  `Change analysis.`,
		Run: func(cmd *cobra.Command, args []string) {
			name, err := cmd.Flags().GetString("name")

			if err != nil {
				fmt.Printf("Error retrieving name: %s\n", err.Error())
			}

			source, err := cmd.Flags().GetString("source")

			if err != nil {
				fmt.Printf("Error retrieving source: %s\n", err.Error())
			}

			_ = source

			handler.ChangeAnalysis(name)
		},
	}

	command.Flags().String("source", "github:pullrequest", "Source of the analysis")
	command.MarkFlagRequired("source")
	viper.BindPFlag("source", command.Flags().Lookup("source"))
	viper.SetDefault("source", "github:pullrequest")

	return command
}
