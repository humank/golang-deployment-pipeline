package main

import (
	"io"
	"log"
	"net/http"
)

const version string = "CLM-2.0.2"

// VersionHandler handles incoming requests to /version
// and just returns a simple version number
func versionHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, version)
}

func main() {
	log.Printf("Listening on port 8000...")
	log.Printf("kim is adding the new comment @9/11.")
	http.HandleFunc("/version", versionHandler)
	http.ListenAndServe(":8000", nil)
}
