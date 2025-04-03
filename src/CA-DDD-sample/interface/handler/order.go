package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/lee212400/myProject/domain/entity"
	"github.com/lee212400/myProject/usecase"
	"github.com/lee212400/myProject/usecase/dto"
)

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

type OrderHandler struct {
	orderService usecase.OrderService
}

func NewOrderHandler(orderService usecase.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (o *OrderHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestOrder
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		dto := dto.PostInputDto{
			Order: &entity.Order{
				Id:     req.OrderId,
				UserId: req.UserId,
			},
		}

		err := o.orderService.Save(&dto)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusCreated, nil)
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
