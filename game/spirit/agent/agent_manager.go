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
	query := "为" + agent.name + "制定今日计划"
	retrieveMemories := agent.memories
	if len(agent.memories) > RetrieveMemoryLength {
		retrieveMemories = agent.memories[len(agent.memories)-RetrieveMemoryLength:]
	}
	topRankMemories := retrieveTopRankMemories(retrieveMemories, query, RetrieveTopRank)
	memoryContentBuilder := &strings.Builder{}
	for _, memory := range topRankMemories {
		memoryContentBuilder.WriteString(memory.content + ";")
	}
	makePlanPrompt := ai_generation.BuildMakePlanPrompt(currentTime, agent.name, agent.summary.String(), agent.status, memoryContentBuilder.String())
	makePlanResponse := ai_generation.OpenAIChat(ai_generation.BuildChatMessage(makePlanPrompt))
	plans := ai_generation.ParsePlan(makePlanResponse)
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

func NewObservation(agentId int, currentTime, observationContent string) (interface{}, error) {
	agent := GetAgentById(agentId)
	if agent == nil {
		return nil, common.AgentNotExistError
	}
	retrieveMemories := agent.memories
	if len(agent.memories) > RetrieveMemoryLength {
		retrieveMemories = agent.memories[len(agent.memories)-RetrieveMemoryLength:]
	}
	topRankMemories := retrieveTopRankMemories(retrieveMemories, observationContent, RetrieveTopRank)
	memoryContentBuilder := &strings.Builder{}
	for _, memory := range topRankMemories {
		memoryContentBuilder.WriteString(memory.content + ";")
	}
	observationPrompt := ai_generation.BuildObservationReplyPrompt(currentTime, agent.name, agent.summary.String(), agent.status, observationContent, memoryContentBuilder.String())
	observationReplyResponse := ai_generation.OpenAIChat(ai_generation.BuildChatMessage(observationPrompt))
	action := ai_generation.ParseAction(observationReplyResponse)
	// 新增记忆并返回记忆的重要度
	addMemory(agent, observationContent, Observation)
	// 重要度累计,大于500进行反思
	if agent.memoriesImportance > ReflectionImportanceThreshold {
		reflect(agent)
	}
	return action, nil
}

//-----------------------------------------------internal--------------------------------------------------------------

func addMemory(agent *Agent, memoryContent string, memoryType MemoryType) int {
	// 查询重要度
	importancePrompt := ai_generation.BuildQueryImportancePrompt(memoryContent)
	queryImportanceMsg := ai_generation.BuildChatMessage(importancePrompt)
	importanceResponse := ai_generation.OpenAIChat(queryImportanceMsg)
	importance := ai_generation.ParseImportance(importanceResponse)
	// 文本嵌入为后续计算相似度
	embeddingResponse := ai_generation.OpenAIEmbedding(memoryContent)
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
	return importance
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
	reflectionResponse := ai_generation.OpenAIChat(ai_generation.BuildChatMessage(reflectionPrompt))
	reflectionQuestions := ai_generation.ParseReflectionQuestion(reflectionResponse)
	for _, question := range reflectionQuestions {
		// 对每个反思问题检索相关记忆,按retrievalScore排序并提取MemoryTopK条
		topMemories := retrieveTopRankMemories(reflectionMemories, question, ReflectionTopRank)
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
func retrieveTopRankMemories(memories []*MemoryInfo, query string, topRank int) []*MemoryInfo {
	embeddingResponse := ai_generation.OpenAIEmbedding(query)
	questionEmbedding := ai_generation.ParseEmbeddingFromResponse(embeddingResponse)
	// 对每个query检索相关记忆
	for _, memory := range memories {
		importanceScore := util.CalculateNormalizationImportanceScore(memory.importance)
		recencyScore := util.CalculateNormalizationRecencyScore(memory.lastVisitTimestamp, time.Now().Unix())
		relevanceScore := util.CalculateNormalizationRelevanceScore(questionEmbedding, memory.embedding)
		memory.retrievalScore = importanceScore + recencyScore + relevanceScore
	}
	// 按retrievalScore排序并提取topRank条
	sort.SliceStable(memories, func(i, j int) bool {
		return memories[i].retrievalScore < memories[j].retrievalScore
	})
	topMemories := memories[:topRank]
	return topMemories
}
