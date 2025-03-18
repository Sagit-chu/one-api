package model

import "github.com/songquanpeng/one-api/relay/adaptor/openrouter"

type ResponseFormat struct {
	Type       string      `json:"type,omitempty"`
	JsonSchema *JSONSchema `json:"json_schema,omitempty"`
}

type JSONSchema struct {
	Description string                 `json:"description,omitempty"`
	Name        string                 `json:"name"`
	Schema      map[string]interface{} `json:"schema,omitempty"`
	Strict      *bool                  `json:"strict,omitempty"`
}

type Audio struct {
	Voice  string `json:"voice,omitempty"`
	Format string `json:"format,omitempty"`
}

type StreamOptions struct {
	IncludeUsage bool `json:"include_usage,omitempty"`
}

type GeneralOpenAIRequest struct {
	// https://platform.openai.com/docs/api-reference/chat/create
	Messages []Message `json:"messages,omitempty"`
	Model    string    `json:"model,omitempty"`
	Store    *bool     `json:"store,omitempty"`
	Metadata any       `json:"metadata,omitempty"`
	// FrequencyPenalty is a number between -2.0 and 2.0 that penalizes
	// new tokens based on their existing frequency in the text so far,
	// default is 0.
	FrequencyPenalty    *float64 `json:"frequency_penalty,omitempty" binding:"omitempty,min=-2,max=2"`
	LogitBias           any      `json:"logit_bias,omitempty"`
	Logprobs            *bool    `json:"logprobs,omitempty"`
	TopLogprobs         *int     `json:"top_logprobs,omitempty"`
	MaxTokens           int      `json:"max_tokens,omitempty"`
	MaxCompletionTokens *int     `json:"max_completion_tokens,omitempty"`
	// N is how many chat completion choices to generate for each input message,
	// default to 1.
	N *int `json:"n,omitempty" binding:"omitempty,min=1"`
	// ReasoningEffort constrains effort on reasoning for reasoning models, reasoning models only.
	ReasoningEffort *string `json:"reasoning_effort,omitempty" binding:"omitempty,oneof=low medium high"`
	// Modalities currently the model only programmatically allows modalities = [“text”, “audio”]
	Modalities []string `json:"modalities,omitempty"`
	Prediction any      `json:"prediction,omitempty"`
	Audio      *Audio   `json:"audio,omitempty"`
	// PresencePenalty is a number between -2.0 and 2.0 that penalizes
	// new tokens based on whether they appear in the text so far, default is 0.
	PresencePenalty  *float64        `json:"presence_penalty,omitempty" binding:"omitempty,min=-2,max=2"`
	ResponseFormat   *ResponseFormat `json:"response_format,omitempty"`
	Seed             float64         `json:"seed,omitempty"`
	ServiceTier      *string         `json:"service_tier,omitempty" binding:"omitempty,oneof=default auto"`
	Stop             any             `json:"stop,omitempty"`
	Stream           bool            `json:"stream,omitempty"`
	StreamOptions    *StreamOptions  `json:"stream_options,omitempty"`
	Temperature      *float64        `json:"temperature,omitempty"`
	TopP             *float64        `json:"top_p,omitempty"`
	TopK             int             `json:"top_k,omitempty"`
	Tools            []Tool          `json:"tools,omitempty"`
	ToolChoice       any             `json:"tool_choice,omitempty"`
	ParallelTooCalls *bool           `json:"parallel_tool_calls,omitempty"`
	User             string          `json:"user,omitempty"`
	FunctionCall     any             `json:"function_call,omitempty"`
	Functions        any             `json:"functions,omitempty"`
	// https://platform.openai.com/docs/api-reference/embeddings/create
	Input          any    `json:"input,omitempty"`
	EncodingFormat string `json:"encoding_format,omitempty"`
	Dimensions     int    `json:"dimensions,omitempty"`
	// https://platform.openai.com/docs/api-reference/images/create
	Prompt           string            `json:"prompt,omitempty"`
	Quality          *string           `json:"quality,omitempty"`
	Size             string            `json:"size,omitempty"`
	Style            *string           `json:"style,omitempty"`
	WebSearchOptions *WebSearchOptions `json:"web_search_options,omitempty"`

	// Others
	Instruction string `json:"instruction,omitempty"`
	NumCtx      int    `json:"num_ctx,omitempty"`
	// -------------------------------------
	// Openrouter
	// -------------------------------------
	Provider         *openrouter.RequestProvider `json:"provider,omitempty"`
	IncludeReasoning *bool                       `json:"include_reasoning,omitempty"`
	// -------------------------------------
	// Anthropic
	// -------------------------------------
	Thinking *Thinking `json:"thinking,omitempty"`
}

// WebSearchOptions is the tool searches the web for relevant results to use in a response.
type WebSearchOptions struct {
	// SearchContextSize is the high level guidance for the amount of context window space to use for the search,
	// default is "medium".
	SearchContextSize *string       `json:"search_context_size,omitempty" binding:"omitempty,oneof=low medium high"`
	UserLocation      *UserLocation `json:"user_location,omitempty"`
}

// UserLocation is a struct that contains the location of the user.
type UserLocation struct {
	// Approximate is the approximate location parameters for the search.
	Approximate UserLocationApproximate `json:"approximate" binding:"required"`
	// Type is the type of location approximation.
	Type string `json:"type" binding:"required,oneof=approximate"`
}

// UserLocationApproximate is a struct that contains the approximate location of the user.
type UserLocationApproximate struct {
	// City is the city of the user, e.g. San Francisco.
	City *string `json:"city,omitempty"`
	// Country is the country of the user, e.g. US.
	Country *string `json:"country,omitempty"`
	// Region is the region of the user, e.g. California.
	Region *string `json:"region,omitempty"`
	// Timezone is the IANA timezone of the user, e.g. America/Los_Angeles.
	Timezone *string `json:"timezone,omitempty"`
}

// https://docs.anthropic.com/en/docs/build-with-claude/extended-thinking#implementing-extended-thinking
type Thinking struct {
	Type         string `json:"type"`
	BudgetTokens int    `json:"budget_tokens" binding:"omitempty,min=1024"`
}

func (r GeneralOpenAIRequest) ParseInput() []string {
	if r.Input == nil {
		return nil
	}
	var input []string
	switch r.Input.(type) {
	case string:
		input = []string{r.Input.(string)}
	case []any:
		input = make([]string, 0, len(r.Input.([]any)))
		for _, item := range r.Input.([]any) {
			if str, ok := item.(string); ok {
				input = append(input, str)
			}
		}
	}
	return input
}
