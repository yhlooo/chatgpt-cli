package chat

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/keybrl/chatgpt-cli/pkg/chat"
	"github.com/keybrl/chatgpt-cli/pkg/openai"
)

var (
	flagServer          string
	flagSecretKey       string
	flagOrgID           string
	flagModel           string
	flagTimeoutPerRound time.Duration
	flagProxy           string
)

var Cmd = &cobra.Command{
	Use:   "chat",
	Short: "Start a chart",
	RunE: func(cmd *cobra.Command, _ []string) error {
		c, err := chat.NewChat(&openai.Config{
			Server:         flagServer,
			SecretKey:      flagSecretKey,
			OrganizationID: flagOrgID,
			Proxy:          flagProxy,
		}, chat.Options{
			Model:           flagModel,
			TimeoutPerRound: flagTimeoutPerRound,
		})
		if err != nil {
			return fmt.Errorf("create chat error: %w", err)
		}
		return c.Start(cmd.Context())
	},
}

func init() {
	Cmd.Flags().StringVar(&flagServer, "server", "", "OpenAI API Server (default: https://api.openai.com)")
	Cmd.Flags().StringVar(&flagSecretKey, "secret-key", "", "OpenAI API Secret")
	Cmd.Flags().StringVar(&flagOrgID, "org-id", "", "OpenAI Organization ID")
	Cmd.Flags().StringVar(&flagModel, "model", "", "GPT Model for chat (default: gpt-3.5-turbo)")
	Cmd.Flags().DurationVar(&flagTimeoutPerRound, "timeout-per-round", 0, "Timeout for chat per round")
	Cmd.Flags().StringVarP(&flagProxy, "proxy", "x", "", "HTTP proxy for request OpenAI API")
}
