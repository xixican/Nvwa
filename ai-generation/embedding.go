package ai_generation

import (
	"Nvwa/logger"
	"Nvwa/util"
	"bytes"
	"encoding/json"
)

const (
	embeddingURL   = "https://api.openai.com/v1/embeddings"
	embeddingModel = "text-embedding-ada-002"
)

func OpenAIEmbedding(input string) *EmbeddingResponse {
	request := &EmbeddingRequest{
		Model: embeddingModel,
		Input: input,
	}
	requestJson, err := json.Marshal(request)
	if err != nil {
		logger.NvwaLog.Errorf("OpenAIEmbedding json.Marshal error")
		return nil
	}
	responseData := util.HttpPost(embeddingURL, APIkey, bytes.NewBuffer(requestJson), true)
	if responseData == nil {
		logger.NvwaLog.Errorf("OpenAIEmbedding response data nil")
		return nil
	}
	logger.NvwaLog.Debugf("原始返回结果为%v", string(responseData))
	response := &EmbeddingResponse{}
	err = json.Unmarshal(responseData, response)
	if err != nil {
		logger.NvwaLog.Errorf("OpenAIEmbedding json.Marshal error")
		return nil
	}
	return response
}
