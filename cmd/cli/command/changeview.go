package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zdrgeo/osmium/pkg/view"
)

func NewChangeViewCommand(handler *view.ChangeViewHandler) *cobra.Command {
	command := &cobra.Command{
		Use:   "change",
		Short: "Change view",
		Long:  `Change view.`,
		Run: func(cmd *cobra.Command, args []string) {
			analysisName, err := cmd.Flags().GetString("analysis-name")

			if err != nil {
				fmt.Printf("Error retrieving analysis name: %s\n", err.Error())
			}

			viewName, err := cmd.Flags().GetString("view-name")

			if err != nil {
				fmt.Printf("Error retrieving view name: %s\n", err.Error())
			}

			nodeNames, err := cmd.Flags().GetStringArray("node-name")

			if err != nil {
				fmt.Printf("Error retrieving node names: %s\n", err.Error())
			}

			builder, err := cmd.Flags().GetString("builder")

			if err != nil {
				fmt.Printf("Error retrieving builder: %s\n", err.Error())
			}

			builderOptions, err := cmd.Flags().GetStringToString("builder-option")

			if err != nil {
				fmt.Printf("Error retrieving builder options: %s\n", err.Error())
			}

			handler.ChangeView(analysisName, viewName, nodeNames, builder, builderOptions)
		},
	}

	command.Flags().StringArrayP("node-name", "n", []string{}, "Names of the nodes")
	viper.BindPFlag("nodenames", command.Flags().Lookup("node-name"))

	command.Flags().StringP("builder", "b", "", "Builder of the view")
	// command.Flags().VarP(&builder, "builder", "b", "Builder of the view")
	command.MarkFlagRequired("builder")
	viper.BindPFlag("builder", command.Flags().Lookup("builder"))
	viper.SetDefault("builder", "filepath")

	command.Flags().StringToStringP("builder-option", "o", map[string]string{}, "Options of the builder")
	viper.BindPFlag("builderoptions", command.Flags().Lookup("builder-option"))

	return command
}
