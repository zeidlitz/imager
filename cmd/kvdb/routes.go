package main

import (
	"net/http"

	"github.com/zeidlitz/kvdb/assets"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	mux := httprouter.New()
	mux.NotFound = http.HandlerFunc(app.notFound)

	fileServer := http.FileServer(http.FS(assets.EmbeddedFiles))
	mux.Handler("GET", "/static/*filepath", fileServer)

	mux.HandlerFunc("GET", "/", app.home)

	return app.logAccess(app.recoverPanic(app.securityHeaders(mux)))
}
