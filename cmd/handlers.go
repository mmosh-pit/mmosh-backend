package main

import (
	"net/http"
	"regexp"
	"strings"

	aiHttp "github.com/mmosh-pit/mmosh_backend/pkg/ai/http"
	appleHttp "github.com/mmosh-pit/mmosh_backend/pkg/apple/http"
	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/app"
	authHttp "github.com/mmosh-pit/mmosh_backend/pkg/auth/http"
	botsHttp "github.com/mmosh-pit/mmosh_backend/pkg/bots/http"
	chatHttp "github.com/mmosh-pit/mmosh_backend/pkg/chat/http"
	commonHttp "github.com/mmosh-pit/mmosh_backend/pkg/common/http"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
	googleHttp "github.com/mmosh-pit/mmosh_backend/pkg/google/http"
	mailHttp "github.com/mmosh-pit/mmosh_backend/pkg/mail/http"
	membersHttp "github.com/mmosh-pit/mmosh_backend/pkg/members/http"
	postsHttp "github.com/mmosh-pit/mmosh_backend/pkg/posts/http"
	receiptHttp "github.com/mmosh-pit/mmosh_backend/pkg/receipt/http"
	subscriptionsHttp "github.com/mmosh-pit/mmosh_backend/pkg/subscriptions/http"
	walletHttp "github.com/mmosh-pit/mmosh_backend/pkg/wallet/http"
)

var regexUUID = "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}"

var routes = []route{
	newRoute("POST", "/early", authHttp.AddEarlyAccessHandler, false, false),

	newRoute("POST", "/login", authHttp.LoginHandler, false, false),
	newRoute("POST", "/request-code", authHttp.RequestCodeHandler, false, false),
	newRoute("POST", "/signup", authHttp.SignUpHandler, false, false),
	newRoute("POST", "/forgot-password-verification", authHttp.ForgotPasswordVerificationHandler, false, false),
	newRoute("POST", "/change-password", authHttp.ChangePasswordHandler, false, false),

	newRoute("GET", "/address", authHttp.GetWalletAddressHandler, true, false),
	newRoute("POST", "/sign", authHttp.SignTransactionHandler, true, false),

	newRoute("GET", "/is-auth", authHttp.IsAuthHandler, true, false),
	newRoute("DELETE", "/logout", authHttp.LogoutHandler, true, false),

	newRoute("GET", "/agents", botsHttp.GetAgentsHandler, false, false),
	newRoute("POST", "/agents/activate", botsHttp.ToggleActivateHandler, true, false),
	newRoute("GET", "/agents/active", botsHttp.GetActiveAgentsHandler, true, false),

	newRoute("GET", "/chat", chatHttp.WsHandler, true, true),
	newRoute("GET", "/chats/active", chatHttp.GetActiveChatsHandler, true, false),

	newRoute("GET", "/all-tokens", walletHttp.GetAllCoinAddressHandler, true, false),
	newRoute("POST", "/mail", mailHttp.IncomingEmailHandler, false, false),

	newRoute("POST", "/google-notifications", googleHttp.WebhookHandler, false, false),
	newRoute("POST", "/apple-notifications/v2", appleHttp.WebhookHandler, false, false),

	newRoute("GET", "/subscriptions", subscriptionsHttp.GetSubscriptionsHandler, true, false),

	newRoute("GET", "/posts", postsHttp.GetAllPostsHandler, true, false),
	newRoute("POST", "/posts", postsHttp.CreatePostHandler, true, false),
	newRoute("GET", "/posts/author", postsHttp.HandlePostsByAuthor, true, false),
	newRoute("GET", "/posts/slug", postsHttp.HandlePostBySlug, true, false),

	newRoute("PUT", "/guest-data", authHttp.CreateGuestUserDataHandler, true, false),
	newRoute("PUT", "/referred", authHttp.AddReferredToUserHandler, true, false),
	newRoute("PUT", "/onboarding-step", authHttp.SetOnboardingStepHandler, true, false),

	newRoute("GET", "/my-bots", botsHttp.GetMyBotsHandler, true, false),
	newRoute("GET", "/bots", botsHttp.GetBotsHandler, true, false),
	newRoute("POST", "/bots", botsHttp.CreateBotHandler, true, false),

	newRoute("PUT", "/update-profile-data", authHttp.UpdateProfileDataHandler, true, false),

	newRoute("GET", "/ai/realtime-token", aiHttp.GetRealtimeTokenHandler, true, false),

	newRoute("POST", "/bluesky", authHttp.AddBlueskyHandler, true, false),
	newRoute("DELETE", "/bluesky", authHttp.DeleteBlueskyHandler, true, false),

	newRoute("POST", "/telegram", authHttp.AddTelegramHandler, true, false),
	newRoute("DELETE", "/telegram", authHttp.DeleteTelegramHandler, true, false),

	newRoute("GET", "/members", membersHttp.GetMembersHandler, true, false),

	newRoute("POST", "/save-receipt", receiptHttp.SaveReceiptHandler, true, false),
	newRoute("POST", "/verify-receipt", receiptHttp.VerifyReceiptHandler, true, false),

	newRoute("GET", "/health", commonHttp.HealthCheckHandler, false, false),
}

type route struct {
	method                 string
	regex                  *regexp.Regexp
	handler                http.HandlerFunc
	requiresAuthentication bool
	isWebSocket            bool
}

func newRoute(method string, pattern string, handler http.HandlerFunc, requiresAuthentication bool, isWebSocket bool) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler, requiresAuthentication, isWebSocket}
}

func SetRouteParams(r *http.Request, params map[string]string) {
	for key, value := range params {
		r.SetPathValue(key, value)
	}
}

// serve Listen to all the http requests and looks if matches with one of the configured routes.
// Handles unauthorized users by rejecting the request and distributes the request to the corresponding service for
// Authenticated Users
func serve(w http.ResponseWriter, r *http.Request) {
	var allow []string
	var invalidAuthentication []string
	var invalidPayload []string
	isOptions := false

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Vary", "Origin")
	w.Header().Add("Vary", "Access-Control-Request-Method")
	w.Header().Add("Vary", "Access-Control-Request-Headers")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token, withCredentials, authorization")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH,OPTIONS")

	for _, route := range routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			continue
		}

		params := make(map[string]string)
		for i, name := range route.regex.SubexpNames() {
			if i > 0 && i <= len(matches) {
				params[name] = matches[i]
			}
		}

		if r.Method == "OPTIONS" {
			isOptions = true
			continue
		}

		if r.Method != route.method {
			allow = append(allow, route.method)
			continue
		}

		if route.requiresAuthentication {
			reqToken := r.Header.Get("Authorization")
			token := strings.Replace(reqToken, "Bearer ", "", 1)

			if route.isWebSocket {
				token = r.URL.Query().Get("token")
			}

			userId, isAuthorized := auth.ValidateAuth(token)
			if !isAuthorized {
				invalidAuthentication = append(invalidAuthentication, route.method)
				continue
			}

			r.Header.Set("userId", userId)
		}

		SetRouteParams(r, params)
		route.handler(w, r)

		return
	}

	if len(invalidAuthentication) > 0 {
		common.SendErrorResponse(w, http.StatusUnauthorized, "user_unauthorized")
		return
	}

	if len(invalidPayload) > 0 {
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid_payload")
		return
	}

	if len(allow) > 0 {
		common.SendErrorResponse(w, http.StatusMethodNotAllowed, "")
		return
	}

	if isOptions {
		common.SendSuccessResponse(w, http.StatusOK, nil)
		return
	}

	http.NotFound(w, r)
}
