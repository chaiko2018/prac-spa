package infrastructure

import(
  gin "gopkg.in/gin-gonic/gin.v1"
  "../interfaces/controllers"
)

var Router *gin.Engine

func init() {
  router := gin.Default()

  itemController := controllers.NewItemController(NewSqlHandler())

  router.GET("/", func(c *gin.Context) { itemController.Index(c) }
  router.GET("/", func(c *gin.Context) { itemController.Show(c) }

  Router = router
}
