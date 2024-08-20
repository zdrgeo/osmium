package command

import (
	"github.com/spf13/cobra"
)

func NewWebBrowserCommand(renderWebBrowserCommand, listenWebBrowserCommand *cobra.Command) *cobra.Command {
	command := &cobra.Command{
		Use:   "web-browser",
		Short: "Web browser",
		Long:  `Web browser.`,
	}

	command.AddCommand(renderWebBrowserCommand)
	command.AddCommand(listenWebBrowserCommand)

	return command
}
