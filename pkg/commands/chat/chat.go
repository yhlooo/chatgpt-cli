package chat

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "chat",
	Short: "Start a chart",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
