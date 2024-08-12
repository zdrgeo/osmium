package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zdrgeo/osmium/pkg/view"
)

func NewListenViewCommand(handler *view.ListenViewHandler) *cobra.Command {
	command := &cobra.Command{
		Use:   "listen",
		Short: "Listen view",
		Long:  `Listen view.`,
		Run: func(cmd *cobra.Command, args []string) {
			analysisName, err := cmd.Flags().GetString("analysis-name")

			if err != nil {
				fmt.Printf("Error retrieving analysis name: %s\n", err.Error())
			}

			name, err := cmd.Flags().GetString("name")

			if err != nil {
				fmt.Printf("Error retrieving name: %s\n", err.Error())
			}

			address, err := cmd.Flags().GetString("address")

			if err != nil {
				fmt.Printf("Error retrieving address: %s\n", err.Error())
			}

			handler.ListenView(analysisName, name, address)
		},
	}

	command.Flags().String("address", ":3000", "Address")

	return command
}
