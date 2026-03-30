package chat

import (
	"encoding/json"

	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"github.com/jackc/pgx/v5"
)

func scanChat(row pgx.Row) (chatDomain.Chat, error) {
	var c chatDomain.Chat
	var agentJSON, lastMessageJSON, participantsJSON []byte

	err := row.Scan(&c.ID, &c.Owner, &agentJSON, &c.Deactivated, &lastMessageJSON, &participantsJSON)
	if err != nil {
		return c, err
	}

	unmarshalChatJSON(&c, agentJSON, lastMessageJSON, participantsJSON)
	return c, nil
}

func scanChatRows(rows pgx.Rows) []chatDomain.Chat {
	var chats []chatDomain.Chat

	for rows.Next() {
		var c chatDomain.Chat
		var agentJSON, lastMessageJSON, participantsJSON []byte

		if err := rows.Scan(&c.ID, &c.Owner, &agentJSON, &c.Deactivated, &lastMessageJSON, &participantsJSON); err != nil {
			continue
		}

		unmarshalChatJSON(&c, agentJSON, lastMessageJSON, participantsJSON)
		chats = append(chats, c)
	}

	return chats
}

func unmarshalChatJSON(c *chatDomain.Chat, agentJSON, lastMessageJSON, participantsJSON []byte) {
	if len(agentJSON) > 0 {
		var agent chatDomain.ChatAgent
		if json.Unmarshal(agentJSON, &agent) == nil {
			c.Agent = &agent
		}
	}
	if len(lastMessageJSON) > 0 {
		var msg chatDomain.Message
		if json.Unmarshal(lastMessageJSON, &msg) == nil {
			c.LastMessage = &msg
		}
	}
	if len(participantsJSON) > 0 {
		json.Unmarshal(participantsJSON, &c.Participants)
	}
}

const chatSelectColumns = `id, owner, chat_agent, deactivated, last_message, participants`
