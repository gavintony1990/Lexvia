$src = "I:/ADM/Lexvia/controller"
$dst = "I:/ADM/Lexvia/controller/billing"
[System.IO.Directory]::CreateDirectory($dst) | Out-Null

$patterns = @("billing.go", "checkin.go", "pricing.go", "ratio_config.go", "ratio_sync.go",
              "redemption.go", "return_path.go", "subscription.go", "subscription_payment_creem.go",
              "subscription_payment_epay.go", "subscription_payment_stripe.go", "subscription_payment_waffo_pancake.go",
              "topup.go", "topup_creem.go", "topup_stripe.go", "topup_waffo.go", "topup_waffo_pancake.go",
              "topup_waffo_pancake_test.go", "payment_compliance.go", "payment_webhook_availability.go",
              "payment_webhook_availability_test.go", "channel-billing.go")

foreach ($f in [System.IO.Directory]::GetFiles($src, "*.go")) {
    $name = [System.IO.Path]::GetFileName($f)
    if ($name -in $patterns) {
        $content = [System.IO.File]::ReadAllText($f)
        $content = $content -replace "package controller", "package billing"
        [System.IO.File]::WriteAllText([System.IO.Path]::Combine($dst, $name), $content)
        Write-Host "OK: $name"
    }
}