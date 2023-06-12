package ai_generation

import (
	"Nvwa/logger"
	"Nvwa/util"
	"bytes"
	"encoding/json"
)

const (
	openAIChatURL   = "https://api.openai.com/v1/chat/completions"
	openAIChatModel = "gpt-3.5-turbo"

	azureChatURL = "https://worldsim.openai.azure.com/openai/deployments/chatgpt/chat/completions?api-version=2023-05-15"
)

func OpenAIChat(message []*ChatMessage) *ChatResponse {
	request := &ChatRequest{
		Model:    openAIChatModel,
		Messages: message,
	}
	requestJson, err := json.Marshal(request)
	if err != nil {
		logger.NvwaLog.Errorf("OpenAIChat json.Marshal error")
		return nil
	}
	responseData := util.HttpPost(openAIChatURL, OpenAIAPIkey, bytes.NewBuffer(requestJson), true)
	if responseData == nil {
		logger.NvwaLog.Errorf("OpenAIChat response data nil")
		return nil
	}
	//logger.NvwaLog.Debugf("原始返回结果为%v", string(responseData))
	response := &ChatResponse{}
	err = json.Unmarshal(responseData, response)
	if err != nil {
		logger.NvwaLog.Errorf("OpenAIChat json.Marshal error")
		return nil
	}
	return response
}

func AzureChat(message []*ChatMessage) *ChatResponse {
	request := &ChatRequest{
		Messages: message,
	}
	requestJson, err := json.Marshal(request)
	if err != nil {
		logger.NvwaLog.Errorf("OpenAIChat json.Marshal error")
		return nil
	}
	headerMap := make(map[string]string)
	headerMap["api-key"] = AzureAPIKey
	responseData := util.HttpPostSetHeader(azureChatURL, headerMap, bytes.NewBuffer(requestJson), true)
	if responseData == nil {
		logger.NvwaLog.Errorf("OpenAIChat response data nil")
		return nil
	}
	//logger.NvwaLog.Debugf("原始返回结果为%v", string(responseData))
	response := &ChatResponse{}
	err = json.Unmarshal(responseData, response)
	if err != nil {
		logger.NvwaLog.Errorf("OpenAIChat json.Marshal error")
		return nil
	}
	return response
}

func BuildChatMessage(prompt string) []*ChatMessage {
	userMessage := &ChatMessage{
		Role:    User,
		Content: prompt,
	}
	return []*ChatMessage{userMessage}
}

func BuildChatMessageWithSystem(system, prompt string) []*ChatMessage {
	systemMessage := &ChatMessage{
		Role:    System,
		Content: system,
	}
	userMessage := &ChatMessage{
		Role:    User,
		Content: prompt,
	}
	return []*ChatMessage{systemMessage, userMessage}
}

func BuildChatMessageWithSystemAssistant(system, exampleUser, assistant, user string) []*ChatMessage {
	systemMessage := &ChatMessage{
		Role:    System,
		Content: system,
	}
	exampleUserMessage := &ChatMessage{
		Role:    User,
		Content: exampleUser,
	}
	assistantMessage := &ChatMessage{
		Role:    Assistant,
		Content: assistant,
	}
	userMessage := &ChatMessage{
		Role:    User,
		Content: user,
	}
	return []*ChatMessage{systemMessage, exampleUserMessage, assistantMessage, userMessage}
}
