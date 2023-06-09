package agent

import (
	"Nvwa/common"
	"sync"
)

var WorldAgentMap sync.Map

func Init() {
	// 创建10个agent
	for i, summary := range common.Summaries {
		WorldAgentMap.Store(i, NewAgent(i, summary))
	}
}

func GetAgentById(agentId int) *Agent {
	agentAny, ok := WorldAgentMap.Load(agentId)
	if !ok {
		return nil
	}
	return agentAny.(*Agent)
}
