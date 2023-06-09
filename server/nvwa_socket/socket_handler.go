package nvwa_socket

import (
	jsonCodec "Nvwa/codec/json-codec"
	"Nvwa/game"
	"Nvwa/logger"
	"Nvwa/session"
	socketRoute "Nvwa/socket-route"
	"reflect"
)

var (
	connectionChan chan Connection
)

func InitSocketHandler() {
	connectionChan = make(chan Connection)
	handleSocket()
}

func handleSocket() {
	for connection := range connectionChan {
		go handleMessage(connection)
	}
}

type Connection interface {
	GetNextMessage() (b []byte, err error)
	Close()
}

var (
	nvwaCodec       = &jsonCodec.JsonCodec{}
	controllerValue = reflect.TypeOf(game.Controller{})
)

type Service struct {
	playerSession *session.Session
}

func handleMessage(conn Connection) {
	defer func() {
		conn.Close()
	}()
	for {
		bytes, err := conn.GetNextMessage()
		if err != nil {
			logger.NvwaLog.Errorf("GetNextMessage error: %v", err)
			return
		}
		route, msg, err := nvwaCodec.Decode(bytes)
		if err != nil || msg == nil {
			logger.NvwaLog.Errorf("Message decode error or msg nil, %v", err)
			continue
		}
		methodName, ok := socketRoute.RouteMap[route]
		if !ok {
			logger.NvwaLog.Errorf("No method match socket-route=%d", route)
			continue
		}
		method, ok := controllerValue.MethodByName(methodName)
		if !ok {
			logger.NvwaLog.Errorf("Method %s not found", methodName)
			continue
		}
		arg := reflect.ValueOf(msg)
		method.Func.Call([]reflect.Value{arg})
	}
}
