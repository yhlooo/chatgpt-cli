package commands

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/keybrl/chatgpt-cli/pkg/commands/chat"
	"github.com/keybrl/chatgpt-cli/pkg/commands/login"
	"github.com/keybrl/chatgpt-cli/pkg/commands/logout"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	flagDebug bool
)

// rootCmd 根命令
var rootCmd = &cobra.Command{
	Use:   "chartgpt-cli",
	Short: "CLI for OpenAI ChatGPT",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if flagDebug {
			logrus.SetLevel(logrus.DebugLevel)
			logrus.Debug("run in debug mode")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, _ []string) error {
		return cmd.Help()
	},
}

// 执行命令
func Execute() {
	ctx := notifyContext(context.Background())

	startTime := time.Now()
	err := rootCmd.ExecuteContext(ctx)
	logrus.Debugf("duration: %s", time.Since(startTime))
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&flagDebug, "debug", false, "run in debug mode")
	rootCmd.AddCommand(
		chat.Cmd,
		login.Cmd,
		logout.Cmd,
	)
}

// notifyContext 返回一个上下文，该上下文会在进程接收到 INT 或 TERM 信号时被取消
func notifyContext(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer cancel()
		defer close(ch)
		// 监听第一次事件
		select {
		case <-ctx.Done():
			return
		case <-ch:
			// 通知上下文取消
			cancel()
		}
		// 接收到第二次事件时直接强行退出
		<-ch
		os.Exit(1)
	}()
	return ctx
}
