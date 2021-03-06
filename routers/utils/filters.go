package utils

import (
	"fmt"
	"github.com/beego/beego/v2/server/web/context"
	"net/http"
)

var FilterUser = func(ctx *context.Context) {
	email := ctx.Input.Session("email")
	ruolo := ctx.Input.Session("ruolo")

	if email == nil || ruolo == nil {
		ctx.Redirect(302, "/login")
	}
}

var FilterRuolo = func(ctx *context.Context) {
	filtro := map[string][]string{
		"privato":        []string{"/dashboard/pagamenti", "/dashboard/prenotazioni", "/dashboard/pazienti", "/dashboard/dipendenti"},
		"laboratorio":    []string{"/dashboard/prenota", "/dashboard/guida", "/dashboard/pazienti", "/dashboard/dipendenti"},
		"organizzazione": []string{"/dashboard/pazienti", "/dashboard/prenotazioni", "/dashboard/pagamenti"},
		"medico":         []string{"/dashboard/dipendenti", "/dashboard/prenotazioni", "/dashboard/pagamenti"},
	}
	ruolo := fmt.Sprint(ctx.Input.Session("ruolo"))
	path := fmt.Sprint(ctx.Request.URL.Path)

	permission := filtro[ruolo]
	ok := contains(permission, path)
	if ok {
		ctx.Abort(http.StatusForbidden, "Non sei autorizzato ad accedere alla seguente pagina")
	}
}

func contains(slice []string, element string) bool {
	for _, value := range slice {
		if value == element {
			return true
		}
	}
	return false
}
