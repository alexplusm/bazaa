package dao

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

// ImageDao image dao
type ImageDao struct {
	URL      string `json:"url"`
	Category string `json:"category"`
}

// InsertImage insert image
func InsertImage(p *pgxpool.Pool, image ImageDao) error {
	conn, err := p.Acquire(context.Background())
	if err != nil {
		fmt.Printf("Unable to acquire a database connection: %v\n", err) // todo: logger
		// w.WriteHeader(500)
		return err
	}
	defer conn.Release()

	row := conn.QueryRow(context.Background(),
		`INSERT INTO images ("url", category) VALUES ($1, $2) RETURNING image_id`,
		image.URL, image.Category)

	var id string
	err = row.Scan(&id)
	if err != nil {
		fmt.Printf("Unable to INSERT: %v\n", err)
		// w.WriteHeader(500)
		return err
	}
	fmt.Println("created image", id)

	return nil
}
