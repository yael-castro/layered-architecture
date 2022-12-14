package main

import (
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	log.SetFlags(log.Flags() | log.Lshortfile)

	log.Printf(`http server is running on port "%v" %v`, port, "🤘\n")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
