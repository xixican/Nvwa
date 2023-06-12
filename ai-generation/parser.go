package ai_generation

import (
	"Nvwa/logger"
	"regexp"
	"strconv"
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
	logger.NvwaLog.Debugf("importance content:%s", importanceContent)
	importance, err := strconv.Atoi(importanceContent)
	if err != nil {
		logger.NvwaLog.Errorf("ParseImportance error:%v", err)
		return importance
	}
	return importance
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

func ParsePlan(response *ChatResponse) string {
	if response == nil {
		return ""
	}
	content := response.Choices[0].Message.Content
	//logger.NvwaLog.Debugf("make plan 的返回内容为\n：%s", content)
	replaceContent := strings.ReplaceAll(content, "，", ",")
	reg := regexp.MustCompile("\\s+") // 正则表达式中的 "\s+" 匹配的是所有的空格字符，包括空格、制表符以及换行符等。
	planContent := reg.ReplaceAllString(replaceContent, "")
	//反序列化，不需要make,因为make操作被封装到Unmarshal函数
	//var plan []*common.AgentPlan
	//err := json.Unmarshal([]byte(planContent), &plan)
	//if err != nil {
	//	logger.NvwaLog.Errorf("parse plan error:%v", err)
	//}
	//return plan
	return planContent
}

func ParseAction(response *ChatResponse) string {
	if response == nil {
		return ""
	}
	content := response.Choices[0].Message.Content
	replaceContent := strings.ReplaceAll(content, "，", ",")
	reg := regexp.MustCompile("\\s+") // 正则表达式中的 "\s+" 匹配的是所有的空格字符，包括空格、制表符以及换行符等。
	actionContent := reg.ReplaceAllString(replaceContent, "")
	//action := &common.AgentAction{}
	//json.Unmarshal([]byte(actionContent), action)
	return actionContent
}
