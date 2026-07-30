# Lexvia 免费部署方案

## 核心思路

```
GitHub Actions (免费 2000 min/月)
    │
    │  ┌─ Docker 多阶段构建 (前端 + Go 后端)
    │  └─ docker save → tar.gz (~250MB)
    │
    ↓ SSH SCP (加密传输)
    │
云服务器 (你自己准备)
    │
    └─ docker load → docker compose up
    └─ curl healthz 健康检查
    └─ 完成

镜像全程不经过 Docker Hub / GHCR，不会被任何人拉取。
```

---

## 服务器需求

| 资源 | 最低 | 推荐 |
|------|------|------|
| CPU | 1 核 | 2 核+ |
| 内存 | 2 GB | 4 GB+ |
| 磁盘 | 10 GB | 20 GB+ (Docker 镜像 + 数据库) |
| 系统 | Ubuntu 20.04+ / Debian 11+ | Ubuntu 22.04+ |
| 端口 | 3000, 22 开放 | - |

**推荐免费/低成本方案**：
- 阿里云/腾讯云 轻量应用服务器：约 ¥8-20/月
- 甲骨文云 Always Free：4 台 ARM 24GB（免费）
- 华为云/百度云 试用主机

---

## 服务器端准备（一次性操作）

### 1. 安装 Docker

```bash
# Ubuntu / Debian
curl -fsSL https://get.docker.com | sh
sudo systemctl enable --now docker
sudo usermod -aG docker $USER
newgrp docker

# 安装 Docker Compose v2
COMPOSE_VER="v2.29.2"
ARCH=$(uname -m)
[ "$ARCH" = "aarch64" ] && COMPOSE_BIN="docker-compose-linux-aarch64" || COMPOSE_BIN="docker-compose-linux-x86_64"
sudo curl -fsSL "https://github.com/docker/compose/releases/download/${COMPOSE_VER}/${COMPOSE_BIN}" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
```

### 2. 创建部署目录

```bash
sudo mkdir -p /opt/new-api/{data,logs}
sudo chown -R $USER:$USER /opt/new-api
```

### 3. 生成 SSH 密钥对（用于 GitHub Actions 连接）

```bash
mkdir -p ~/.ssh
ssh-keygen -t ed25519 -f ~/.ssh/deploy_key -N '' -C 'github-deploy'
cat ~/.ssh/deploy_key.pub >> ~/.ssh/authorized_keys
chmod 600 ~/.ssh/authorized_keys
```

### 4. 放置 docker-compose.yml

```bash
curl -fsSL "https://raw.githubusercontent.com/gavintony1990/Lexvia/Lexvia/docker-compose.yml" \
  -o /opt/new-api/docker-compose.yml
```

---

## GitHub Secrets 配置

在仓库设置 → Actions secrets 中添加 **6 个密钥**：

| Secret | 值示例 | 说明 |
|--------|--------|------|
| `SERVER_HOST` | `1.2.3.4` | 服务器 IP |
| `SERVER_USER` | `root` | SSH 用户名 |
| `SERVER_PORT` | `22` | SSH 端口（默认可不填） |
| `SERVER_SSH_KEY` | `-----BEGIN OPENSSH PRIVATE KEY----- ...` | **`~/.ssh/deploy_key` 完整内容** |
| `SERVER_DEPLOY_DIR` | `/opt/new-api` | 部署目录 |
| `SERVER_IMAGE_NAME` | `new-api` | 服务器本地镜像名 |

SSH 密钥配置：
```bash
# 在服务器上
cat ~/.ssh/deploy_key
# 把完整输出（含 BEGIN/END 行）粘贴到 GitHub Secrets
```

---

## 部署操作

1. 打开 GitHub Actions 页面
2. 选择 **Deploy to Server**
3. 点击 **Run workflow**
4. 参数：
   - **Tag name**: `latest`（首次）
   - **Environment**: `production`
5. 点击 **Run**

**流程**：Actions 拉代码 → 多阶段 Docker 构建（~5min） → tar.gz SCP（~2min） → `docker load` + compose up → healthcheck → 完成

**总耗时**：首次约 10 分钟，后续增量约 5-7 分钟。

---

## 访问服务

```
http://<服务器IP>:3000
```
首次访问进入初始化向导。

---

## 成本

| 项目 | 费用 |
|------|------|
| GitHub Actions（开源仓库） | **免费**（2000 min/月） |
| 云服务器（1C2G） | 约 ¥8-20/月 |
| **总计** | **约 ¥100-240/年** |
