package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/blue-army-2017/knight/util"
	"github.com/zitadel/zitadel-go/v3/pkg/authentication"
	openid "github.com/zitadel/zitadel-go/v3/pkg/authentication/oidc"
	"github.com/zitadel/zitadel-go/v3/pkg/zitadel"
)

var routes = map[string]func(w http.ResponseWriter, r *http.Request){
	"/health":                              getHealth,
	"/":                                    getHome,
	"/members":                             getMembers,
	"/members/new":                         newMember,
	"POST /members/new":                    postNewMember,
	"/members/{id}/edit":                   editMember,
	"POST /members/{id}/edit":              postEditMember,
	"/members/{id}/delete":                 deleteMember,
	"/seasons":                             getSeasons,
	"/seasons/{id}":                        getSeason,
	"/seasons/new":                         newSeason,
	"POST /seasons/new":                    postNewSeason,
	"/seasons/{id}/edit":                   editSeason,
	"POST /seasons/{id}/edit":              postEditSeason,
	"/seasons/{id}/delete":                 deleteSeason,
	"/seasons/{s_id}/games":                getSeasonGames,
	"/seasons/{s_id}/games/new":            newSeasonGame,
	"POST /seasons/{s_id}/games/new":       postNewSeasonGame,
	"/seasons/{s_id}/games/{id}/edit":      editSeasonGame,
	"POST /seasons/{s_id}/games/{id}/edit": postEditSeasonGame,
	"/seasons/{s_id}/games/{id}/delete":    deleteSeasonGame,
	"/seasons/{s_id}/presence":             getSeasonPresence,
}

func GetRoutesMux() (*http.ServeMux, error) {
	authRedirect := fmt.Sprintf("%s/auth/callback", util.GetServerUri())
	authN, err := authentication.New(
		context.Background(),
		zitadel.New(util.Config.AuthDomain),
		util.Config.AuthKey,
		openid.DefaultAuthentication(util.Config.ClientID, authRedirect, util.Config.AuthKey),
	)
	if err != nil {
		return nil, err
	}
	authMw := authentication.Middleware(authN)

	mux := http.NewServeMux()

	mux.Handle("/auth/", authN)

	mux.Handle("/static/", assetsHandler())

	for route, fn := range routes {
		handleFunc := fn
		if util.Config.Environment == "development" {
			handleFunc = logMiddleware(fn)
		}

		handler := authMw.RequireAuthentication()(http.HandlerFunc(handleFunc))
		mux.Handle(route, handler)
	}

	return mux, nil
}

func assetsHandler() http.Handler {
	dir := http.Dir("assets/")
	fs := http.FileServer(dir)
	return http.StripPrefix("/static/", fs)
}
