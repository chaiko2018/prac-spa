package infrastructure

import(
  gin "github.com/gin-gonic/gin"
  "../interfaces/controllers"
)

var Router *gin.Engine

func init() {
  router := gin.Default()

  itemController := controllers.NewItemController(NewSqlHandler())

  router.GET("/api/v1/items", func(c *gin.Context) { itemController.Index(c) }
  router.GET("/api/v1/items/:id", func(c *gin.Context) { itemController.Show(c) }

  Router = router
}

