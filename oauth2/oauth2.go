package oauth2

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

type Handler struct {
	OauthConfig *oauth2.Config
	UserInfoURL string
}

type ProviderData struct {
	RedirectURL  string
	ClientID     string
	ClientSecret string
	Scopes       []string
	Endpoint     oauth2.Endpoint
	UserInfoURL  string
}

func NewHandler(data ProviderData) *Handler {
	return &Handler{
		OauthConfig: &oauth2.Config{
			RedirectURL:  data.RedirectURL,
			ClientID:     data.ClientID,
			ClientSecret: data.ClientSecret,
			Scopes:       data.Scopes,
			Endpoint:     data.Endpoint,
		},
		UserInfoURL: data.UserInfoURL,
	}
}

func (h *Handler) GetAuthURL() string {
	return h.OauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
}

func (h *Handler) GetToken(ctx context.Context, authCode string) (*oauth2.Token, error) {
	token, err := h.OauthConfig.Exchange(ctx, authCode)
	if err != nil {
		return nil, fmt.Errorf("error during authorization code exchange. %v", err)
	}

	return token, nil
}

func (h *Handler) GetAccountInfo(ctx context.Context, token *oauth2.Token) (*http.Response, error) {
	var err error

	c := oauth2.NewClient(ctx, h.OauthConfig.TokenSource(ctx, token))
	resp, err := c.Get(h.UserInfoURL)
	if err != nil {
		return nil, fmt.Errorf("could not get user information. %v", err)
	}

	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	return resp, err
}

/* func (h *Handler) HandleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := h.OauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

func (h *Handler) HandleCallback() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.FormValue("code")
		token, err := h.OauthConfig.Exchange(context.Background(), code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		client := oauth2.NewClient(context.Background(), h.OauthConfig.TokenSource(context.Background(), token))
		response, err := client.Get(h.UserInfoURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer response.Body.Close()

		content, err := io.ReadAll(response.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Response: %s", content)
	}
} */
