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

			name, err := cmd.Flags().GetString("name")

			if err != nil {
				fmt.Printf("Error retrieving name: %s\n", err.Error())
			}

			nodeNames, err := cmd.Flags().GetStringArray("node-name")

			if err != nil {
				fmt.Printf("Error retrieving node names: %s\n", err.Error())
			}

			handler.ChangeView(analysisName, name, nodeNames)
		},
	}

	command.Flags().StringArray("node-name", []string{}, "Names of the nodes")
	viper.BindPFlag("nodenames", command.Flags().Lookup("node-name"))

	return command
}
