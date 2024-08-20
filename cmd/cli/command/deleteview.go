package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zdrgeo/osmium/pkg/view"
)

func NewDeleteViewCommand(handler *view.DeleteViewHandler) *cobra.Command {
	command := &cobra.Command{
		Use:   "delete",
		Short: "Delete view",
		Long:  `Delete view.`,
		Run: func(cmd *cobra.Command, args []string) {
			analysisName, err := cmd.Flags().GetString("analysis-name")

			if err != nil {
				fmt.Printf("Error retrieving analysis name: %s\n", err.Error())
			}

			viewName, err := cmd.Flags().GetString("view-name")

			if err != nil {
				fmt.Printf("Error retrieving view name: %s\n", err.Error())
			}

			handler.DeleteView(analysisName, viewName)
		},
	}

	return command
}
