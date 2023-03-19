package chat

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/keybrl/chatgpt-cli/pkg/openai"
)

const (
	defaultModel           = "gpt-3.5-turbo"
	defaultTimeoutPerRound = 30 * time.Second
)

// Options 对话配置
type Options struct {
	// 对话使用的模型
	Model string
	// 每回合超时时间
	TimeoutPerRound time.Duration
	// 是否使用彩色输出
	OutputColor bool
	// 输出是否含修饰符
	OutputDecoration bool
}

// Chat 对话
type Chat interface {
	// Start 开始对话，并阻塞直到对话结束
	Start(ctx context.Context) error
}

// NewChat 创建一个 Chat
func NewChat(config *openai.Config, opts Options) (Chat, error) {
	// 创建客户端
	client, err := openai.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("new openai client error: %w", err)
	}

	// 设置默认值
	if opts.Model == "" {
		opts.Model = defaultModel
	}
	if opts.TimeoutPerRound == 0 {
		opts.TimeoutPerRound = defaultTimeoutPerRound
	}

	return &consoleChat{
		opts:   opts,
		input:  bufio.NewReader(os.Stdin),
		client: client,
	}, nil
}

// consoleChat Chat 的一个命令行实现
type consoleChat struct {
	ctx    context.Context
	cancel context.CancelFunc

	opts Options

	client   openai.Client
	input    *bufio.Reader
	messages []openai.ChatMessage
}

var _ Chat = &consoleChat{}

// Start TODO: 开始对话，并阻塞直到对话结束
func (chat *consoleChat) Start(ctx context.Context) error {
	// 设置上下文
	chat.ctx, chat.cancel = context.WithCancel(ctx)
	defer chat.cancel()

	for {
		// 检查上下文
		select {
		case <-chat.ctx.Done():
			return nil
		default:
		}

		var msg *openai.ChatMessage
		var err error

		// 接收下游输入
		if msg, err = chat.recvFromDownstream(); err != nil {
			return fmt.Errorf("get input from console error: %w", err)
		}
		if msg == nil || msg.Content == "" {
			continue
		}

		// 转发给上游，并接收上游输出
		if msg, err = chat.forwardToUpstream(msg); err != nil {
			msg = &openai.ChatMessage{
				Role: openai.SystemChatMessageRole,
				Content: fmt.Sprintf("ERROR: %s", err.Error()),
			}
		}

		// 转发给下游
		if err = chat.sendToDownstream(msg); err != nil {
			return fmt.Errorf("send output to console error: %w", err)
		}
	}
}
