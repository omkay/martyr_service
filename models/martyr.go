package models

type Martyr struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	CauseOfDeath string `json:"cause_of_death"`
	DateOfDeath string `json:"date_of_death"`
	PlaceOfDeath string `json:"place_of_death"`
	Description string `json:"description"`
	ImageUrl string `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

