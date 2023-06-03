package http

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type OAuthHandler interface {
	GetAuthURL() string
	GetAccountInfo(ctx context.Context, authCode string) (*http.Response, error)
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	var html = `<html><body><a href="/handleLogin">Log In</a></body></html>`
	fmt.Fprint(w, html)
}

// HandleLogin is called on the initial login after the user presses "Log in with <oauth2_provider>" and redirects them to
// the right autohorization server based on the OAuthHandler config.
func HandleLogin(h OAuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l, _, cancel := SetupHandler(w, r.Context())
		defer cancel()

		url := h.GetAuthURL()

		l.Info().Msgf("redirecting user to %s", url)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

func HandleCallback(h OAuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l, ctx, cancel := SetupHandler(w, r.Context())
		defer cancel()

		code := r.FormValue("code")
		info, err := h.GetAccountInfo(ctx, code)
		if err != nil {
			l.Error().Msgf("error fetching account info %v", err)
			JSON(w, msg{"error": "fetching account info"}, http.StatusInternalServerError)
			return
		}

		//TODO: remove this after debugging
		content, err := io.ReadAll(info.Body)
		if err != nil {
			l.Error().Msgf("error reading account info %v", err)
			JSON(w, msg{"error": "reading account info"}, http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Response :%s", content)
	}
}
