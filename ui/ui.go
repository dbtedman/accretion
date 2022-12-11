package ui

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"regexp"
)

//go:embed dist
var embeddedFS embed.FS

func HandleStaticAssets() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		isNotUIRoute, _ := regexp.MatchString("^(assets|api)/*", r.URL.Path)

		if isNotUIRoute {
			return
		}

		uiIndexHTML, err := template.ParseFS(embeddedFS, "dist/index.html")

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Cache-Control", "no-cache")
		w.WriteHeader(http.StatusOK)

		if err := uiIndexHTML.Execute(w, map[string]interface{}{}); err != nil {
			return
		}
	})

	distDirectory, err := fs.Sub(embeddedFS, "dist")
	if err != nil {
		log.Fatal(err)
	}

	// TODO: We want to assign headers to the static assets results
	http.Handle("/assets/", http.FileServer(http.FS(distDirectory)))
}
