package relay

import (
	"strconv"

	"github.com/QuantumNous/new-api/constant"
	"github.com/QuantumNous/new-api/relay/channel"
	"github.com/QuantumNous/new-api/relay/channel/advancedcustom"
	"github.com/QuantumNous/new-api/relay/channel/ali"
	"github.com/QuantumNous/new-api/relay/channel/aws"
	"github.com/QuantumNous/new-api/relay/channel/baidu"
	"github.com/QuantumNous/new-api/relay/channel/baidu_v2"
	"github.com/QuantumNous/new-api/relay/channel/claude"
	"github.com/QuantumNous/new-api/relay/channel/cloudflare"
	"github.com/QuantumNous/new-api/relay/channel/codex"
	"github.com/QuantumNous/new-api/relay/channel/cohere"
	"github.com/QuantumNous/new-api/relay/channel/coze"
	"github.com/QuantumNous/new-api/relay/channel/deepseek"
	"github.com/QuantumNous/new-api/relay/channel/dify"
	"github.com/QuantumNous/new-api/relay/channel/gemini"
	"github.com/QuantumNous/new-api/relay/channel/jimeng"
	"github.com/QuantumNous/new-api/relay/channel/jina"
	"github.com/QuantumNous/new-api/relay/channel/minimax"
	"github.com/QuantumNous/new-api/relay/channel/mistral"
	"github.com/QuantumNous/new-api/relay/channel/mokaai"
	"github.com/QuantumNous/new-api/relay/channel/moonshot"
	"github.com/QuantumNous/new-api/relay/channel/ollama"
	"github.com/QuantumNous/new-api/relay/channel/openai"
	"github.com/QuantumNous/new-api/relay/channel/palm"
	"github.com/QuantumNous/new-api/relay/channel/perplexity"
	"github.com/QuantumNous/new-api/relay/channel/replicate"
	"github.com/QuantumNous/new-api/relay/channel/siliconflow"
	"github.com/QuantumNous/new-api/relay/channel/submodel"
	taskali "github.com/QuantumNous/new-api/relay/channel/task/ali"
	taskdoubao "github.com/QuantumNous/new-api/relay/channel/task/doubao"
	taskGemini "github.com/QuantumNous/new-api/relay/channel/task/gemini"
	"github.com/QuantumNous/new-api/relay/channel/task/hailuo"
	taskjimeng "github.com/QuantumNous/new-api/relay/channel/task/jimeng"
	"github.com/QuantumNous/new-api/relay/channel/task/kling"
	tasksora "github.com/QuantumNous/new-api/relay/channel/task/sora"
	"github.com/QuantumNous/new-api/relay/channel/task/suno"
	taskvertex "github.com/QuantumNous/new-api/relay/channel/task/vertex"
	taskVidu "github.com/QuantumNous/new-api/relay/channel/task/vidu"
	"github.com/QuantumNous/new-api/relay/channel/tencent"
	"github.com/QuantumNous/new-api/relay/channel/vertex"
	"github.com/QuantumNous/new-api/relay/channel/volcengine"
	"github.com/QuantumNous/new-api/relay/channel/xai"
	"github.com/QuantumNous/new-api/relay/channel/xunfei"
	"github.com/QuantumNous/new-api/relay/channel/zhipu"
	"github.com/QuantumNous/new-api/relay/channel/zhipu_4v"
	"github.com/QuantumNous/new-api/relay/registry"
	"github.com/gin-gonic/gin"
)

func init() {
	registry.RegisterAdaptor(constant.APITypeAli, func() channel.Adaptor { return &ali.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeAnthropic, func() channel.Adaptor { return &claude.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeBaidu, func() channel.Adaptor { return &baidu.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeGemini, func() channel.Adaptor { return &gemini.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeOpenAI, func() channel.Adaptor { return &openai.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypePaLM, func() channel.Adaptor { return &palm.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeTencent, func() channel.Adaptor { return &tencent.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeXunfei, func() channel.Adaptor { return &xunfei.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeZhipu, func() channel.Adaptor { return &zhipu.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeZhipuV4, func() channel.Adaptor { return &zhipu_4v.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeOllama, func() channel.Adaptor { return &ollama.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypePerplexity, func() channel.Adaptor { return &perplexity.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeAws, func() channel.Adaptor { return &aws.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeCohere, func() channel.Adaptor { return &cohere.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeDify, func() channel.Adaptor { return &dify.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeJina, func() channel.Adaptor { return &jina.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeCloudflare, func() channel.Adaptor { return &cloudflare.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeSiliconFlow, func() channel.Adaptor { return &siliconflow.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeVertexAi, func() channel.Adaptor { return &vertex.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeMistral, func() channel.Adaptor { return &mistral.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeDeepSeek, func() channel.Adaptor { return &deepseek.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeMokaAI, func() channel.Adaptor { return &mokaai.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeVolcEngine, func() channel.Adaptor { return &volcengine.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeBaiduV2, func() channel.Adaptor { return &baidu_v2.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeOpenRouter, func() channel.Adaptor { return &openai.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeXinference, func() channel.Adaptor { return &openai.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeXai, func() channel.Adaptor { return &xai.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeCoze, func() channel.Adaptor { return &coze.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeJimeng, func() channel.Adaptor { return &jimeng.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeMoonshot, func() channel.Adaptor { return &moonshot.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeSubmodel, func() channel.Adaptor { return &submodel.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeMiniMax, func() channel.Adaptor { return &minimax.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeReplicate, func() channel.Adaptor { return &replicate.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeCodex, func() channel.Adaptor { return &codex.Adaptor{} })
	registry.RegisterAdaptor(constant.APITypeAdvancedCustom, func() channel.Adaptor { return &advancedcustom.Adaptor{} })
}

func GetAdaptor(apiType int) channel.Adaptor {
	return registry.GetAdaptor(apiType)
}

func GetTaskPlatform(c *gin.Context) constant.TaskPlatform {
	channelType := c.GetInt("channel_type")
	if channelType > 0 {
		return constant.TaskPlatform(strconv.Itoa(channelType))
	}
	return constant.TaskPlatform(c.GetString("platform"))
}

func GetTaskAdaptor(platform constant.TaskPlatform) channel.TaskAdaptor {
	switch platform {
	//case constant.APITypeAIProxyLibrary:
	//	return &aiproxy.Adaptor{}
	case constant.TaskPlatformSuno:
		return &suno.TaskAdaptor{}
	}
	if channelType, err := strconv.ParseInt(string(platform), 10, 64); err == nil {
		switch channelType {
		case constant.ChannelTypeAli:
			return &taskali.TaskAdaptor{}
		case constant.ChannelTypeKling:
			return &kling.TaskAdaptor{}
		case constant.ChannelTypeJimeng:
			return &taskjimeng.TaskAdaptor{}
		case constant.ChannelTypeVertexAi:
			return &taskvertex.TaskAdaptor{}
		case constant.ChannelTypeVidu:
			return &taskVidu.TaskAdaptor{}
		case constant.ChannelTypeDoubaoVideo, constant.ChannelTypeVolcEngine:
			return &taskdoubao.TaskAdaptor{}
		case constant.ChannelTypeSora, constant.ChannelTypeOpenAI:
			return &tasksora.TaskAdaptor{}
		case constant.ChannelTypeGemini:
			return &taskGemini.TaskAdaptor{}
		case constant.ChannelTypeMiniMax:
			return &hailuo.TaskAdaptor{}
		}
	}
	return nil
}
