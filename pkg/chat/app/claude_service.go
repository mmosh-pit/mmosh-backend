package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	domain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

// Service calls the Claude API to interpret the user's message and decide
// which Twitter action (if any) should be performed.
type Service struct {
	client anthropic.Client
}

func NewService(apiKey string) *Service {
	return &Service{
		client: anthropic.NewClient(option.WithAPIKey(apiKey)),
	}
}

func (s *Service) Chat(ctx context.Context, req *domain.ChatRequest) (*domain.ChatResponse, error) {
	messages := buildMessages(req)

	resp, err := s.client.Messages.New(ctx, anthropic.MessageNewParams{
		Model:     anthropic.ModelClaudeSonnet4_6,
		MaxTokens: 1024,
		System: []anthropic.TextBlockParam{
			{Text: buildSystemPrompt(req.PageContext)},
		},
		Messages: messages,
		Tools:    twitterTools(),
	})
	if err != nil {
		return nil, fmt.Errorf("claude api: %w", err)
	}

	return parseResponse(resp, req.ConversationID)
}

// ─── Message builder ──────────────────────────────────────────────────────────

func buildMessages(req *domain.ChatRequest) []anthropic.MessageParam {
	msgs := make([]anthropic.MessageParam, 0, len(req.History)+1)
	for _, h := range req.History {
		if h.Content != "" {
			if h.Role == "user" {
				msgs = append(msgs, anthropic.NewUserMessage(anthropic.NewTextBlock(h.Content)))
			} else {
				msgs = append(msgs, anthropic.NewAssistantMessage(anthropic.NewTextBlock(h.Content)))
			}
		}

	}
	msgs = append(msgs, anthropic.NewUserMessage(anthropic.NewTextBlock(req.Message)))
	return msgs
}

// ─── System prompt ────────────────────────────────────────────────────────────

func buildSystemPrompt(ctx domain.PageContext) string {
	var sb strings.Builder

	sb.WriteString("You are Kinship, an AI assistant that helps users interact with X (Twitter). ")
	sb.WriteString("You can perform actions on the page on behalf of the user: post tweets, reply to tweets, like tweets, retweet, and navigate to URLs on X.\n\n")
	sb.WriteString(fmt.Sprintf("Current page URL: %s\n", ctx.URL))
	sb.WriteString(fmt.Sprintf("Page type: %s\n", ctx.PageType))

	if len(ctx.Tweets) > 0 {
		sb.WriteString(fmt.Sprintf("\nVisible tweets (%d total):\n", len(ctx.Tweets)))
		for i, t := range ctx.Tweets {
			sb.WriteString(fmt.Sprintf("  [%d] @%s — %q  (url: %s)\n", i+1, t.Handle, t.Text, t.URL))
		}
	} else {
		sb.WriteString("\nNo tweets are visible on the current page.\n")
	}

	sb.WriteString("\nWhen the user asks you to perform an action, call the matching tool AND include a short, friendly text explanation of what you are doing. ")
	sb.WriteString("If no action is needed, just reply conversationally.")

	return sb.String()
}

// ─── Tool definitions ─────────────────────────────────────────────────────────

