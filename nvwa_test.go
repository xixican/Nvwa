package main

import (
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
	fmt.Printf("结果为%v", res)
}
