package main

import (
	"encoding/json"
	"net/http"

	"github.com/padeshaies/pkmdamagecalculator/helpers"
	"github.com/padeshaies/pkmdamagecalculator/types"
)

func handleGetCalc(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		LeftPokemon  types.ShowdownPokemon `json:"left_pokemon"`
		RightPokemon types.ShowdownPokemon `json:"right_pokemon"`
		Field        types.Field           `json:"field"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding request; Please provide valid JSON"))
		return
	}

	leftPokemon, err := helpers.ImportPokemonFromShowdownSet(params.LeftPokemon)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding request; Please provide valid Pokemon"))
		return
	}

	rightPokemon, err := helpers.ImportPokemonFromShowdownSet(params.RightPokemon)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding request; Please provide valid Pokemon"))
		return
	}

	field := types.Field{}

	// change this to check all moves
	damage := helpers.CalculateDamage(leftPokemon, rightPokemon, leftPokemon.Moves[0], field)

	type response struct {
		Damage []int `json:"damage"`
	}
	res := response{Damage: damage}

	responseData, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error encoding response"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}
