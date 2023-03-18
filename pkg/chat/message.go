package chat

// Message 上下游之间传递的消息
type Message struct {
	Content string
}

// String 转为字符串表示形式
func (msg *Message) String() string {
	if msg == nil {
		return ""
	}
	return msg.Content
}
