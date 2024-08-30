package command

import (
	"github.com/spf13/cobra"
)

func NewPNGCommand(renderPNGCommand *cobra.Command) *cobra.Command {
	command := &cobra.Command{
		Use:   "png",
		Short: "PNG",
		Long:  `PNG.`,
	}

	command.AddCommand(renderPNGCommand)

	return command
}
