package ui

import (
	"bytes"
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"regexp"
	"strings"
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
		// https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP
		w.Header().Set(
			"Content-Security-Policy",
			"default-src 'self'; style-src 'self' 'nonce-1234' https://fonts.googleapis.com; font-src 'self' https://fonts.gstatic.com;",
		)
		w.WriteHeader(http.StatusOK)

		// TODO: nonce (https://content-security-policy.com/nonce/)
		// TODO: Cleanup this code and use a nonce with actual random value to it
		var tpl bytes.Buffer
		if err := uiIndexHTML.Execute(&tpl, map[string]interface{}{}); err != nil {
			return
		}

		var index string = tpl.String()
		replacedIndex := strings.ReplaceAll(index, "PLACEHOLDER_NONCE", "1234")

		_, _ = w.Write([]byte(replacedIndex))
	})

	distDirectory, err := fs.Sub(embeddedFS, "dist")
	if err != nil {
		log.Fatal(err)
	}

	assetsFS := http.FileServer(http.FS(distDirectory))
	http.Handle("/assets/", handleAssets(assetsFS))
}

func handleAssets(assetsFS http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "max-age=604800")

		assetsFS.ServeHTTP(w, r)
	}
}
