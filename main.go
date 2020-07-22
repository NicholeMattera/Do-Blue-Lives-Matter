package main

import (
	"log"
	"net/http"
	"text/template"
)

type TemplateData struct {
	Title		string
	Description	string
	URL			string
}

func main() {
	mainTemplate, err := template.ParseFiles("./templates/main.tmpl")
	if err != nil {
		log.Fatal("Error: Unable to parse the template file.")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/smatter" {
			data := TemplateData {
				Title: "Do Blue Lives Matter?",
				Description: "No",
				URL: "https://doblue.live/smatter",
			}

			mainTemplate.Execute(w, data)
		} else if req.URL.Path == "/sreallymatter" {
			data := TemplateData {
				Title: "Do Blue Lives Really Matter?",
				Description: "Still No",
				URL: "https://doblue.live/sreallymatter",
			}

			mainTemplate.Execute(w, data)
		} else {
			scheme := "http"
			if req.TLS != nil {
				scheme += "s"
			}

			http.Redirect(w, req, scheme + "://" + req.Host + "/smatter", http.StatusPermanentRedirect)
			
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}