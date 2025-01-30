package main

import (
	"encoding/json"
	"net/http"

	"github.com/padeshaies/pkmdamagecalculator/helpers"
	"github.com/padeshaies/pkmdamagecalculator/types"
)

func handleGetCalc(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		LeftPokemon  types.Pokemon `json:"left_pokemon"`
		RightPokemon types.Pokemon `json:"right_pokemon"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding request; Please provide valid JSON"))
		return
	}

	err = helpers.CalculateStats(&params.LeftPokemon)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding request; Please provide valid Pokemon"))
		return
	}

	err = helpers.CalculateStats(&params.RightPokemon)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding request; Please provide valid Pokemon"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("calculating..."))
}
