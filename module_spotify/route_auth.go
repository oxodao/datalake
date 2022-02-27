package module_spotify

import (
	"encoding/json"
	"net/http"

	"github.com/oxodao/datalake/services"
	"github.com/zmb3/spotify/v2"
)

// @TODO: Authenticate this request so that no one can change the account
func routeAuthenticate(dp *DataProvider) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, dp.Authenticator.AuthURL("@TODO"), http.StatusFound)
	}
}

func routeCallback(dp *DataProvider) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tok, err := dp.Authenticator.Token(r.Context(), "@TODO", r)
		if err != nil {
			http.Error(w, "Could not get token: "+err.Error(), http.StatusInternalServerError)
			return
		}

		user, err := services.Get().ORM.Auth.FindByUsername("admin")
		if err != nil {
			http.Error(w, "Could not find user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		data, _ := json.Marshal(tok)

		client := spotify.New(dp.Authenticator.Client(r.Context(), tok))
		su, err := client.CurrentUser(r.Context())
		if err != nil {
			http.Error(w, "The Spotify API returned an error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		err = services.Get().ORM.Auth.UpsertModuleAuthentication(user, dp.GetName(), su.ID, string(data))
		if err != nil {
			http.Error(w, "Could not insert authentication: "+err.Error(), http.StatusInternalServerError)
		}
	}
}
