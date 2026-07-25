#!/bin/bash
set -e

# opencode-minimal 构建脚本
# 从 opencode 源码中移除非必要模块，构建精简二进制

REPO="https://github.com/opencode-ai/opencode.git"
TAG="${1:-v1.17.11}"
WORKDIR=$(pwd)/_build_opencode
OUT="opencode.minimal"

echo "=== opencode-minimal 构建 ==="
echo "目标版本: $TAG"

# 1. 克隆源码
echo "[1/5] 克隆源码..."
rm -rf "$WORKDIR"
git clone --depth 1 --branch "$TAG" "$REPO" "$WORKDIR"
cd "$WORKDIR"

# 2. 移除非必要模块
echo "[2/5] 移除非必要模块..."
# TUI 是最大的模块
rm -rf internal/tui
# Shell completions
rm -rf internal/completions
# LSP (Language Server Protocol)
rm -rf internal/lsp

# 3. 修补 import 引用
echo "[3/5] 修补 import..."
# app.go 引用了 tui/theme，需要 stub 化
# 创建最小 theme stub
mkdir -p internal/tui/theme
cat > internal/tui/theme/theme.go << 'THEMEEOF'
package theme

// Stub: 原版 theme 被 app 模块引用，精简版不需要 UI 主题
// 保留包声明以满足 import，所有导出类型返回零值

type Theme struct{}

func (Theme) Get() Theme { return Theme{} }
func Init(_ string)      {}
THEMEEOF

# 4. 修补 app/app.go 中的 tui/theme 引用
echo "[4/5] 修补 app 模块..."
# 替换 import 路径（如果 theme stub 不够，直接删 import）
if grep -q 'tui/theme' internal/app/app.go 2>/dev/null; then
    # theme stub 已就位，import 应该能通过
    echo "  tui/theme import 将由 stub 包满足"
fi

# 5. 构建
echo "[5/5] 构建精简二进制..."
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 \
  go build -trimpath -ldflags="-s -w" \
  -o "../../$OUT" . || {
    echo "构建失败，尝试不移除模块的纯构建优化..."
    cd "$WORKDIR"
    rm -rf internal/tui internal/completions internal/lsp
    git checkout -- .
    CGO_ENABLED=0 GOOS=linux GOARCH=arm64 \
      go build -trimpath -ldflags="-s -w" \
      -o "../../$OUT" .
}

# 清理
cd /
rm -rf "$WORKDIR"

# 结果
SIZE=$(ls -lh "../../$OUT" 2>/dev/null | awk '{print $5}')
echo ""
echo "=== 构建完成 ==="
echo "输出: $OUT"
echo "体积: $SIZE"
