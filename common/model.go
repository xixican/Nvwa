package common

type AgentPlan struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Content   string `json:"content"`
}

type AgentAction struct {
	ActionType int    `json:"action_type"`
	From       string `json:"from"`
	To         string `json:"to"`
	Content    string `json:"content"`
}
