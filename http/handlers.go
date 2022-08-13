package http

import (
	"fmt"
	"net/http"
)

func handleTemplate(w http.ResponseWriter, r *http.Request) {
        urlPath := r.URL.RequestURI()

        if urlPath == "/" {
                urlPath = "/index"
        }

        content := getTemplateContent(urlPath)

        if content == "" {
                w.WriteHeader(http.StatusNotFound)
                w.Write([]byte("404 page not found"))
        } else {
                fmt.Fprintf(w, content)
        }
}

func handleStart(w http.ResponseWriter, r *http.Request) {
	message, err := startWorkflow()

	if err != nil {
		fmt.Fprintf(w, "Error: " + message)
	} else {
		fmt.Fprintf(w, message)
	}
}

func handleList(w http.ResponseWriter, r *http.Request) {

}
