package infrastructures

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

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
	dbUrl := getDbURL()

	config, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		return nil, fmt.Errorf("postgres connection: %v", err)
	}

	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		err := initSchema(ctx, conn)
		if err != nil {
			return fmt.Errorf("init postgres: %v", err)
		}
		// TODO: migrate()
		log.Info("Postgres connected:", dbUrl)
		return nil
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)

	return pool, err
}

func migrate() {
	// TODO: migrate DB
}

func initSchema(ctx context.Context, conn *pgx.Conn) error {
	file, err := ioutil.ReadFile("sql/schema.sql")
	if err != nil {
		return fmt.Errorf("init schema: read schema: %v", err)
	}
	statements := strings.Split(string(file), ";")

	for _, statement := range statements {
		_, err := conn.Exec(ctx, statement)
		if err != nil {
			return fmt.Errorf("init schema: %v", err)
		}
	}

	return nil
}

func getDbURL() string {
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

	return dbUrl
}
