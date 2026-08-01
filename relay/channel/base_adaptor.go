package channel

import (
	"errors"
	"io"

	"github.com/gavintony1990/Lexvia/dto"
	relaycommon "github.com/gavintony1990/Lexvia/relay/common"
	"github.com/gin-gonic/gin"
)

// BaseAdaptor provides default no-op implementations for optional Adaptor methods.
// Providers that don't support a feature can embed BaseAdaptor instead of writing
// stub methods that return "not implemented" errors.
type BaseAdaptor struct{}

func (a *BaseAdaptor) ConvertRerankRequest(c *gin.Context, relayMode int, request dto.RerankRequest) (any, error) {
	return nil, errors.New("not implemented")
}

func (a *BaseAdaptor) ConvertEmbeddingRequest(c *gin.Context, info *relaycommon.RelayInfo, request dto.EmbeddingRequest) (any, error) {
	return nil, errors.New("not implemented")
}

func (a *BaseAdaptor) ConvertAudioRequest(c *gin.Context, info *relaycommon.RelayInfo, request dto.AudioRequest) (io.Reader, error) {
	return nil, errors.New("not implemented")
}

func (a *BaseAdaptor) ConvertImageRequest(c *gin.Context, info *relaycommon.RelayInfo, request dto.ImageRequest) (any, error) {
	return nil, errors.New("not implemented")
}

func (a *BaseAdaptor) ConvertOpenAIResponsesRequest(c *gin.Context, info *relaycommon.RelayInfo, request dto.OpenAIResponsesRequest) (any, error) {
	return nil, errors.New("not implemented")
}

func (a *BaseAdaptor) ConvertClaudeRequest(c *gin.Context, info *relaycommon.RelayInfo, request *dto.ClaudeRequest) (any, error) {
	return nil, errors.New("not implemented")
}

func (a *BaseAdaptor) ConvertGeminiRequest(c *gin.Context, info *relaycommon.RelayInfo, request *dto.GeminiChatRequest) (any, error) {
	return nil, errors.New("not implemented")
}

// Note: BaseAdaptor is not a complete Adaptor implementation.
// It only provides default "not implemented" stubs for the optional methods.
// Provider adaptors embed BaseAdaptor and implement the mandatory methods
// (Init, GetRequestURL, SetupRequestHeader, ConvertOpenAIRequest, DoRequest,
// DoResponse, GetModelList, GetChannelName) themselves.