$files = @{
    "I:/ADM/Lexvia/router/api-router.go" = @(
        "controller.GetPricing", "billing.GetPricing",
        "controller.GetRatioConfig", "billing.GetRatioConfig",
        "controller.StripeWebhook", "billing.StripeWebhook",
        "controller.CreemWebhook", "billing.CreemWebhook",
        "controller.WaffoWebhook", "billing.WaffoWebhook",
        "controller.WaffoPancakeWebhook", "billing.WaffoPancakeWebhook",
        "controller.EpayNotify", "billing.EpayNotify",
        "controller.GetAffCode", "billing.GetAffCode",
        "controller.GetTopUpInfo", "billing.GetTopUpInfo",
        "controller.GetUserTopUps", "billing.GetUserTopUps",
        "controller.TopUp", "billing.TopUp",
        "controller.RequestEpay", "billing.RequestEpay",
        "controller.RequestAmount", "billing.RequestAmount",
        "controller.RequestStripePay", "billing.RequestStripePay",
        "controller.RequestStripeAmount", "billing.RequestStripeAmount",
        "controller.RequestCreemPay", "billing.RequestCreemPay",
        "controller.RequestWaffoAmount", "billing.RequestWaffoAmount",
        "controller.RequestWaffoPay", "billing.RequestWaffoPay",
        "controller.RequestWaffoPancakeAmount", "billing.RequestWaffoPancakeAmount",
        "controller.RequestWaffoPancakePay", "billing.RequestWaffoPancakePay",
        "controller.TransferAffQuota", "billing.TransferAffQuota",
        "controller.GetCheckinStatus", "billing.GetCheckinStatus",
        "controller.DoCheckin", "billing.DoCheckin",
        "controller.GetSubscriptionPlans", "billing.GetSubscriptionPlans",
        "controller.GetSubscriptionSelf", "billing.GetSubscriptionSelf",
        "controller.UpdateSubscriptionPreference", "billing.UpdateSubscriptionPreference",
        "controller.SubscriptionRequestBalancePay", "billing.SubscriptionRequestBalancePay",
        "controller.SubscriptionRequestEpay", "billing.SubscriptionRequestEpay",
        "controller.SubscriptionRequestStripePay", "billing.SubscriptionRequestStripePay",
        "controller.SubscriptionRequestCreemPay", "billing.SubscriptionRequestCreemPay",
        "controller.SubscriptionRequestWaffoPancakePay", "billing.SubscriptionRequestWaffoPancakePay",
        "controller.AdminListSubscriptionPlans", "billing.AdminListSubscriptionPlans",
        "controller.AdminCreateSubscriptionPlan", "billing.AdminCreateSubscriptionPlan",
        "controller.AdminUpdateSubscriptionPlan", "billing.AdminUpdateSubscriptionPlan",
        "controller.AdminUpdateSubscriptionPlanStatus", "billing.AdminUpdateSubscriptionPlanStatus",
        "controller.AdminBindSubscription", "billing.AdminBindSubscription",
        "controller.AdminListUserSubscriptions", "billing.AdminListUserSubscriptions",
        "controller.AdminCreateUserSubscription", "billing.AdminCreateUserSubscription",
        "controller.AdminInvalidateUserSubscription", "billing.AdminInvalidateUserSubscription",
        "controller.AdminDeleteUserSubscription", "billing.AdminDeleteUserSubscription",
        "controller.SubscriptionEpayNotify", "billing.SubscriptionEpayNotify",
        "controller.SubscriptionEpayReturn", "billing.SubscriptionEpayReturn",
        "controller.ResetModelRatio", "billing.ResetModelRatio",
        "controller.MigrateConsoleSetting", "billing.MigrateConsoleSetting",
        "controller.ListWaffoPancakeCatalog", "billing.ListWaffoPancakeCatalog",
        "controller.CreateWaffoPancakePair", "billing.CreateWaffoPancakePair",
        "controller.SaveWaffoPancake", "billing.SaveWaffoPancake",
        "controller.CreateWaffoPancakeSubscriptionProduct", "billing.CreateWaffoPancakeSubscriptionProduct",
        "controller.ListWaffoPancakeSubscriptionProductOptions", "billing.ListWaffoPancakeSubscriptionProductOptions",
        "controller.GetSyncableChannels", "billing.GetSyncableChannels",
        "controller.FetchUpstreamRatios", "billing.FetchUpstreamRatios",
        "controller.GetAllRedemptions", "billing.GetAllRedemptions",
        "controller.SearchRedemptions", "billing.SearchRedemptions",
        "controller.GetRedemption", "billing.GetRedemption",
        "controller.AddRedemption", "billing.AddRedemption",
        "controller.UpdateRedemption", "billing.UpdateRedemption",
        "controller.DeleteInvalidRedemption", "billing.DeleteInvalidRedemption",
        "controller.DeleteRedemption", "billing.DeleteRedemption",
        "controller.GetAllTopUps", "billing.GetAllTopUps",
        "controller.AdminCompleteTopUp", "billing.AdminCompleteTopUp"
    );
    "I:/ADM/Lexvia/router/dashboard.go" = @(
        "controller.GetSubscription", "billing.GetSubscription",
        "controller.GetUsage", "billing.GetUsage"
    )
}

foreach ($file in $files.Keys) {
    $content = [System.IO.File]::ReadAllText($file)
    $pairs = $files[$file]
    for ($i = 0; $i -lt $pairs.Length; $i += 2) {
        $content = $content -replace $pairs[$i], $pairs[$i+1]
    }
    [System.IO.File]::WriteAllText($file, $content)
    Write-Host "Updated: $file"
}