package controllers

import(
  "../../interfaces/database"
  "../../usecase"
  "strconv"
  "net/http"
)

type ItemController struct {
  Interactor  usecase.ItemInteractor
}

func NewItemController(sqlHandler database.SqlHandler) *ItemController {
  return &ItemController{
    Interactor: usecase.ItemInteractor{
      ItemRepository: &database.ItemRepository{
        SqlHandler: sqlHandler,
      },
    },
  }
}


func (controller *ItemController) Index(c Context) {
  items, err := controller.Interactor.Items()
  if err != nil{
    c.JSON(500, NewError(err))
    return
  }
  c.JSON(200, items)
}

func (controller *ItemController) Show(c Context) {
  id, _ := strconv.Atoi(c.Param("id"))
  item, err := controller.Interactor.ItemById(id)
  if err != nil {
    c.JSON(500, NewError(err))
    return
  }
  c.JSON(200, item)
}
