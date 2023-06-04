package oauth

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

// Handler holds the Oauth configuration and the UserInfoURL that has the information about a particular user
type Handler struct {
	OauthConfig *oauth2.Config
	UserInfoURL string
}

// ProviderData has all the necessary information that allows ez-auth to get profile information
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

func parseScopes(scopes string) []string {
	var s []string

}

// getToken returns the Token struct that holds the Access and Refresh tokens for OAuth
func (h *Handler) getToken(ctx context.Context, authCode string) (*oauth2.Token, error) {
	token, err := h.OauthConfig.Exchange(ctx, authCode)
	if err != nil {
		return nil, fmt.Errorf("error during authorization code to token exchange. %v", err)
	}

	return token, nil
}

// GetAuthURL returns the authentication URL for providers
func (h *Handler) GetAuthURL() string {
	return h.OauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
}

// GetAccountInfo returns all the information about a specified account
func (h *Handler) GetAccountInfo(ctx context.Context, authCode string) (*http.Response, error) {
	var err error

	t, err := h.getToken(ctx, authCode)
	if err != nil {
		return nil, err
	}

	c := oauth2.NewClient(ctx, h.OauthConfig.TokenSource(ctx, t))
	resp, err := c.Get(h.UserInfoURL)
	if err != nil {
		return nil, fmt.Errorf("could not get user information. %v", err)
	}

	/* 	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil && err == nil {
			err = closeErr
		}
	}() */

	return resp, err
}
