package main

import (
	"github.com/Antonious-Stewart/Go-CDN/cmd/origin"
	"github.com/Antonious-Stewart/Go-CDN/internal/middleware"
)

func main() {
	logMiddleware := middleware.Log{}

	origin.Server(logMiddleware)
}
