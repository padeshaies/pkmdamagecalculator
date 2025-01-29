package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/padeshaies/pkmdamagecalculator/types"
)

func handleGetCalc(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		LeftPokemon  types.Pokemon `json:"left_pokemon"`
		RightPokemon types.Pokemon `json:"right_pokemon"`
	}

	fmt.Println("Handling GET /api/calc")

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding request; Please provide valid JSON"))
		return
	}

	err = calculateStats(&params.LeftPokemon)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding request; Please provide valid Pokemon"))
		return
	}

	err = calculateStats(&params.RightPokemon)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding request; Please provide valid Pokemon"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("calculating..."))
}
