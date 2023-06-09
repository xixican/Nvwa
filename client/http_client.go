package main

import (
	"Nvwa/logger"
	"Nvwa/util"
	"Nvwa/web"
	"bytes"
	"encoding/json"
)

func main() {
	logger.InitNvwaLogger("info")
	testRecordEvent()
}

func testRecordEvent() {
	url := "http://127.0.0.1:8888/RecordEvent"
	recordEvent := &web.RecordEventRequest{
		AgentId: 1,
		Event:   "hello",
	}
	data, _ := json.Marshal(recordEvent)
	util.HttpPost(url, "", bytes.NewBuffer(data), false)
}
