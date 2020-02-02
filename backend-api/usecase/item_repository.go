package usecase

import(
  "../domain"
)

type ItemRepository interface {
  FindAll() (domain.Items, error)
  FindById(int) (domain.Item, error)
}
