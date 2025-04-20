package main

import (
	"net/http"

	"github.com/justinas/alice"
	"snippetbox.shrishail.dev/ui"
)

func (app *application) registerRoutes() http.Handler {
	mux := http.NewServeMux()

	// unprotected routes
	// static files handler
	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	dynamicHandlerChain := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

	mux.Handle("GET /{$}", dynamicHandlerChain.ThenFunc(app.defaultRouteHandler))
	mux.HandleFunc("GET /ping", ping)
	mux.Handle("GET /snippet/view/{id}", dynamicHandlerChain.ThenFunc(app.snippetView))

	// Add the five new routes, all of which use our 'dynamic' middleware chain.
	mux.Handle("GET /user/signup", dynamicHandlerChain.ThenFunc(app.userSignup))
	mux.Handle("POST /user/signup", dynamicHandlerChain.ThenFunc(app.userSignupPost))
	mux.Handle("GET /user/login", dynamicHandlerChain.ThenFunc(app.userLogin))
	mux.Handle("POST /user/login", dynamicHandlerChain.ThenFunc(app.userLoginPost))
	mux.Handle("POST /user/logout", dynamicHandlerChain.ThenFunc(app.userLogoutPost))

	protectedHandlerChain := dynamicHandlerChain.Append(app.requireAuthentication)
	mux.Handle("GET /snippet/create", protectedHandlerChain.ThenFunc(app.snippetCreateForm))
	mux.Handle("POST /snippet/create", protectedHandlerChain.ThenFunc(app.snippetCreate))

	standardChain := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standardChain.Then(mux)
}
