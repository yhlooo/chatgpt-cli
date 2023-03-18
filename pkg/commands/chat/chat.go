package chat

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/keybrl/chatgpt-cli/pkg/chat"
)

var Cmd = &cobra.Command{
	Use:   "chat",
	Short: "Start a chart",
	RunE: func(cmd *cobra.Command, _ []string) error {
		c, err := chat.NewChat(chat.Options{})
		if err != nil {
			return fmt.Errorf("create chat error: %w", err)
		}
		return c.Start(cmd.Context())
	},
}
