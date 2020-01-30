package database

type SqlHandler interface {
  Query(string, ...interface{}) (Result, error)
}

type Row interface {
  Scan(...interface{}) error
  Next() bool
  Close() error
}

