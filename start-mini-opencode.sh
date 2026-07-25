#!/bin/bash
set -e

# start-mini-opencode.sh — 启动精简版 opencode (Web Server Only)
# 通过 SSH 端口 8022 连接后执行此脚本

BINARY="opencode.minimal-arm64"
REPO="openclawdev-agent/opencode-minimal"
INSTALL_DIR="$HOME/.local/bin"
CONFIG_DIR="$HOME/.config/opencode"
PORT="${OPENCODE_PORT:-8787}"

echo "🚀 启动 mini-opencode..."
echo "   二进制: $BINARY"
echo "   端口: $PORT"
echo ""

# 1. 检查/下载二进制
if [ ! -f "$INSTALL_DIR/$BINARY" ]; then
    echo "📥 下载 $BINARY..."
    mkdir -p "$INSTALL_DIR"
    
    # 尝试从 GitHub Release 下载
    DOWNLOAD_URL="https://github.com/$REPO/releases/latest/download/${BINARY}"
    if command -v curl &>/dev/null; then
        curl -L -o "$INSTALL_DIR/$BINARY" "$DOWNLOAD_URL" || {
            echo "❌ 下载失败，请手动下载并放置到 $INSTALL_DIR/$BINARY"
            exit 1
        }
    elif command -v wget &>/dev/null; then
        wget -O "$INSTALL_DIR/$BINARY" "$DOWNLOAD_URL" || {
            echo "❌ 下载失败，请手动下载并放置到 $INSTALL_DIR/$BINARY"
            exit 1
        }
    else
        echo "❌ 需要 curl 或 wget，请手动下载"
        echo "   URL: $DOWNLOAD_URL"
        exit 1
    fi
    chmod +x "$INSTALL_DIR/$BINARY"
    echo "✅ 下载完成"
fi

# 2. 创建配置目录
mkdir -p "$CONFIG_DIR"

# 3. 检查端口占用
if lsof -i :"$PORT" &>/dev/null 2>&1 || ss -tlnp 2>/dev/null | grep -q ":$PORT "; then
    echo "⚠️  端口 $PORT 已被占用"
    echo "   使用 OPENCODE_PORT=xxxx 指定其他端口"
    exit 1
fi

# 4. 启动 web 服务器
echo ""
echo "🌐 启动 Web Server..."
echo "   访问地址: http://localhost:$PORT"
echo "   或通过 SSH 隧道: ssh -L $PORT:localhost:$PORT -p 8022 user@host"
echo ""
echo "   按 Ctrl+C 停止"
echo ""

exec "$INSTALL_DIR/$BINARY" web --port "$PORT"
