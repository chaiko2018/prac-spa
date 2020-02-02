package usecase

import(
  "../domain"
)


type ItemInteractor struct {
  ItemRepository  ItemRepository
}

func (interactor *ItemInteractor) Items() (items *domain.Items, err error) {
  items, err := interactor.ItemRepository.FindAll()
  return
}

func (interactor *ItemInteractor) ItemById(identifier int) (item *domain.Item, err error) {
  item, err := interactor.ItemRepository.FindById(identifier)
  return
}
