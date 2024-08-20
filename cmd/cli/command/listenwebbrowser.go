package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zdrgeo/osmium/pkg/view"
)

func NewListenWebBrowserCommand(handler *view.ListenWebBrowserHandler) *cobra.Command {
	command := &cobra.Command{
		Use:   "listen",
		Short: "Listen web browser",
		Long:  `Listen web browser.`,
		Run: func(cmd *cobra.Command, args []string) {
			analysisName, err := cmd.Flags().GetString("analysis-name")

			if err != nil {
				fmt.Printf("Error retrieving analysis name: %s\n", err.Error())
			}

			viewName, err := cmd.Flags().GetString("view-name")

			if err != nil {
				fmt.Printf("Error retrieving view name: %s\n", err.Error())
			}

			address, err := cmd.Flags().GetString("address")

			if err != nil {
				fmt.Printf("Error retrieving address: %s\n", err.Error())
			}

			handler.ListenWebBrowser(analysisName, viewName, address)
		},
	}

	command.Flags().String("address", ":3000", "Address")

	return command
}
