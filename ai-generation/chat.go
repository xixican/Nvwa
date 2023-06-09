package ai_generation

import (
	"Nvwa/logger"
	"Nvwa/util"
	"bytes"
	"encoding/json"
)

const (
	chatURL   = "https://api.openai.com/v1/chat/completions"
	chatModel = "gpt-3.5-turbo"
)

func OpenAIChat(message []*ChatMessage) *ChatResponse {
	request := &ChatRequest{
		Model:    chatModel,
		Messages: message,
	}
	requestJson, err := json.Marshal(request)
	if err != nil {
		logger.NvwaLog.Errorf("OpenAIChat json.Marshal error")
		return nil
	}
	responseData := util.HttpPost(chatURL, APIkey, bytes.NewBuffer(requestJson), true)
	if responseData == nil {
		logger.NvwaLog.Errorf("OpenAIChat response data nil")
		return nil
	}
	logger.NvwaLog.Debugf("原始返回结果为%v", string(responseData))
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
