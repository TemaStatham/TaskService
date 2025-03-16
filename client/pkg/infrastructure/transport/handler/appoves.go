package handler

import (
	"github.com/TemaStatham/TaskService/client/pkg/app/approve/data"
	"github.com/TemaStatham/TaskService/client/pkg/infrastructure/middleware/auth"
	"github.com/TemaStatham/TaskService/client/pkg/infrastructure/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) addApproves(c *gin.Context) {
	_, err := auth.GetUserId(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input data.CreateApprove

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, InvalidInputBodyErr)
		return
	}

	err = h.approveService.Create(c.Request.Context(), input)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
	})
}
