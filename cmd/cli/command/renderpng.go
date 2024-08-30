package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zdrgeo/osmium/pkg/view"
)

func NewRenderPNGCommand(handler *view.RenderPNGHandler) *cobra.Command {
	command := &cobra.Command{
		Use:   "render",
		Short: "Render PNG view",
		Long:  `Render PNG view.`,
		Run: func(cmd *cobra.Command, args []string) {
			analysisName, err := cmd.Flags().GetString("analysis-name")

			if err != nil {
				fmt.Printf("Error retrieving analysis name: %s\n", err.Error())
			}

			viewName, err := cmd.Flags().GetString("view-name")

			if err != nil {
				fmt.Printf("Error retrieving view name: %s\n", err.Error())
			}

			spanName, err := cmd.Flags().GetString("span-name")

			if err != nil {
				fmt.Printf("Error retrieving span name: %s\n", err.Error())
			}

			handler.RenderPNG(analysisName, viewName, spanName)
		},
	}

	command.Flags().StringP("span-name", "s", "", "Name of the span")

	return command
}
