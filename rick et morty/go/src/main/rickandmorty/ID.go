package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getCharacterByID(id int) {
	url := fmt.Sprintf("https://rickandmortyapi.com/api/character/%d", id)

	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Erreur lors de la requête : %v", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture de la réponse : %v", err)
	}

	var character Character
	err = json.Unmarshal(body, &character)
	if err != nil {
		log.Fatalf("Erreur lors du parsing JSON : %v", err)
	}

	fmt.Printf("Nom: %s\nStatut: %s\nEspèce: %s\n", character.Name, character.Status, character.Species)
}

func id() {
	fmt.Println("Personnage spécifique :")
	getCharacterByID(1)
}
