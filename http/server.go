package http

import (
	"net/http"

	"go.temporal.io/sdk/client"
)

type AppState struct {
    client client.Client
}

var appState AppState

func Server() {
	// Attempt to establish Temporal client
	c, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})

	if err != nil {
		panic("Unable to create Temporal client")
	}

	appState.client = c

        defer c.Close()

	// Start HTTP server
	fileServer := http.FileServer(http.Dir(getServerDirectory() + "/static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))
	http.HandleFunc("/start/", handleStart)
	http.HandleFunc("/", handleTemplate)

	http.ListenAndServe(":5050", nil)
}
