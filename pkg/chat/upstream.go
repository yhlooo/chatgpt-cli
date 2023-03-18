package chat

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/keybrl/chatgpt-cli/pkg/openai"
	"github.com/sirupsen/logrus"
)

// forwardToUpstream 转发消息到上游
func (chat *consoleChat) forwardToUpstream(msg *openai.ChatMessage) (*openai.ChatMessage, error) {
	if msg == nil {
		return nil, nil
	}
	// 准备输入
	msgs := append(chat.messages, *msg)
	input := &openai.CreateChatCompletionInput{
		Model:    chat.opts.Model,
		Messages: msgs,
		// TODO: 探索更多参数的可能
	}
	logrus.Debugf("chat input: %s", string(mustJSONMarshal(input)))

	// 设置上下文超时
	ctx, cancel := context.WithTimeout(chat.ctx, chat.opts.TimeoutPerRound)
	defer cancel()

	// 请求 OpenAI
	output, err := chat.client.CreateChatCompletion(ctx, input)
	logrus.Debugf("chat output: %s", string(mustJSONMarshal(output)))
	if err != nil {
		return nil, fmt.Errorf("create chat to openai error: %w", err)
	}

	// 获取输出
	if len(output.Choices) == 0 {
		return &openai.ChatMessage{
			Content: "(it said nothing ...)",
		}, nil
	}
	respMsg := output.Choices[0].Message
	chat.messages = append(msgs, respMsg)

	return &respMsg, nil
}

// mustJSONMarshal 序列化 obj 为 JSON ，忽略错误
func mustJSONMarshal(obj interface{}) []byte {
	data, _ := json.Marshal(obj)
	return data
}
