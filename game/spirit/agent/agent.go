package agent

import (
	"Nvwa/common"
	"sync"
)

type MemoryType int

const (
	Observation MemoryType = 1
	Reflection  MemoryType = 2
	Plan        MemoryType = 3
	Action      MemoryType = 4
)

type Agent struct {
	id                 int
	name               string
	summary            *common.AgentSummary // 摘要信息描述
	status             string               // 当前状态
	memories           []*MemoryInfo        // 记忆体
	memoryLock         sync.Mutex           // 记忆体操作锁
	memoriesImportance int                  // 记忆体重要度总数
	memoryTotalCount   int64                // 记录总条数作为新增记忆的id
	EventChan          chan EventHandler
}

func NewAgent(agentId int, agentSummary *common.AgentSummary) *Agent {
	agent := &Agent{
		id:        agentId,
		name:      agentSummary.Name,
		summary:   agentSummary,
		EventChan: make(chan EventHandler),
	}
	//go agent.handle()
	return agent
}

func (a *Agent) handle() {
	for {
		select {
		case handler := <-a.EventChan:
			handler.handleEvent()
		}
	}
}

type MemoryInfo struct {
	id                 int64 // 唯一Id
	memoryType         MemoryType
	content            string    // 记忆内容
	importance         int       // 重要度
	embedding          []float32 // 文本嵌入向量
	creatTimestamp     int64     // 创建时间
	lastVisitTimestamp int64     // 上次访问时间

	//relevance int // 相似度
	//recency   int // 时间相近度

	retrievalScore float64 //检索分数(importance+relevance+recency)
}

type EventHandler interface {
	handleEvent()
}
