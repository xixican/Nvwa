package ai_generation

import (
	"Nvwa/logger"
	"Nvwa/util"
	"bytes"
	"encoding/json"
)

const (
	openAIEmbeddingURL   = "https://api.openai.com/v1/embeddings"
	openAIEmbeddingModel = "text-embedding-ada-002"

	azureEmbeddingURL = "https://worldsim.openai.azure.com/openai/deployments/embedding/embeddings?api-version=2022-12-01"
)

func OpenAIEmbedding(input string) *EmbeddingResponse {
	request := &EmbeddingRequest{
		Model: openAIEmbeddingModel,
		Input: input,
	}
	requestJson, err := json.Marshal(request)
	if err != nil {
		logger.NvwaLog.Errorf("OpenAIEmbedding json.Marshal error")
		return nil
	}
	responseData := util.HttpPost(openAIEmbeddingURL, OpenAIAPIkey, bytes.NewBuffer(requestJson), true)
	if responseData == nil {
		logger.NvwaLog.Errorf("OpenAIEmbedding response data nil")
		return nil
	}
	//logger.NvwaLog.Debugf("原始返回结果为%v", string(responseData))
	response := &EmbeddingResponse{}
	err = json.Unmarshal(responseData, response)
	if err != nil {
		logger.NvwaLog.Errorf("OpenAIEmbedding json.Marshal error")
		return nil
	}
	return response
}

func AzureEmbedding(input string) *EmbeddingResponse {
	request := &EmbeddingRequest{
		Input: input,
	}
	requestJson, err := json.Marshal(request)
	if err != nil {
		logger.NvwaLog.Errorf("OpenAIEmbedding json.Marshal error")
		return nil
	}
	headerMap := make(map[string]string)
	headerMap["api-key"] = AzureAPIKey
	responseData := util.HttpPostSetHeader(azureEmbeddingURL, headerMap, bytes.NewBuffer(requestJson), true)
	if responseData == nil {
		logger.NvwaLog.Errorf("OpenAIEmbedding response data nil")
		return nil
	}
	//logger.NvwaLog.Debugf("原始返回结果为%v", string(responseData))
	response := &EmbeddingResponse{}
	err = json.Unmarshal(responseData, response)
	if err != nil {
		logger.NvwaLog.Errorf("OpenAIEmbedding json.Marshal error")
		return nil
	}
	return response
}
