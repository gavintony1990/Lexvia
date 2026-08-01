$src = "I:/ADM/Lexvia/service"
$dst = "I:/ADM/Lexvia/service/httputil"

# Create dir if missing
if (-not (Test-Path $dst)) { New-Item -ItemType Directory -Path $dst -Force | Out-Null }

# Copy files with package name change
$files = @("http_client.go", "http.go", "download.go")
foreach ($f in $files) {
    $content = [System.IO.File]::ReadAllText("$src/$f")
    $content = $content -replace '^package service', 'package httputil'
    [System.IO.File]::WriteAllText("$dst/$f", $content)
    Write-Host "Copied $f"
}
Write-Host "Done"