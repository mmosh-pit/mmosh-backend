package main

import (
	"net/http"
	"testing"
)

// ---------------------------------------------------------------------------
// Public endpoints – no auth required
// ---------------------------------------------------------------------------

func TestGetAgents_Public(t *testing.T) {
	resp, raw := doRequest("GET", "/agents", nil, "")
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d — body: %s", resp.StatusCode, raw)
	}
}

// ---------------------------------------------------------------------------
// Bots
// ---------------------------------------------------------------------------

func TestGetBots_MissingPageParam(t *testing.T) {
	email := uniqueEmail("getbots")
	_, token := createTestUser(email, "member")

	// Missing page param → handler expects ?page=<int>
	resp, _ := doRequest("GET", "/bots", nil, token)
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400 for missing page, got %d", resp.StatusCode)
	}
}

func TestGetBots_ValidRequest(t *testing.T) {
	email := uniqueEmail("getbots-valid")
	_, token := createTestUser(email, "member")

	resp, raw := doRequest("GET", "/bots?page=0", nil, token)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d — body: %s", resp.StatusCode, raw)
	}
}

func TestGetMyBots(t *testing.T) {
	email := uniqueEmail("mybots")
	_, token := createTestUser(email, "member")

	resp, raw := doRequest("GET", "/my-bots", nil, token)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d — body: %s", resp.StatusCode, raw)
	}
}

func TestGetActiveAgents(t *testing.T) {
	email := uniqueEmail("active-agents")
	_, token := createTestUser(email, "member")

	resp, raw := doRequest("GET", "/agents/active", nil, token)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d — body: %s", resp.StatusCode, raw)
	}
}

func TestCreateBot_InvalidPayload(t *testing.T) {
	email := uniqueEmail("createbot")
	_, token := createTestUser(email, "member")

	resp, _ := doRequest("POST", "/bots", map[string]any{}, token)
	// Empty payload - agent type/key not found → expect error.
	if resp.StatusCode == http.StatusOK {
		t.Fatal("expected error for invalid bot payload, got 200")
	}
}

func TestToggleActivateAgent_InvalidPayload(t *testing.T) {
	email := uniqueEmail("activate")
	_, token := createTestUser(email, "member")

	resp, _ := doRequest("POST", "/agents/activate", map[string]any{}, token)
	if resp.StatusCode == http.StatusOK {
		t.Fatal("expected error for invalid activate payload, got 200")
	}
}

// ---------------------------------------------------------------------------
// Chats
// ---------------------------------------------------------------------------

func TestGetActiveChats(t *testing.T) {
	email := uniqueEmail("activechats")
	_, token := createTestUser(email, "member")

	resp, raw := doRequest("GET", "/chats/active", nil, token)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d — body: %s", resp.StatusCode, raw)
	}
}

func TestSendBotMessage_InvalidPayload(t *testing.T) {
	// send-bot-message is public but requires a valid payload.
	resp, _ := doRequest("POST", "/send-bot-message", map[string]any{}, "")
	if resp.StatusCode == http.StatusOK {
		t.Fatal("expected error for empty payload, got 200")
	}
}

// ---------------------------------------------------------------------------
// Posts
// ---------------------------------------------------------------------------

func TestGetPosts(t *testing.T) {
	email := uniqueEmail("getposts")
	_, token := createTestUser(email, "member")

	resp, raw := doRequest("GET", "/posts", nil, token)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d — body: %s", resp.StatusCode, raw)
	}
}

func TestGetPostsByAuthor(t *testing.T) {
	email := uniqueEmail("postsbyauthor")
	_, token := createTestUser(email, "member")

	resp, raw := doRequest("GET", "/posts/author?author=someauthor", nil, token)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d — body: %s", resp.StatusCode, raw)
	}
}

func TestGetPostBySlug_NotFound(t *testing.T) {
	email := uniqueEmail("postbyslug")
	_, token := createTestUser(email, "member")

	resp, _ := doRequest("GET", "/posts/slug?slug=nonexistent-slug", nil, token)
	// 400 or 404 expected when slug doesn't exist.
	if resp.StatusCode == http.StatusOK {
		t.Fatal("expected error for nonexistent slug, got 200")
	}
}

func TestCreatePost_MissingPrompt(t *testing.T) {
	email := uniqueEmail("createpost")
	_, token := createTestUser(email, "member")

	resp, _ := doRequest("POST", "/posts", map[string]any{
		"header": "Test Header",
		"body":   "Test body",
	}, token)
	// Missing prompt field → 400.
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400 for missing prompt, got %d", resp.StatusCode)
	}
}

// ---------------------------------------------------------------------------
// Subscriptions
// ---------------------------------------------------------------------------

