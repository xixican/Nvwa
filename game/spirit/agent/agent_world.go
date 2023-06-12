package agent

import (
	"Nvwa/common"
	"Nvwa/logger"
	"encoding/json"
	"strings"
	"sync"
)

var (
	NvWaWorld     *World
	WorldPrompt   string
	WorldAgentMap sync.Map
	LocationMap   map[string]int
)

func Init() {
	// 创建世界
	NvWaWorld = &World{
		objectMap: make(map[int]*common.Object),
	}
	for _, object := range common.WorldObjects {
		NvWaWorld.objectMap[object.Id] = object
	}
	LocationMap = make(map[string]int)
	// 创建10个agent
	for i, summary := range common.Summaries {
		newAgent := NewAgent(i+1, summary)
		WorldAgentMap.Store(i, newAgent)
		// 初始化agent位置
		//LocationMap[newAgent.name] = i + 1
	}
	WorldPrompt = NvWaWorld.description()
	logger.NvwaLog.Debugf("systemt prompt:%s", WorldPrompt)
	// 生成agent的计划
	WorldAgentMap.Range(func(key, value any) bool {
		agent := value.(*Agent)
		MakePlan(agent.id, "2023-06-10 8:00")
		return true
	})
}

type World struct {
	objectMap map[int]*common.Object
}

func (w *World) description() string {
	objectJson, _ := json.Marshal(w.objectMap)
	locationBuilder := &strings.Builder{}
	for agentName, objectId := range LocationMap {
		locationName := common.WorldObjects[objectId].Name
		locationBuilder.WriteString(agentName + "在" + locationName + "。")
	}
	return "模拟一个现代虚拟小镇，小镇内有10个不同角色的人，小镇内有如下位置：" + string(objectJson) +
		"返回结果中如果需要place参数，需要从上述给定的位置的id中选取，如果需要targetLocation参数，需要从上述给定的位置的name中选取"
	//"10个人当前在小镇的位置描述如下:" + locationBuilder.String()
}

func GetAgentById(agentId int) *Agent {
	agentAny, ok := WorldAgentMap.Load(agentId)
	if !ok {
		return nil
	}
	return agentAny.(*Agent)
}
