package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	adminHttp "github.com/mmosh-pit/mmosh_backend/pkg/admin/http"
	aiHttp "github.com/mmosh-pit/mmosh_backend/pkg/ai/http"
	appleHttp "github.com/mmosh-pit/mmosh_backend/pkg/apple/http"
	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/app"
	authHttp "github.com/mmosh-pit/mmosh_backend/pkg/auth/http"
	botsHttp "github.com/mmosh-pit/mmosh_backend/pkg/bots/http"
	chatHttp "github.com/mmosh-pit/mmosh_backend/pkg/chat/http"
	commonHttp "github.com/mmosh-pit/mmosh_backend/pkg/common/http"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
	configHttp "github.com/mmosh-pit/mmosh_backend/pkg/config/http"
	googleHttp "github.com/mmosh-pit/mmosh_backend/pkg/google/http"
	mailHttp "github.com/mmosh-pit/mmosh_backend/pkg/mail/http"
	membersHttp "github.com/mmosh-pit/mmosh_backend/pkg/members/http"
	postsHttp "github.com/mmosh-pit/mmosh_backend/pkg/posts/http"
	receiptHttp "github.com/mmosh-pit/mmosh_backend/pkg/receipt/http"
	subscriptionsHttp "github.com/mmosh-pit/mmosh_backend/pkg/subscriptions/http"
	walletHttp "github.com/mmosh-pit/mmosh_backend/pkg/wallet/http"

	stripeHttp "github.com/mmosh-pit/mmosh_backend/pkg/stripe/http"
)

var regexObjectID = `/^[a-f\d]{24}$/i`

var routes = []route{
	newRoute("POST", "/early", authHttp.AddEarlyAccessHandler, false, false, false),

	newRoute("POST", "/login", authHttp.LoginHandler, false, false, false),
	newRoute("POST", "/request-code", authHttp.RequestCodeHandler, false, false, false),
	newRoute("POST", "/signup", authHttp.SignUpHandler, false, false, false),
	newRoute("POST", "/forgot-password-verification", authHttp.ForgotPasswordVerificationHandler, false, false, false),
	newRoute("POST", "/change-password", authHttp.ChangePasswordHandler, false, false, false),

	newRoute("GET", "/address", authHttp.GetWalletAddressHandler, true, false, false),
	newRoute("POST", "/sign", authHttp.SignTransactionHandler, true, false, false),

	newRoute("GET", "/is-auth", authHttp.IsAuthHandler, true, false, false),
	newRoute("DELETE", "/logout", authHttp.LogoutHandler, true, false, false),

	newRoute("GET", "/agents", botsHttp.GetAgentsHandler, false, false, false),
	newRoute("POST", "/agents/activate", botsHttp.ToggleActivateHandler, true, false, false),
	newRoute("GET", "/agents/active", botsHttp.GetActiveAgentsHandler, true, false, false),

	newRoute("GET", "/chat", chatHttp.WsHandler, true, true, false),
	newRoute("GET", "/chats/active", chatHttp.GetActiveChatsHandler, true, false, false),

	newRoute("GET", "/all-tokens", walletHttp.GetAllCoinAddressHandler, true, false, false),
	newRoute("POST", "/mail", mailHttp.IncomingEmailHandler, false, false, false),

	newRoute("POST", "/google-notifications", googleHttp.WebhookHandler, false, false, false),
	newRoute("POST", "/apple-notifications/v2", appleHttp.WebhookHandler, false, false, false),

	newRoute("GET", "/subscriptions", subscriptionsHttp.GetSubscriptionsHandler, true, false, false),

	newRoute("GET", "/posts", postsHttp.GetAllPostsHandler, true, false, false),
	newRoute("POST", "/posts", postsHttp.CreatePostHandler, true, false, false),
	newRoute("GET", "/posts/author", postsHttp.HandlePostsByAuthor, true, false, false),
	newRoute("GET", "/posts/slug", postsHttp.HandlePostBySlug, true, false, false),

	newRoute("PUT", "/guest-data", authHttp.CreateGuestUserDataHandler, true, false, false),
	newRoute("PUT", "/referred", authHttp.AddReferredToUserHandler, true, false, false),
	newRoute("PUT", "/onboarding-step", authHttp.SetOnboardingStepHandler, true, false, false),

	newRoute("GET", "/my-bots", botsHttp.GetMyBotsHandler, true, false, false),
	newRoute("GET", "/bots", botsHttp.GetBotsHandler, true, false, false),
	newRoute("POST", "/bots", botsHttp.CreateBotHandler, true, false, false),

	newRoute("PUT", "/update-profile-data", authHttp.UpdateProfileDataHandler, true, false, false),

	newRoute("GET", "/ai/realtime-token", aiHttp.GetRealtimeTokenHandler, true, false, false),

	newRoute("POST", "/bluesky", authHttp.AddBlueskyHandler, true, false, false),
	newRoute("DELETE", "/bluesky", authHttp.DeleteBlueskyHandler, true, false, false),

	newRoute("POST", "/telegram", authHttp.AddTelegramHandler, true, false, false),
	newRoute("DELETE", "/telegram", authHttp.DeleteTelegramHandler, true, false, false),

	newRoute("GET", "/members", membersHttp.GetMembersHandler, true, false, false),

	newRoute("POST", "/save-receipt", receiptHttp.SaveReceiptHandler, true, false, false),
	newRoute("POST", "/verify-receipt", receiptHttp.VerifyReceiptHandler, true, false, false),

	newRoute("GET", "/health", commonHttp.HealthCheckHandler, false, false, false),

	newRoute("POST", "/account-deletion", authHttp.AccountDeletionHandler, false, false, false),

	newRoute("POST", "/send-bot-message", chatHttp.SendBotMessageHandler, false, false, false),

	newRoute("POST", "/stripe-session", stripeHttp.GetStripeOnboardURLHandler, true, false, false),

	newRoute("GET", "/available-themes", configHttp.GetAppThemesHandler, true, false, false),
	newRoute("POST", "/theme", configHttp.CreateAppThemeHandler, true, false, false),

	// ADMIN ROUTES

	// USERS
	newRoute("GET", "/admin/users", adminHttp.GetAllUsersHandler, true, false, true),
	newRoute("PATCH", fmt.Sprintf("/admin/user/(?P<userId>%s)/deactivate", regexObjectID), adminHttp.DeactivateUserHandler, true, false, true),
	newRoute("PATCH", fmt.Sprintf("/admin/user/(?P<userId>%s)/activate", regexObjectID), adminHttp.DeactivateUserHandler, true, false, true),
	newRoute("DELETE", fmt.Sprintf("/admin/user/(?P<userId>%s)/delete", regexObjectID), adminHttp.DeleteUserHandler, true, false, true),
	// newRoute("PATCH", fmt.Sprintf("/admin/user/(?P<userId>%s)/deactivate", regexObjectID), , true, false, true),

	// BOTS
	newRoute("GET", "/admin/bots", adminHttp.GetAllBotsHandler, true, false, true),
}

type route struct {
	method                 string
	regex                  *regexp.Regexp
	handler                http.HandlerFunc
	requiresAuthentication bool
	isWebSocket            bool
	isAdmin                bool
}

func newRoute(method string, pattern string, handler http.HandlerFunc, requiresAuthentication, isWebSocket, isAdmin bool) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler, requiresAuthentication, isWebSocket, isAdmin}
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

			userId, isAuthorized := auth.ValidateAuth(token, route.isAdmin)
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
