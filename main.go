package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kpunith8/yt-monitor/websocket"
)

func stats(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrader(w, r)

	if err != nil {
		fmt.Fprintf(w, "+%v\n", err)
	}

	go websocket.Writer(ws)
}

func setupRoutes() {
	fs := http.FileServer(http.Dir("static"))

	// Route / to access index.html
	http.Handle("/", fs)

	// Route /stats to access websocket
	http.HandleFunc("/stats", stats)
	log.Fatal(http.ListenAndServe(":3030", nil))
}

func main() {
	fmt.Println("Youtube Subscriber Monitor")

	setupRoutes()
}
