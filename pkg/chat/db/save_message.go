package chat

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func SaveMessage(message *chatDomain.Message, chatAgentKey string, userContent string, authToken string) error {
	pool := config.GetPool()
	ctx := context.Background()

	_, err := pool.Exec(ctx,
		`INSERT INTO messages (chat_id, content, type, created_at, sender, agent_id)
		 VALUES ($1, $2, $3, $4, $5, $6)`,
		message.ChatId, message.Content, message.Type, time.Now(), message.Sender, message.AgentId,
	)

	if err != nil {
		log.Printf("Error trying to save message: %v\n", err)
	}

	data := map[string]any{
		"chatId":       message.ChatId,
		"namespaces":   []string{chatAgentKey, "PUBLIC"},
		"systemPrompt": message.SystemPrompt,
		"agentID":      message.AgentId,
		"aiModel":      "gpt-5.1",
		"botContent":   message.Content,
		"userContent":  userContent,
	}

	encoded, err := json.Marshal(data)

	if err != nil {
		log.Printf("Error trying to encode following payload: %v\n, %v\n", data, err)
		return err
	}

	baseUrl := config.GetAIBaseUrl()
	url := fmt.Sprintf("%s/save-chat", baseUrl)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(encoded))
	if err != nil {
		log.Printf("[SAVE MESSAGE] Error creating request: %v\n", err)
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authToken))

	httpClient := http.Client{
		Timeout: time.Second * 500,
	}
	resp, err := httpClient.Do(req)

	if err != nil {
		log.Printf("[SAVE MESSAGE] Error in request: %v\n", err)
		return err
	}

	body, err := io.ReadAll(resp.Body)
	log.Printf("[SAVE MESSAGE] response: %v\n", string(body))

	return nil
}
