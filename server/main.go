package server

import (
	"fmt"
	"net/http"

	"github.com/bertbr/just-go/games"
)

func Start() {
	http.HandleFunc("/games", games.GetGames)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
