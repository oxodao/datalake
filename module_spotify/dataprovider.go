package module_spotify

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oxodao/datalake/config"
	"github.com/oxodao/datalake/services"
	"github.com/spf13/cobra"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

type DataProvider struct {
	Config        *config.Config
	ORM           *ORM
	Authenticator *spotifyauth.Authenticator
}

func New() DataProvider {
	cfg := services.Get().Config

	return DataProvider{
		Config: cfg,
		ORM:    newORM(services.Get().DB),
		Authenticator: spotifyauth.New(
			spotifyauth.WithRedirectURL(cfg.Web.Url+"/api/spotify/callback"),
			spotifyauth.WithScopes(
				spotifyauth.ScopeUserReadPrivate,
			),
			spotifyauth.WithClientID(cfg.Spotify.ClientID),
			spotifyauth.WithClientSecret(cfg.Spotify.ClientSecret),
		),
	}
}

func (dp DataProvider) IsEnabled() bool {
	return dp.Config.Spotify.Enabled
}

func (dp DataProvider) GetName() string {
	return "spotify"
}

func (dp DataProvider) Process() {
	fmt.Println("Checking spotify")
}

func (dp DataProvider) RegisterRoutes(r *mux.Router) {
	fmt.Println("Registering spotify routes")
	r.HandleFunc("/auth", routeAuthenticate(&dp)).Methods(http.MethodGet)
	r.HandleFunc("/callback", routeCallback(&dp)).Methods(http.MethodGet)
}

func (dp DataProvider) RegisterCustomCommands(cmd *cobra.Command) {
	cmd.AddCommand(importCmd)
}
