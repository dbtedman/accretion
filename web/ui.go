package web

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/dbtedman/accretion/internal/domain/csp"
)

//go:embed dist
var embeddedFS embed.FS

const (
	HTTPContentType           = "Content-Type"
	HTTPCacheControl          = "Cache-Control"
	HTTPContentSecurityPolicy = "Content-Security-Policy"
)

func UI() {
	distDirectory, err := fs.Sub(embeddedFS, "dist")
	if err != nil {
		log.Fatal(err)
	}

	assetsFS := http.FileServer(http.FS(distDirectory))

	http.HandleFunc("/", indexHandler)
	http.Handle("/assets/", assetHandlerMiddleware(assetsFS))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	isNotUIRoute, _ := regexp.MatchString("^(assets|api)/*", r.URL.Path)

	if isNotUIRoute {
		// This function is does not handle static asset routes, or api routes.
		return
	}

	indexHTMLBytes, err := embeddedFS.ReadFile("dist/index.html")
	indexHTML := string(indexHTMLBytes)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	nonce := csp.GenerateNonce()

	// Add nonce into our HTML by replacing placeholder string.
	indexHTML = strings.ReplaceAll(indexHTML, "PLACEHOLDER_NONCE", nonce)

	w.Header().Set(HTTPContentType, "text/html; charset=utf-8")
	w.Header().Set(HTTPCacheControl, "max-age=0, private, must-revalidate")
	w.Header().Set(HTTPContentSecurityPolicy, csp.Generate(nonce))

	w.WriteHeader(http.StatusOK)

	// Respond with our index.html file.
	_, err = w.Write([]byte(indexHTML))

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func assetHandlerMiddleware(assetsFS http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(HTTPCacheControl, "max-age=604800")

		assetsFS.ServeHTTP(w, r)
	}
}
