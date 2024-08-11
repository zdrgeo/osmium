package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewViewCommand(createViewCommand, changeViewCommand, deleteViewCommand, renderViewCommand, listenViewCommand *cobra.Command) *cobra.Command {
	command := &cobra.Command{
		Use:   "view",
		Short: "View",
		Long:  `View.`,
	}

	command.PersistentFlags().String("analysis-name", "", "Name of the analysis")
	command.MarkPersistentFlagRequired("analysis-name")
	viper.BindPFlag("analysisname", command.PersistentFlags().Lookup("analysis-name"))
	command.PersistentFlags().String("name", "", "Name of the view")
	command.MarkPersistentFlagRequired("name")
	viper.BindPFlag("viewname", command.PersistentFlags().Lookup("name"))

	command.AddCommand(createViewCommand)
	command.AddCommand(changeViewCommand)
	command.AddCommand(deleteViewCommand)
	command.AddCommand(renderViewCommand)
	command.AddCommand(listenViewCommand)

	return command
}
