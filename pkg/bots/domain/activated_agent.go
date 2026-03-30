package bots

type ActivatedAgent struct {
	AgentId string `json:"agent_id"`
	UserId  string `json:"user_id"`
}

type ActivatedAgentResponse struct {
	AgentId string `json:"agentId"`
}
