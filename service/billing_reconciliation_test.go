package service

import (
	"testing"

	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAttachBillingReconciliation(t *testing.T) {
	other := map[string]interface{}{
		"admin_info": map[string]interface{}{"use_channel": []string{"7"}},
	}
	info := &relaycommon.RelayInfo{
		FinalPreConsumedQuota: 120,
		BillingSource:         BillingSourceWallet,
		RetryIndex:            1,
	}

	attachBillingReconciliation(other, info, 150, "")

	adminInfo, ok := other["admin_info"].(map[string]interface{})
	require.True(t, ok)
	assert.Equal(t, []string{"7"}, adminInfo["use_channel"])
	reconciliation, ok := adminInfo["billing_reconciliation"].(map[string]interface{})
	require.True(t, ok)
	assert.Equal(t, 120, reconciliation["pre_consumed_quota"])
	assert.Equal(t, 150, reconciliation["actual_quota"])
	assert.Equal(t, 30, reconciliation["delta_quota"])
	assert.Equal(t, "settled", reconciliation["settlement_status"])
	assert.Equal(t, 1, reconciliation["retry_index"])
}

func TestAttachBillingReconciliationSurfacesSettlementError(t *testing.T) {
	other := map[string]interface{}{}
	info := &relaycommon.RelayInfo{
		FinalPreConsumedQuota:  200,
		BillingSettlementError: "wallet update failed",
	}

	attachBillingReconciliation(other, info, 175, "")

	adminInfo := other["admin_info"].(map[string]interface{})
	reconciliation := adminInfo["billing_reconciliation"].(map[string]interface{})
	assert.Equal(t, "error", reconciliation["settlement_status"])
	assert.Equal(t, "wallet update failed", reconciliation["settlement_error"])
	assert.Equal(t, BillingSourceWallet, reconciliation["funding_source"])
}
