package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zdrgeo/osmium/pkg/view"
)

func NewCreateViewCommand(handler *view.CreateViewHandler) *cobra.Command {
	command := &cobra.Command{
		Use:   "create",
		Short: "Create view",
		Long:  `Create view.`,
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

			handler.CreateView(analysisName, viewName, nodeNames)
		},
	}

	command.Flags().StringArrayP("node-name", "nn", []string{}, "Names of the nodes")
	viper.BindPFlag("nodenames", command.Flags().Lookup("node-name"))

	return command
}