func TestGetSubscriptions(t *testing.T) {
	email := uniqueEmail("subs")
	_, token := createTestUser(email, "member")

	resp, raw := doRequest("GET", "/subscriptions", nil, token)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d — body: %s", resp.StatusCode, raw)
	}
}

// ---------------------------------------------------------------------------
// Members
// ---------------------------------------------------------------------------

func TestGetMembers(t *testing.T) {
	email := uniqueEmail("members")
	_, token := createTestUser(email, "member")

	resp, raw := doRequest("GET", "/members", nil, token)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d — body: %s", resp.StatusCode, raw)
	}
}

// ---------------------------------------------------------------------------
// Wallet / tokens
// ---------------------------------------------------------------------------

func TestGetWalletAddress(t *testing.T) {
	email := uniqueEmail("wallet")
	_, token := createTestUser(email, "member")

	resp, raw := doRequest("GET", "/address", nil, token)
	// 200 if wallet exists; error otherwise. Server must not crash.
	if resp.StatusCode == 0 {
		t.Fatalf("no response — body: %s", raw)
	}
}

func TestGetAllTokens(t *testing.T) {
	email := uniqueEmail("alltokens")
	_, token := createTestUser(email, "member")

	resp, raw := doRequest("GET", "/all-tokens", nil, token)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d — body: %s", resp.StatusCode, raw)
	}
}

// ---------------------------------------------------------------------------
// Themes / Config
// ---------------------------------------------------------------------------

func TestGetAvailableThemes(t *testing.T) {
	email := uniqueEmail("themes")
	_, token := createTestUser(email, "member")

	resp, raw := doRequest("GET", "/available-themes", nil, token)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d — body: %s", resp.StatusCode, raw)
	}
}

func TestCreateTheme(t *testing.T) {
	email := uniqueEmail("createtheme")
	_, token := createTestUser(email, "member")

	resp, raw := doRequest("POST", "/theme", map[string]any{
		"name":             "Test Theme",
		"code_name":        "test-theme",
		"background_color": "#000000",
		"primary_color":    "#ffffff",
		"secondary_color":  "#aaaaaa",
		"logo":             "https://example.com/logo.png",
	}, token)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d — body: %s", resp.StatusCode, raw)
	}
}

// ---------------------------------------------------------------------------
// Receipts
// ---------------------------------------------------------------------------

func TestSaveReceipt_InvalidPayload(t *testing.T) {
	email := uniqueEmail("receipt")
	_, token := createTestUser(email, "member")

	resp, _ := doRequest("POST", "/save-receipt", map[string]any{}, token)
	// Empty receipt payload → validation error.
	if resp.StatusCode == http.StatusOK {
		t.Fatal("expected error for empty receipt, got 200")
	}
}

func TestVerifyReceipt_InvalidPayload(t *testing.T) {
	email := uniqueEmail("verify-receipt")
	_, token := createTestUser(email, "member")

	resp, _ := doRequest("POST", "/verify-receipt", map[string]any{}, token)
	if resp.StatusCode == http.StatusOK {
		t.Fatal("expected error for empty verify-receipt, got 200")
	}
}

// ---------------------------------------------------------------------------
// AI chat (Claude)
// ---------------------------------------------------------------------------

func TestClaudeHandler_InvalidPayload(t *testing.T) {
	email := uniqueEmail("claude")
	_, token := createTestUser(email, "member")

	resp, _ := doRequest("POST", "/api/chat", map[string]any{}, token)
	// Empty payload or missing required fields → error.
	if resp.StatusCode == http.StatusOK {
		t.Fatal("expected error for empty claude payload, got 200")
	}
}

// ---------------------------------------------------------------------------
// OPTIONS preflight (CORS)
// ---------------------------------------------------------------------------

func TestOptionsPreflightRequest(t *testing.T) {
	req, _ := http.NewRequest("OPTIONS", testServer.URL+"/login", nil)
	req.Header.Set("Origin", "http://example.com")
	req.Header.Set("Access-Control-Request-Method", "POST")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("options request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 for OPTIONS, got %d", resp.StatusCode)
	}
	if resp.Header.Get("Access-Control-Allow-Origin") != "*" {
		t.Fatalf("expected Access-Control-Allow-Origin: *, got %q", resp.Header.Get("Access-Control-Allow-Origin"))
	}
}

// ---------------------------------------------------------------------------
// 404 for unknown routes
// ---------------------------------------------------------------------------

func TestUnknownRoute_Returns404(t *testing.T) {
	resp, _ := doRequest("GET", "/this/does/not/exist", nil, "")
	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", resp.StatusCode)
	}
}

// ---------------------------------------------------------------------------
// Method not allowed
// ---------------------------------------------------------------------------

func TestMethodNotAllowed(t *testing.T) {
	// /health is GET only; DELETE should return 405.
	resp, _ := doRequest("DELETE", "/health", nil, "")
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405, got %d", resp.StatusCode)
	}
}
