#!/usr/bin/env bash
# ================================================================
#  Lexvia 服务器初始化脚本（一次性执行）
#
#  用法:  bash setup-server.sh
#  可选:  bash setup-server.sh /opt/new-api  custom-image-name
# ================================================================
set -euo pipefail

DEPLOY_DIR="${1:-/opt/new-api}"
IMAGE_NAME="${2:-new-api}"

echo "=== Lexvia Server Setup ==="
echo "DEPLOY_DIR=${DEPLOY_DIR}"
echo "IMAGE_NAME=${IMAGE_NAME}"

# 1. 安装 Docker
if ! command -v docker &>/dev/null; then
  echo "[1/5] Installing Docker..."
  curl -fsSL https://get.docker.com | sh
  sudo systemctl enable --now docker
  sudo usermod -aG docker "${SUDO_USER:-$USER}" 2>/dev/null || true
else
  echo "[1/5] Docker already installed: $(docker --version)"
fi

# 2. 安装 Docker Compose v2
if ! docker compose version &>/dev/null; then
  echo "[2/5] Installing Docker Compose v2..."
  COMPOSE_VER="v2.29.2"
  case "$(uname -m)" in
    x86_64)  BIN="docker-compose-linux-x86_64" ;;
    aarch64) BIN="docker-compose-linux-aarch64" ;;
    *)       echo "Unsupported arch"; exit 1 ;;
  esac
  sudo curl -fsSL "https://github.com/docker/compose/releases/download/${COMPOSE_VER}/${BIN}" \
    -o /usr/local/bin/docker-compose
  sudo chmod +x /usr/local/bin/docker-compose
else
  echo "[2/5] Docker Compose already installed: $(docker compose version)"
fi

# 3. 创建部署目录
echo "[3/5] Creating deployment directory..."
sudo mkdir -p "${DEPLOY_DIR}"/{data,logs}
sudo chown -R "$(whoami):$(whoami)" "${DEPLOY_DIR}"

# 4. 下载 docker-compose.yml
echo "[4/5] Downloading docker-compose.yml..."
curl -fsSL "https://raw.githubusercontent.com/gavintony1990/Lexvia/Lexvia/docker-compose.yml" \
  -o "${DEPLOY_DIR}/docker-compose.yml"

# 5. 生成 SSH 密钥（供 GitHub Actions 使用）
echo "[5/5] Generating SSH key..."
mkdir -p ~/.ssh
KEY_PATH="${HOME}/.ssh/deploy_key"
if [ ! -f "${KEY_PATH}" ]; then
  ssh-keygen -t ed25519 -f "${KEY_PATH}" -N '' -C 'lexvia-deploy'
fi
# 添加到 authorized_keys（如果还没加）
PUB_KEY=$(cat "${KEY_PATH}.pub")
if ! grep -q "$PUB_KEY" ~/.ssh/authorized_keys 2>/dev/null; then
  echo "$PUB_KEY" >> ~/.ssh/authorized_keys
  chmod 600 ~/.ssh/authorized_keys
fi
echo ""
echo "=========================================="
echo "  初始化完成！"
echo "=========================================="
echo ""
echo "  部署目录:    ${DEPLOY_DIR}"
echo "  镜像名称:    ${IMAGE_NAME}"
echo ""
echo "  请在 GitHub 仓库 Settings → Actions → Secrets"
echo "  添加以下 6 个密钥："
echo ""
echo "    SERVER_HOST        = 本服务器 IP/域名"
echo "    SERVER_USER        = $(whoami)"
echo "    SERVER_PORT        = 22（如已修改请填对应值）"
echo "    SERVER_DEPLOY_DIR  = ${DEPLOY_DIR}"
echo "    SERVER_IMAGE_NAME  = ${IMAGE_NAME}"
echo "    SERVER_SSH_KEY     = 下面这条命令的完整输出"
echo ""
echo "  # 复制私钥到 GitHub Secrets:"
echo "  cat ~/.ssh/deploy_key"
echo ""
echo "  # 然后去 GitHub Actions 触发 'Deploy to Server' workflow"
