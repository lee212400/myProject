package handler

import (
	"net/http"

	"github.com/lee212400/myProject/application"
	"github.com/lee212400/myProject/domain/model"

	"github.com/labstack/echo"
)

type UserHandler interface {
	Post() echo.HandlerFunc
	Get() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	Put() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type OrderHandler struct {
	orderService *application.OrderService
}

func NewOrderHandler(orderService *application.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

type requestOrderItem struct {
	ProductID string  `json:"ProductID"`
	Quantity  int     `json:"Quantity"`
	Price     float64 `json:"Price"`
}

type requestOrder struct {
	OrderId    string             `json:"orderId"`
	UserId     string             `json:"userId"`
	OrderItems []requestOrderItem `json:"orderItems"`
}

type responsOrder struct {
	OrderId    string             `json:"orderId"`
	UserId     string             `json:"userId"`
	OrderItems []requestOrderItem `json:"orderItems"`
}

func (o *OrderHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestOrder
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		oItems := []*model.OrderItem{}

		err := o.orderService.Save(req.OrderId, req.UserId, oItems)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responsOrder{}

		return c.JSON(http.StatusCreated, res)
	}
}

func (o *OrderHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		order, err := o.orderService.FindByID(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responsOrder{
			OrderId: order.Id,
			UserId:  order.UserId,
		}

		return c.JSON(http.StatusOK, res)
	}
}
