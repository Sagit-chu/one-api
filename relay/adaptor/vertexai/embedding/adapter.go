package vertexai

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/songquanpeng/one-api/relay/adaptor/gemini"
	"github.com/songquanpeng/one-api/relay/adaptor/openai"
	model2 "github.com/songquanpeng/one-api/relay/adaptor/vertexai/model"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/songquanpeng/one-api/relay/meta"
	"github.com/songquanpeng/one-api/relay/model"
)

var ModelList = []string{
	"textembedding-gecko-multilingual@001", "text-multilingual-embedding-002",
}

type Adaptor struct {
	model string
	task  EmbeddingTaskType
}

var _ model2.InnerAIAdapter = (*Adaptor)(nil)

func (a *Adaptor) ConvertRequest(c *gin.Context, relayMode int, request *model.GeneralOpenAIRequest) (any, error) {
	if request == nil {
		return nil, errors.New("request is nil")
	}
	inputs := request.ParseInput()
	if len(inputs) == 0 {
		return nil, errors.New("request is nil")
	}
	parts := strings.Split(request.Model, "|")
	if len(parts) >= 2 {
		a.task = EmbeddingTaskType(parts[1])
	} else {
		a.task = EmbeddingTaskTypeSemanticSimilarity
	}
	a.model = parts[0]
	instances := make([]EmbeddingInstance, len(inputs))
	for i, input := range inputs {
		instances[i] = EmbeddingInstance{
			Content:  input,
			TaskType: a.task,
		}
	}

	embeddingRequest := EmbeddingRequest{
		Instances: instances,
		Parameters: EmbeddingParams{
			OutputDimensionality: request.Dimensions,
		},
	}

	return embeddingRequest, nil
}

func (a *Adaptor) DoResponse(c *gin.Context, resp *http.Response, meta *meta.Meta) (usage *model.Usage, err *model.ErrorWithStatusCode) {
	err, usage = EmbeddingHandler(c, a.model, resp)
	return
}

func EmbeddingHandler(c *gin.Context, modelName string, resp *http.Response) (*model.ErrorWithStatusCode, *model.Usage) {
	var vertexEmbeddingResponse EmbeddingResponse
	responseBody, err := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return openai.ErrorWrapper(err, "read_response_body_failed", http.StatusInternalServerError), nil
	}
	if err != nil {
		return openai.ErrorWrapper(err, "read_response_body_failed", http.StatusInternalServerError), nil
	}
	err = resp.Body.Close()
	if err != nil {
		return openai.ErrorWrapper(err, "close_response_body_failed", http.StatusInternalServerError), nil
	}
	err = json.Unmarshal(responseBody, &vertexEmbeddingResponse)
	if err != nil {
		return openai.ErrorWrapper(err, "unmarshal_response_body_failed", http.StatusInternalServerError), nil
	}

	openaiResp := &openai.EmbeddingResponse{
		Model: modelName,
		Data:  make([]openai.EmbeddingResponseItem, 0, len(vertexEmbeddingResponse.Predictions)),
		Usage: model.Usage{
			TotalTokens: 0,
		},
	}

	for i, pred := range vertexEmbeddingResponse.Predictions {
		openaiResp.Data = append(openaiResp.Data, openai.EmbeddingResponseItem{
			Index:     i,
			Embedding: pred.Embeddings.Values,
		})
	}

	for _, pred := range vertexEmbeddingResponse.Predictions {
		openaiResp.Usage.TotalTokens += pred.Embeddings.Statistics.TokenCount
	}

	return gemini.EmbeddingResponseHandler(c, resp.StatusCode, openaiResp)
}
