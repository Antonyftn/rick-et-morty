package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Structure pour les personnages
type Character struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Species string `json:"species"`
}

// Structure de la réponse de l'API
type ApiResponse struct {
	Results []Character `json:"results"`
}

func getAllCharacters() {
	// URL de l'API
	url := "https://rickandmortyapi.com/api/character"

	// Envoyer une requête GET à l'API
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Erreur lors de la requête : %v", err)
	}
	defer response.Body.Close()

	// Lire la réponse
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture de la réponse : %v", err)
	}

	// Parser la réponse JSON
	var apiResponse ApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		log.Fatalf("Erreur lors du parsing JSON : %v", err)
	}

	// Afficher les informations des personnages
	for _, character := range apiResponse.Results {
		fmt.Printf("Nom: %s, Statut: %s, Espèce: %s\n", character.Name, character.Status, character.Species)
	}
}

func main() {
	fmt.Println("Liste des personnages de Rick and Morty :")
	getAllCharacters()
}
