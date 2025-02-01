package main

import (
	"github.com/go-chi/chi/v5"
	"UpptimeKumaProbeAPI/endpoints"
)

func main() {
	r := chi.NewRouter()

	r.Get("/editor", serveEditor)
	r.Get("/status/{scan_name}", serveStatus)
}
