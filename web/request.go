package web

type MakePlanRequest struct {
	AgentId     int    `json:"agentId"`
	CurrentTime string `json:"currentTime"`
}

type SetStatusRequest struct {
	AgentId int    `json:"agentId"`
	Status  string `json:"status"`
}

type NewObservationRequest struct {
	AgentId     int    `json:"agentId"`
	CurrentTime string `json:"currentTime"`
	Observation string `json:"observation"`
}
