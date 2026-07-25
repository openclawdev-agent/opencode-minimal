package lsp

// 精简版 stub：移除LSP依赖
// 原版LSP用于代码智能提示，Web Server不需要

import (
	"context"
	"sync"
)

// Client stub
type Client struct {
	mu sync.Mutex
}

// NewClient stub
func NewClient(ctx context.Context, command string, args ...string) (*Client, error) {
	return &Client{}, nil
}

// Shutdown stub
func (c *Client) Shutdown() error {
	return nil
}

// LSPClients type
type LSPClients = map[string]*Client
