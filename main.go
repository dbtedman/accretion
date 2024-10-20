package main

import "log"

// version, commit, and date are populated during build
var (
	version = "latest"
	commit  = "n/a"
	date    = "n/a"
)

func main() {
	log.Println("Hello, World!")
}
