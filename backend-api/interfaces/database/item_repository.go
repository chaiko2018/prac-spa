package database

import(
  "prac-spa/backend-api/domain"
)

type ItemRepository struct {
  SqlHandler
}

func (repo *ItemRepository) FindById(identifier int) (item *domain.Item, err error) {
  row, err := repo.Query("SELECT id, name, description, amount FROM items WHERE id = ?", identifier)
  if err != nil{
    return
  }
  defer row.Close()

  var id int
  var name string
  var description string
  var amount int64
  row.Next()
  if err = row.Scan(&id, &name, &description, &amount); err != nil {
    return
  }
  item.ID = id
  item.Name = name
  item.Description = description
  item.Amount = amount
  return
}

