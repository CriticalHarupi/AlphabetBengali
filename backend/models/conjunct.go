package models

type Conjunct struct {
	ID           int      `json:"id"`
	Char         string   `json:"char"`
	Components   []string `json:"components"`
	Romanization string   `json:"romanization"`
	ExampleWord  string   `json:"example_word"`
}
