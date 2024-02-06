package web

import (
	"fmt"
	"io/fs"
	"net/http"
	"time"
)

func StartServer(htmlFS fs.FS, staticFS fs.FS) {
	HandleDiscoverUI(htmlFS)
	HandleStaticAssets(staticFS)

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
