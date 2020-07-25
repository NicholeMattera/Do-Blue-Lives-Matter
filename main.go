package main

import (
    "log"
    "net/http"
    "text/template"
)

type TemplateData struct {
    Title		            string
    Description	            string
    URL			            string
    IncludeServiceWorker    bool
}

func main() {
    mainTemplate, err := template.ParseFiles("./templates/main.tmpl")
    if err != nil {
        log.Fatal("Error: Unable to parse the template file.")
    }

    http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        if req.URL.Path == "/smatter" || req.URL.Path == "/smatter/offline.html" {
            data := TemplateData {
                Title:                  "Do Blue Lives Matter?",
                Description:            "No",
                URL:                    "https://doblue.live/smatter",
                IncludeServiceWorker:   req.URL.Path == "/smatter",
            }

            mainTemplate.Execute(w, data)
        } else if req.URL.Path == "/sreallymatter" || req.URL.Path == "/sreallymatter/offline.html"  {
            data := TemplateData {
                Title:                  "Do Blue Lives Really Matter?",
                Description:            "Still No",
                URL:                    "https://doblue.live/sreallymatter",
                IncludeServiceWorker:   req.URL.Path == "/sreallymatter",
            }

            mainTemplate.Execute(w, data)
        } else if req.URL.Path == "/robots.txt" || req.URL.Path == "/service-worker.js" || req.URL.Path == "/manifest.json" {
            http.ServeFile(w, req, "./static" + req.URL.Path)
        } else {
            scheme := "http"
            if req.TLS != nil {
                scheme += "s"
            }

            http.Redirect(w, req, scheme + "://" + req.Host + "/smatter", http.StatusPermanentRedirect)
        }
    })

    log.Fatal(http.ListenAndServe(":8081", nil))
}