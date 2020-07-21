package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/smatter" {
			scheme := "http"
			if req.TLS != nil {
				scheme += "s"
			}

			http.Redirect(w, req, scheme + "://" + req.Host + "/smatter", http.StatusPermanentRedirect)
			
			return
		}
	
		http.ServeFile(w, req, "./static/index.html")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}