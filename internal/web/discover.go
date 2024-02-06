package web

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"regexp"

	"github.com/dbtedman/accretion/internal/security"
)

const (
	HTTPContentType           = "Content-Type"
	HTTPCacheControl          = "Cache-Control"
	HTTPContentSecurityPolicy = "Content-Security-Policy"
)

func HandleDiscoverUI(htmlFS fs.FS) {
	htmlTemplates := template.Must(template.ParseFS(htmlFS, "html/*.gohtml"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		isNotUIRoute, _ := regexp.MatchString("^(static|api)/*", r.URL.Path)

		if isNotUIRoute {
			return
		}

		nonce := security.PseudoRandomNonceArgon2()

		w.Header().Set(HTTPContentType, "text/html; charset=utf-8")
		w.Header().Set(HTTPCacheControl, "max-age=0, private, must-revalidate")
		w.Header().Set(HTTPContentSecurityPolicy, fmt.Sprintf("default-src 'none'; style-src 'nonce-%s'; font-src https://fonts.gstatic.com", nonce))

		type DiscoverData struct {
			Description string
			Title       string
			Nonce       string
			Message     string
		}

		err := htmlTemplates.ExecuteTemplate(w, "discover.gohtml", DiscoverData{
			Description: "",
			Title:       "",
			Nonce:       nonce,
			Message:     "Hello, World!",
		})

		if err != nil {
			fmt.Println(err)
		}
	})
}
