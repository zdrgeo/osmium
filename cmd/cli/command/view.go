package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewViewCommand(createViewCommand, changeViewCommand, deleteViewCommand, terminalCommand, webBrowserCommand, csvCommand, pngCommand *cobra.Command) *cobra.Command {
	command := &cobra.Command{
		Use:   "view",
		Short: "View",
		Long:  `View.`,
	}

	command.PersistentFlags().StringP("analysis-name", "a", "", "Name of the analysis")
	command.MarkPersistentFlagRequired("analysis-name")
	viper.BindPFlag("analysisname", command.PersistentFlags().Lookup("analysis-name"))
	command.PersistentFlags().StringP("view-name", "v", "", "Name of the view")
	command.MarkPersistentFlagRequired("view-name")
	viper.BindPFlag("viewname", command.PersistentFlags().Lookup("view-name"))

	command.AddCommand(createViewCommand)
	command.AddCommand(changeViewCommand)
	command.AddCommand(deleteViewCommand)
	command.AddCommand(terminalCommand)
	command.AddCommand(webBrowserCommand)
	command.AddCommand(csvCommand)
	command.AddCommand(pngCommand)

	return command
}
