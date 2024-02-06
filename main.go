package main

import (
	"embed"

	"github.com/dbtedman/accretion/internal/web"
)

//go:embed static/*
var staticFS embed.FS

//go:embed html/*
var htmlFS embed.FS

func main() {
	web.StartServer(htmlFS, staticFS)
}
