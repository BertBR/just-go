package games

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Games []struct {
	Console  string `json:"console"`
	FileURL  string `json:"file_url"`
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	Sorted   int    `json:"sorted"`
	ImageURL string `json:"image_url"`
	Active   bool   `json:"active"`
}

func loadEnv() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetGames(w http.ResponseWriter, req *http.Request) {
	loadEnv()
	resp, err := http.Get(os.Getenv("API_URL"))

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var result Games
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
