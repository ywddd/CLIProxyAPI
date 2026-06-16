# CPA 自用版

这是基于 [router-for-me/CLIProxyAPI](https://github.com/router-for-me/CLIProxyAPI) 的自用构建，重点服务 Codex/Responses 稳定性、多账号运行、NAS/Docker 部署和日常 CPA 管理。

当前同步基线：上游 `v7.2.12`。自用版本建议标记为：

```text
v7.2.12-selfuse.20260617
```

## 本构建保留的改动

### 1. Codex 上下文过长直接交回客户端

当 Codex 上游以 `context_too_large` / `context_length_exceeded` 结束流式响应时，本构建不在 CPA 中间层自行压缩历史、生成 `history.txt` 或移除 reasoning 后继续重试。

这样可以避免 CPA 把历史会话改写成新的请求后再次喂给模型，降低长会话里重复读工作区、重复确认状态、重复规划的风险。

### 2. 加密 reasoning 上下文降级重试

部分 Codex/Responses 请求会携带 `input[*].encrypted_content`。当上游明确拒绝这段加密 reasoning 上下文时，本构建会移除无效的加密 reasoning 上下文，并重试一次。

同时，当上游返回 `Item with id 'rs_...' not found` 且提示 `store=false` 时，也会移除 stale reasoning item 并重试一次。

### 3. Codex 响应头超时

Codex 上游请求有时会在返回响应头前卡住。本构建增加只作用于响应头阶段的超时：

```yaml
codex-response-header-timeout-seconds: 180
```

响应头到达后的流式正文不受该超时限制。设置为负数可关闭：

```yaml
codex-response-header-timeout-seconds: -1
```

也支持环境变量覆盖：

```bash
CPA_CODEX_RESPONSE_HEADER_TIMEOUT_SECONDS=180
```

### 4. OpenAI-compatible JSON 预检

Kimi K2.7 Code 等走 `openai-compatibility` 的模型在请求体包含未转义反斜杠时，上游可能返回 Cloudflare 侧的 `invalid escaped character in string`。

本构建会在入口路由前和发往 OpenAI-compatible 上游前对 JSON 做兼容处理：

- 对 `C:\Users\...` 这类常见未转义 Windows 路径，自动修复字符串里的非法反斜杠转义后继续请求。
- `/v1/chat/completions` 和 `/v1/completions` 会先修复/校验请求体，再读取 `model` 做 provider 路由。
- 对缺引号、结构损坏等不可恢复的非法 JSON，仍然在 CPA 本地返回 `400`。

### 5. 管理 UI 增强

管理页保留 selfuse 的运维增强：

- 可视化配置 `codex-response-header-timeout-seconds`。
- auth 文件单独测试模型。
- 当前页批量测试 auth 文件。
- 每个账号显示测试结果和延迟。

## 上游同步摘要

本轮从 `v7.2.5` 合并到 `v7.2.12`，重点包括：

- 管理日志 API 增加 cursor/tail/轮转续读能力。
- 插件删除、配置修改、生命周期和 stream callback 的异步 reload/race 修复。
- 新增插件 ModelRouter，可在鉴权前做模型路由。
- Claude/Anthropic 兼容增强，包括 web search tool domain 清洗、tool_result 规范化、Codex web_search_call 流式转换修复、namespace/function call 映射增强。
- 视频输入增强，增加 `video_url` 提取和校验。
- 插件默认 `Enabled` 行为改为 `false`，已有插件配置需要显式启用。

上游已经覆盖的通用修复尽量使用官方实现；上游尚未覆盖的 selfuse 运行补丁继续保留。

## 推荐配置

```yaml
request-retry: 3
max-retry-credentials: 3
max-retry-interval: 30

routing:
  session-affinity: true

nonstream-keepalive-interval: 15
codex-response-header-timeout-seconds: 180

streaming:
  keepalive-seconds: 15
  bootstrap-retries: 1
```

## Docker Compose 使用

构建并启动：

```bash
docker compose up -d --build
```

管理页和 API 端口取决于你的 compose 文件。参考部署中：

```text
CPA API:    http://<host>:8317
CPA Plus:   http://<host>:18317/management.html
```

## 版本规则

本仓库的自用发布版本固定使用 `selfuse` 后缀，例如：

```text
v7.2.12-selfuse.20260617
```

NAS 本地 Docker 镜像建议使用稳定标签：

```text
cli-proxy-api:v7.2.12-selfuse.20260617
```

## 安全说明

不要提交真实 auth 文件、refresh token、access token、id token、management key 或 API key。推荐作为运行态文件保留在仓库外或 `.gitignore` 中：

```text
auth-dir/
auths/
logs/
*.sqlite
*.db
config.yaml
```

公开 fork 或发布前，建议扫描敏感信息：

```bash
rg -n "github_pat_|refresh_token|access_token|id_token|sk-[A-Za-z0-9]|secret-key:" .
```
