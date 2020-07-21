package main

import (
	"log"
	"net/http"
)

func index(w http.ResponseWriter, req * http.Request) {
	if req.URL.Path != "/smatter" {
		http.Redirect(w, req, "https://" + req.Host + "/smatter", http.StatusPermanentRedirect)
		return
	}

	http.ServeFile(w, req, "./static/index.html")
}

func main() {
	http.HandleFunc("/", index)

	log.Fatal(http.ListenAndServe(":8081", nil))
}