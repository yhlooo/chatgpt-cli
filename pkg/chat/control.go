package chat

import (
    "github.com/keybrl/chatgpt-cli/pkg/openai"
    "strings"
)

// handleControlMessage 处理控制消息
func (chat *consoleChat) handleControlMessage(msg *openai.ChatMessage) (bool, error) {
	if msg == nil {
		return false, nil
	}
	if !strings.HasPrefix(msg.Content, ".") {
		return false, nil
	}
	switch msg.Content {
	case ".quit", ".exit":
		chat.cancel()
		chat.printDecoration("(exit.)\n")
	case ".new":
		chat.messages = nil
		chat.printDecoration("\n(new chat.)\n")
	case ".help":
		fallthrough
	default:
		chat.printDecoration("Commands:\n")
		chat.printDecoration("    .new          Create a new chat\n")
		chat.printDecoration("    .quit .exit   Exit\n")
		chat.printDecoration("    .help         Show this message\n\n")
	}
	return true, nil
}
