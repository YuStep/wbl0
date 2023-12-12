package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Get All Orders
// @Tags Orders
// @Description get all orders
// @ID get-all-orders
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllOrdersResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /order [get]
func (h *Handler) getAllOrders(c *gin.Context) {
	strLimit := c.Query("limit")
	strOffset := c.Query("offset")
	var (
		limit     int = 10
		errLimit  error
		offset    int = 0
		errOffset error
	)
	if strLimit != "" {
		limit, errLimit = strconv.Atoi(strLimit)
	}
	if strOffset != "" {
		offset, errOffset = strconv.Atoi(c.Query("offset"))
	}
	if errLimit != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid limit param")
		return
	}
	if errOffset != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid offset param")
		return
	}
	orders := h.services.Order.GetAll(limit, offset)
	allOrdersCount := h.services.Order.GetOrdersCount()
	//c.JSON(http.StatusOK, getAllOrdersResponse{
	//	Data: orders,
	//})

	//c.HTML(http.StatusOK, "index.tmpl", gin.H{"orderList": getAllOrdersResponse{
	//	Data: orders,
	//}})
	c.JSON(http.StatusOK, getAllOrdersResponse{
		Data:  orders,
		Count: allOrdersCount,
	})
}

// Login godoc
//
//	@Summary Get Order By Id
//	@Tags Orders
//	@Description get order by id
//	@ID get-order-by-id
//	@Accept  json
//	@Produce  json
//	@Success 200 {object} models.Order
//	@Failure 400,404 {object} errorResponse
//	@Failure 500 {object} errorResponse
//	@Failure default {object} errorResponse
//	@Router /order/:id [get]
func (h *Handler) getOrderById(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.Order.GetById(id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getOrderByIdResponse{
		Order: list,
	})
}
