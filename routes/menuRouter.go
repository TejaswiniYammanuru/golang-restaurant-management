package routes

import (
	controller "github.com/TejaswiniYammanuru/golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)



func MenuRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/menus", controller.GetMenus())
	incomingRoutes.GET("/menus/:menu_id",controller.GetMenu())
	incomingRoutes.POST("/menus",controller.CreateMenu())
	incomingRoutes.PATCH("/menus/:menu_id",controller.UpdateMenu())
	incomingRoutes.DELETE("/menus/:menu_id", controller.DeleteMenu())

}                  