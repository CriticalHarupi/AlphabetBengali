package models

type Character struct {
	ID                   int    `json:"id"`
	Type                 string `json:"type"`
	Group                string `json:"group"`
	Char                 string `json:"char"`
	Romanization         string `json:"romanization"`
	ExampleWord          string `json:"example_word"`
	VowelSign            string `json:"vowel_sign,omitempty"`
	VowelSignExampleWord string `json:"vowel_sign_example_word,omitempty"`
}

// test
