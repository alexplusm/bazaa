package dbcon

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Connect to db
func Connect() (*pgxpool.Pool, error) {
	dbURL := os.Getenv("DATABASE_URL")
	fmt.Println("DatabaseUrl", dbURL)

	return pgxpool.Connect(context.Background(), dbURL)
}
