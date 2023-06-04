package main

import (
	"adykaaa/ez-auth/config"
	"adykaaa/ez-auth/db"
	"adykaaa/ez-auth/http"
	"adykaaa/ez-auth/logger"
	"adykaaa/ez-auth/oauth"
	"log"
	"os"

	"github.com/go-chi/chi"
	"golang.org/x/oauth2/google"
)

func main() {
	config, err := config.Load(".")
	if err != nil {
		log.Fatalf("Could not load config. %v", err)
	}

	l := logger.New(config.LogLevel)

	h := oauth.NewHandler(oauth.ProviderData{
		RedirectURL:  config.RedirectURL,
		ClientID:     config.ClientID,
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
		UserInfoURL:  "https://www.googleapis.com/oauth2/v3/userinfo",
	})

	r := chi.NewRouter()
	r.Get("/", http.HandleHome)
	r.Get("/handleLogin", http.HandleLogin(h))
	r.Get("/auth/callback", http.HandleCallback(h))

	_, err = db.NewSQL("postgres", config.DBConnString, &l)
	if err != nil {
		l.Fatal().Err(err).Send()
	}

	err = db.MigrateDB(config.DBConnString, "file://db/migrations/", &l)
	if err != nil {
		l.Fatal().Err(err).Send()
	}

	srv, err := http.NewServer(r, config.HTTPServerAddress, &l)
	if err != nil {
		l.Fatal().Msgf("could not initiate new HTTP server %v", err)
	}

	err = srv.Shutdown()
	if err != nil {
		l.Fatal().Msgf("error during server shutdown %v", err)
	}
}
