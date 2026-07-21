package doubao

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/QuantumNous/new-api/constant"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateNativeSeedanceRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	body := `{
		"model":"doubao-seedance-2-0-260128",
		"content":[
			{"type":"video_url","video_url":{"url":"https://example.com/input.mp4"}},
			{"type":"text","text":"A cinematic sunrise"}
		],
		"resolution":"1080p",
		"duration":8,
		"generate_audio":true,
		"safety_identifier":"tenant-123",
		"priority":2
	}`
	request := httptest.NewRequest(http.MethodPost, "/api/v3/contents/generations/tasks", strings.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	context.Request = request
	info := &relaycommon.RelayInfo{TaskRelayInfo: &relaycommon.TaskRelayInfo{}}

	taskErr := (&TaskAdaptor{}).ValidateRequestAndSetAction(context, info)

	require.Nil(t, taskErr)
	req, err := relaycommon.GetTaskRequest(context)
	require.NoError(t, err)
	assert.Equal(t, "doubao-seedance-2-0-260128", req.Model)
	assert.Equal(t, "A cinematic sunrise", req.Prompt)
	assert.Equal(t, 8, req.Duration)
	assert.Equal(t, constant.TaskActionGenerate, info.Action)
	assert.Equal(t, "1080p", req.Metadata["resolution"])
	assert.True(t, hasVideoInMetadata(req.Metadata))
}

func TestValidateNativeSeedanceRequestRejectsMissingTextAndUnboundedDuration(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tests := []struct {
		name string
		body string
		code string
	}{
		{
			name: "missing text content",
			body: `{"model":"doubao-seedance-2-0-260128","content":[{"type":"image_url","image_url":{"url":"https://example.com/a.png"}}]}`,
			code: "invalid_request",
		},
		{
			name: "duration exceeds billing bound",
			body: `{"model":"doubao-seedance-2-0-260128","content":[{"type":"text","text":"hello"}],"duration":999999}`,
			code: "invalid_seconds",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/api/v3/contents/generations/tasks", strings.NewReader(tt.body))
			request.Header.Set("Content-Type", "application/json")
			context, _ := gin.CreateTestContext(httptest.NewRecorder())
			context.Request = request
			info := &relaycommon.RelayInfo{TaskRelayInfo: &relaycommon.TaskRelayInfo{}}

			taskErr := (&TaskAdaptor{}).ValidateRequestAndSetAction(context, info)

			require.NotNil(t, taskErr)
			assert.Equal(t, tt.code, taskErr.Code)
		})
	}
}

func TestSeedanceTwoBillingRatios(t *testing.T) {
	tests := []struct {
		name       string
		model      string
		resolution string
		hasVideo   bool
		want       float64
	}{
		{"base", "doubao-seedance-2-0-260128", "720p", false, 1},
		{"video input", "doubao-seedance-2-0-260128", "720p", true, 28.0 / 46.0},
		{"1080p", "doubao-seedance-2-0-260128", "1080P", false, 51.0 / 46.0},
		{"4k video input", "doubao-seedance-2-0-260128", " 4K ", true, 16.0 / 46.0},
		{"fast video input", "doubao-seedance-2-0-fast-260128", "720p", true, 22.0 / 37.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := GetVideoInputRatio(tt.model, tt.resolution, tt.hasVideo)
			require.True(t, ok)
			assert.InDelta(t, tt.want, got, 1e-12)
		})
	}
}
