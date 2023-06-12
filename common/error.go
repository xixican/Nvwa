package common

import "errors"

var (
	ParamIllegalError  = errors.New("param illegal")
	AgentNotExistError = errors.New("agent not exist")
	ResponseNilError   = errors.New("response nil")
	EmbeddingError     = errors.New("embedding error")
	MakePlanError      = errors.New("make plan error")
)
