$files = Get-ChildItem "controller/billing/*.go"
foreach ($f in $files) {
    $content = Get-Content $f.FullName -Raw
    $content = $content -replace "package controller", "package billing"
    Set-Content $f.FullName $content -NoNewline
    Write-Host "Fixed: $($f.Name)"
}