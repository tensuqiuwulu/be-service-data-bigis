package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/tensuqiuwulu/be-service-data-bigis/controller"
)

func BigisRoute(e *echo.Echo, bigisControllerInterface controller.BigisControllerInterface) {
	group := e.Group("api/v1")
	group.POST("/response", bigisControllerInterface.FindResponseByNik)
}
