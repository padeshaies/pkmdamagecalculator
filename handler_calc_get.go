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

	move := types.Move{
		Name:        "shadow ball",
		Type:        "ghost",
		Power:       80,
		DamageClass: "special",
		Target:      "selected-pokemon",
		CriticalHit: false,
	}

	field := types.Field{}

	damage := helpers.CalculateDamage(params.LeftPokemon, params.RightPokemon, move, field)

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
