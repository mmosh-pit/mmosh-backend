package main

import (
	"log"
	"net/http"
	"regexp"
	"strings"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/app"
	authHttp "github.com/mmosh-pit/mmosh_backend/pkg/auth/http"
	chatHttp "github.com/mmosh-pit/mmosh_backend/pkg/chat/http"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
	mailHttp "github.com/mmosh-pit/mmosh_backend/pkg/mail/http"
	walletHttp "github.com/mmosh-pit/mmosh_backend/pkg/wallet/http"
)

var regexUUID = "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}"

var routes = []route{
	newRoute("POST", "/login", authHttp.LoginHandler, false, false),
	newRoute("POST", "/request-code", authHttp.RequestCodeHandler, false, false),
	newRoute("POST", "/signup", authHttp.SignUpHandler, false, false),

	newRoute("GET", "/is-auth", authHttp.IsAuthHandler, true, false),
	newRoute("GET", "/get-or-create-chat", chatHttp.GetChatHandler, true, false),
	newRoute("DELETE", "/logout", authHttp.LogoutHandler, true, false),

	newRoute("GET", "/chat", chatHttp.WsHandler, true, true),

	newRoute("GET", "/all-tokens", walletHttp.GetAllCoinAddressHandler, true, false),
	newRoute("POST", "/mail", mailHttp.IncomingEmailHandler, false, false),
	newRoute("GET", "/private-key", authHttp.GetPrivateKeyHandler, true, false),
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
			log.Println("Requires authentication?????")
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

		log.Println("Sending to handler")
		SetRouteParams(r, params)
		route.handler(w, r)

		return
	}

	if len(invalidAuthentication) > 0 {
		common.SendErrorResponse(w, http.StatusUnauthorized, []string{"user_unauthorized"})
		return
	}

	if len(invalidPayload) > 0 {
		common.SendErrorResponse(w, http.StatusBadRequest, []string{"invalid_payload"})
		return
	}

	if len(allow) > 0 {
		common.SendErrorResponse(w, http.StatusMethodNotAllowed, nil)
		return
	}

	if isOptions {
		common.SendSuccessResponse(w, http.StatusOK, nil)
		return
	}

	http.NotFound(w, r)
}
