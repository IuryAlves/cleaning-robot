package main

import (
	"fmt"
	"github.com/IuryAlves/cleaning-robot/server"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/tibber-developer-test/enter-path", server.EnterPathHandler)
	fmt.Printf("listening on port: %s\n", port)
	http.ListenAndServe(":8080", nil)
}
