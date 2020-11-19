package infrastructures

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PSQLHandler struct {
	Conn *pgxpool.Pool
}

func (handler *PSQLHandler) GetPool() *pgxpool.Pool {
	return handler.Conn
}

func initPostgresql() (*pgxpool.Pool, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	poolMaxConns := os.Getenv("DB_POOL_MAX_CONNS")

	/*
	*	Example DSN
	* 		user=jack password=secret host=pg.example.com
	*		port=5432 dbname=mydb sslmode=verify-ca pool_max_conns=10
	*	Example URL
	*		postgres://jack:secret@pg.example.com:5432/mydb?sslmode=verify-ca&pool_max_conns=10
	 */
	dbUrl := "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + name
	dbUrl += "?pool_max_conns=" + poolMaxConns

	config, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		return nil, fmt.Errorf("postgres connection: %v", err)
	}

	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		migrate()
		initSchema()

		fmt.Println("Postgres connected:")
		fmt.Println(
			"HOST:", host, "USER:", user, "PASSWORD:", password,
			"NAME:", name, "PORT:", port, "POOL_MAX_CONNS:", poolMaxConns,
		)
		return nil
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)

	return pool, err
}

func migrate() {
	// TODO: migrate DB
}
func initSchema() {
	// TODO: init schema
}
