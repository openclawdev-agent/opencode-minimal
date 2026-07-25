package lsp

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/opencode-ai/opencode/internal/lsp/protocol"
)

type ServerState int

const (
	StateStarting ServerState = iota
	StateReady
	StateError
)

type NotificationHandler func(json.RawMessage)

type Client struct {
	notificationHandlers map[string]NotificationHandler
	notificationMu       sync.RWMutex
	diagnostics          map[protocol.DocumentUri][]protocol.Diagnostic
	diagnosticsMu        sync.RWMutex
	serverState          ServerState
}

func NewClient(ctx context.Context, command string, args ...string) (*Client, error) {
	return &Client{
		notificationHandlers: make(map[string]NotificationHandler),
		diagnostics:          make(map[protocol.DocumentUri][]protocol.Diagnostic),
	}, nil
}

func (c *Client) InitializeLSPClient(ctx context.Context, workspaceDir string) (*protocol.InitializeResult, error) {
	return nil, nil
}

func (c *Client) Close() error                              { return nil }
func (c *Client) Shutdown(ctx context.Context) error         { return nil }
func (c *Client) WaitForServerReady(ctx context.Context) error { return nil }
func (c *Client) GetServerState() ServerState               { return c.serverState }
func (c *Client) SetServerState(state ServerState)          { c.serverState = state }
func (c *Client) OpenFile(ctx context.Context, filepath string) error    { return nil }
func (c *Client) NotifyChange(ctx context.Context, filepath string) error { return nil }
func (c *Client) IsFileOpen(filepath string) bool              { return false }

func (c *Client) GetDiagnostics() map[protocol.DocumentUri][]protocol.Diagnostic {
	c.diagnosticsMu.RLock()
	defer c.diagnosticsMu.RUnlock()
	return c.diagnostics
}

func (c *Client) RegisterNotificationHandler(method string, handler NotificationHandler) {
	c.notificationMu.Lock()
	defer c.notificationMu.Unlock()
	c.notificationHandlers[method] = handler
}

func HandleDiagnostics(client *Client, params json.RawMessage) {}

type LSPClients = map[string]*Client
