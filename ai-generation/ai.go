package ai_generation

const (
	//OpenAIAPIkey = "Bearer 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"
	OpenAIAPIkey = "Bearer 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"
	AzureAPIKey  = "8644ead1d14445e5aed55731e9e072ac"
)

type ChatRole string

const (
	User      ChatRole = "user"
	System    ChatRole = "system"
	Assistant ChatRole = "assistant"
)

type ChatMessage struct {
	Role    ChatRole `json:"role"`
	Content string   `json:"content"`
}

type ChatRequest struct {
	Model    string         `json:"model"`
	Messages []*ChatMessage `json:"messages"`
}

type ChatResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Choices []struct {
		Index   int64 `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int64 `json:"prompt_tokens"`
		CompletionTokens int64 `json:"completion_tokens"`
		TotalTokens      int64 `json:"total_tokens"`
	} `json:"usage"`
}

type EmbeddingRequest struct {
	Model string `json:"model"`
	Input string `json:"input"`
}

type EmbeddingResponse struct {
	Object string `json:"object"`
	Data   []struct {
		Object    string    `json:"object"`
		Embedding []float32 `json:"embedding"`
	} `json:"data"`
	Model string `json:"model"`
	Usage struct {
		PromptTokens int `json:"prompt_tokens"`
		TotalTokens  int `json:"total_tokens"`
	} `json:"usage"`
}

type InteractionModel struct {
	From    string `json:"from"`
	Action  string `json:"action"`
	Target  string `json:"target"`
	Info    string `json:"info"`
	Content string `json:"content"`
}
