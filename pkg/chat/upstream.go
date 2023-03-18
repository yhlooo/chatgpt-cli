package chat

import "time"

// forwardToUpstream 转发消息到上游
func (chat *consoleChat) forwardToUpstream(msg *Message) (*Message, error) {
	time.Sleep(5 * time.Second)
	return msg, nil
}
