package openai

import (
	"context"
	"encoding/json"
	"strconv"
	"time"
)

// Client OpenAI 客户端
type Client interface {
	// 创建对话
	CreateChatCompletion(ctx context.Context, input *CreateChatCompletionInput) (output *CreateChatCompletionOutput, err error)
}

// Config OpenAI 客户端配置
type Config struct {
	// OpenAI API 地址
	// 默认为 https://api.openai.com
	Server string
	// OpenAI API 密钥
	SecretKey string
	// OpenAI 组织 ID
	OrganizationID string
	// Proxy 访问 OpenAI 接口时使用的 http 代理地址
	Proxy string
}

// CreateChatCompletionInput 创建对话接口输入
type CreateChatCompletionInput struct {
	// 使用的模型
	// 参考 https://platform.openai.com/docs/models/model-endpoint-compatibility
	Model string `json:"model"`

	// 输入用于生成对话结果的信息
	Messages []ChatMessage `json:"messages"`

	// 采样温度，介于 0 到 2 之间
	// 越高的值（比如 0.8 ）会使输出更随机
	// 较低的值（比如 0.2 ）会使输出更集中和确定
	// 建议仅调整这个值或 TopP 二者之一，不要同时调整二者
	// 默认 1.0
	Temperature *float64 `json:"temperature,omitempty"`

	// 一种替代 Temperature 的方法，称为“核采样”，其中模型参考 TopP 概率质量的结果
	// 因此 0.1 意味着只有组成前 10% 概率质量的 token 会被参考
	// 默认 1.0
	TopP *float64 `json:"top_p,omitempty"`

	// 对于每个输入信息选择多少个生成结果
	// 默认 1
	N *int64 `json:"n,omitempty"`

	// 是否流式传输结果
	// 如果设置了，输出信息会被一部分一部分得发送，就像 ChatGPT 一样，
	// 输出格式参考 https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events#Event_stream_format
	Stream bool `json:"stream,omitempty"`

	// 停止标志
	// 可以指定最多 4 个字符串，当结果出现这些字符串时模型将停止生成更多 token
	Stop []string `json:"stop,omitempty"`

	// 最大生成 token 数
	// 默认无限
	MaxTokens *int64 `json:"max_tokens,omitempty"`

	// 重复惩罚
	// -2 到 2 的值。正值表示会根据是否已经在结果中出现过而抑制模型生成同样的 token ，从而增加模型谈论新主题的可能性
	// 参考 https://platform.openai.com/docs/api-reference/parameter-details
	// 默认 0
	PresencePenalty float64 `json:"presence_penalty,omitempty"`

	// 频率惩罚
	// -2 到 2 的值。正值表示会根据在结果中出现的频率而抑制模型生成同样的 token ，从而减少模型生成的重复字词
	// 参考 https://platform.openai.com/docs/api-reference/parameter-details
	// 默认 0
	FrequencyPenalty float64 `json:"frequency_penalty,omitempty"`

	// 对数偏差，修改指定 token 出现的可能性
	// 从 token （通过 token ID 指定）到一个 -100 到 100 之间的关联偏差值的映射
	// 具体效果因模型而异，但是 -1 到 1 之间的值应该会减少或增加选择的可能性，像 -100 或 100 这样的值应该会导致相关 token 被完全禁止或完全独占使用
	// token ID 参考 https://platform.openai.com/tokenizer
	LogitBias map[int64]float64 `json:"logit_bias,omitempty"`

	// 最终用户的唯一 ID
	// 可以帮助 OpenAI 监控和检测滥用行为
	// 参考 https://platform.openai.com/docs/guides/safety-best-practices/end-user-ids
	User string `json:"user,omitempty"`
}

// CreateChatCompletionOutput TODO: 创建对话接口输出
type CreateChatCompletionOutput struct {
	// 请求 ID
	ID string `json:"id"`

	// 请求对象
	Object string `json:"object"`

	// 结果创建时间
	Created Time `json:"created"`

	// 结果选择
	Choices []ChatCompletionResult `json:"choices"`

	// token 使用量
	Usage TokenUsage `json:"usage"`
}

// ChatMessageRole 对话消息角色
type ChatMessageRole string

// ChatMessageRole 的可选值
const (
	SystemChatMessageRole    ChatMessageRole = "system"
	UserChatMessageRole      ChatMessageRole = "user"
	AssistantChatMessageRole ChatMessageRole = "assistant"
)

// ChatMessage 对话消息
type ChatMessage struct {
	// 输出该信息的角色
	Role ChatMessageRole `json:"role"`
	// 消息内容
	Content string `json:"content"`
}

// ChatCompletionChoice 对话完成结果
type ChatCompletionResult struct {
	Index        int64       `json:"index"`
	Message      ChatMessage `json:"message"`
	FinishReason string      `json:"finish_reason"`
}

// TokenUsage token 使用量
type TokenUsage struct {
	PromptTokens     int64 `json:"prompt_tokens,omitempty"`
	CompletionTokens int64 `json:"completion_tokens,omitempty"`
	TotalTokens      int64 `json:"total_tokens,omitempty"`
}

// Time 时间
type Time time.Time

var _ json.Marshaler = &Time{}
var _ json.Unmarshaler = &Time{}

// MarshalJSON 序列化为 JSON
func (t *Time) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte("0"), nil
	}
	return []byte(strconv.FormatInt(time.Time(*t).Unix(), 10)), nil
}

// UnmarshalJSON 从 JSON 反序列化
func (t *Time) UnmarshalJSON(data []byte) error {
	ts, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*t = Time(time.Unix(ts, 0))
	return nil
}
