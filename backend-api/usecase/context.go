package usecase

type Context interface {
  Param(string) string
  Bind(interface{}) error
  Status(int64)
  JSON(int64, interface{})
}
