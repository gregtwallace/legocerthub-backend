package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) Routes() http.Handler {
	router := httprouter.New()

	// app
	router.HandlerFunc(http.MethodGet, "/api/status", app.statusHandler)

	// private_keys
	router.HandlerFunc(http.MethodGet, "/api/v1/privatekeys", app.keys.GetAllKeys)
	router.HandlerFunc(http.MethodPost, "/api/v1/privatekeys", app.keys.PostNewKey)
	router.HandlerFunc(http.MethodGet, "/api/v1/privatekeys/:id", app.keys.GetOneKey)
	router.HandlerFunc(http.MethodPut, "/api/v1/privatekeys/:id", app.keys.PutNameDescKey)
	router.HandlerFunc(http.MethodDelete, "/api/v1/privatekeys/:id", app.keys.DeleteKey)

	// acme_accounts
	router.HandlerFunc(http.MethodGet, "/api/v1/acmeaccounts", app.accounts.GetAllAccounts)
	router.HandlerFunc(http.MethodPost, "/api/v1/acmeaccounts", app.accounts.PostNewAccount)
	router.HandlerFunc(http.MethodGet, "/api/v1/acmeaccounts/:id", app.accounts.GetOneAccount)
	router.HandlerFunc(http.MethodPut, "/api/v1/acmeaccounts/:id", app.accounts.PutNameDescAccount)
	router.HandlerFunc(http.MethodDelete, "/api/v1/acmeaccounts/:id", app.accounts.DeleteAccount)

	return app.enableCORS(router)
}
