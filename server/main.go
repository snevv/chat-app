package main

import (
	"log"
	"net/http"
	"server/ws"
)

func main() {
	hub := ws.NewHub()
	go hub.Run()

	handler := ws.NewHandler(hub)

	http.HandleFunc("/ws", handler.ServeWS)

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
