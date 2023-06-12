package agent

import (
	ai_generation "Nvwa/ai-generation"
	"Nvwa/common"
	"Nvwa/logger"
	"Nvwa/util"
	"sort"
	"strings"
	"time"
)

const (
	RetrieveMemoryLength          = 100 // 检索记忆长度
	RetrieveTopRank               = 10  // 检索记忆TopRank
	ReflectionImportanceThreshold = 500 // 执行反思需要的重要度阈值
	ReflectionTopRank             = 5
)

func MakePlan(agentId int, currentTime string) (interface{}, error) {
	agent := GetAgentById(agentId)
	if agent == nil {
		return nil, common.AgentNotExistError
	}
	startTime := time.Now()
	query := "为" + agent.name + "制定今日计划"
	// embedding 1rate
	//queryEmbeddingResponse := ai_generation.OpenAIEmbedding(query)
	queryEmbeddingResponse := ai_generation.AzureEmbedding(query)
	queryEmbedding := ai_generation.ParseEmbeddingFromResponse(queryEmbeddingResponse)
	logger.NvwaLog.Debugf(query)
	retrieveMemories := agent.memories
	if len(agent.memories) > RetrieveMemoryLength {
		retrieveMemories = agent.memories[len(agent.memories)-RetrieveMemoryLength:]
	}
	topRankMemories := retrieveTopRankMemories(retrieveMemories, queryEmbedding, RetrieveTopRank)
	memoryContentBuilder := &strings.Builder{}
	for _, memory := range topRankMemories {
		memoryContentBuilder.WriteString(memory.content + ";")
	}
	makePlanPrompt := ai_generation.BuildMakePlanPrompt(currentTime, agent.name, agent.summary.String(), agent.status, memoryContentBuilder.String())
	logger.NvwaLog.Debugf("make plan的prompt为：%s", makePlanPrompt)
	exampleUserPrompt := "当前时间：2023-06-10 8:00\n以下是关于李明的描述\n基本信息：年龄45岁，职业是医生，妻子是王丽，女儿是李梦，兴趣爱好是钓鱼和阅读医学杂志\n性格：内敛、负责、善良\n背景故事：李明是一个在当地医院工作的医生，他曾经是一名优秀的外科医生，但因为一次手术失误，他的信心受到了打击。现在他主要负责门诊工作。他和妻子关系融洽，但对女儿的教育问题有些担忧。\n李明当前状态：\n相关记忆如下：\n为李明制定今日计划,返回json格式的数据"
	assistantPrompt := "[{\"startTime\": \"8:00\",\"endTime\": \"12:00\",\"place\": 10,\"content\":\"上午工作时间，处理门诊病人\"},{\"startTime\": \"12:00\", \"endTime\": \"14:00\", \"place\": 5, \"content\": \"和朋友吃饭\"}]"
	// chat 1rate
	//makePlanResponse := ai_generation.OpenAIChat(ai_generation.BuildChatMessageWithSystemAssistant(WorldPrompt, exampleUserPrompt, assistantPrompt, makePlanPrompt))
	makePlanResponse := ai_generation.AzureChat(ai_generation.BuildChatMessageWithSystemAssistant(WorldPrompt, exampleUserPrompt, assistantPrompt, makePlanPrompt))
	plans := ai_generation.ParsePlan(makePlanResponse)
	if plans == "" {
		return nil, common.MakePlanError
	}
	// 将plan作为记忆存储
	planContent := agent.name + "在" + currentTime + "计划" + plans
	// chat 1rate, embedding 1rate
	addMemory(agent, planContent, Plan)
	logger.NvwaLog.Debugf(agent.name+"的今日计划为：%+v, 耗时 %s", plans, time.Since(startTime).String())
	return plans, nil
}

func SetStatus(agentId int, status string) error {
	agent := GetAgentById(agentId)
	if agent == nil {
		return common.AgentNotExistError
	}
	agent.status = status
	return nil
}

func NewObservation(agentId int, currentTime, at string, peopleNearBy []string) (interface{}, error) {
	agent := GetAgentById(agentId)
	if agent == nil {
		return nil, common.AgentNotExistError
	}
	if len(peopleNearBy) == 0 {
		return nil, nil
	}
	startTime := time.Now()
	contentBuilder := &strings.Builder{}
	for _, p := range peopleNearBy {
		contentBuilder.WriteString(p)
	}
	content := agent.name + "在" + at + "看到" + contentBuilder.String()
	logger.NvwaLog.Debugf("%s NewObservation, content=%s", agent.name, content)
	retrieveMemories := agent.memories
	if len(agent.memories) > RetrieveMemoryLength {
		retrieveMemories = agent.memories[len(agent.memories)-RetrieveMemoryLength:]
	}
	// 将observation作为记忆存储
	observationContent := currentTime + content
	// chat 1rate, embedding 1rate
	observationMemory := addMemory(agent, observationContent, Observation)
	topRankMemories := retrieveTopRankMemories(retrieveMemories, observationMemory.embedding, RetrieveTopRank)
	memoryContentBuilder := &strings.Builder{}
	for _, memory := range topRankMemories {
		memoryContentBuilder.WriteString(memory.content + ";")
	}
	// chat 1rate
	observationPrompt := ai_generation.BuildObservationReplyPrompt(currentTime, agent.name, agent.summary.String(), agent.status, observationContent, memoryContentBuilder.String())
	logger.NvwaLog.Debugf("%s NewObservation prompt=%s", agent.name, observationPrompt)
	//observationReplyResponse := ai_generation.OpenAIChat(ai_generation.BuildChatMessage(observationPrompt))
	observationReplyResponse := ai_generation.AzureChat(ai_generation.BuildChatMessage(observationPrompt))
	action := ai_generation.ParseAction(observationReplyResponse)
	// 将action作为记忆存储 chat 1rate
	actionContent := agent.name + "在" + currentTime + "执行" + action
	addMemory(agent, actionContent, Action)
	// 重要度累计,大于500进行反思
	if agent.memoriesImportance > ReflectionImportanceThreshold {
		logger.NvwaLog.Debugf("%s 执行reflect, 当前记忆条数为%d", agent.name, len(agent.memories))
		reflect(agent)
	}
	logger.NvwaLog.Infof("%s NewObservation 的response:%s, 耗时%s", agent.name, action, time.Since(startTime).String())
	// format
	//s1 := strings.ReplaceAll(action, "1", "move")
	//s2 := strings.ReplaceAll(s1, "2", "chat")
	//s3 := strings.ReplaceAll(s2, "3", "idle")
	return action, nil
}

