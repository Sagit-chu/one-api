package controller

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/songquanpeng/one-api/common/ctxkey"
	"github.com/songquanpeng/one-api/middleware"
	"github.com/songquanpeng/one-api/model"
	relay "github.com/songquanpeng/one-api/relay"
	"github.com/songquanpeng/one-api/relay/adaptor/openai"
	"github.com/songquanpeng/one-api/relay/apitype"
	"github.com/songquanpeng/one-api/relay/channeltype"
	"github.com/songquanpeng/one-api/relay/meta"
	relaymodel "github.com/songquanpeng/one-api/relay/model"
)

// https://platform.openai.com/docs/api-reference/models/list

type OpenAIModelPermission struct {
	Id                 string  `json:"id"`
	Object             string  `json:"object"`
	Created            int     `json:"created"`
	AllowCreateEngine  bool    `json:"allow_create_engine"`
	AllowSampling      bool    `json:"allow_sampling"`
	AllowLogprobs      bool    `json:"allow_logprobs"`
	AllowSearchIndices bool    `json:"allow_search_indices"`
	AllowView          bool    `json:"allow_view"`
	AllowFineTuning    bool    `json:"allow_fine_tuning"`
	Organization       string  `json:"organization"`
	Group              *string `json:"group"`
	IsBlocking         bool    `json:"is_blocking"`
}

type OpenAIModels struct {
	// Id model's name
	//
	// BUG: Different channels may have the same model name
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	// OwnedBy is the channel's adaptor name
	OwnedBy    string                  `json:"owned_by"`
	Permission []OpenAIModelPermission `json:"permission"`
	Root       string                  `json:"root"`
	Parent     *string                 `json:"parent"`
}

// BUG: 更新 custom channel 时，应该同步更新所有自定义的 models 到 allModels
var allModels []OpenAIModels
var modelsMap map[string]OpenAIModels
var channelId2Models map[int][]string

func init() {
	var permission []OpenAIModelPermission
	permission = append(permission, OpenAIModelPermission{
		Id:                 "modelperm-LwHkVFn8AcMItP432fKKDIKJ",
		Object:             "model_permission",
		Created:            1626777600,
		AllowCreateEngine:  true,
		AllowSampling:      true,
		AllowLogprobs:      true,
		AllowSearchIndices: false,
		AllowView:          true,
		AllowFineTuning:    false,
		Organization:       "*",
		Group:              nil,
		IsBlocking:         false,
	})
	// https://platform.openai.com/docs/models/model-endpoint-compatibility
	for i := 0; i < apitype.Dummy; i++ {
		if i == apitype.AIProxyLibrary {
			continue
		}
		adaptor := relay.GetAdaptor(i)
		channelName := adaptor.GetChannelName()
		modelNames := adaptor.GetModelList()
		for _, modelName := range modelNames {
			allModels = append(allModels, OpenAIModels{
				Id:         modelName,
				Object:     "model",
				Created:    1626777600,
				OwnedBy:    channelName,
				Permission: permission,
				Root:       modelName,
				Parent:     nil,
			})
		}
	}
	for _, channelType := range openai.CompatibleChannels {
		if channelType == channeltype.Azure {
			continue
		}
		channelName, channelModelList := openai.GetCompatibleChannelMeta(channelType)
		for _, modelName := range channelModelList {
			allModels = append(allModels, OpenAIModels{
				Id:         modelName,
				Object:     "model",
				Created:    1626777600,
				OwnedBy:    channelName,
				Permission: permission,
				Root:       modelName,
				Parent:     nil,
			})
		}
	}
	modelsMap = make(map[string]OpenAIModels)
	for _, model := range allModels {
		modelsMap[model.Id] = model
	}
	channelId2Models = make(map[int][]string)
	for i := 1; i < channeltype.Dummy; i++ {
		adaptor := relay.GetAdaptor(channeltype.ToAPIType(i))
		meta := &meta.Meta{
			ChannelType: i,
		}
		adaptor.Init(meta)
		channelId2Models[i] = adaptor.GetModelList()
	}
}

func DashboardListModels(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    channelId2Models,
	})
}

func ListAllModels(c *gin.Context) {
	c.JSON(200, gin.H{
		"object": "list",
		"data":   allModels,
	})
}

// ListModels lists all models available to the user.
func ListModels(c *gin.Context) {
	userId := c.GetInt(ctxkey.Id)
	userGroup, err := model.CacheGetUserGroup(userId)
	if err != nil {
		middleware.AbortWithError(c, http.StatusBadRequest, err)
		return
	}

	// Get available models with their channel names
	availableAbilities, err := model.CacheGetGroupModelsV2(c.Request.Context(), userGroup)
	if err != nil {
		middleware.AbortWithError(c, http.StatusBadRequest, err)
		return
	}

	// Create a map for quick lookup of enabled model+channel combinations
	// Only store the exact model:channel combinations from abilities
	abilityMap := make(map[string]bool)
	for _, ability := range availableAbilities {
		// Store as "modelName:channelName" for exact matching
		adaptor := relay.GetAdaptor(channeltype.ToAPIType(ability.ChannelType))
		key := ability.Model + ":" + adaptor.GetChannelName()
		abilityMap[key] = true

		// for custom channels, store the model name only
		if ability.ChannelType == channeltype.Custom {
			abilityMap[ability.Model] = true
		}
	}

	// Filter models that match user's abilities with EXACT model+channel matches
	userAvailableModels := make([]OpenAIModels, 0)

	// Only include models that have a matching model+channel combination
	for _, model := range allModels {
		key := model.Id + ":" + model.OwnedBy
		if abilityMap[key] {
			userAvailableModels = append(userAvailableModels, model)
		} else if abilityMap[model.Id] {
			// for custom channels, store the model name only
			userAvailableModels = append(userAvailableModels, model)
		}
	}

	// Sort models alphabetically for consistent presentation
	sort.Slice(userAvailableModels, func(i, j int) bool {
		return userAvailableModels[i].Id < userAvailableModels[j].Id
	})

	c.JSON(200, gin.H{
		"object": "list",
		"data":   userAvailableModels,
	})
}
func RetrieveModel(c *gin.Context) {
	modelId := c.Param("model")
	if model, ok := modelsMap[modelId]; ok {
		c.JSON(200, model)
	} else {
		Error := relaymodel.Error{
			Message: fmt.Sprintf("The model '%s' does not exist", modelId),
			Type:    "invalid_request_error",
			Param:   "model",
			Code:    "model_not_found",
		}
		c.JSON(200, gin.H{
			"error": Error,
		})
	}
}

func GetUserAvailableModels(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.GetInt(ctxkey.Id)
	userGroup, err := model.CacheGetUserGroup(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	models, err := model.CacheGetGroupModelsV2(ctx, userGroup)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	var modelNames []string
	modelsMap := map[string]bool{}
	for _, model := range models {
		modelsMap[model.Model] = true
	}
	for modelName := range modelsMap {
		modelNames = append(modelNames, modelName)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    modelNames,
	})
	return
}
