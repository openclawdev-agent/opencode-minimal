package watcher

import (
	"context"
	"github.com/opencode-ai/opencode/internal/lsp"
)

type WorkspaceWatcher struct {
	client *lsp.Client
}

func NewWorkspaceWatcher(client *lsp.Client) *WorkspaceWatcher {
	return &WorkspaceWatcher{client: client}
}

func (w *WorkspaceWatcher) WatchWorkspace(ctx context.Context, workspacePath string) {}
