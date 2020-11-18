package infrastructures

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func initPostgresql() (*pgxpool.Pool, error) {
	dbUrl := os.Getenv("DATABASE_URL")
	config, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		return nil, nil
	}

	// config.MaxConns = 10 // todo: From config
	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		fmt.Println("Connected:", dbUrl)
		return nil
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)

	return pool, err
}

//func initDatabase(dbHost string, dbUser string, dbPass string, dbName string, dbPort uint16, maxConnectionsInPool int) (*pgx.ConnPool, error) {
//var successOrFailure string = "OK"
//
//var config pgx.ConnPoolConfig
//
//config.Host = dbHost
//config.User = dbUser
//config.Password = dbPass
//config.Database = dbName
//config.Port = dbPort
//
//config.MaxConnections = maxConnectionsInPool
//
//config.AfterConnect = func(conn *pgx.Conn) error {
//	worldSelectStmt = mustPrepare(conn, "worldSelectStmt", "SELECT id, randomNumber FROM World WHERE id = $1")
//	worldUpdateStmt = mustPrepare(conn, "worldUpdateStmt", "UPDATE World SET randomNumber = $1 WHERE id = $2")
//	fortuneSelectStmt = mustPrepare(conn, "fortuneSelectStmt", "SELECT id, message FROM Fortune")
//
//	// Disable synchronous commit for the current db connection
//	// as a performance optimization.
//	// See http://www.postgresql.org/docs/current/static/runtime-config-wal.html
//	// for details.
//	if _, err := conn.Exec("SET synchronous_commit TO OFF"); err != nil {
//		log.Fatalf("Error when disabling synchronous commit")
//	}
//
//	return nil
//}
//
//fmt.Println("--------------------------------------------------------------------------------------------")
//
//connPool, err := pgx.NewConnPool(config)
//if err != nil {
//	successOrFailure = "FAILED"
//	log.Println("Connecting to database ", dbName, " as user ", dbUser, " ", successOrFailure, ": \n ", err)
//} else {
//	log.Println("Connecting to database ", dbName, " as user ", dbUser, ": ", successOrFailure)
//
//	log.Println("Fetching one record to test if db connection is valid...")
//	var w World
//	n := randomWorldNum()
//	if errPing := connPool.QueryRow("worldSelectStmt", n).Scan(&w.Id, &w.RandomNumber); errPing != nil {
//		log.Fatalf("Error scanning world row: %s", errPing)
//	}
//	log.Println("OK")
//}
//
//fmt.Println("--------------------------------------------------------------------------------------------")
//
//return connPool, err
//}

// ----------------

type PSQLHandler struct {
	Conn *pgxpool.Pool
}

// TODO: Temporary: look at interfaces.IDBHandler
func (handler *PSQLHandler) GetPool() *pgxpool.Pool {
	return handler.Conn
}

// TODO: try to create some generic methods: see SQLiteHandler example
//// Example for SQLiteHandler
//func (handler *SQLiteHandler) Query(statement string) (interfaces.IRow, error) {
//	//fmt.Println(statement)
//	rows, err := handler.Conn.Query(statement)
//
//	if err != nil {
//		fmt.Println(err)
//		return new(SqliteRow),err
//	}
//	row := new(SqliteRow)
//	row.Rows = rows
//
//	return row, nil
//}
//
//type SqliteRow struct {
//	Rows *sql.Rows
//}
//
//func (r SqliteRow) Scan(dest ...interface{}) error {
//	err := r.Rows.Scan(dest...)
//	if err != nil {
//		return err
//	}
//
//	return  nil
//}
//
//func (r SqliteRow) Next() bool {
//	return r.Rows.Next()
//}
