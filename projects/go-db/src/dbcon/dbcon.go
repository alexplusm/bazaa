package dbcon

import (
	// "context"
	"fmt"
	"os"

	// "github.com/jackc/pgx"
)

// Connect to db
func Connect() {
	dbURL := os.Getenv("DATABASE_URL")
	fmt.Println("DatabaseUrl", dbURL)
	// conn, err := pgx.Connect(context.Background(), dataBaseUrl)
	// fmt.Println("db connect", conn, err)
}
