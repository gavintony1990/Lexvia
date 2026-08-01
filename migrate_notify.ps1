$modulePath = "github.com/gavintony1990/Lexvia"
$httputilImport = "`"$modulePath/service/httputil`""

# 1. Copy 3 notify files to service/notify/ with package notify + httputil refs
$files = @("notify-limit.go", "user_notify.go", "webhook.go")
foreach ($f in $files) {
    $src = "I:/ADM/Lexvia/service/$f"
    $dst = "I:/ADM/Lexvia/service/notify/$f"
    $content = [System.IO.File]::ReadAllText($src)
    $content = $content -replace '^package service', 'package notify'
    # Replace bare httputil function refs (they were same-package calls in service)
    $content = $content -replace '\bWorkerRequest\{', 'httputil.WorkerRequest{'
    $content = $content -replace '&WorkerRequest\{', '&httputil.WorkerRequest{'
    $content = $content -replace '\bDoWorkerRequest\(', 'httputil.DoWorkerRequest('
    $content = $content -replace '\bGetHttpClient\(\)', 'httputil.GetHttpClient()'
    # Add httputil import after system_setting import if the file uses httputil
    if ($content.Contains('httputil.')) {
        $anchor = "`"$modulePath/setting/system_setting`""
        if ($content.Contains($anchor)) {
            $content = $content -replace [regex]::Escape($anchor), ($httputilImport + "`n`t" + $anchor)
        }
    }
    [System.IO.File]::WriteAllText($dst, $content)
    Write-Host "CREATED $dst"
}

# 2. Update controller callers (service.Notify* -> notify.Notify*)
$controllerFiles = @(
    "I:/ADM/Lexvia/controller/channel_upstream_update.go",
    "I:/ADM/Lexvia/controller/channelctrl/channel_upstream_update.go",
    "I:/ADM/Lexvia/controller/channel-test.go",
    "I:/ADM/Lexvia/controller/channelctrl/channel-test.go"
)
foreach ($f in $controllerFiles) {
    if (-not (Test-Path $f)) { Write-Host "SKIP $f"; continue }
    $content = [System.IO.File]::ReadAllText($f)
    $changed = $false
    if ($content.Contains('service.NotifyUpstreamModelUpdateWatchers')) {
        $content = $content -replace 'service\.NotifyUpstreamModelUpdateWatchers', 'notify.NotifyUpstreamModelUpdateWatchers'
        $changed = $true
    }
    if ($content.Contains('service.NotifyRootUser')) {
        $content = $content -replace 'service\.NotifyRootUser', 'notify.NotifyRootUser'
        $changed = $true
    }
    if ($changed) {
        # Add notify import alongside service import
        $anchor = "`"$modulePath/service`""
        if ($content.Contains($anchor) -and -not $content.Contains("`"$modulePath/service/notify`"")) {
            $content = $content -replace [regex]::Escape($anchor), ("notify `"$modulePath/service/notify`"`n`t" + $anchor)
        }
        [System.IO.File]::WriteAllText($f, $content)
        Write-Host "UPDATED $f"
    }
}

# 3. Update internal service callers: channel.go (NotifyRootUser), quota.go (NotifyUser)
$serviceCallers = @(
    @{ File = "I:/ADM/Lexvia/service/channel.go"; Old = 'NotifyRootUser('; New = 'notify.NotifyRootUser(' },
    @{ File = "I:/ADM/Lexvia/service/quota.go"; Old = 'err := NotifyUser('; New = 'err := notify.NotifyUser(' }
)
foreach ($c in $serviceCallers) {
    if (-not (Test-Path $c.File)) { Write-Host "SKIP $($c.File)"; continue }
    $content = [System.IO.File]::ReadAllText($c.File)
    if ($content.Contains($c.Old)) {
        $content = $content -replace [regex]::Escape($c.Old), $c.New
        # Add notify import
        $anchor = "`"$modulePath/service/httputil`""
        if ($content.Contains($anchor) -and -not $content.Contains("`"$modulePath/service/notify`"")) {
            $content = $content -replace [regex]::Escape($anchor), ("notify `"$modulePath/service/notify`"`n`t" + $anchor)
        } elseif (-not $content.Contains("`"$modulePath/service/notify`"")) {
            # fallback: add after model import
            $modelAnchor = "`"$modulePath/model`""
            if ($content.Contains($modelAnchor)) {
                $content = $content -replace [regex]::Escape($modelAnchor), ($modelAnchor + "`n`tnotify `"$modulePath/service/notify`"")
            }
        }
        [System.IO.File]::WriteAllText($c.File, $content)
        Write-Host "UPDATED $($c.File)"
    } else {
        Write-Host "NO MATCH in $($c.File)"
    }
}
Write-Host "DONE"