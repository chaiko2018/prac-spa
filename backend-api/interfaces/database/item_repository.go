package database

import(
  "../../domain"
)

type ItemRepository struct {
  SqlHandler
}

func (repo *ItemRepository) FindById(identifier int) (item domain.Item, err error) {
  row, err := repo.Query("SELECT id, name, description, amount FROM items WHERE id = ?", identifier)
  if err != nil{
    return
  }
  defer row.Close()

  var id int64
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

func (repo *ItemRepository) FindAll() (items domain.Items, err error) {
  rows, err := repo.Query("SELECT id, name, description, amount FROM items")
  if err != nil {
    return
  }
  defer rows.Close()

  for rows.Next() {
    var id int64
    var name string
    var description string
    var amount int64

    if err := rows.Scan(&id, &name, &description, &amount); err != nil{
      continue
    }

    item := domain.Item{
      ID:           id,
      Name:         name,
      Description:  description,
      Amount:       amount,
    }

    items = append(items, item)
  }
  return
}


