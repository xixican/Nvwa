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
	CurrentTime  string   `json:"currentTime"`
	Who          int      `json:"who"`
	Content      string   `json:"content"`
	At           string   `json:"at"`
	PeopleNearBy []string `json:"peopleNearBy"`
}
