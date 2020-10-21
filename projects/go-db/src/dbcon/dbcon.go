package dbcon

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/pgxpool"
)

// Connect to db
func Connect() (*pgxpool.Pool, error) {
	dbURL := os.Getenv("DATABASE_URL")
	fmt.Println("DatabaseUrl", dbURL)

	return pgxpool.Connect(context.Background(), dbURL)
}
