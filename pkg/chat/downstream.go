package chat

import (
    "context"
    "errors"
	"fmt"
	"io"
	"strings"
    "time"

    "github.com/keybrl/chatgpt-cli/pkg/openai"
)

// recvFromDownstream 从下游接收消息
func (chat *consoleChat) recvFromDownstream() (*openai.ChatMessage, error) {
	var content string

	chat.printDecoration("You: \n> ")
	for {
		line, err := chat.input.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, fmt.Errorf("read from stdin error: %w", err)
		}

		if line == "\n" {
			break
		}
		content += line
		if content[0] == '.' {
			break
		}

		chat.printDecoration("> ")
	}

	content = strings.TrimRight(content, "\n")

	return &openai.ChatMessage{
		Role:    openai.UserChatMessageRole,
		Content: content,
	}, nil
}

// sendToDownstream 发送消息到下游
func (chat *consoleChat) sendToDownstream(msg *openai.ChatMessage) error {
	if msg == nil {
		return nil
	}
	content := msg.Content
	content = strings.Trim(content, " \n")
	content += "\n"

	var prefix, suffix string
	if msg.Role == openai.SystemChatMessageRole {
		prefix, suffix = chat.setTextEffect(redEffect)
	} else {
		prefix, suffix = chat.setTextEffect(greenEffect)
		chat.printDecoration("ChatGPT:\n")
	}
	fmt.Println(prefix+content+suffix)
	return nil
}

// printDecoration 打印装饰符
func (chat *consoleChat) printDecoration(str string) {
	if !chat.opts.OutputDecoration {
		return
	}
	prefix, suffix := chat.setTextEffect(weakEffect)
	fmt.Print(prefix+str+suffix)
}

// printDecorationAfter 一段时间后打印装饰符
func (chat *consoleChat) printDecorationAfter(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-ctx.Done():
		return
	case <-time.After(d):
		chat.printDecoration(msg)
	}
}

const (
	weakEffect = "weak"
	greenEffect = "green"
	redEffect = "red"
)

func (chat *consoleChat) setTextEffect(mode string) (prefix, suffix string) {
	if !chat.opts.OutputColor {
		return "", ""
	}
    switch mode {
	case weakEffect:
		return "\033[90m", "\033[0m"
	case greenEffect:
		return "\033[32m", "\033[0m"
	case redEffect:
		return "\033[31m", "\033[0m"
    }
	return "", ""
}
