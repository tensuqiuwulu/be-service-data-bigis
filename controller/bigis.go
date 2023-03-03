package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-data-bigis/models/request"
	"github.com/tensuqiuwulu/be-service-data-bigis/models/response"
	"github.com/tensuqiuwulu/be-service-data-bigis/service"
)

type BigisControllerInterface interface {
	FindResponseByNik(c echo.Context) error
}

type BigisControllerImplementation struct {
	Logger                *logrus.Logger
	BigisServiceInterface service.BigisServiceInterface
}

func NewBigisController(
	logger *logrus.Logger,
	bigisServiceInterface service.BigisServiceInterface,
) BigisControllerInterface {
	return &BigisControllerImplementation{
		BigisServiceInterface: bigisServiceInterface,
	}
}

func (controller *BigisControllerImplementation) FindResponseByNik(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	request := request.ReadFromFindBigisResponsesRequestBody(c, requestId, controller.Logger)
	bigisResponse := controller.BigisServiceInterface.FindResponseByNik(requestId, request)
	responses := response.Response{Code: 200, Mssg: "success", Data: bigisResponse, Error: []string{}}
	return c.JSON(http.StatusOK, responses)
}
