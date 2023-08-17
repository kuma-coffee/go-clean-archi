package entities

type Book struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Year string `json:"year"`
}
