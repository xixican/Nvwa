package web

import (
	"Nvwa/game/spirit/agent"
	"github.com/gin-gonic/gin"
)

type AgentService struct {
}

func (s *AgentService) MakePlan(ctx *gin.Context) {
	var request MakePlanRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		responseFailedWithMessage(err.Error(), ctx)
	}
	agentId := request.AgentId
	currentTime := request.CurrentTime
	var plans interface{}
	plans, err = agent.MakePlan(agentId, currentTime)
	if err != nil {
		responseFailedWithMessage(err.Error(), ctx)
	}
	responseSuccessWithData(plans, ctx)
}

func (s *AgentService) SetStatus(ctx *gin.Context) {
	var request SetStatusRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		responseFailedWithMessage(err.Error(), ctx)
	}
	agentId := request.AgentId
	status := request.Status
	err = agent.SetStatus(agentId, status)
	if err != nil {
		responseFailedWithMessage(err.Error(), ctx)
	}
	responseSuccessWithData(nil, ctx)
}

func (s *AgentService) NewObservation(ctx *gin.Context) {
	var request NewObservationRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		responseFailedWithMessage(err.Error(), ctx)
	}
	agentId := request.Who
	currentTime := request.CurrentTime
	at := request.At
	peopleNearBy := request.PeopleNearBy
	var actions interface{}
	actions, err = agent.NewObservation(agentId, currentTime, at, peopleNearBy)
	if err != nil {
		responseFailedWithMessage(err.Error(), ctx)
	}
	responseSuccessWithData(actions, ctx)
}
