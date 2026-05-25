package model

// Medicine mirrors the medicines table in the database
// each field maps to a column — same names, same types
type Medicine struct {
	ID     int     `json:"id"`     // maps to the id column — json tag controls what the field is called in API responses
	Name   string  `json:"name"`   // maps to the name column
	Dosage string  `json:"dosage"` // maps to the dosage column
	Stock  int     `json:"stock"`  // maps to the stock column
	Price  float64 `json:"price"`  // maps to the price column
}
