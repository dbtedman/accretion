package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"time"
)

//go:embed static/*
var staticFS embed.FS

//go:embed html/*
var htmlFS embed.FS

const (
	HTTPContentType           = "Content-Type"
	HTTPCacheControl          = "Cache-Control"
	HTTPContentSecurityPolicy = "Content-Security-Policy"
)

func main() {
	htmlTemplates := template.Must(template.ParseFS(htmlFS, "html/*.gohtml"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		isNotUIRoute, _ := regexp.MatchString("^(static|api)/*", r.URL.Path)

		if isNotUIRoute {
			return
		}

		nonce := "placeholder" // TODO: Generate a validate nonce.

		w.Header().Set(HTTPContentType, "text/html; charset=utf-8")
		w.Header().Set(HTTPCacheControl, "max-age=0, private, must-revalidate")
		w.Header().Set(HTTPContentSecurityPolicy, fmt.Sprintf("default-src 'none'; style-src 'nonce-%s'", nonce))

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

	http.Handle("/static/", staticHandlerMiddleware(http.FileServer(http.FS(staticFS))))

	server := &http.Server{
		Addr:         ":3000",
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
		IdleTimeout:  time.Minute,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
	}
}

func staticHandlerMiddleware(assetsFS http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		assetsFS.ServeHTTP(w, r)
	}
}
