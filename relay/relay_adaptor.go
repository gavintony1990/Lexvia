package relay

import (
	"strconv"

	"github.com/gavintony1990/Lexvia/constant"
	"github.com/gavintony1990/Lexvia/relay/channel"
	"github.com/gavintony1990/Lexvia/relay/channel/advancedcustom"
	"github.com/gavintony1990/Lexvia/relay/channel/ali"
	"github.com/gavintony1990/Lexvia/relay/channel/aws"
	"github.com/gavintony1990/Lexvia/relay/channel/baidu"
	"github.com/gavintony1990/Lexvia/relay/channel/baidu_v2"
	"github.com/gavintony1990/Lexvia/relay/channel/claude"
	"github.com/gavintony1990/Lexvia/relay/channel/cloudflare"
	"github.com/gavintony1990/Lexvia/relay/channel/codex"
	"github.com/gavintony1990/Lexvia/relay/channel/cohere"
	"github.com/gavintony1990/Lexvia/relay/channel/coze"
	"github.com/gavintony1990/Lexvia/relay/channel/deepseek"
	"github.com/gavintony1990/Lexvia/relay/channel/dify"
	"github.com/gavintony1990/Lexvia/relay/channel/gemini"
	"github.com/gavintony1990/Lexvia/relay/channel/jimeng"
	"github.com/gavintony1990/Lexvia/relay/channel/jina"
	"github.com/gavintony1990/Lexvia/relay/channel/minimax"
	"github.com/gavintony1990/Lexvia/relay/channel/mistral"
	"github.com/gavintony1990/Lexvia/relay/channel/mokaai"
	"github.com/gavintony1990/Lexvia/relay/channel/moonshot"
	"github.com/gavintony1990/Lexvia/relay/channel/ollama"
	"github.com/gavintony1990/Lexvia/relay/channel/openai"
	"github.com/gavintony1990/Lexvia/relay/channel/palm"
	"github.com/gavintony1990/Lexvia/relay/channel/perplexity"
	"github.com/gavintony1990/Lexvia/relay/channel/replicate"
	"github.com/gavintony1990/Lexvia/relay/channel/siliconflow"
	"github.com/gavintony1990/Lexvia/relay/channel/submodel"
	taskali "github.com/gavintony1990/Lexvia/relay/channel/task/ali"
	taskdoubao "github.com/gavintony1990/Lexvia/relay/channel/task/doubao"
	taskGemini "github.com/gavintony1990/Lexvia/relay/channel/task/gemini"
	"github.com/gavintony1990/Lexvia/relay/channel/task/hailuo"
	taskjimeng "github.com/gavintony1990/Lexvia/relay/channel/task/jimeng"
	"github.com/gavintony1990/Lexvia/relay/channel/task/kling"
	tasksora "github.com/gavintony1990/Lexvia/relay/channel/task/sora"
	"github.com/gavintony1990/Lexvia/relay/channel/task/suno"
	taskvertex "github.com/gavintony1990/Lexvia/relay/channel/task/vertex"
	taskVidu "github.com/gavintony1990/Lexvia/relay/channel/task/vidu"
	"github.com/gavintony1990/Lexvia/relay/channel/tencent"
	"github.com/gavintony1990/Lexvia/relay/channel/vertex"
	"github.com/gavintony1990/Lexvia/relay/channel/volcengine"
	"github.com/gavintony1990/Lexvia/relay/channel/xai"
	"github.com/gavintony1990/Lexvia/relay/channel/xunfei"
	"github.com/gavintony1990/Lexvia/relay/channel/zhipu"
	"github.com/gavintony1990/Lexvia/relay/channel/zhipu_4v"
	"github.com/gavintony1990/Lexvia/relay/registry"
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
