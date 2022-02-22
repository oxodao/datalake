package module_spotify

import (
	"encoding/json"
	"net/http"

	"github.com/oxodao/datalake/services"
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

		err = services.Get().ORM.Auth.InsertModuleAuthentication(user, dp.GetName(), string(data))
		if err != nil {
			http.Error(w, "Could not insert authentication: "+err.Error(), http.StatusInternalServerError)
		}
	}
}
