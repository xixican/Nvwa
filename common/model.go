package common

type AgentPlan struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Place     int    `json:"place"`
	Content   string `json:"content"`
}

func (p *AgentPlan) String() string {
	if p.Place == 0 {
		return ""
	}
	placeName := WorldObjects[p.Place].Name
	return "从" + p.StartTime + "到" + p.EndTime + "，在" + placeName + p.Content
}
