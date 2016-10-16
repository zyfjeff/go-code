package main

import (
	"golang.org/x/net/websocket"
	"http"
	"log"
	"os"
)

func main() {
	http.Handle("/", websocket.Handler(WSHandler))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("listen and serve failure")
		os.Exit(-1)
	}
}
