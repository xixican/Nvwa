package web

import "github.com/gin-gonic/gin"

type AgentRoute struct {
	AgentService
}

func (a *AgentRoute) InitAgentRoute(engine *gin.Engine) {
	agentGroup := engine.Group("")

	agentGroup.POST("MakePlan", a.AgentService.MakePlan)
	agentGroup.POST("SetStatus", a.SetStatus)
	agentGroup.POST("NewObservation", a.NewObservation)
}
