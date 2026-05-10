# API Reference

Base URL: `http://localhost:8080`

---

## Characters

### `GET /api/characters`

Returns the full list of Bengali alphabet characters (vowels and consonants).

**Response** `200 OK`
```json
[
  {
    "id": 1,
    "type": "vowel",
    "group": "vowel",
    "char": "অ",
    "romanization": "o",
    "example_word": "অজ (oj) — goat",
    "vowel_sign": "া",
    "vowel_sign_example_word": "কাজ (kaj) — work"
  }
]
```

| Field | Type | Description |
|---|---|---|
| `id` | int | Unique identifier (1–50) |
| `type` | string | `"vowel"` or `"consonant"` |
| `group` | string | Phonetic group: `vowel`, `velar`, `palatal`, `retroflex`, `dental`, `labial`, `other` |
| `char` | string | Bengali Unicode character |
| `romanization` | string | Approximate Latin transliteration |
| `example_word` | string | Example word using this character |
| `vowel_sign` | string | Vowel diacritic form — vowels only, omitted for অ |
| `vowel_sign_example_word` | string | Example word using the vowel sign — omitted when `vowel_sign` absent |

---

### `GET /api/characters/:id`

Returns a single character by ID.

**Path parameter:** `id` — integer, must be ≥ 1

**Response** `200 OK` — same shape as a single element from the list above

**Errors**

| Status | Condition | Body |
|---|---|---|
| `400 Bad Request` | `id` is not a valid positive integer | `{"error": "invalid id"}` |
| `404 Not Found` | No character with that ID | `{"error": "character not found"}` |

---

## Conjuncts

### `GET /api/conjuncts`

Returns the full list of Bengali conjunct consonants (যুক্তব্যঞ্জন), sorted by first component in Bengali alphabet order.

**Response** `200 OK`
```json
[
  {
    "id": 3,
    "char": "ক্ত",
    "components": ["ক", "ত"],
    "romanization": "kta",
    "example_word": "রক্ত (rakto) — blood"
  }
]
```

| Field | Type | Description |
|---|---|---|
| `id` | int | Unique identifier (1–170) |
| `char` | string | Rendered conjunct glyph |
| `components` | []string | Constituent consonants in order (2 or 3 elements) |
| `romanization` | string | Approximate Latin transliteration |
| `example_word` | string | Example word containing this conjunct |

---

### `GET /api/conjuncts/:id`

Returns a single conjunct by ID.

**Path parameter:** `id` — integer, must be ≥ 1

**Response** `200 OK` — same shape as a single element from the list above

**Errors**

| Status | Condition | Body |
|---|---|---|
| `400 Bad Request` | `id` is not a valid positive integer | `{"error": "invalid id"}` |
| `404 Not Found` | No conjunct with that ID | `{"error": "conjunct not found"}` |

---

## Frontend

### `GET /`

Serves the embedded test frontend (`web/index.html`). Calls `/api/characters` and `/api/conjuncts` on load.
