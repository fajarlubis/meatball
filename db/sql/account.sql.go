package sql

type Customer struct {
	ID          string `json:"_id"`
	ExternalID  string `json:"external_id"`
	FirstName   string `json:"first_name"`
	MiddleName  string `json:"middle_name"`
	LastName    string `json:"last_name"`
	LegalName   string `json:"legal_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func InsertOne() error {
	NewConnection()

	_, err := db.pool.Exec(`INSERT INTO customer(
		external_id, first_name, middle_name, last_name, legal_name, email, phone_number) 
		VALUES($1, $2, $3, $4, $5, $6, $7)`,
		"cust-01",
		"Fajar",
		"Sidiq",
		"Lubis",
		"Fajar Sidiq Lubis",
		"fajarsidiqlubis13@gmail.com",
		"082363680865",
	)

	return err
}
