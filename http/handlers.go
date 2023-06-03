package http

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

type OAuthHandler interface {
	GetAuthURL() string
	GetAccountInfo(ctx context.Context, authCode string, token *oauth2.Token) (*http.Response, error)
}

// HandleLogin is called on the initial login after the user presses "Log in with <oauth2_provider>" and redirects them to
// the right autohorization server.
func HandleLogin(h OAuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := h.GetAuthURL()
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

func HandleCallback() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
