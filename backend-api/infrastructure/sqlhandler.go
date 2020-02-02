package infrastructure

import(
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "../interfaces/database"
)

type SqlHandler struct{
  Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
  conn, err := sql.Open("mysql", "root:@tcp(db:3306)/sample")
  if err != nil{
    panic(err.Error)
  }
  
  sqlHandler := new(SqlHandler)
  sqlHandler.Conn = conn
  return sqlHandler
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
  rows, err := handler.Conn.Query(statement, args...)
  if err != nil{
    return new(SqlRow), err
  }
  row := new(SqlRow)
  row.Rows = rows
  return row, nil
}

type SqlRow struct{
  Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) error {
  return r.Rows.Scan(dest)
}

func (r SqlRow) Next() bool {
  return r.Rows.Next()
}

func (r SqlRow) Close() error {
  return r.Rows.Close()
}
