package chat

import (
	"bufio"
	"context"
	"fmt"
	"os"
)

// Options 对话配置
type Options struct{}

// Chat 对话
type Chat interface {
	// Start 开始对话，并阻塞直到对话结束
	Start(ctx context.Context) error
}

// NewChat 创建一个 Chat
func NewChat(opts Options) (Chat, error) {
	return &consoleChat{
		opts: opts,
	}, nil
}

// consoleChat Chat 的一个命令行实现
type consoleChat struct {
	ctx    context.Context
	cancel context.CancelFunc

	opts Options

	input *bufio.Reader
}

var _ Chat = &consoleChat{}

// Start TODO: 开始对话，并阻塞直到对话结束
func (chat *consoleChat) Start(ctx context.Context) error {
	// 设置上下文
	chat.ctx, chat.cancel = context.WithCancel(ctx)
	defer chat.cancel()

	chat.input = bufio.NewReader(os.Stdin)

	for {
		// 检查上下文
		select {
		case <-chat.ctx.Done():
			return nil
		default:
		}

		var msg *Message
		var err error

		// 接收下游输入
		if msg, err = chat.recvFromDownstream(); err != nil {
			return fmt.Errorf("get input from console error: %w", err)
		}
		if msg.String() == "" {
			continue
		}

		// 转发给上游，并接收上游输出
		if msg, err = chat.forwardToUpstream(msg); err != nil {
			return fmt.Errorf("forward message to upstream error: %w", err)
		}

		// 转发给下游
		if err = chat.sendToDownstream(msg); err != nil {
			return fmt.Errorf("send output to console error: %w", err)
		}
	}
}
