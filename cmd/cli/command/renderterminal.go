package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zdrgeo/osmium/pkg/view"
)

func NewRenderTerminalCommand(handler *view.RenderTerminalHandler) *cobra.Command {
	command := &cobra.Command{
		Use:   "render",
		Short: "Render terminal view",
		Long:  `Render terminal view.`,
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

			xNodeStart, err := cmd.Flags().GetInt("x-node-start")

			if err != nil {
				fmt.Printf("Error retrieving X node start: %s\n", err.Error())
			}

			yNodeStart, err := cmd.Flags().GetInt("y-node-start")

			if err != nil {
				fmt.Printf("Error retrieving Y node start: %s\n", err.Error())
			}

			nodeCount, err := cmd.Flags().GetInt("node-count")

			if err != nil {
				fmt.Printf("Error retrieving node count: %s\n", err.Error())
			}

			handler.RenderTerminal(analysisName, viewName, spanName, xNodeStart, yNodeStart, nodeCount)
		},
	}

	command.Flags().StringP("span-name", "s", "", "Name of the span")
	command.Flags().Int("x-node-start", 0, "Start of the X nodes")
	command.Flags().Int("y-node-start", 0, "Start of the Y nodes")
	command.Flags().Int("node-count", 100, "Count of the nodes")

	return command
}
