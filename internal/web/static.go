package web

import (
	"io/fs"
	"net/http"
)

func HandleStaticAssets(staticFS fs.FS) {
	http.Handle("/static/", staticHandlerMiddleware(http.FileServer(http.FS(staticFS))))
}

func staticHandlerMiddleware(assetsFS http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		assetsFS.ServeHTTP(w, r)
	}
}
