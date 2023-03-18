package login

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "login",
	Short: "Login to OpenAI",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
