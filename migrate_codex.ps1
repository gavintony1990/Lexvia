$modulePath = "github.com/gavintony1990/Lexvia"
$httputilImport = "`"$modulePath/service/httputil`""
$codexImport = "codexSvc `"$modulePath/service/codex`""

New-Item -ItemType Directory -Force -Path "I:/ADM/Lexvia/service/codex" | Out-Null

# 1. Copy 4 codex files to service/codex/ with package codex + httputil refs
$files = @("codex_oauth.go", "codex_credential_refresh.go", "codex_credential_refresh_task.go", "codex_wham_usage.go")
foreach ($f in $files) {
    $src = "I:/ADM/Lexvia/service/$f"
    $dst = "I:/ADM/Lexvia/service/codex/$f"
    $content = [System.IO.File]::ReadAllText($src)
    $content = $content -replace '^package service', 'package codex'
    # Replace bare httputil function refs (were same-package calls)
    $content = $content -replace '\bGetHttpClientWithProxy\(', 'httputil.GetHttpClientWithProxy('
    $content = $content -replace '\bGetHttpClient\(\)', 'httputil.GetHttpClient()'
    $content = $content -replace '\bResetProxyClientCache\(\)', 'httputil.ResetProxyClientCache()'
    # Add httputil import after common import if file uses httputil
    if ($content.Contains('httputil.')) {
        $anchor = "`"$modulePath/common`""
        if ($content.Contains($anchor)) {
            $content = $content -replace [regex]::Escape($anchor), ($anchor + "`n`t" + $httputilImport)
        }
    }
    [System.IO.File]::WriteAllText($dst, $content)
    Write-Host "CREATED $dst"
}

# 2. Update external callers
$callers = @(
    "I:/ADM/Lexvia/main.go",
    "I:/ADM/Lexvia/controller/channel.go",
    "I:/ADM/Lexvia/controller/channelctrl/channel.go",
    "I:/ADM/Lexvia/controller/codex_usage.go"
)
foreach ($f in $callers) {
    if (-not (Test-Path $f)) { Write-Host "SKIP $f"; continue }
    $content = [System.IO.File]::ReadAllText($f)
    $changed = $false
    foreach ($fn in @('StartCodexCredentialAutoRefreshTask', 'RefreshCodexChannelCredential', 'CodexCredentialRefreshOptions', 'RefreshCodexOAuthTokenWithProxy', 'RefreshCodexOAuthToken', 'FetchCodexWhamUsage', 'CodexOAuthTokenResult', 'CodexOAuthKey')) {
        if ($content.Contains("service.$fn")) {
            $content = $content -replace "service\.$fn", "codexSvc.$fn"
            $changed = $true
        }
    }
    if ($changed) {
        $anchor = "`"$modulePath/service`""
        if ($content.Contains($anchor) -and -not $content.Contains("service/codex")) {
            $content = $content -replace [regex]::Escape($anchor), ($codexImport + "`n`t" + $anchor)
        }
        [System.IO.File]::WriteAllText($f, $content)
        Write-Host "UPDATED $f"
    } else {
        Write-Host "NO CHANGE $f"
    }
}
Write-Host "DONE"