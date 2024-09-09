package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewGitCommand(createGitCommand, changeGitCommand *cobra.Command) *cobra.Command {
	command := &cobra.Command{
		Use:   "git",
		Short: "Git",
		Long:  `Git.`,
	}

	command.Flags().StringP("change", "c", "commit", "Change of the analysis")
	// command.Flags().VarP(&change, "change", "c", "Change of the analysis")
	command.MarkFlagRequired("change")
	viper.BindPFlag("change", command.Flags().Lookup("change"))
	viper.SetDefault("change", "commit")

	command.Flags().StringToStringP("change-option", "o", map[string]string{}, "Options of the change")
	viper.BindPFlag("changeoptions", command.Flags().Lookup("change-option"))

	command.PersistentFlags().IntP("span-size", "s", 0, "Size of the span")
	viper.BindPFlag("spansize", command.PersistentFlags().Lookup("span-size"))

	command.AddCommand(createGitCommand)
	command.AddCommand(changeGitCommand)

	return command
}
