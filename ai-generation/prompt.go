package ai_generation

import (
	"strings"
)

func BuildQueryImportancePrompt(content string) string {
	promptBuilder := &strings.Builder{}
	promptBuilder.WriteString("在数字1到10的范围内，1表示非常平凡的事（比如：刷牙、吃早饭），10表示非常极其深刻的事（比如：分手、考上大学），评估下面这件事的重要度，返回一个1到10的整数")
	return ""
}

func BuildMakePlanPrompt(time, agentName, agentSummary, agentStatus, topRankMemories string) string {
	promptBuilder := &strings.Builder{}
	promptBuilder.WriteString(buildAgentDescription(time, agentName, agentSummary, agentStatus, "", topRankMemories))
	promptBuilder.WriteString("为" + agentName + "制定今日计划，返回内容格式为：{'startTime': '12:0:0', 'endTIme': '12:30:00', 'content': '吃午饭'}")
	return promptBuilder.String()
}

func BuildObservationReplyPrompt(time, agentName, agentSummary, agentStatus, observation, topRankMemories string) string {
	promptBuilder := &strings.Builder{}
	promptBuilder.WriteString(buildAgentDescription(time, agentName, agentSummary, agentStatus, observation, topRankMemories))
	promptBuilder.WriteString(agentName + "是否需要对观察到的事情做出反应，如果需要则返回反应行为，格式为：{'actionType': 'talk', 'location': '超市', 'from': '李明', 'to': '张伟', 'content': '最近怎么样'}" +
		"如果不需要则返回{'actionType': 'none', 'location': '', 'from': '', 'to': '', 'content': ''}")
	return promptBuilder.String()
}

func BuildReflectionPrompt(agentName string, memoriesMap map[int64]string) string {
	promptBuilder := &strings.Builder{}
	promptBuilder.WriteString(agentName + "有如下记忆：\n")
	for _, memory := range memoriesMap {
		promptBuilder.WriteString(memory + ";")
	}
	promptBuilder.WriteString("\n")
	promptBuilder.WriteString("根据这些信息，我们可以提出的三个最突出的高层次问题是什么，返回内容用英文分号分隔")
	return promptBuilder.String()
}

func BuildAbstractMemoryPrompt(agentName string, memoryContent string) string {
	promptBuilder := &strings.Builder{}
	promptBuilder.WriteString("我们可以对" + agentName + "提出如下问题，原因是以往记忆中有如下内容：" + "\n")
	promptBuilder.WriteString(memoryContent + "\n")
	promptBuilder.WriteString("对上述记忆用一句话总结")
	return promptBuilder.String()
}

// ---------------------------------------------------internal----------------------------------------------------------
func buildAgentDescription(time, agentName, agentSummary, agentStatus, observation, topRankMemories string) string {
	descBuilder := &strings.Builder{}
	descBuilder.WriteString("当前时间：" + time + "\n")
	descBuilder.WriteString(agentSummary)
	descBuilder.WriteString(agentName + "当前状态：" + agentStatus + "\n")
	if observation != "" {
		descBuilder.WriteString(agentName + "观察到：" + observation + "\n")
	}
	descBuilder.WriteString("相关记忆如下：" + topRankMemories + "\n")
	return descBuilder.String()
}

//// BuildInitiateChatPrompt 发起聊天prompt
//func BuildInitiateChatPrompt(agentName, agentStatus, observation string, relevantMemories []string) string {
//	promptBuilder := &strings.Builder{}
//	promptBuilder.WriteString("现在时间是：")
//	promptBuilder.WriteString(time.Now().String())
//	promptBuilder.WriteString("\n")
//	promptBuilder.WriteString(agentName)
//	promptBuilder.WriteString("的当前状态为：")
//	promptBuilder.WriteString(agentStatus)
//	promptBuilder.WriteString("\n")
//	promptBuilder.WriteString("观察到的内容为：")
//	promptBuilder.WriteString(observation)
//	promptBuilder.WriteString("\n")
//	promptBuilder.WriteString(agentName)
//	promptBuilder.WriteString("记忆中的相关内容如下：")
//	for _, memory := range relevantMemories {
//		promptBuilder.WriteString(memory)
//		promptBuilder.WriteString("。")
//	}
//	promptBuilder.WriteString("\n")
//	promptBuilder.WriteString("他会说什么")
//
//	return promptBuilder.String()
//}
//
//// BuildContinueChatPrompt 回复聊天prompt
//func BuildContinueChatPrompt(agentName, agentStatus, observation string, relevantMemories []string, chatHistories []string) string {
//	promptBuilder := &strings.Builder{}
//	promptBuilder.WriteString("现在时间是：")
//	promptBuilder.WriteString(time.Now().String())
//	promptBuilder.WriteString("\n")
//	promptBuilder.WriteString(agentName)
//	promptBuilder.WriteString("的当前状态为：")
//	promptBuilder.WriteString(agentStatus)
//	promptBuilder.WriteString("\n")
//	promptBuilder.WriteString("观察到的内容为：")
//	promptBuilder.WriteString(observation)
//	promptBuilder.WriteString("\n")
//	promptBuilder.WriteString(agentName)
//	promptBuilder.WriteString("记忆中的相关内容如下：")
//	for _, memory := range relevantMemories {
//		promptBuilder.WriteString(memory)
//		promptBuilder.WriteString("。")
//	}
//	promptBuilder.WriteString("\n")
//	promptBuilder.WriteString("下面是他们的对话历史记录：")
//	for _, history := range chatHistories {
//		promptBuilder.WriteString(history)
//		promptBuilder.WriteString("。")
//	}
//	promptBuilder.WriteString("\n")
//	promptBuilder.WriteString("他该回应什么")
//	return promptBuilder.String()
//}
