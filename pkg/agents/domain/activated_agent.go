package agents

type ActivatedAgent struct {
	AgentId string `bson:"agentId" json:"agent_id"`
	UserId  string `bson:"userId" json:"user_id"`
}

type ActivatedAgentResponse struct {
	AgentId string `bson:"agentId" json:"agentId"`
}
