package json_codec

import (
	"Nvwa/logger"
	"Nvwa/socket-route"
	"encoding/json"
	"errors"
	"fmt"
)

type JsonCodec struct {
}

func (j *JsonCodec) Decode(b []byte) (int, interface{}, error) {
	jsonMsg := &JsonMessage{}
	var err error
	err = json.Unmarshal(b, jsonMsg)
	if err != nil {
		logger.NvwaLog.Errorf("json.Unmarshal JsonMessage error: %v", err)
		return 0, nil, err
	}
	var msg interface{}
	switch jsonMsg.Route {
	case socket_route.JoinWorldReq:
		msg = &JoinWorldRequest{}
		err = json.Unmarshal(jsonMsg.Message, msg)
	default:
		err = errors.New(fmt.Sprintf("socket-route=%d undefined", jsonMsg.Route))
	}
	return jsonMsg.Route, msg, err
}

func (j *JsonCodec) Encode() []byte {
	return nil
}

type JsonMessage struct {
	Route   int    `json:"socket-route"`
	Message []byte `json:"message"`
}
