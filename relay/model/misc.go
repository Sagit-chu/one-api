package model

// Usage is the token usage information returned by OpenAI API.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
	// PromptTokensDetails may be empty for some models
	PromptTokensDetails *usagePromptTokensDetails `json:"prompt_tokens_details,omitempty"`
	// CompletionTokensDetails may be empty for some models
	CompletionTokensDetails *usageCompletionTokensDetails `json:"completion_tokens_details,omitempty"`
	ServiceTier             string                        `json:"service_tier,omitempty"`
	SystemFingerprint       string                        `json:"system_fingerprint,omitempty"`

	// -------------------------------------
	// Custom fields
	// -------------------------------------
	// ToolsCost is the cost of using tools, in quota.
	ToolsCost int64 `json:"tools_cost,omitempty"`
}

type Error struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Param   string `json:"param"`
	Code    any    `json:"code"`
}

type ErrorWithStatusCode struct {
	Error
	StatusCode int `json:"status_code"`
}

type usagePromptTokensDetails struct {
	CachedTokens int `json:"cached_tokens"`
	AudioTokens  int `json:"audio_tokens"`
	// TextTokens could be zero for pure text chats
	TextTokens  int `json:"text_tokens"`
	ImageTokens int `json:"image_tokens"`
}

type usageCompletionTokensDetails struct {
	ReasoningTokens          int `json:"reasoning_tokens"`
	AudioTokens              int `json:"audio_tokens"`
	AcceptedPredictionTokens int `json:"accepted_prediction_tokens"`
	RejectedPredictionTokens int `json:"rejected_prediction_tokens"`
	// TextTokens could be zero for pure text chats
	TextTokens int `json:"text_tokens"`
}
