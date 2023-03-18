package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	defaultServer = "https://api.openai.com"
)

// restClient 是 Client 的一个实现
type restClient struct {
	config     *Config
	httpClient *http.Client
}

var _ Client = &restClient{}

// NewForConfig 使用指定客户端配置创建一个客户端
func NewForConfig(config *Config) (Client, error) {
	if config == nil {
		return nil, fmt.Errorf("config can not be nil")
	}
	if config.Server == "" {
		config.Server = defaultServer
	}

	// 设置 http 客户端
	httpClient := &http.Client{}
	if config.Proxy != "" {
		proxyURL, err := url.Parse(config.Proxy)
		if err != nil {
			return nil, fmt.Errorf("parse proxy url error: %w", err)
		}
		httpClient.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	}

	return &restClient{
		config:     config,
		httpClient: httpClient,
	}, nil
}

// CreateChatCompletion 创建对话
func (c *restClient) CreateChatCompletion(ctx context.Context, input *CreateChatCompletionInput) (*CreateChatCompletionOutput, error) {
	var output CreateChatCompletionOutput
	return &output, c.doRequest(ctx, http.MethodPost, "/v1/chat/completions", input, &output)
}

// doRequest 请求
func (c *restClient) doRequest(ctx context.Context, method, path string, input interface{}, output interface{}) error {
	// 生成请求体
	var body io.Reader
	if input != nil {
		raw, err := json.Marshal(input)
		if err != nil {
			return fmt.Errorf("marshal input error: %w", err)
		}
		body = bytes.NewReader(raw)
	}

	// 生成请求
	req, err := http.NewRequestWithContext(ctx, method, c.config.Server+path, body)
	if err != nil {
		return fmt.Errorf("make request error: %w", err)
	}

	// 添加认证头
	if c.config.SecretKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.config.SecretKey)
	}
	if c.config.OrganizationID != "" {
		req.Header.Set("OpenAI-Organization", c.config.OrganizationID)
	}
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("send request error: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response erro: %w", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("received response with unexpected status code: %d, body: %s", resp.StatusCode, string(respBody))
	}

	if output != nil {
		if err := json.Unmarshal(respBody, output); err != nil {
			return fmt.Errorf("unmarshal response body error: %w", err)
		}
	}

	return nil
}
