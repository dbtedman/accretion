package ui

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed dist
var embeddedFS embed.FS

func HandleStaticAssets() {
	distDirectory, err := fs.Sub(embeddedFS, "dist")
	if err != nil {
		log.Fatal(err)
	}
	distAssetsDirectory, err := fs.Sub(distDirectory, "assets")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.FS(distAssetsDirectory))))
}

func HandleIndexHTML() {
	// TODO: Support the index.html
}
