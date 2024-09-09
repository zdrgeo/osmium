package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zdrgeo/osmium/pkg/view"
)

func NewRenderWebBrowserCommand(handler *view.RenderWebBrowserHandler) *cobra.Command {
	command := &cobra.Command{
		Use:   "render",
		Short: "Render web browser view",
		Long:  `Render web browser view.`,
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

			handler.RenderWebBrowser(analysisName, viewName, spanName)
		},
	}

	command.Flags().StringP("span-name", "s", "", "Name of the span")
	viper.BindPFlag("spanname", command.PersistentFlags().Lookup("span-name"))

	return command
}
