# opencode-minimal

精简版 opencode — 仅保留 Web Server 能力，移除 TUI 等非必要模块。

## 目标

| 指标 | 原版 | 精简版 |
|------|------|--------|
| 二进制体积 | ~154MB | **~30-50MB** |
| 符号数量 | ~556K | ~100K |
| 架构 | 全功能 | Web Server Only |

## 精简策略

1. **移除 TUI 模块** (`internal/tui/` 及其依赖)
2. **移除 Shell Completions** (`internal/completions/`)
3. **移除 LSP** (`internal/lsp/`)
4. **保留核心**：app、config、db、llm、session、pubsub、permission、logging
5. **Go 构建优化**：`-ldflags="-s -w" -trimpath`
6. **交叉编译**：GitHub Actions ARM64 静态二进制

## 构建

GitHub Actions 自动构建，产物在 Releases 页面下载。

手动构建（需要 Go 1.24+）：
```bash
./scripts/build-strip.sh
```

## 使用

与原版 `opencode serve` 完全兼容：
```bash
./opencode.minimal serve --port 8787
```

## 上游

基于 [opencode-ai/opencode](https://github.com/opencode-ai/opencode) v1.17.x。
