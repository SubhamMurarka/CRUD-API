package models

// creating Movie struct
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// creating director struct
type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
