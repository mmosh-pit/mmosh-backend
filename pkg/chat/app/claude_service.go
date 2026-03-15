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
// which browser action (if any) should be performed.
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
		MaxTokens: 2048,
		System: []anthropic.TextBlockParam{
			{Text: buildSystemPrompt(req.PageContext)},
		},
		Messages: messages,
		Tools:    browserTools(),
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

	sb.WriteString("You are Kinship, an AI browser assistant that can interact with any website on behalf of the user. ")
	sb.WriteString("You can read the current page and perform actions like clicking buttons, filling forms, navigating, scrolling, and more.\n\n")

	sb.WriteString(fmt.Sprintf("Current page URL: %s\n", ctx.URL))
	if ctx.Title != "" {
		sb.WriteString(fmt.Sprintf("Page title: %s\n", ctx.Title))
	}

	if len(ctx.InteractiveElements) > 0 {
		sb.WriteString(fmt.Sprintf("\nInteractive elements on page (%d total):\n", len(ctx.InteractiveElements)))
		for i, el := range ctx.InteractiveElements {
			line := fmt.Sprintf("  [%d] type=%s selector=%q", i+1, el.Type, el.Selector)
			if el.Text != "" {
				line += fmt.Sprintf(" text=%q", el.Text)
			}
			if el.AriaLabel != "" {
				line += fmt.Sprintf(" aria-label=%q", el.AriaLabel)
			}
			if el.Placeholder != "" {
				line += fmt.Sprintf(" placeholder=%q", el.Placeholder)
			}
			if el.Value != "" {
				line += fmt.Sprintf(" value=%q", el.Value)
			}
			if el.Href != "" {
				line += fmt.Sprintf(" href=%q", el.Href)
			}
			sb.WriteString(line + "\n")
		}
	} else {
		sb.WriteString("\nNo interactive elements detected on this page.\n")
	}

	if ctx.PageText != "" {
		sb.WriteString("\nVisible page text (truncated):\n")
		text := ctx.PageText
		if len(text) > 1500 {
			text = text[:1500] + "…"
		}
		sb.WriteString(text + "\n")
	}

	sb.WriteString("\nGuidelines:\n")
	sb.WriteString("- Always use the exact selector strings from the element list above when targeting elements.\n")
	sb.WriteString("- For multi-step tasks (e.g. fill a form and submit), use the execute_steps tool to batch all actions.\n")
	sb.WriteString("- Always include a short, friendly text explanation of what you are doing.\n")
	sb.WriteString("- If you cannot complete a task because the required element is not visible, explain why and suggest what the user should do.\n")
	sb.WriteString("- If no action is needed, just reply conversationally.")

	return sb.String()
}

// ─── Tool definitions ─────────────────────────────────────────────────────────

