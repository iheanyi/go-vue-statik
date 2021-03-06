//go:generate npm run build
package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/iheanyi/go-vue-statik/cmd/sampleapp/statik"
	"github.com/rakyll/statik/fs"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "1337"
	}
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	staticHandler := http.FileServer(statikFS)
	// Serves up the index.html file regardless of the path.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/"
		staticHandler.ServeHTTP(w, r)
	})
	http.Handle("/static/", staticHandler)
	http.ListenAndServe(":"+port, nil)
}
