package ai_generation

import (
	"strings"
)

func BuildQueryImportancePrompt(content string) string {
	promptBuilder := &strings.Builder{}
	promptBuilder.WriteString("在数字1到10的范围内，1表示非常平凡的事（比如：刷牙、吃早饭），10表示非常极其深刻的事（比如：分手、考上大学），评估下面这件事的重要度，事件：" + content + "，返回一个1到10的整数,不需要多余信息和标点符号")
	return promptBuilder.String()
}

func BuildMakePlanPrompt(time, agentName, agentSummary, agentStatus, topRankMemories string) string {
	promptBuilder := &strings.Builder{}
	promptBuilder.WriteString(buildAgentDescription(time, agentName, agentSummary, agentStatus, "", topRankMemories))
	promptBuilder.WriteString("以上为" + agentName + "的人物描述和记忆，为人物制定未来24小时的计划。返回指定json格式，不要有多余内容。")
	return promptBuilder.String()
}

func BuildObservationReplyPrompt(time, agentName, agentSummary, agentStatus, observation, topRankMemories string) string {
	promptBuilder := &strings.Builder{}
	promptBuilder.WriteString(buildAgentDescription(time, agentName, agentSummary, agentStatus, observation, topRankMemories))
	promptBuilder.WriteString("根据上述人物描述和记忆内容，判断" + agentName + "是否需要对观察到的事情做出反应，如果需要则返回一个反应行为，actionType参数只有1，2和3，1表示移动，2表示对话聊天，3表示其他，其他时优先做自己的当天计划，：" +
		"{\"actionType\": 1, \"targetLocation\":10, \"emoji\": \"🚶\"} {\"actionType\": 2, \"talkTo\":\"李梦\",\"content\":\"你今天要去学校吗？\",\"emoji\": \"😊\"}" +
		"{\"actionType\": 3,\"content\":\"在家里看书\", \"emoji\": \"📖\"}" +
		"尽可能做出反应，如果不需要做出反应则返回移动行为，只返回json数据")
	//"如果不需要做出反应则返回{\"actionType\": 3, \"content\":\"保持\", \"emoji\": \"😐\"}，只返回json数据")
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
