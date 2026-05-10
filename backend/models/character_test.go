package models_test

import (
	"encoding/json"
	"testing"

	"alphabetbengali/models"
)

func TestVowelRoundtrip(t *testing.T) {
	vowel := models.Character{
		ID:                   2,
		Type:                 "vowel",
		Group:                "vowel",
		Char:                 "আ",
		Romanization:         "a",
		ExampleWord:          "আম (am) — mango",
		VowelSign:            "া",
		VowelSignExampleWord: "কাজ (kaj) — work",
	}

	data, err := json.Marshal(vowel)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}

	var got models.Character
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}

	if got != vowel {
		t.Errorf("roundtrip mismatch: got %+v, want %+v", got, vowel)
	}
}

func TestConsonantOmitsVowelFields(t *testing.T) {
	consonant := models.Character{
		ID:           12,
		Type:         "consonant",
		Group:        "velar",
		Char:         "ক",
		Romanization: "k",
		ExampleWord:  "কলা (kola) — banana",
	}

	raw, err := json.Marshal(consonant)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}

	var m map[string]interface{}
	if err := json.Unmarshal(raw, &m); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}

	if _, ok := m["vowel_sign"]; ok {
		t.Error("consonant JSON must not contain vowel_sign")
	}
	if _, ok := m["vowel_sign_example_word"]; ok {
		t.Error("consonant JSON must not contain vowel_sign_example_word")
	}
}

func TestConsonantRoundtrip(t *testing.T) {
	consonant := models.Character{
		ID:           12,
		Type:         "consonant",
		Group:        "velar",
		Char:         "ক",
		Romanization: "k",
		ExampleWord:  "কলা (kola) — banana",
	}

	data, err := json.Marshal(consonant)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}

	var got models.Character
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}

	if got != consonant {
		t.Errorf("roundtrip mismatch: got %+v, want %+v", got, consonant)
	}
}