func browserTools() []anthropic.ToolUnionParam {
	return []anthropic.ToolUnionParam{
		{OfTool: &anthropic.ToolParam{
			Name:        "click",
			Description: anthropic.String("Click an element on the page. Use 'selector' (CSS selector from the element list) or 'text' (visible text to find the element by)."),
			InputSchema: anthropic.ToolInputSchemaParam{
				Properties: map[string]any{
					"selector": map[string]any{
						"type":        "string",
						"description": "CSS selector of the element to click.",
					},
					"text": map[string]any{
						"type":        "string",
						"description": "Visible text of the element to click (used if selector is not provided).",
					},
				},
			},
		}},
		{OfTool: &anthropic.ToolParam{
			Name:        "type_text",
			Description: anthropic.String("Type text into an input, textarea, or contenteditable element."),
			InputSchema: anthropic.ToolInputSchemaParam{
				Properties: map[string]any{
					"selector": map[string]any{
						"type":        "string",
						"description": "CSS selector of the input element.",
					},
					"text": map[string]any{
						"type":        "string",
						"description": "The text to type into the element.",
					},
				},
				Required: []string{"selector", "text"},
			},
		}},
		{OfTool: &anthropic.ToolParam{
			Name:        "clear_field",
			Description: anthropic.String("Clear the value of an input or textarea element."),
			InputSchema: anthropic.ToolInputSchemaParam{
				Properties: map[string]any{
					"selector": map[string]any{
						"type":        "string",
						"description": "CSS selector of the field to clear.",
					},
				},
				Required: []string{"selector"},
			},
		}},
		{OfTool: &anthropic.ToolParam{
			Name:        "select_option",
			Description: anthropic.String("Select an option from a <select> dropdown by its text or value."),
			InputSchema: anthropic.ToolInputSchemaParam{
				Properties: map[string]any{
					"selector": map[string]any{
						"type":        "string",
						"description": "CSS selector of the <select> element.",
					},
					"value": map[string]any{
						"type":        "string",
						"description": "The option text or value to select.",
					},
				},
				Required: []string{"selector", "value"},
			},
		}},
		{OfTool: &anthropic.ToolParam{
			Name:        "scroll_to",
			Description: anthropic.String("Scroll the page to a specific element or Y position."),
			InputSchema: anthropic.ToolInputSchemaParam{
				Properties: map[string]any{
					"selector": map[string]any{
						"type":        "string",
						"description": "CSS selector of the element to scroll into view. If omitted, scrolls to the Y coordinate.",
					},
					"y": map[string]any{
						"type":        "number",
						"description": "Vertical scroll position in pixels (used when no selector is provided).",
					},
				},
			},
		}},
		{OfTool: &anthropic.ToolParam{
			Name:        "press_key",
			Description: anthropic.String("Press a keyboard key, optionally on a specific element. Useful for submitting forms (Enter), closing dialogs (Escape), or tabbing between fields."),
			InputSchema: anthropic.ToolInputSchemaParam{
				Properties: map[string]any{
					"key": map[string]any{
						"type":        "string",
						"description": "Key name, e.g. 'Enter', 'Escape', 'Tab', 'ArrowDown'.",
					},
					"selector": map[string]any{
						"type":        "string",
						"description": "CSS selector of the element to dispatch the key event on. Defaults to the currently focused element.",
					},
				},
				Required: []string{"key"},
			},
		}},
		{OfTool: &anthropic.ToolParam{
			Name:        "hover_element",
			Description: anthropic.String("Hover the mouse over an element to reveal tooltips or dropdown menus."),
			InputSchema: anthropic.ToolInputSchemaParam{
				Properties: map[string]any{
					"selector": map[string]any{
						"type":        "string",
						"description": "CSS selector of the element to hover.",
					},
					"text": map[string]any{
						"type":        "string",
						"description": "Visible text of the element to hover (used if selector is not provided).",
					},
				},
			},
		}},
		{OfTool: &anthropic.ToolParam{
			Name:        "navigate_to",
			Description: anthropic.String("Navigate the browser to a URL."),
			InputSchema: anthropic.ToolInputSchemaParam{
				Properties: map[string]any{
					"url": map[string]any{
						"type":        "string",
						"description": "The full URL to navigate to.",
					},
				},
				Required: []string{"url"},
			},
		}},
		{OfTool: &anthropic.ToolParam{
			Name:        "execute_steps",
			Description: anthropic.String("Execute multiple browser actions in sequence. Use this for multi-step tasks like filling a form and submitting it, or performing a series of clicks."),
			InputSchema: anthropic.ToolInputSchemaParam{
				Properties: map[string]any{
					"steps": map[string]any{
						"type":        "array",
						"description": "Ordered list of actions to perform.",
						"items": map[string]any{
							"type": "object",
							"properties": map[string]any{
								"action":   map[string]any{"type": "string", "description": "Action type: click, type_text, clear_field, select_option, scroll_to, press_key, hover_element, navigate_to."},
								"selector": map[string]any{"type": "string", "description": "CSS selector."},
								"text":     map[string]any{"type": "string", "description": "Text to type or element text to find."},
								"value":    map[string]any{"type": "string", "description": "Option value for select_option."},
								"url":      map[string]any{"type": "string", "description": "URL for navigate_to."},
								"key":      map[string]any{"type": "string", "description": "Key name for press_key."},
								"y":        map[string]any{"type": "number", "description": "Y position for scroll_to."},
							},
							"required": []string{"action"},
						},
					},
				},
				Required: []string{"steps"},
			},
		}},
	}
}

// ─── Response parser ──────────────────────────────────────────────────────────

func parseResponse(resp *anthropic.Message, conversationID string) (*domain.ChatResponse, error) {
	var textParts []string
	var action *domain.BrowserAction

	for _, block := range resp.Content {
		switch block.Type {
		case "text":
			textParts = append(textParts, block.Text)
		case "tool_use":
			if action == nil {
				a, err := toolUseToBrowserAction(block.Name, block.Input)
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

// toolUseToBrowserAction maps a Claude tool_use block to a BrowserAction.
func toolUseToBrowserAction(name string, rawInput json.RawMessage) (*domain.BrowserAction, error) {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(rawInput, &raw); err != nil {
		return nil, fmt.Errorf("unmarshal tool input: %w", err)
	}

	actionType := map[string]string{
		"click":         "click",
		"type_text":     "type",
		"clear_field":   "clear",
		"select_option": "select_option",
		"scroll_to":     "scroll",
		"press_key":     "press_key",
		"hover_element": "hover",
		"navigate_to":   "navigate",
		"execute_steps": "steps",
	}[name]

	if actionType == "" {
		return nil, fmt.Errorf("unknown tool: %s", name)
	}

	params := &domain.ActionParams{}

	if v, ok := raw["selector"]; ok {
		var s string
		if json.Unmarshal(v, &s) == nil {
			params.Selector = &s
		}
	}
	if v, ok := raw["text"]; ok {
		var s string
		if json.Unmarshal(v, &s) == nil {
			params.Text = &s
		}
	}
	if v, ok := raw["value"]; ok {
		var s string
		if json.Unmarshal(v, &s) == nil {
			params.Value = &s
		}
	}
	if v, ok := raw["url"]; ok {
		var s string
		if json.Unmarshal(v, &s) == nil {
			params.URL = &s
		}
	}
	if v, ok := raw["key"]; ok {
		var s string
		if json.Unmarshal(v, &s) == nil {
			params.Key = &s
		}
	}
	if v, ok := raw["x"]; ok {
		var f float64
		if json.Unmarshal(v, &f) == nil {
			params.X = &f
		}
	}
	if v, ok := raw["y"]; ok {
		var f float64
		if json.Unmarshal(v, &f) == nil {
			params.Y = &f
		}
	}
	if v, ok := raw["steps"]; ok {
		var steps []domain.BrowserStep
		if json.Unmarshal(v, &steps) == nil {
			params.Steps = steps
		}
	}

	return &domain.BrowserAction{
		Type:   actionType,
		Params: params,
	}, nil
}
