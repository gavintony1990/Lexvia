package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsVideoGenerationTaskPath(t *testing.T) {
	tests := []struct {
		path string
		want bool
	}{
		{"/v1/video/generations", true},
		{"/v1/video/generations/task_123", true},
		{"/api/v3/contents/generations/tasks", true},
		{"/api/v3/contents/generations/tasks/task_123", true},
		{"/v1/chat/completions", false},
		{"/api/v3/contents/generations/other", false},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			assert.Equal(t, tt.want, isVideoGenerationTaskPath(tt.path))
		})
	}
}
