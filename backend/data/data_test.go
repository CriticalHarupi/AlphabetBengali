package data_test

import (
	"testing"

	"alphabetbengali/data"
)

func TestCharactersLoaded(t *testing.T) {
	if len(data.Characters) == 0 {
		t.Fatal("Characters slice is empty — JSON failed to load")
	}
}

func TestAllCharactersHaveRequiredFields(t *testing.T) {
	const wantTotal = 50
	if got := len(data.Characters); got != wantTotal {
		t.Fatalf("expected %d characters, got %d", wantTotal, got)
	}
	for _, ch := range data.Characters {
		if ch.ID == 0 {
			t.Errorf("character %q has zero ID", ch.Char)
		}
		if ch.Char == "" {
			t.Errorf("character with id=%d has empty Char", ch.ID)
		}
		if ch.Romanization == "" {
			t.Errorf("character %q has empty Romanization", ch.Char)
		}
		if ch.ExampleWord == "" {
			t.Errorf("character %q has empty ExampleWord", ch.Char)
		}
		if ch.Type != "vowel" && ch.Type != "consonant" {
			t.Errorf("character %q has invalid type %q", ch.Char, ch.Type)
		}
		validGroups := map[string]bool{
			"vowel": true, "velar": true, "palatal": true,
			"retroflex": true, "dental": true, "labial": true, "other": true,
		}
		if !validGroups[ch.Group] {
			t.Errorf("character %q has invalid group %q", ch.Char, ch.Group)
		}
	}
}

func TestVowelSignConsistency(t *testing.T) {
	// Every vowel that has a vowel_sign must also have a vowel_sign_example_word.
	for _, ch := range data.Characters {
		if ch.Type == "vowel" && ch.VowelSign != "" && ch.VowelSignExampleWord == "" {
			t.Errorf("vowel %q has vowel_sign but missing vowel_sign_example_word", ch.Char)
		}
	}
}

func TestCharacterIDsAreUnique(t *testing.T) {
	seen := map[int]bool{}
	for _, ch := range data.Characters {
		if seen[ch.ID] {
			t.Errorf("duplicate ID %d (char: %q)", ch.ID, ch.Char)
		}
		seen[ch.ID] = true
	}
}
