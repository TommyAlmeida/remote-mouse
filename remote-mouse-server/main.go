package main

import (
	"fmt"
	"net/http"

	"github.com/tommyalmeida/remote-mouse/server"
)

func main() {
	if count := server.GetActiveConnectionCount(); count > 0 {
		fmt.Printf("WARNING: %d active connections detected on startup. This may cause unexpected mouse movement.\n", count)
	}

	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	
	http.HandleFunc("/ws", server.WSHandler)
	
	fmt.Println("Remote Mouse Server started on localhost:8080")
	fmt.Println("Connect at http://localhost:8080 to control the mouse")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}