package main

import (
	"Nvwa/common"
	"Nvwa/logger"
	"Nvwa/util"
	"Nvwa/web"
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestBuildPrompt(t *testing.T) {
	logger.InitNvwaLogger("info")
	req := &web.MakePlanRequest{
		AgentId:     0,
		CurrentTime: time.Now().String(),
	}
	data, _ := json.Marshal(req)
	res := util.HttpPost("http://127.0.0.1:8888/MakePlan", "", bytes.NewReader(data), false)
	fmt.Printf("结果为%s", string(res))
}

func TestNewObservation(t *testing.T) {
	logger.InitNvwaLogger("info")
	req := &web.NewObservationRequest{
		Who:          0,
		CurrentTime:  "18:15:00",
		At:           "李明的家",
		PeopleNearBy: []string{"王丽"},
	}
	data, _ := json.Marshal(req)
	res := util.HttpPost("http://127.0.0.1:8888/NewObservation", "", bytes.NewReader(data), false)
	fmt.Printf("结果为%s", string(res))
}

func TestJsonUnMarshal(t *testing.T) {
	var arr []*common.AgentPlan
	arr = append(arr, &common.AgentPlan{
		StartTime: "111",
		EndTime:   "222",
		Place:     0,
		Content:   "333",
	})
	arr = append(arr, &common.AgentPlan{
		StartTime: "444",
		EndTime:   "555",
		Place:     0,
		Content:   "666",
	})
	jsonData, _ := json.Marshal(arr)
	fmt.Printf("%s", string(jsonData))
}
