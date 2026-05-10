# Design

## Package Structure

```mermaid
graph TD
    main["main\n(entry point)"]
    handlers["handlers\n(HTTP layer)"]
    data["data\n(embedded JSON)"]
    models["models\n(data types)"]
    web["web/index.html\n(embedded frontend)"]

    main --> handlers
    main --> web
    handlers --> data
    handlers --> models
    data --> models
```

---

## Data Models

```mermaid
classDiagram
    class Character {
        +int ID
        +string Type
        +string Group
        +string Char
        +string Romanization
        +string ExampleWord
        +string VowelSign
        +string VowelSignExampleWord
    }

    class Conjunct {
        +int ID
        +string Char
        +[]string Components
        +string Romanization
        +string ExampleWord
    }
```

`Character.Type` is `"vowel"` or `"consonant"`.

`Character.Group` is one of: `vowel`, `velar`, `palatal`, `retroflex`, `dental`, `labial`, `other`.

`Conjunct.Components` has 2 or 3 elements, each a single Bengali consonant. Sorted by first component following Bengali alphabet order.

`VowelSign` and `VowelSignExampleWord` are omitted (`omitempty`) for consonants and for the vowel অ (which has no separate diacritic form).

---

## Request Flow

```mermaid
sequenceDiagram
    participant Client
    participant main
    participant handlers
    participant data

    Note over data: init() runs at startup
    data->>data: json.Unmarshal(characters.json)
    data->>data: json.Unmarshal(conjuncts.json)

    Client->>main: GET /api/characters
    main->>handlers: GetCharacters(c)
    handlers->>data: data.Characters
    handlers-->>Client: 200 []Character

    Client->>main: GET /api/characters/:id
    main->>handlers: GetCharacterByID(c)
    handlers->>data: linear scan data.Characters
    handlers-->>Client: 200 Character / 400 / 404

    Client->>main: GET /api/conjuncts
    main->>handlers: GetConjuncts(c)
    handlers->>data: data.Conjuncts
    handlers-->>Client: 200 []Conjunct

    Client->>main: GET /api/conjuncts/:id
    main->>handlers: GetConjunctByID(c)
    handlers->>data: linear scan data.Conjuncts
    handlers-->>Client: 200 Conjunct / 400 / 404
```

---

## File Layout

```
backend/
├── main.go                  # server entry point, route registration
├── web/
│   └── index.html           # embedded test frontend
├── models/
│   ├── character.go         # Character struct
│   └── conjunct.go          # Conjunct struct
├── data/
│   ├── data.go              # embeds characters.json → data.Characters
│   ├── conjuncts_data.go    # embeds conjuncts.json  → data.Conjuncts
│   ├── characters.json      # 50 characters (vowels + consonants)
│   └── conjuncts.json       # 170 conjuncts sorted by first component
└── handlers/
    ├── characters.go        # GetCharacters, GetCharacterByID
    └── conjuncts.go         # GetConjuncts, GetConjunctByID
```
