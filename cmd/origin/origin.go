package origin

import (
	"github.com/Antonious-Stewart/Go-CDN/internal/middleware"
	"log"
	"net/http"
)

// Server
// The server that hosts the original content, which will be distributed across the CDN.
func Server(logging middleware.Logger) {
	// where the content is stored
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./static"))

	mux.Handle("/static/", logging.Info(http.StripPrefix("/static/", fs)))

	log.Println("📦 Static CDN running at http://localhost:8080/static/")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
