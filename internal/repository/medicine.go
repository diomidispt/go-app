package repository

import (
	"database/sql" // gives us the DB connection type and query methods

	"github.com/diomidispt/go-app/internal/model" // the Medicine struct we defined
)

// MedicineRepository holds the database connection
// all medicine SQL queries live here
type MedicineRepository struct {
	db *sql.DB // the open connection to Postgres — passed in when we create the repository
}

// NewMedicineRepository creates a new repository with the DB connection
// this is called once at startup and reused for every request
func NewMedicineRepository(db *sql.DB) *MedicineRepository {
	return &MedicineRepository{db: db}
}

// GetAll returns every medicine from the database
func (r *MedicineRepository) GetAll() ([]model.Medicine, error) {
	rows, err := r.db.Query("SELECT id, name, dosage, stock, price FROM medicines")
	if err != nil {
		return nil, err
	}
	defer rows.Close() // always close rows when done to free the DB connection

	var medicines []model.Medicine // start with an empty list
	for rows.Next() {              // loop through each row returned
		var m model.Medicine
		rows.Scan(&m.ID, &m.Name, &m.Dosage, &m.Stock, &m.Price) // read the columns into the struct fields
		medicines = append(medicines, m)                          // add to the list
	}
	return medicines, nil
}

// GetByID returns a single medicine by its ID
func (r *MedicineRepository) GetByID(id int) (model.Medicine, error) {
	var m model.Medicine
	row := r.db.QueryRow("SELECT id, name, dosage, stock, price FROM medicines WHERE id = $1", id) // $1 is a placeholder — prevents SQL injection
	err := row.Scan(&m.ID, &m.Name, &m.Dosage, &m.Stock, &m.Price)
	return m, err
}

// Create inserts a new medicine and returns it with the generated ID
func (r *MedicineRepository) Create(m model.Medicine) (model.Medicine, error) {
	err := r.db.QueryRow(
		"INSERT INTO medicines (name, dosage, stock, price) VALUES ($1, $2, $3, $4) RETURNING id",
		m.Name, m.Dosage, m.Stock, m.Price,
	).Scan(&m.ID) // RETURNING id gives us back the auto-generated ID Postgres assigned
	return m, err
}

// Update modifies an existing medicine by ID
func (r *MedicineRepository) Update(m model.Medicine) error {
	_, err := r.db.Exec(
		"UPDATE medicines SET name=$1, dosage=$2, stock=$3, price=$4 WHERE id=$5",
		m.Name, m.Dosage, m.Stock, m.Price, m.ID,
	)
	return err
}

// Delete removes a medicine by ID
func (r *MedicineRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM medicines WHERE id = $1", id)
	return err
}
