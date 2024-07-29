package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zdrgeo/osmium/internal/view"
)

func NewRenderViewCommand(handler *view.RenderViewHandler) *cobra.Command {
	command := &cobra.Command{
		Use:   "render",
		Short: "Render view",
		Long:  `Render view.`,
		Run: func(cmd *cobra.Command, args []string) {
			analysisName, err := cmd.Flags().GetString("analysis-name")

			if err != nil {
				fmt.Printf("Error retrieving analysis name: %s\n", err.Error())
			}

			name, err := cmd.Flags().GetString("name")

			if err != nil {
				fmt.Printf("Error retrieving name: %s\n", err.Error())
			}

			handler.RenderView(analysisName, name)
		},
	}

	command.Flags().StringArray("node-name", []string{}, "Names of the nodes")

	return command
}
