package logout

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout from OpenAI",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