//-----------------------------------------------internal--------------------------------------------------------------

func addMemory(agent *Agent, memoryContent string, memoryType MemoryType) *MemoryInfo {
	// 查询重要度
	importancePrompt := ai_generation.BuildQueryImportancePrompt(memoryContent)
	logger.NvwaLog.Debugf("查询重要度的prompt:%s", importancePrompt)
	queryImportanceMsg := ai_generation.BuildChatMessage(importancePrompt)
	//importanceResponse := ai_generation.OpenAIChat(queryImportanceMsg)
	importanceResponse := ai_generation.AzureChat(queryImportanceMsg)
	importance := ai_generation.ParseImportance(importanceResponse)
	// 文本嵌入为后续计算相似度
	//embeddingResponse := ai_generation.OpenAIEmbedding(memoryContent)
	embeddingResponse := ai_generation.AzureEmbedding(memoryContent)
	embedding := ai_generation.ParseEmbeddingFromResponse(embeddingResponse)
	// 新增记忆
	agent.memoryTotalCount += 1
	newMemory := &MemoryInfo{
		id:                 agent.memoryTotalCount,
		memoryType:         memoryType,
		content:            memoryContent,
		importance:         importance,
		embedding:          embedding,
		creatTimestamp:     time.Now().Unix(),
		lastVisitTimestamp: time.Now().Unix(),
	}
	agent.memories = append(agent.memories, newMemory)
	// 重要度总计
	agent.memoriesImportance += importance
	return newMemory
}

func reflect(agent *Agent) {
	var reflectionMemories []*MemoryInfo
	// 最大100条
	if len(reflectionMemories) > RetrieveMemoryLength {
		reflectionMemories = reflectionMemories[len(reflectionMemories)-RetrieveMemoryLength:]
		logger.NvwaLog.Infof("agent %s do reflect, memory length = %d ", agent.name, len(reflectionMemories))
	}
	// reflection
	memoryContentMap := make(map[int64]string, len(reflectionMemories))
	for _, memory := range reflectionMemories {
		memoryContentMap[memory.id] = memory.content
	}
	reflectionPrompt := ai_generation.BuildReflectionPrompt(agent.name, memoryContentMap)
	// chat 1rate
	reflectionResponse := ai_generation.OpenAIChat(ai_generation.BuildChatMessage(reflectionPrompt))
	reflectionQuestions := ai_generation.ParseReflectionQuestion(reflectionResponse)
	for _, question := range reflectionQuestions {
		// 对每个反思问题检索相关记忆,按retrievalScore排序并提取MemoryTopK条
		questionEmbeddingResponse := ai_generation.OpenAIEmbedding(question)
		questionEmbedding := ai_generation.ParseEmbeddingFromResponse(questionEmbeddingResponse)
		topMemories := retrieveTopRankMemories(reflectionMemories, questionEmbedding, ReflectionTopRank)
		// 让ChatGPT抽象成见解
		contentBuilder := &strings.Builder{}
		for _, memory := range topMemories {
			contentBuilder.WriteString(memory.content + ";")
			// 修改MemoryTopK记忆的上次访问时间
			memory.lastVisitTimestamp = time.Now().Unix()
		}
		abstractPrompt := ai_generation.BuildAbstractMemoryPrompt(agent.name, contentBuilder.String())
		abstractResponse := ai_generation.OpenAIChat(ai_generation.BuildChatMessage(abstractPrompt))
		abstractContent := ai_generation.ParseAbstractMemory(abstractResponse)
		// reflection存储为新记忆
		addMemory(agent, abstractContent, Reflection)
	}
	agent.memoriesImportance = 0
}

// 检索TopRank的记忆体
func retrieveTopRankMemories(memories []*MemoryInfo, queryEmbedding []float32, topRank int) []*MemoryInfo {
	if len(memories) == 0 {
		return nil
	}
	var topMemories []*MemoryInfo
	// 对每个query检索相关记忆
	for _, memory := range memories {
		importanceScore := util.CalculateNormalizationImportanceScore(memory.importance)
		recencyScore := util.CalculateNormalizationRecencyScore(memory.lastVisitTimestamp, time.Now().Unix())
		relevanceScore := util.CalculateNormalizationRelevanceScore(queryEmbedding, memory.embedding)
		memory.retrievalScore = importanceScore + recencyScore + relevanceScore
	}
	// 按retrievalScore排序并提取topRank条
	sort.SliceStable(memories, func(i, j int) bool {
		return memories[i].retrievalScore < memories[j].retrievalScore
	})
	retrieveCount := topRank
	if len(memories) < topRank {
		retrieveCount = len(memories)
	}
	topMemories = memories[:retrieveCount]
	return topMemories
}
