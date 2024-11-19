package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Character struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Species string `json:"species"`
	Gender  string `json:"gender"`
	Origin  struct {
		Name string `json:"name"`
	} `json:"origin"`
	Image string `json:"image"`
}

type ApiResponse struct {
	Results []Character `json:"results"`
}

func fetchAllCharacters() ([]Character, error) {
	url := "https://rickandmortyapi.com/api/character"
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var apiResponse ApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, err
	}

	return apiResponse.Results, nil
}

func fetchCharacterByID(id int) (*Character, error) {
	url := fmt.Sprintf("https://rickandmortyapi.com/api/character/%d", id)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var character Character
	err = json.Unmarshal(body, &character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

func handleAllCharacters(w http.ResponseWriter, r *http.Request) {
	characters, err := fetchAllCharacters()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(characters)
}

func handleCharacterByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	characterID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	character, err := fetchCharacterByID(characterID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(character)
}

func main() {
	fmt.Println("Liste des personnages de Rick and Morty :")
	http.HandleFunc("/", handleAllCharacters)
	http.HandleFunc("/character", handleCharacterByID)
	fmt.Println(`Serveur démarré sur http://localhost:8081`)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
