package command

import (
	"errors"
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

			_ = source

			handler.CreateAnalysis(name)
		},
	}

	command.Flags().String("source", "", "Source of the analysis")
	// command.Flags().Var(&source, "source", "Source of the analysis")
	command.MarkFlagRequired("source")
	viper.BindPFlag("source", command.Flags().Lookup("source"))
	viper.SetDefault("source", "github:pullrequest")

	return command
}

type source string

const (
	gitHub_pullRequest source = "github:pullrequest"
)

func (s *source) String() string {
	return string(*s)
}

func (s *source) Set(value string) error {
	switch value {
	case "github:pullrequest":
		*s = source(value)
		return nil
	default:
		return errors.New(`must be one of "github:pullrequest"`)
	}
}

func (s *source) Type() string {
	return "source"
}
