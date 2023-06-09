package ai_generation

import (
	"Nvwa/common"
	"Nvwa/logger"
	"encoding/json"
	"strings"
)

//--------------------------------------------ParseResponse----------------------------------------------------------

func ParseEmbeddingFromResponse(response *EmbeddingResponse) []float32 {
	if response == nil {
		return nil
	}
	return response.Data[0].Embedding
}

func ParseImportance(response *ChatResponse) int {
	if response == nil {
		return 0
	}
	importanceContent := response.Choices[0].Message.Content
	model := &ImportanceModel{}
	err := json.Unmarshal([]byte(importanceContent), model)
	if err != nil {
		logger.NvwaLog.Errorf("解析importance错误， err=%s", err.Error())
		return 0
	}
	return model.Importance
}

func ParseReflectionQuestion(response *ChatResponse) []string {
	if response == nil {
		return nil
	}
	questionContent := response.Choices[0].Message.Content
	return strings.Split(strings.TrimSpace(questionContent), ";")
}

func ParseAbstractMemory(response *ChatResponse) string {
	if response == nil {
		return ""
	}
	return response.Choices[0].Message.Content
}

func ParsePlan(response *ChatResponse) []*common.AgentPlan {
	if response == nil {
		return nil
	}
	planContent := strings.TrimSpace(response.Choices[0].Message.Content)
	plan := make([]*common.AgentPlan, 8)
	json.Unmarshal([]byte(planContent), plan)
	return plan
}

func ParseAction(response *ChatResponse) *common.AgentAction {
	if response == nil {
		return nil
	}
	actionContent := strings.TrimSpace(response.Choices[0].Message.Content)
	action := &common.AgentAction{}
	json.Unmarshal([]byte(actionContent), action)
	return action
}
