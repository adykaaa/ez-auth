package main

import (
	"adykaaa/ez-auth/db"
	"adykaaa/ez-auth/logger"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8080/auth/callback",
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

func main() {
	l := logger.New(zerolog.InfoLevel.String())

	sqldb, err := db.NewSQL("postgres", config.DBConnString, &l)
	if err != nil {
		l.Fatal().Err(err).Send()
	}

	err = db.MigrateDB(config.DBConnString, "file://db/migrations/", &l)
	if err != nil {
		l.Fatal().Err(err).Send()
	}

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleGoogleLogin)
	http.HandleFunc("/auth/callback", handleGoogleCallback)
	http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	var html = `<html><body><a href="/login">Google Log In</a></body></html>`
	fmt.Fprint(w, html)
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Now you can access the Google User info with the access token:
	// You can use the `oauth2` package `TokenSource` method and the `http` package to create a new client.
	client := oauth2.NewClient(context.Background(), googleOauthConfig.TokenSource(context.Background(), token))
	// And then use the `Get` method to access the Google userinfo endpoint.
	response, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
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
