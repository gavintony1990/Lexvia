package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeAnalyticsID(t *testing.T) {
	assert.Equal(t, "G-ABC_123", sanitizeAnalyticsID(" G-ABC_123 "))
	assert.Empty(t, sanitizeAnalyticsID(`x');alert(1);//`))
	assert.Empty(t, sanitizeAnalyticsID("line\nbreak"))
}

func TestSafeAnalyticsURL(t *testing.T) {
	value, ok := safeAnalyticsURL("https://analytics.example.com/script.js")
	assert.True(t, ok)
	assert.Equal(t, "https://analytics.example.com/script.js", value)

	_, ok = safeAnalyticsURL("javascript:alert(1)")
	assert.False(t, ok)
	_, ok = safeAnalyticsURL("//example.com/script.js")
	assert.False(t, ok)
}
