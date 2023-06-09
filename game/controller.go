package game

import (
	jsonCodec "Nvwa/codec/json-codec"
	"Nvwa/logger"
)

type Controller struct {
}

func (c *Controller) JoinWorld(request jsonCodec.JoinWorldRequest) {
	logger.NvwaLog.Debugf("处理JoinWorld请求，requset=%v", request)
}
