package data

import (
	_ "embed"
	"encoding/json"

	"alphabetbengali/models"
)

//go:embed conjuncts.json
var rawConjuncts []byte

var Conjuncts = []models.Conjunct{}

func init() {
	if err := json.Unmarshal(rawConjuncts, &Conjuncts); err != nil {
		panic("failed to load conjuncts.json: " + err.Error())
	}
}
