package main

import (
	"net/http"
)

func originServer() {
}

func edgeServer(address string) *http.Client {
	return &http.Client{}
}

func serveStaticFromDisk() {

}

func main() {

}
