package infrastructures

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type PSQLHandler struct {
	Conn *pgxpool.Pool
}

// TODO: Temporary: look at interfaces.IDBHandler
func (handler *PSQLHandler) GetConn() *pgxpool.Pool {
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
