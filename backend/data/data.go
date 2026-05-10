package data

import (
	_ "embed"
	"encoding/json"

	"alphabetbengali/models"
)

//go:embed characters.json
var raw []byte

var Characters = []models.Character{}

func init() {
	if err := json.Unmarshal(raw, &Characters); err != nil {
		panic("failed to load characters.json: " + err.Error())
	}
}
