package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

func handleAction(w http.ResponseWriter, r *http.Request) {
	action := r.URL.Path[len("/"):]

	switch action {
	case "start_motor":
		fmt.Println("Motor gestartet")
		w.Write([]byte("Motor gestartet"))
	case "forward":
		fmt.Println("Vorwärts")
		w.Write([]byte("Vorwärts"))
	case "backward":
		fmt.Println("Rückwärts")
		w.Write([]byte("Rückwärts"))
	case "left":
		fmt.Println("Links")
		w.Write([]byte("Links"))
	case "right":
		fmt.Println("Rechts")
		w.Write([]byte("Rechts"))
	case "stop":
		fmt.Println("Stopp")
		w.Write([]byte("Stopp"))
	default:
		http.NotFound(w, r)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleAction)

	corsHandler := cors.Default().Handler(mux)

	fmt.Println("Die REST-API ist gestartet und hört auf Port 8080")
	http.ListenAndServe(":8000", corsHandler)
}
