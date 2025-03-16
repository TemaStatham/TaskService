package handler

import (
	"github.com/TemaStatham/TaskService/client/pkg/app/comment/data"
	"github.com/TemaStatham/TaskService/client/pkg/infrastructure/middleware/auth"
	"github.com/TemaStatham/TaskService/client/pkg/infrastructure/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getComments(c *gin.Context) {
	_, err := auth.GetUserId(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input data.ShowComment

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, InvalidInputBodyErr)
		return
	}

	pag, err := h.commentQuery.Show(c.Request.Context(), input)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"paginate": pag,
	})
}

func (h *Handler) addComments(c *gin.Context) {
	authUser, err := auth.GetUserId(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input data.CreateComment

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, InvalidInputBodyErr)
		return
	}

	id, err := h.commentService.Create(c.Request.Context(), input, authUser)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