func twitterTools() []anthropic.ToolUnionParam {
	return []anthropic.ToolUnionParam{
		{OfTool: &anthropic.ToolParam{
			Name:        "post_tweet",
			Description: anthropic.String("Post a new tweet on behalf of the user."),
			InputSchema: anthropic.ToolInputSchemaParam{
				Properties: map[string]any{
					"text": map[string]any{
						"type":        "string",
						"description": "The text content of the tweet (max 280 characters).",
					},
				},
				Required: []string{"text"},
			},
		}},
		{OfTool: &anthropic.ToolParam{
			Name:        "reply_to_tweet",
			Description: anthropic.String("Reply to an existing tweet."),
			InputSchema: anthropic.ToolInputSchemaParam{
				Properties: map[string]any{
					"tweet_url": map[string]any{
						"type":        "string",
						"description": "The full URL of the tweet to reply to.",
					},
					"text": map[string]any{
						"type":        "string",
						"description": "The reply text (max 280 characters).",
					},
				},
				Required: []string{"tweet_url", "text"},
			},
		}},
		{OfTool: &anthropic.ToolParam{
			Name:        "delete_tweet",
			Description: anthropic.String("Delete a tweet."),
			InputSchema: anthropic.ToolInputSchemaParam{
				Properties: map[string]any{
					"tweet_url": map[string]any{
						"type":        "string",
						"description": "The full URL of the tweet to reply to.",
					},
					"text": map[string]any{
						"type":        "string",
						"description": "The reply text (max 280 characters).",
					},
				},
				Required: []string{"tweet_url", "text"},
			},
		}},
		{OfTool: &anthropic.ToolParam{
			Name:        "like_tweet",
			Description: anthropic.String("Like a tweet."),
			InputSchema: anthropic.ToolInputSchemaParam{
				Properties: map[string]any{
					"tweet_url": map[string]any{
						"type":        "string",
						"description": "The full URL of the tweet to like.",
					},
				},
				Required: []string{"tweet_url"},
			},
		}},
		{OfTool: &anthropic.ToolParam{
			Name:        "retweet",
			Description: anthropic.String("Retweet a tweet."),
			InputSchema: anthropic.ToolInputSchemaParam{
				Properties: map[string]any{
					"tweet_url": map[string]any{
						"type":        "string",
						"description": "The full URL of the tweet to retweet.",
					},
				},
				Required: []string{"tweet_url"},
			},
		}},
		{OfTool: &anthropic.ToolParam{
			Name:        "navigate",
			Description: anthropic.String("Navigate the browser to a URL on X.com (e.g. a user profile, notifications, or a specific tweet)."),
			InputSchema: anthropic.ToolInputSchemaParam{
				Properties: map[string]any{
					"url": map[string]any{
						"type":        "string",
						"description": "The X.com URL to navigate to.",
					},
				},
				Required: []string{"url"},
			},
		}},
	}
}

// ─── Response parser ──────────────────────────────────────────────────────────

func parseResponse(resp *anthropic.Message, conversationID string) (*domain.ChatResponse, error) {
	var textParts []string
	var action *domain.TwitterAction

	for _, block := range resp.Content {
		switch block.Type {
		case "text":
			textParts = append(textParts, block.Text)
		case "tool_use":
			// Only capture the first tool call
			if action == nil {
				a, err := toolUseToAction(block.Name, block.Input)
				if err == nil {
					action = a
				}
			}
		}
	}

	message := strings.TrimSpace(strings.Join(textParts, " "))
	if message == "" && action != nil {
		message = fmt.Sprintf("Executing: %s.", action.Type)
	}

	return &domain.ChatResponse{
		Message:        message,
		Action:         action,
		ConversationID: conversationID,
	}, nil
}

// toolUseToAction converts a Claude tool_use block into a TwitterAction.
// block.Input is json.RawMessage — the parsed tool arguments from Claude.
func toolUseToAction(name string, rawInput json.RawMessage) (*domain.TwitterAction, error) {
	var params map[string]string
	if err := json.Unmarshal(rawInput, &params); err != nil {
		return nil, fmt.Errorf("unmarshal tool input: %w", err)
	}

	action := &domain.TwitterAction{Type: name}

	text, hasText := params["text"]
	tweetURL, hasTweetURL := params["tweet_url"]
	url, hasURL := params["url"]

	if hasText || hasTweetURL || hasURL {
		action.Params = &domain.ActionParams{}
		if hasText {
			action.Params.Text = &text
		}
		if hasTweetURL {
			action.Params.TweetURL = &tweetURL
		}
		if hasURL {
			action.Params.URL = &url
		}
	}

	return action, nil
}
