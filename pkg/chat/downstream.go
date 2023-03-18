package chat

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

// recvFromDownstream 从下游接收消息
func (chat *consoleChat) recvFromDownstream() (*Message, error) {
	var content string

	fmt.Print(">>> ")
recvLoop:
	for {
		line, err := chat.input.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, fmt.Errorf("read from stdin error: %w", err)
		}

		switch line {
		case "\n":
			break recvLoop
		case "exit\n", "quit\n":
			chat.cancel()
			return nil, nil
		}
		content += line
		fmt.Print("... ")
	}

	content = strings.TrimRight(content, "\n")

	return &Message{
		Content: content,
	}, nil
}

// sendToDownstream 发送消息到下游
func (chat *consoleChat) sendToDownstream(msg *Message) error {
	fmt.Println(msg.String())
	return nil
}
