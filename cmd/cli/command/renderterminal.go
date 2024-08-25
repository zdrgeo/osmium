package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zdrgeo/osmium/pkg/view"
)

func NewRenderTerminalCommand(handler *view.RenderTerminalHandler) *cobra.Command {
	command := &cobra.Command{
		Use:   "render",
		Short: "Render terminal",
		Long:  `Render terminal.`,
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

			nodeStart, err := cmd.Flags().GetInt("node-start")

			if err != nil {
				fmt.Printf("Error retrieving nodes start: %s\n", err.Error())
			}

			edgeNodeStart, err := cmd.Flags().GetInt("edge-node-start")

			if err != nil {
				fmt.Printf("Error retrieving edge nodes start: %s\n", err.Error())
			}

			nodeCount, err := cmd.Flags().GetInt("node-count")

			if err != nil {
				fmt.Printf("Error retrieving nodes count: %s\n", err.Error())
			}

			handler.RenderTerminal(analysisName, viewName, spanName, nodeStart, edgeNodeStart, nodeCount)
		},
	}

	command.Flags().StringP("span-name", "s", "", "Name of the span")
	command.Flags().Int("node-start", 0, "Start of the nodes")
	command.Flags().Int("edge-node-start", 0, "Start of the edge nodes")
	command.Flags().Int("node-count", 100, "Count of the nodes")

	return command
}
