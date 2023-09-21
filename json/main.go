package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	ID      string `json:"id"`
	Balance int    `json:"balance"`
}

func main() {
	// encode to json
	account := Account{ID: "1", Balance: 100}

	res, err := json.Marshal(account)

	if err != nil {
		panic(err)
	}
	println(string(res))

	encoder := json.NewEncoder(os.Stdout)
	err = encoder.Encode(account)

	if err != nil {
		panic(err)
	}

	// decode from json
	from_json := []byte(`{"id": "2", "balance": 250}`)
	var account_from_json Account

	// using & points to memory address of variable,
	// emulates and updates the state of variable with response
	// of Unmarshal method.
	json.Unmarshal(from_json, &account_from_json)

	print(account_from_json.Balance)
}
