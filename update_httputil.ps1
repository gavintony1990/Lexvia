$modulePath = "github.com/gavintony1990/Lexvia"
$httputilImport = "`"$modulePath/service/httputil`""

$replacements = @(
    @{ Old = 'service.CloseResponseBodyGracefully'; New = 'httputil.CloseResponseBodyGracefully' },
    @{ Old = 'service.GetHttpClient()'; New = 'httputil.GetHttpClient()' },
    @{ Old = 'service.GetHttpClientWithProxy'; New = 'httputil.GetHttpClientWithProxy' },
    @{ Old = 'service.NewProxyHttpClient'; New = 'httputil.NewProxyHttpClient' },
    @{ Old = 'service.ResetProxyClientCache'; New = 'httputil.ResetProxyClientCache' },
    @{ Old = 'service.InitHttpClient'; New = 'httputil.InitHttpClient' },
    @{ Old = 'service.IOCopyBytesGracefully'; New = 'httputil.IOCopyBytesGracefully' },
    @{ Old = 'service.ShouldCopyUpstreamHeader'; New = 'httputil.ShouldCopyUpstreamHeader' },
    @{ Old = 'service.DoWorkerRequest'; New = 'httputil.DoWorkerRequest' },
    @{ Old = 'service.WorkerRequest'; New = 'httputil.WorkerRequest' }
)

$files = @(
    "I:/ADM/Lexvia/relay/common_handler/rerank.go",
    "I:/ADM/Lexvia/relay/channel/zhipu_4v/image.go",
    "I:/ADM/Lexvia/relay/channel/zhipu/relay-zhipu.go",
    "I:/ADM/Lexvia/relay/channel/xai/text.go",
    "I:/ADM/Lexvia/relay/channel/vertex/service_account.go",
    "I:/ADM/Lexvia/relay/channel/tencent/relay-tencent.go",
    "I:/ADM/Lexvia/relay/channel/task/vidu/adaptor.go",
    "I:/ADM/Lexvia/relay/channel/task/vertex/adaptor.go",
    "I:/ADM/Lexvia/relay/channel/task/suno/adaptor.go",
    "I:/ADM/Lexvia/relay/channel/task/sora/adaptor.go",
    "I:/ADM/Lexvia/relay/channel/task/kling/adaptor.go",
    "I:/ADM/Lexvia/relay/channel/task/jimeng/adaptor.go",
    "I:/ADM/Lexvia/relay/channel/task/hailuo/adaptor.go",
    "I:/ADM/Lexvia/relay/channel/task/gemini/adaptor.go",
    "I:/ADM/Lexvia/relay/channel/task/doubao/adaptor.go",
    "I:/ADM/Lexvia/relay/channel/task/ali/adaptor.go",
    "I:/ADM/Lexvia/relay/channel/siliconflow/relay-siliconflow.go",
    "I:/ADM/Lexvia/relay/channel/replicate/adaptor.go",
    "I:/ADM/Lexvia/relay/channel/palm/relay-palm.go",
    "I:/ADM/Lexvia/relay/channel/openai/responses_via_chat.go",
    "I:/ADM/Lexvia/relay/channel/openai/relay_responses_compact.go",
    "I:/ADM/Lexvia/relay/channel/openai/relay_responses.go",
    "I:/ADM/Lexvia/relay/channel/openai/relay_image.go",
    "I:/ADM/Lexvia/relay/channel/openai/relay-openai.go",
    "I:/ADM/Lexvia/relay/channel/openai/chat_via_responses.go",
    "I:/ADM/Lexvia/relay/channel/openai/audio.go",
    "I:/ADM/Lexvia/relay/channel/ollama/stream.go",
    "I:/ADM/Lexvia/relay/channel/ollama/relay-ollama.go",
    "I:/ADM/Lexvia/relay/channel/mokaai/relay-mokaai.go",
    "I:/ADM/Lexvia/relay/channel/minimax/tts.go",
    "I:/ADM/Lexvia/relay/channel/minimax/image.go",
    "I:/ADM/Lexvia/relay/channel/jimeng/image.go",
    "I:/ADM/Lexvia/relay/channel/gemini/relay_responses.go",
    "I:/ADM/Lexvia/relay/channel/gemini/relay-gemini.go",
    "I:/ADM/Lexvia/relay/channel/gemini/relay-gemini-native.go",
    "I:/ADM/Lexvia/relay/channel/dify/relay-dify.go",
    "I:/ADM/Lexvia/relay/channel/coze/relay-coze.go",
    "I:/ADM/Lexvia/relay/channel/cohere/relay-cohere.go",
    "I:/ADM/Lexvia/relay/channel/cloudflare/relay_cloudflare.go",
    "I:/ADM/Lexvia/relay/channel/claude/relay-claude.go",
    "I:/ADM/Lexvia/relay/channel/baidu/relay-baidu.go",
    "I:/ADM/Lexvia/relay/channel/aws/relay-aws.go",
    "I:/ADM/Lexvia/relay/channel/ali/rerank.go",
    "I:/ADM/Lexvia/relay/channel/ali/image.go",
    "I:/ADM/Lexvia/relay/channel/api_request.go",
    "I:/ADM/Lexvia/relay/mjproxy_handler.go",
    "I:/ADM/Lexvia/controller/relayctrl/video_proxy.go",
    "I:/ADM/Lexvia/controller/channelctrl/channel_upstream_update.go",
    "I:/ADM/Lexvia/controller/channelctrl/channel.go",
    "I:/ADM/Lexvia/controller/billing/channel-billing.go",
    "I:/ADM/Lexvia/controller/midjourney.go",
    "I:/ADM/Lexvia/controller/codex_usage.go",
    "I:/ADM/Lexvia/controller/channel_upstream_update.go",
    "I:/ADM/Lexvia/controller/channel.go",
    "I:/ADM/Lexvia/controller/channel-billing.go",
    "I:/ADM/Lexvia/main.go"
)

$serviceFiles = @(
    "I:/ADM/Lexvia/service/error.go",
    "I:/ADM/Lexvia/service/midjourney.go",
    "I:/ADM/Lexvia/service/notify/user_notify.go",
    "I:/ADM/Lexvia/service/notify/webhook.go"
)

foreach ($file in $files) {
    if (-not (Test-Path $file)) {
        Write-Host "SKIP $file (not found)"
        continue
    }
    $content = [System.IO.File]::ReadAllText($file)
    $changed = $false
    foreach ($r in $replacements) {
        if ($content.Contains($r.Old)) {
            $content = $content -replace [regex]::Escape($r.Old), $r.New
            $changed = $true
        }
    }
    if ($changed) {
        # Add httputil import if it doesn't already have it
        $searchStr = "`"$modulePath/service`""
        $replaceStr = "`"$modulePath/service/httputil`"`n`t`"$modulePath/service`""
        if ($content.Contains($searchStr)) {
            $content = $content -replace [regex]::Escape($searchStr), $replaceStr
        }
        [System.IO.File]::WriteAllText($file, $content)
        Write-Host "UPDATED $file"
    } else {
        Write-Host "NO CHANGE $file"
    }
}

# Update service/error.go and service/midjourney.go (they import service/ and use httputil functions)
foreach ($file in $serviceFiles) {
    if (-not (Test-Path $file)) {
        Write-Host "SKIP $file (not found)"
        continue
    }
    $content = [System.IO.File]::ReadAllText($file)
    $changed = $false
    foreach ($r in $replacements) {
        if ($content.Contains($r.Old)) {
            $content = $content -replace [regex]::Escape($r.Old), $r.New
            $changed = $true
        }
    }
    if ($changed) {
        # Add httputil import
        $searchStr = "`"$modulePath/service`""
        $replaceStr = "`"$modulePath/service/httputil`"`n`t`"$modulePath/service`""
        if ($content.Contains($searchStr)) {
            $content = $content -replace [regex]::Escape($searchStr), $replaceStr
        }
        [System.IO.File]::WriteAllText($file, $content)
        Write-Host "UPDATED $file"
    }
}

Write-Host "DONE"