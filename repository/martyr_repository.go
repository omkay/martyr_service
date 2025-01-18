package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MartyrRepository struct {
	dbpool *pgxpool.Pool
}

func NewMartyrRepository(dbpool *pgxpool.Pool) *MartyrRepository {
	return &MartyrRepository{dbpool}
}

func (pr *MartyrRepository) CreateMartyr(martyr *models.Martyr) error {
	query := "INSERT INTO martyrs (first_name, last_name, date_of_birth, cause_of_death, date_of_death, place_of_death, description, image_url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	return pr.dbpool.QueryRow(context.Background(), query, martyr.FirstName, martyr.LastName, martyr.DateOfBirth, martyr.CauseOfDeath, martyr.DateOfDeath, martyr.PlaceOfDeath, martyr.Description, martyr.ImageUrl).
		Scan(&martyr.ID)
}

func (pr *MartyrRepository) GetMartyrs(ctx context.Context) ([]models.Martyr, error) {
	rows, err := pr.dbpool.Query(
		ctx,
		"SELECT id, first_name, last_name, date_of_birth, cause_of_death, date_of_death, place_of_death, description, image_url FROM martyrs",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var martyr models.Martyr
		if err := rows.Scan(&martyr.ID, &martyr.FirstName, &martyr.LastName, &martyr.DateOfBirth, &martyr.CauseOfDeath, &martyr.DateOfDeath, &martyr.PlaceOfDeath, &martyr.Description, &martyr.ImageUrl); err != nil {
			return nil, err
		}
		martyrs = append(martyrs, martyr)
	}

	return martyrs, nil
}

func (pr *MartyrRepository) GetMartyr(id int) (*models.Martyr, error) {
	var martyr models.Martyr
	query := "SELECT id, first_name, last_name, date_of_birth, cause_of_death, date_of_death, place_of_death, description, image_url FROM martyrs WHERE id = $1"
	err := pr.dbpool.QueryRow(context.Background(), query, id).
		Scan(&martyr.ID, &martyr.FirstName, &martyr.LastName, &martyr.DateOfBirth, &martyr.CauseOfDeath, &martyr.DateOfDeath, &martyr.PlaceOfDeath, &martyr.Description, &martyr.ImageUrl)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil // Or a custom "not found" error
		}
		return nil, err
	}

	return &martyr, nil
}

func (pr *MartyrRepository) UpdateMartyr(martyr *models.Martyr) error {
	query := "UPDATE martyrs SET first_name = $1, last_name = $2, date_of_birth = $3, cause_of_death = $4, date_of_death = $5, place_of_death = $6, description = $7, image_url = $8 WHERE id = $9"
	_, err := pr.dbpool.Exec(
		context.Background(),
		query,
		martyr.FirstName,
		martyr.LastName,
		martyr.DateOfBirth,
		martyr.CauseOfDeath,
		martyr.DateOfDeath,
		martyr.PlaceOfDeath,
		martyr.Description,
		martyr.ImageUrl,
		martyr.ID,
	)
	return err
}

func (pr *MartyrRepository) DeleteMartyr(id int) error {
	query := "DELETE FROM martyrs WHERE id = $1"
	_, err := pr.dbpool.Exec(context.Background(), query, id)
	return err
}

