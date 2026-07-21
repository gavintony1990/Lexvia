package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterExcludedChannelIDs(t *testing.T) {
	channels := []int{1, 2, 3, 2}
	excluded := map[int]struct{}{2: {}}

	filtered := filterExcludedChannelIDs(channels, excluded)

	assert.Equal(t, []int{1, 3}, filtered)
	assert.Equal(t, []int{1, 2, 3, 2}, channels, "cached candidate slice must not be mutated")
}
